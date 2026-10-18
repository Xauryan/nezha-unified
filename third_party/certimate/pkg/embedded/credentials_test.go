package embedded

import (
	"context"
	"encoding/json"
	"errors"
	"testing"
)

func TestCredentialResolutionRequiresCurrentScopedAuthority(t *testing.T) {
	e := testEngine(t)
	ctx := context.Background()
	r, err := e.CreateCredential(ctx, "1", CredentialInput{Name: "scoped", Provider: "cloudflare", Config: json.RawMessage(`{"apiToken":"synthetic-v1"}`)})
	if err != nil {
		t.Fatal(err)
	}
	use := CredentialUse{ActorID: 2, CredentialID: r.ID, Version: 1, Purpose: "dns.manage", ResourceKind: "dns_account", ResourceID: "7"}
	if _, err = e.ResolveCredential(ctx, use); err == nil {
		t.Fatal("resolver accepted missing authorizer")
	}
	revoked := false
	e.SetCredentialAuthorizer(func(_ context.Context, u CredentialUse) error {
		if revoked || u.ActorID != 2 || u.Purpose != "dns.manage" || u.ResourceID != "7" {
			return errors.New("denied")
		}
		return nil
	})
	cfg, err := e.ResolveCredential(ctx, use)
	if err != nil || cfg["apiToken"] != "synthetic-v1" {
		t.Fatalf("resolution: %v", err)
	}
	wrong := use
	wrong.ResourceID = "8"
	if _, err = e.ResolveCredential(ctx, wrong); err == nil {
		t.Fatal("target substitution")
	}
	revoked = true
	if _, err = e.ResolveCredential(ctx, use); err == nil {
		t.Fatal("revoked grant resolved")
	}
	revoked = false
	if _, err = e.UpdateCredential(ctx, "1", r.ID, 1, true, json.RawMessage(`{"apiToken":"synthetic-v2"}`)); err != nil {
		t.Fatal(err)
	}
	if _, err = e.ResolveCredential(ctx, use); err == nil {
		t.Fatal("old version resolved")
	}
	use.Version = 2
	cfg, err = e.ResolveCredential(ctx, use)
	if err != nil || cfg["apiToken"] != "synthetic-v2" {
		t.Fatalf("rotation: %v", err)
	}
	if _, err = e.UpdateCredential(ctx, "1", r.ID, 1, true, json.RawMessage(`{"apiToken":"stale"}`)); err == nil {
		t.Fatal("stale rotation won")
	}
	if _, err = e.UpdateCredential(ctx, "1", r.ID, 2, false, nil); err != nil {
		t.Fatal(err)
	}
	if _, err = e.ResolveCredential(ctx, use); err == nil {
		t.Fatal("disabled credential resolved")
	}
}

func TestCredentialResolutionRejectsReplacedAuthorizer(t *testing.T) {
	e := testEngine(t)
	ctx := context.Background()
	r, err := e.CreateCredential(ctx, "synthetic", CredentialInput{Name: "synthetic", Provider: "synthetic", Config: json.RawMessage(`{"token":"synthetic"}`)})
	if err != nil {
		t.Fatal(err)
	}
	e.SetCredentialAuthorizer(func(context.Context, CredentialUse) error {
		e.SetCredentialAuthorizer(nil)
		return nil
	})
	_, err = e.ResolveCredential(ctx, CredentialUse{ActorID: 1, CredentialID: r.ID, Version: 1, Purpose: "synthetic", ResourceKind: "credential", ResourceID: r.ID})
	if err == nil {
		t.Fatal("replaced authorizer still released material")
	}
}
