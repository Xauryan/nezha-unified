package embeddedhooks

import (
	"context"
	"errors"
	"testing"
)

func TestBoundExecutionCannotLoseAuthorization(t *testing.T) {
	Configure(&Hooks{})
	t.Cleanup(func() { Configure(nil) })
	ctx := Bind(context.Background())
	Configure(nil)
	if !errors.Is(Check(ctx, Operation{Action: "workflow.run"}), ErrDenied) {
		t.Fatal("in-flight embedded execution lost fail-closed policy")
	}
}

func TestCredentialAuthorizationAndExecutionContext(t *testing.T) {
	var seen Operation
	var resolved bool
	Configure(&Hooks{Authorize: func(_ context.Context, op Operation) error { seen = op; return nil }, ResolveCredential: func(_ context.Context, id, purpose string) (map[string]any, error) {
		resolved = true
		return map[string]any{"token": "synthetic-test-value"}, nil
	}})
	t.Cleanup(func() { Configure(nil) })
	ctx := WithOperation(Bind(context.Background()), Operation{WorkflowID: "workflow", RunID: "run", NodeID: "node", GraphDigest: "digest"})
	if _, err := Credential(ctx, "credential", "dns"); err != nil {
		t.Fatal(err)
	}
	if !resolved || seen.WorkflowID != "workflow" || seen.RunID != "run" || seen.NodeID != "node" || seen.GraphDigest != "digest" || seen.ResourceID != "credential" || seen.Action != "credential.use" {
		t.Fatalf("authorization context not preserved: %#v", seen)
	}
}

func TestRevocationAndMissingCallbacksDeny(t *testing.T) {
	revoked := false
	resolved := 0
	Configure(&Hooks{Authorize: func(context.Context, Operation) error {
		if revoked {
			return errors.New("synthetic private diagnostic")
		}
		return nil
	}, ResolveCredential: func(context.Context, string, string) (map[string]any, error) {
		resolved++
		return map[string]any{}, nil
	}})
	t.Cleanup(func() { Configure(nil) })
	ctx := Bind(context.Background())
	if _, err := Credential(ctx, "credential", "dns"); err != nil {
		t.Fatal(err)
	}
	revoked = true
	if _, err := Credential(ctx, "credential", "dns"); !errors.Is(err, ErrDenied) {
		t.Fatalf("revocation error: %v", err)
	}
	if resolved != 1 {
		t.Fatal("revoked credential reached resolver")
	}
	if err := DeployTarget(ctx, TargetDeployment{}); !errors.Is(err, ErrDenied) {
		t.Fatal("missing deployment callback did not deny")
	}
	if err := RunCreated(ctx, "workflow", "run"); !errors.Is(err, ErrDenied) {
		t.Fatal("missing run binding callback did not deny")
	}
	if got := SafeError(ctx, errors.New("synthetic private diagnostic")); got != ErrExecution {
		t.Fatal("provider error escaped")
	}
}

func TestGraphDigestMapOrdering(t *testing.T) {
	a, err := GraphDigest(map[string]any{"nodes": []any{}, "version": 1})
	if err != nil {
		t.Fatal(err)
	}
	b, err := GraphDigest(map[string]any{"version": 1, "nodes": []any{}})
	if err != nil {
		t.Fatal(err)
	}
	if a != b {
		t.Fatal("equivalent map graph hashes differ")
	}
}
