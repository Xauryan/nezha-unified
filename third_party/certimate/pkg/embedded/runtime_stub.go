//go:build !certimate_full

package embedded

import (
	"context"
	"github.com/pocketbase/pocketbase/core"
)

type runtimeState struct{}

func RuntimeAvailable() bool { return false }

func (e *Engine) NativeWorkflowStats(context.Context) (WorkflowStatistics, error) {
	return WorkflowStatistics{}, ErrRuntimeUnavailable
}

func (e *Engine) StartRuntime(context.Context, RuntimeHooks) error { return ErrRuntimeUnavailable }
func (e *Engine) StopRuntime(context.Context) error                { return nil }
func (e *Engine) StartWorkflow(context.Context, string) (string, error) {
	return "", ErrRuntimeUnavailable
}
func (e *Engine) CancelWorkflow(context.Context, string, string) error                   { return ErrRuntimeUnavailable }
func (e *Engine) reconcileNativeWrite(context.Context, string, *core.Record, bool) error { return nil }
