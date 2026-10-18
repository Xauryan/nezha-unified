package rpc

import (
	"context"
	"encoding/json"
	"github.com/nezhahq/nezha/pkg/agentproto"
	pb "github.com/nezhahq/nezha/proto"
	"github.com/nezhahq/nezha/service/agentcall"
	"github.com/nezhahq/nezha/service/singleton"
	"net/netip"
)

// UnifiedAgentCalls has independent capacity and cancellation from MCP calls.
var UnifiedAgentCalls = agentcall.New(func(ctx context.Context, id uint64, task *pb.Task) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}
	server, _ := singleton.ServerShared.Get(id)
	if server == nil || server.GetTaskStream() == nil {
		return agentcall.ErrOffline
	}
	return server.SendTask(task)
}, 64)

func deliverUnifiedResult(reporter uint64, result *pb.TaskResult) bool {
	if result.Type == agentproto.TaskHostAddresses {
		var report agentproto.HostAddresses
		if reporter == 0 || len(result.Data) > 65536 || json.Unmarshal([]byte(result.Data), &report) != nil || report.Protocol != agentproto.Protocol || len(report.Addresses) > 128 {
			return true
		}
		for _, a := range report.Addresses {
			if len(a.Interface) > 128 || a.Interface == "" {
				return true
			}
			ip, err := netip.ParseAddr(a.Address)
			if err != nil || !ip.IsGlobalUnicast() {
				return true
			}
		}
		notifyUnifiedHostAddresses(reporter, report.Addresses)
		return true
	}
	if result.Type != agentproto.TaskCapabilities && result.Type != agentproto.TaskCall {
		return false
	}
	UnifiedAgentCalls.Deliver(reporter, result)
	return true
}
