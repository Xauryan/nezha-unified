package main

import (
	"context"
	"encoding/json"
	"sync"
	"time"

	"github.com/nezhahq/agent/pkg/agentproto"
	"github.com/nezhahq/agent/pkg/certdeploy"
	pb "github.com/nezhahq/agent/proto"
)

var certificateExecutor = &certdeploy.Executor{Allowed: func() bool { return !loadRuntimeConfig().DisableCommandExecute }}
var certificateCalls = struct {
	sync.Mutex
	active  map[string]context.CancelFunc
	permits map[string]chan struct{}
}{active: map[string]context.CancelFunc{}, permits: map[string]chan struct{}{}}
var certificateSlots = make(chan struct{}, 4)

func (e agentTaskExecution) dispatchUnified(task *pb.Task) bool {
	if task.Type != agentproto.TaskCapabilities && task.Type != agentproto.TaskCall {
		return false
	}
	if task.Type == agentproto.TaskCapabilities {
		startUnifiedAddressReports(e.parent, e.send)
		b, _ := json.Marshal(agentproto.Capabilities{Protocol: agentproto.Protocol, CertificateDeploy: certificateExecutor.Allowed(), Online: true})
		_ = e.send(&pb.TaskResult{Id: task.Id, Type: task.Type, Successful: true, Data: string(b)})
		return true
	}
	// The optional extension has its own bounded worker pool; heartbeats and
	// original terminal/file tasks never wait behind a certificate command.
	go e.runUnified(task)
	return true
}
func (e agentTaskExecution) runUnified(task *pb.Task) {
	var req agentproto.Request
	response := func(phase, code string, final, success bool, payload any) {
		data, _ := json.Marshal(payload)
		out, _ := json.Marshal(agentproto.Response{Protocol: agentproto.Protocol, OperationID: req.OperationID, Phase: phase, Code: code, Final: final, Payload: data})
		_ = e.send(&pb.TaskResult{Id: task.Id, Type: task.Type, Data: string(out), Successful: success})
	}
	// This path never logs task.Data, including on panic or malformed input.
	defer func() {
		if recover() != nil {
			response("unknown", "unknown", true, false, nil)
		}
	}()
	if len(task.Data) > agentproto.MaxPayload || json.Unmarshal([]byte(task.Data), &req) != nil || req.Protocol != agentproto.Protocol || req.OperationID == "" {
		response("failed", "invalid", true, false, nil)
		return
	}
	if req.Method == "cancel" {
		certificateCalls.Lock()
		cancel := certificateCalls.active[req.OperationID]
		if cancel != nil {
			cancel()
		}
		certificateCalls.Unlock()
		response("cancel_requested", "", true, true, nil)
		return
	}
	if req.Method == "permit" {
		certificateCalls.Lock()
		ch := certificateCalls.permits[req.OperationID]
		if ch != nil {
			select {
			case ch <- struct{}{}:
			default:
			}
		}
		certificateCalls.Unlock()
		response("authorized", "", true, ch != nil, nil)
		return
	}
	if req.Method != "deploy" && req.Method != "reconcile" && req.Method != "inspect" && req.Method != "retire" {
		response("failed", "invalid", true, false, nil)
		return
	}
	select {
	case certificateSlots <- struct{}{}:
		defer func() { <-certificateSlots }()
	default:
		response("failed", "capacity", true, false, nil)
		return
	}
	var d agentproto.Deployment
	if json.Unmarshal(req.Payload, &d) != nil || d.OperationID != req.OperationID {
		response("failed", "invalid", true, false, nil)
		return
	}
	ctx, cancel := context.WithTimeout(e.parent, 110*time.Second)
	defer cancel()
	certificateCalls.Lock()
	if certificateCalls.active[req.OperationID] != nil {
		certificateCalls.Unlock()
		response("unknown", "conflict", true, false, nil)
		return
	}
	certificateCalls.active[req.OperationID] = cancel
	permit := make(chan struct{}, 1)
	certificateCalls.permits[req.OperationID] = permit
	certificateCalls.Unlock()
	defer func() {
		certificateCalls.Lock()
		delete(certificateCalls.active, req.OperationID)
		delete(certificateCalls.permits, req.OperationID)
		certificateCalls.Unlock()
	}()
	var receipt agentproto.Receipt
	var err error
	if req.Method == "inspect" {
		receipt, err = certificateExecutor.Inspect(ctx, d)
	} else if req.Method == "retire" {
		receipt, err = certificateExecutor.Retire(ctx, d)
	} else if req.Method == "reconcile" {
		receipt, err = certificateExecutor.Reconcile(ctx, d)
	} else {
		receipt, err = certificateExecutor.Execute(ctx, d, func(r agentproto.Receipt) { response(r.Phase, "", false, true, r) }, func(ctx context.Context) error {
			response("ready", "", false, true, nil)
			select {
			case <-permit:
				return ctx.Err()
			case <-ctx.Done():
				return ctx.Err()
			}
		})
	}
	code := ""
	if err != nil {
		code = "failed"
	}
	response(receipt.Phase, code, true, err == nil, receipt)
}
