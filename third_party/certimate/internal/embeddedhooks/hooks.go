// Package embeddedhooks supplies the host authorization boundary to the original
// workflow engine without exposing a second identity or management server.
package embeddedhooks

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"maps"
	"sync/atomic"
)

type Operation struct {
	GraphDigest  string
	WorkflowID   string
	RunID        string
	NodeID       string
	Action       string
	ResourceKind string
	ResourceID   string
	Domains      []string
}

type TargetDeployment struct {
	WorkflowID, RunID, NodeID, TargetID                         string
	TargetVersion                                               uint64
	CertificateID, CertificateVersion, Fingerprint, OperationID string
	CertificatePEM, PrivateKeyPEM                               string `json:"-"`
	Domains                                                     []string
}

type Hooks struct {
	OnRunCreated      func(context.Context, string, string) error
	DeployTarget      func(context.Context, TargetDeployment) error
	Authorize         func(context.Context, Operation) error
	ResolveCredential func(context.Context, string, string) (map[string]any, error)
}

var configured atomic.Pointer[Hooks]
var ErrDenied = errors.New("embedded workflow authorization denied")
var ErrExecution = errors.New("embedded workflow execution failed")

type contextKey struct{}
type boundHooks struct{ hooks *Hooks }
type operationKey struct{}

// Configure must be called before starting execution. An empty Hooks deliberately
// enables fail-closed embedded mode; nil restores the standalone upstream mode.
func Configure(h *Hooks) {
	if h == nil {
		configured.Store(nil)
		return
	}
	copied := *h
	configured.Store(&copied)
}
func Enabled() bool { return configured.Load() != nil }

// Bind captures the configured host for the whole invocation. Shutting down or
// replacing a runtime must never downgrade authorization for an in-flight run.
func Bind(ctx context.Context) context.Context {
	if _, ok := ctx.Value(contextKey{}).(boundHooks); ok {
		return ctx
	}
	return context.WithValue(ctx, contextKey{}, boundHooks{configured.Load()})
}
func current(ctx context.Context) *Hooks {
	if h, ok := ctx.Value(contextKey{}).(boundHooks); ok {
		return h.hooks
	}
	return configured.Load()
}
func EnabledFor(ctx context.Context) bool { return current(ctx) != nil }
func WithOperation(ctx context.Context, op Operation) context.Context {
	return context.WithValue(ctx, operationKey{}, op)
}
func OperationFromContext(ctx context.Context) (Operation, bool) {
	op, ok := ctx.Value(operationKey{}).(Operation)
	return op, ok
}
func Check(ctx context.Context, op Operation) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	h := current(ctx)
	if h == nil {
		return nil
	}
	if base, ok := ctx.Value(operationKey{}).(Operation); ok {
		if op.GraphDigest == "" {
			op.GraphDigest = base.GraphDigest
		}
		if op.WorkflowID == "" {
			op.WorkflowID = base.WorkflowID
		}
		if op.RunID == "" {
			op.RunID = base.RunID
		}
		if op.NodeID == "" {
			op.NodeID = base.NodeID
		}
	}
	if h.Authorize == nil {
		return ErrDenied
	}
	if err := h.Authorize(ctx, op); err != nil {
		return ErrDenied
	}
	return nil
}
func Credential(ctx context.Context, id, purpose string) (map[string]any, error) {
	h := current(ctx)
	if h == nil || h.ResolveCredential == nil {
		return nil, ErrDenied
	}
	if err := Check(ctx, Operation{Action: "credential.use", ResourceKind: "credential", ResourceID: id}); err != nil {
		return nil, err
	}
	config, err := h.ResolveCredential(ctx, id, purpose)
	if err != nil {
		return nil, ErrDenied
	}
	if config == nil {
		return nil, ErrDenied
	}
	return maps.Clone(config), nil
}

// SafeError prevents provider responses, URLs, scripts, and secrets from being
// copied to the persistent workflow error field or error template variables.
func SafeError(ctx context.Context, err error) error {
	if err == nil || !EnabledFor(ctx) {
		return err
	}
	if errors.Is(err, context.Canceled) {
		return context.Canceled
	}
	if errors.Is(err, context.DeadlineExceeded) {
		return context.DeadlineExceeded
	}
	if errors.Is(err, ErrDenied) {
		return ErrDenied
	}
	return ErrExecution
}

// GraphDigest uses canonical map JSON, independent of struct field order.
func GraphDigest(graph any) (string, error) {
	raw, err := json.Marshal(graph)
	if err != nil {
		return "", err
	}
	var canonical any
	if err := json.Unmarshal(raw, &canonical); err != nil {
		return "", err
	}
	raw, err = json.Marshal(canonical)
	if err != nil {
		return "", err
	}
	sum := sha256.Sum256(raw)
	return hex.EncodeToString(sum[:]), nil
}
func RunCreated(ctx context.Context, workflowID, runID string) error {
	h := current(ctx)
	if h == nil {
		return nil
	}
	if h.OnRunCreated == nil {
		return ErrDenied
	}
	if err := h.OnRunCreated(ctx, workflowID, runID); err != nil {
		return ErrDenied
	}
	return nil
}
func DeployTarget(ctx context.Context, request TargetDeployment) error {
	h := current(ctx)
	if h == nil || h.DeployTarget == nil {
		return ErrDenied
	}
	if err := Check(ctx, Operation{WorkflowID: request.WorkflowID, RunID: request.RunID, NodeID: request.NodeID, Action: "deployment.execute", ResourceKind: "deployment_target", ResourceID: request.TargetID, Domains: request.Domains}); err != nil {
		return err
	}
	return SafeError(ctx, h.DeployTarget(ctx, request))
}
