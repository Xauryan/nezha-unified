package agentcall

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/nezhahq/nezha/pkg/agentproto"
	pb "github.com/nezhahq/nezha/proto"
	"sync/atomic"
	"testing"
	"time"
)

func permit(context.Context) error { return nil }
func TestResultIdentityPurposeAndCorrelation(t *testing.T) {
	dispatched := make(chan *pb.Task, 1)
	b := New(func(_ context.Context, _ uint64, task *pb.Task) error { dispatched <- task; return nil }, 2)
	done := make(chan error, 1)
	go func() {
		_, err := b.Call(context.Background(), 42, agentproto.TaskCall, "certificate", agentproto.Request{Protocol: 1, OperationID: "op", Method: "deploy"}, permit, nil)
		done <- err
	}()
	task := <-dispatched
	data, _ := json.Marshal(agentproto.Response{Protocol: 1, OperationID: "op", Final: true})
	result := &pb.TaskResult{Id: task.Id, Type: task.Type, Successful: true, Data: string(data)}
	if b.Deliver(41, result) {
		t.Fatal("foreign reporter accepted")
	}
	result.Type++
	if b.Deliver(42, result) {
		t.Fatal("wrong task type accepted")
	}
	result.Type--
	if b.CancelPurpose("mcp") != 0 {
		t.Fatal("MCP cancels certificate calls")
	}
	if !b.Deliver(42, result) {
		t.Fatal("valid response rejected")
	}
	if err := <-done; err != nil {
		t.Fatal(err)
	}
	if b.Deliver(42, result) {
		t.Fatal("late replay accepted")
	}
}
func TestRevocationBeatsReplyAndWaitIsBounded(t *testing.T) {
	var allowed atomic.Bool
	allowed.Store(true)
	dispatched := make(chan *pb.Task, 1)
	b := New(func(_ context.Context, _ uint64, task *pb.Task) error { dispatched <- task; return nil }, 1)
	auth := func(context.Context) error {
		if !allowed.Load() {
			return ErrDenied
		}
		return nil
	}
	done := make(chan error, 1)
	go func() {
		_, err := b.Call(context.Background(), 1, agentproto.TaskCall, "certificate", agentproto.Request{OperationID: "x"}, auth, nil)
		done <- err
	}()
	task := <-dispatched
	_, err := b.Call(context.Background(), 1, agentproto.TaskCall, "certificate", agentproto.Request{}, permit, nil)
	if !errors.Is(err, ErrCapacity) {
		t.Fatal("unbounded inflight")
	}
	allowed.Store(false)
	data, _ := json.Marshal(agentproto.Response{Protocol: 1, OperationID: "x", Final: true})
	b.Deliver(1, &pb.TaskResult{Id: task.Id, Type: task.Type, Successful: true, Data: string(data)})
	if err = <-done; !errors.Is(err, ErrDenied) {
		t.Fatalf("revoked result: %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
	defer cancel()
	_, err = b.Call(ctx, 1, agentproto.TaskCapabilities, "probe", nil, permit, nil)
	if !errors.Is(err, ErrTimeout) {
		t.Fatalf("timeout: %v", err)
	}
}
