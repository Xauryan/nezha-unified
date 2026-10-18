package main

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"github.com/nezhahq/agent/pkg/agentproto"
	pb "github.com/nezhahq/agent/proto"
	"net"
	"net/netip"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var unifiedAddressSessions sync.Map

func startUnifiedAddressReports(ctx context.Context, send func(*pb.TaskResult) error) {
	if ctx.Done() == nil {
		return
	}
	if _, loaded := unifiedAddressSessions.LoadOrStore(ctx, true); loaded {
		return
	}
	go func() { defer unifiedAddressSessions.Delete(ctx); reportUnifiedAddresses(ctx, send) }()
}

func unifiedAddresses() []agentproto.HostAddress {
	// Linux exposes temporary, deprecated, tentative and DAD-failed address
	// flags. IPv6 without known flags is unsuitable for automated DDNS.
	flags := map[string]uint64{}
	if data, err := os.ReadFile("/proc/net/if_inet6"); err == nil {
		for _, line := range strings.Split(string(data), "\n") {
			fields := strings.Fields(line)
			if len(fields) != 6 {
				continue
			}
			raw, e1 := hex.DecodeString(fields[0])
			f, e2 := strconv.ParseUint(fields[4], 16, 32)
			if e1 != nil || e2 != nil || len(raw) != 16 {
				continue
			}
			var addr [16]byte
			copy(addr[:], raw)
			flags[fields[5]+":"+netip.AddrFrom16(addr).String()] = f
		}
	}
	out := []agentproto.HostAddress{}
	interfaces, _ := net.Interfaces()
	for _, nic := range interfaces {
		if nic.Flags&net.FlagUp == 0 || nic.Flags&net.FlagLoopback != 0 {
			continue
		}
		addresses, _ := nic.Addrs()
		for _, address := range addresses {
			prefix, err := netip.ParsePrefix(address.String())
			if err != nil {
				continue
			}
			ip := prefix.Addr().Unmap()
			if !ip.IsGlobalUnicast() {
				continue
			}
			a := agentproto.HostAddress{Interface: nic.Name, Address: ip.String()}
			if ip.Is6() {
				f, known := flags[nic.Name+":"+ip.String()]
				a.Temporary = !known || f&1 != 0
				a.Deprecated = f&(0x20|0x40|0x08) != 0
			}
			out = append(out, a)
			if len(out) == 128 {
				return out
			}
		}
	}
	return out
}
func reportUnifiedAddresses(ctx context.Context, send func(*pb.TaskResult) error) {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()
	for {
		data, _ := json.Marshal(agentproto.HostAddresses{Protocol: agentproto.Protocol, Addresses: unifiedAddresses()})
		if send(&pb.TaskResult{Type: agentproto.TaskHostAddresses, Successful: true, Data: string(data)}) != nil {
			return
		}
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
		}
	}
}
