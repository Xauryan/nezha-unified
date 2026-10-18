package certacme

import (
	"context"
	"errors"
	"net"
	"strings"
	"sync"
	"time"

	"github.com/miekg/dns"
)

// A provider commonly calls GetChallengeInfo again after the host has approved
// the effective FQDN. Pin CNAME responses for this issuance so that a DNS change
// cannot redirect that second lookup to an unapproved record. Other query types
// remain live for TXT propagation and provider zone discovery.
type pinnedResolver struct {
	ctx         context.Context
	nameservers []string
	mu          sync.Mutex
	cnames      map[string]*dns.Msg
}

func (p *pinnedResolver) ServeDNS(w dns.ResponseWriter, query *dns.Msg) {
	if len(query.Question) != 1 {
		reply := new(dns.Msg)
		reply.SetRcode(query, dns.RcodeFormatError)
		_ = w.WriteMsg(reply)
		return
	}
	question := query.Question[0]
	if question.Qtype == dns.TypeCNAME {
		p.mu.Lock()
		defer p.mu.Unlock()
		key := strings.ToLower(question.Name)
		if cached := p.cnames[key]; cached != nil {
			reply := cached.Copy()
			reply.Id = query.Id
			_ = w.WriteMsg(reply)
			return
		}
		reply := p.forward(query)
		// Cache failures as well: a failed policy lookup must not turn into a
		// successful redirection when the provider repeats the query.
		p.cnames[key] = reply.Copy()
		_ = w.WriteMsg(reply)
		return
	}
	_ = w.WriteMsg(p.forward(query))
}

func (p *pinnedResolver) forward(query *dns.Msg) *dns.Msg {
	for _, server := range p.nameservers {
		ctx, cancel := context.WithTimeout(p.ctx, 5*time.Second)
		client := &dns.Client{Timeout: 5 * time.Second}
		reply, _, err := client.ExchangeContext(ctx, query, server)
		if err == nil && reply.Truncated {
			client.Net = "tcp"
			reply, _, err = client.ExchangeContext(ctx, query, server)
		}
		cancel()
		if err == nil {
			return reply
		}
	}
	reply := new(dns.Msg)
	reply.SetRcode(query, dns.RcodeServerFailure)
	return reply
}

func newPinnedResolver(ctx context.Context, servers []string) (string, func(), error) {
	if len(servers) == 0 {
		config, err := dns.ClientConfigFromFile("/etc/resolv.conf")
		if err != nil {
			return "", nil, errors.New("DNS resolver configuration unavailable")
		}
		for _, ip := range config.Servers {
			servers = append(servers, net.JoinHostPort(ip, config.Port))
		}
	}
	if len(servers) == 0 {
		return "", nil, errors.New("DNS resolver configuration is empty")
	}
	upstream := make([]string, 0, len(servers))
	for _, server := range servers {
		if _, _, err := net.SplitHostPort(server); err != nil {
			server = net.JoinHostPort(strings.Trim(server, "[]"), "53")
		}
		upstream = append(upstream, server)
	}
	resolverCtx, cancel := context.WithCancel(ctx)
	lc := &net.ListenConfig{}
	udp, err := lc.ListenPacket(resolverCtx, "udp", "127.0.0.1:0")
	if err != nil {
		cancel()
		return "", nil, err
	}
	address := udp.LocalAddr().String()
	tcp, err := lc.Listen(resolverCtx, "tcp", address)
	if err != nil {
		_ = udp.Close()
		cancel()
		return "", nil, err
	}
	resolver := &pinnedResolver{ctx: resolverCtx, nameservers: upstream, cnames: make(map[string]*dns.Msg)}
	udpServer := &dns.Server{PacketConn: udp, Handler: resolver}
	tcpServer := &dns.Server{Listener: tcp, Handler: resolver}
	go func() { _ = udpServer.ActivateAndServe() }()
	go func() { _ = tcpServer.ActivateAndServe() }()
	return address, func() {
		cancel()
		closeCtx, done := context.WithTimeout(context.Background(), time.Second)
		defer done()
		_ = udpServer.ShutdownContext(closeCtx)
		_ = tcpServer.ShutdownContext(closeCtx)
		_ = udp.Close()
		_ = tcp.Close()
	}, nil
}
