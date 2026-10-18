//go:build certimate_full

package embedded

import (
	"context"
	"errors"
	"sync/atomic"
	"testing"
	"time"

	"github.com/certimate-go/certimate/internal/domain"
	"github.com/certimate-go/certimate/internal/embedcrypto"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/cron"
)

func runtimeTestWorkflow(t *testing.T, e *Engine, wait int) string {
	t.Helper()
	c, err := e.app.FindCollectionByNameOrId("workflow")
	if err != nil {
		t.Fatal(err)
	}
	r := core.NewRecord(c)
	r.Set("name", "synthetic runtime contract")
	r.Set("enabled", true)
	r.Set("trigger", "manual")
	r.Set("hasContent", true)
	nodes := []*domain.WorkflowNode{
		{Id: "start", Type: domain.WorkflowNodeTypeStart, Data: domain.WorkflowNodeData{Name: "start"}},
		{Id: "delay", Type: domain.WorkflowNodeTypeDelay, Data: domain.WorkflowNodeData{Name: "delay", Config: domain.WorkflowNodeConfig{"wait": wait}}},
		{Id: "end", Type: domain.WorkflowNodeTypeEnd, Data: domain.WorkflowNodeData{Name: "end"}},
	}
	if err := embedcrypto.SetJSON(e.app, r, "graphContent", &domain.WorkflowGraph{Nodes: nodes}); err != nil {
		t.Fatal(err)
	}
	if err := e.app.Save(r); err != nil {
		t.Fatal(err)
	}
	return r.Id
}

func runtimeTestHooks() RuntimeHooks {
	return RuntimeHooks{
		Authorize: func(context.Context, RuntimeOperation) error { return nil },
		ResolveCredential: func(context.Context, string, string) (map[string]any, error) {
			return nil, errors.New("synthetic test has no credentials")
		},
		OnRunCreated: func(context.Context, string, string) error { return nil },
	}
}

func waitRuntimeStatus(t *testing.T, e *Engine, id, want string) {
	t.Helper()
	deadline := time.Now().Add(5 * time.Second)
	for time.Now().Before(deadline) {
		r, err := e.app.FindRecordById("workflow_run", id)
		if err == nil && r.GetString("status") == want {
			return
		}
		time.Sleep(10 * time.Millisecond)
	}
	r, _ := e.app.FindRecordById("workflow_run", id)
	if r == nil {
		t.Fatalf("run %s missing", id)
	}
	t.Fatalf("run status %s, want %s", r.GetString("status"), want)
}

func TestRuntimeBindingPrecedesNodesAndDenialIsTerminal(t *testing.T) {
	e, err := Open(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	defer e.Close()
	id := runtimeTestWorkflow(t, e, 0)
	var bound atomic.Bool
	var nodes atomic.Int32
	h := runtimeTestHooks()
	h.OnRunCreated = func(context.Context, string, string) error { bound.Store(true); return nil }
	h.Authorize = func(_ context.Context, op RuntimeOperation) error {
		if op.Action == "workflow.run" && op.GraphDigest == "" {
			return errors.New("missing approved graph digest")
		}
		if op.NodeID != "" {
			if !bound.Load() {
				t.Error("node authorization ran before host binding")
			}
			nodes.Add(1)
			return errors.New("synthetic revoked grant")
		}
		return nil
	}
	if err := e.StartRuntime(context.Background(), h); err != nil {
		t.Fatal(err)
	}
	run, err := e.StartWorkflow(context.Background(), id)
	if err != nil {
		t.Fatal(err)
	}
	waitRuntimeStatus(t, e, run, "failed")
	if nodes.Load() == 0 || !bound.Load() {
		t.Fatal("execution never crossed the node authorization boundary")
	}
}

func TestRuntimeRejectsStartBeforePersistence(t *testing.T) {
	e, err := Open(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	defer e.Close()
	id := runtimeTestWorkflow(t, e, 0)
	h := runtimeTestHooks()
	h.Authorize = func(context.Context, RuntimeOperation) error { return errors.New("denied") }
	if err := e.StartRuntime(context.Background(), h); err != nil {
		t.Fatal(err)
	}
	if _, err := e.StartWorkflow(context.Background(), id); err == nil {
		t.Fatal("unauthorized run started")
	}
	count, err := e.app.CountRecords("workflow_run")
	if err != nil || count != 0 {
		t.Fatalf("unauthorized start persisted %d runs: %v", count, err)
	}
}

func TestRuntimeStopDrainsWorkersAndCanRestart(t *testing.T) {
	e, err := Open(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	defer e.Close()
	id := runtimeTestWorkflow(t, e, 60)
	entered := make(chan struct{}, 1)
	h := runtimeTestHooks()
	h.Authorize = func(_ context.Context, op RuntimeOperation) error {
		if op.NodeID == "delay" && op.Action == "node.execute" {
			select {
			case entered <- struct{}{}:
			default:
			}
		}
		return nil
	}
	if err := e.StartRuntime(context.Background(), h); err != nil {
		t.Fatal(err)
	}
	run, err := e.StartWorkflow(context.Background(), id)
	if err != nil {
		t.Fatal(err)
	}
	select {
	case <-entered:
	case <-time.After(5 * time.Second):
		t.Fatal("delay never started")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := e.StopRuntime(ctx); err != nil {
		t.Fatal(err)
	}
	waitRuntimeStatus(t, e, run, "canceled")
	if _, err := e.StartWorkflow(context.Background(), id); err == nil {
		t.Fatal("stopped dispatcher accepted work")
	}
	if err := e.StartRuntime(context.Background(), runtimeTestHooks()); err != nil {
		t.Fatal(err)
	}
	next := runtimeTestWorkflow(t, e, 0)
	nextRun, err := e.StartWorkflow(context.Background(), next)
	if err != nil {
		t.Fatal(err)
	}
	waitRuntimeStatus(t, e, nextRun, "succeeded")
}

func TestRuntimeScheduledCallbackRebindsAfterRestart(t *testing.T) {
	e, err := Open(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	defer e.Close()
	id := runtimeTestWorkflow(t, e, 0)
	r, err := e.app.FindRecordById("workflow", id)
	if err != nil {
		t.Fatal(err)
	}
	r.Set("trigger", "scheduled")
	r.Set("triggerCron", "* * * * *")
	if err := e.app.Save(r); err != nil {
		t.Fatal(err)
	}
	if err := e.StartRuntime(context.Background(), runtimeTestHooks()); err != nil {
		t.Fatal(err)
	}
	var oldJob *cron.Job
	for _, job := range e.app.Cron().Jobs() {
		if job.Id() == "workflow#"+id {
			oldJob = job
		}
	}
	if oldJob == nil {
		t.Fatal("scheduled job missing")
	}
	if err := e.StopRuntime(context.Background()); err != nil {
		t.Fatal(err)
	}
	created := make(chan string, 2)
	h := runtimeTestHooks()
	h.OnRunCreated = func(_ context.Context, _, runID string) error { created <- runID; return nil }
	if err := e.StartRuntime(context.Background(), h); err != nil {
		t.Fatal(err)
	}
	for _, job := range e.app.Cron().Jobs() {
		if job.Id() == "workflow#"+id {
			job.Run()
		}
	}
	select {
	case runID := <-created:
		waitRuntimeStatus(t, e, runID, "succeeded")
	case <-time.After(5 * time.Second):
		t.Fatal("restarted cron retained a stopped service callback")
	}
	if err := e.Close(); err != nil {
		t.Fatal(err)
	}
	// A callback captured before shutdown must safely no-op after DB release.
	oldJob.Run()
}
