package certacme

import (
	"context"
	"github.com/miekg/dns"
	"net"
	"sync/atomic"
	"testing"
)

func TestPinnedResolverPreventsCNAMEChangeDuringProviderCall(t *testing.T) {
	conn, err := net.ListenPacket("udp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	var alternate atomic.Bool
	server := &dns.Server{PacketConn: conn, Handler: dns.HandlerFunc(func(w dns.ResponseWriter, r *dns.Msg) {
		reply := new(dns.Msg)
		reply.SetReply(r)
		q := r.Question[0]
		if q.Qtype == dns.TypeCNAME {
			target := "approved.example.test."
			if alternate.Load() {
				target = "unauthorized.example.test."
			}
			reply.Answer = []dns.RR{&dns.CNAME{Hdr: dns.RR_Header{Name: q.Name, Rrtype: dns.TypeCNAME, Class: dns.ClassINET, Ttl: 0}, Target: target}}
		} else if q.Qtype == dns.TypeTXT {
			value := "before"
			if alternate.Load() {
				value = "after"
			}
			reply.Answer = []dns.RR{&dns.TXT{Hdr: dns.RR_Header{Name: q.Name, Rrtype: dns.TypeTXT, Class: dns.ClassINET, Ttl: 0}, Txt: []string{value}}}
		}
		_ = w.WriteMsg(reply)
	})}
	go func() { _ = server.ActivateAndServe() }()
	t.Cleanup(func() { _ = server.Shutdown(); _ = conn.Close() })
	address, closeResolver, err := newPinnedResolver(context.Background(), []string{conn.LocalAddr().String()})
	if err != nil {
		t.Fatal(err)
	}
	defer closeResolver()
	query := func(kind uint16) *dns.Msg {
		t.Helper()
		request := new(dns.Msg)
		request.SetQuestion("_acme-challenge.example.test.", kind)
		reply, _, err := (&dns.Client{}).Exchange(request, address)
		if err != nil {
			t.Fatal(err)
		}
		if len(reply.Answer) != 1 {
			t.Fatalf("unexpected answer count: %d", len(reply.Answer))
		}
		return reply
	}
	if got := query(dns.TypeCNAME).Answer[0].(*dns.CNAME).Target; got != "approved.example.test." {
		t.Fatal("initial CNAME mismatch")
	}
	if got := query(dns.TypeTXT).Answer[0].(*dns.TXT).Txt[0]; got != "before" {
		t.Fatal("initial TXT mismatch")
	}
	alternate.Store(true)
	if got := query(dns.TypeCNAME).Answer[0].(*dns.CNAME).Target; got != "approved.example.test." {
		t.Fatal("CNAME changed after authorization")
	}
	if got := query(dns.TypeTXT).Answer[0].(*dns.TXT).Txt[0]; got != "after" {
		t.Fatal("TXT propagation was incorrectly cached")
	}
}
