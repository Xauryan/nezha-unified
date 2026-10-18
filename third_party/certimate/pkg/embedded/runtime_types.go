package embedded

import (
	"context"
	"errors"
	"github.com/certimate-go/certimate/internal/domain"
	"github.com/certimate-go/certimate/internal/embedcrypto"
	"github.com/certimate-go/certimate/internal/embeddedhooks"
)

var ErrRuntimeUnavailable = errors.New("Certimate runtime requires a certimate_full build")

// RuntimeOperation carries execution identity and the immutable graph digest to
// the host's server-side authorization boundary before every side effect.
type RuntimeOperation = embeddedhooks.Operation
type RuntimeHooks = embeddedhooks.Hooks
type TargetDeployment = embeddedhooks.TargetDeployment

type WorkflowStatistics struct {
	Concurrency      int      `json:"concurrency"`
	PendingRunIDs    []string `json:"pendingRunIds"`
	ProcessingRunIDs []string `json:"processingRunIds"`
}

func RuntimeOperationFromContext(ctx context.Context) (RuntimeOperation, bool) {
	return embeddedhooks.OperationFromContext(ctx)
}

// WorkflowGraphDigest hashes the same published graph used by StartWorkflow.
// Draft changes do not silently alter the approved execution graph.
func (e *Engine) WorkflowGraphDigest(ctx context.Context, id string) (string, error) {
	e.mu.RLock()
	defer e.mu.RUnlock()
	if err := e.check(ctx); err != nil {
		return "", err
	}
	record, err := e.app.FindRecordById("workflow", id)
	if err != nil {
		return "", ErrNotFound
	}
	var graph *domain.WorkflowGraph
	if err := embedcrypto.ReadJSON(e.app, record, "graphContent", &graph); err != nil {
		return "", err
	}
	if graph == nil {
		return "", ErrInvalid
	}
	return embeddedhooks.GraphDigest(graph)
}
