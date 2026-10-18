package unified

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/certimate-go/certimate/pkg/embedded"
)

// A real provider client may retain a decrypted configuration while waiting on
// ACME. The public runtime hook must recheck every approved credential before
// the next effect, even when the changed CA/HTTP credential is not op.ResourceID.
func TestIdentityRuntimeRechecksResolvedCredentialsBeforeEveryEffect(t *testing.T) {
	for _, tc := range []struct {
		name      string
		secondary bool
		rotate    bool
	}{
		{name: "primary-disabled"},
		{name: "primary-rotated", rotate: true},
		{name: "retained-ca-disabled", secondary: true},
		{name: "retained-http-client-rotated", secondary: true, rotate: true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			s := newIdentityTestService(t)
			ctx := context.Background()
			admin := Principal{UserID: 1, Admin: true}
			credentials := make([]embedded.Resource, 2)
			for i, name := range []string{"synthetic-effect-provider", "synthetic-retained-ca-or-http-provider"} {
				var err error
				credentials[i], err = s.engine.CreateCredential(ctx, "1", embedded.CredentialInput{Name: name, Provider: "cloudflare", Config: json.RawMessage(`{"apiToken":"synthetic-retained-configuration"}`)})
				if err != nil {
					t.Fatal(err)
				}
			}
			const workflowID = "synthetic-effect-workflow"
			const runID = "synthetic-effect-accepted-run"
			wp := WorkflowPolicy{WorkflowID: workflowID, Version: 1, Enabled: true, RunAsUserID: 1, GraphDigest: "synthetic-effect-graph", Credentials: map[string]int{credentials[0].ID: credentials[0].Version, credentials[1].ID: credentials[1].Version}}
			if err := s.db.Create(&wp).Error; err != nil {
				t.Fatal(err)
			}
			if err := s.db.Create(&ResourceBinding{Kind: "workflow", ResourceID: workflowID, OwnerID: 1, Version: 1, Active: true}).Error; err != nil {
				t.Fatal(err)
			}
			if err := s.bindRun(WithPrincipal(ctx, admin), workflowID, runID); err != nil {
				t.Fatal(err)
			}
			retained := make([]map[string]any, 2)
			for i, credential := range credentials {
				var err error
				retained[i], err = s.engine.ResolveCredential(WithPrincipal(ctx, admin), embedded.CredentialUse{ActorID: 1, CredentialID: credential.ID, Version: credential.Version, Purpose: "synthetic.runtime.provider", ResourceKind: "workflow", ResourceID: workflowID, WorkflowID: workflowID, RunID: runID})
				if err != nil || len(retained[i]) == 0 {
					t.Fatal("approved workflow could not resolve its synthetic provider configuration")
				}
			}
			operations := []embedded.RuntimeOperation{
				{WorkflowID: workflowID, RunID: runID, GraphDigest: wp.GraphDigest, Action: "certificate.issue", Domains: []string{"allowed.example.com"}},
				{WorkflowID: workflowID, RunID: runID, GraphDigest: wp.GraphDigest, Action: "certificate.deploy", ResourceKind: "certificate", ResourceID: "synthetic-certificate", Domains: []string{"allowed.example.com"}},
				{WorkflowID: workflowID, RunID: runID, GraphDigest: wp.GraphDigest, Action: "notification.send", ResourceKind: "credential", ResourceID: credentials[0].ID},
				{WorkflowID: workflowID, RunID: runID, GraphDigest: wp.GraphDigest, Action: "dns.challenge", ResourceKind: "credential", ResourceID: credentials[0].ID, Domains: []string{"_acme-challenge.allowed.example.com"}},
			}
			hooks := s.RuntimeHooks()
			for _, op := range operations {
				if err := hooks.Authorize(ctx, op); err != nil {
					t.Fatalf("approved baseline operation rejected: %s", op.Action)
				}
			}
			if err := s.authorizeChallenge(ctx, admin, wp, credentials[0].ID, []string{"_acme-challenge.allowed.example.com"}); err != nil {
				t.Fatal("approved baseline challenge rejected")
			}
			index := 0
			if tc.secondary {
				index = 1
			}
			var replacement json.RawMessage
			if tc.rotate {
				replacement = json.RawMessage(`{"apiToken":"synthetic-replacement-configuration"}`)
			}
			if _, err := s.UpdateCredential(ctx, admin, credentials[index].ID, credentials[index].Version, tc.rotate, replacement); err != nil {
				t.Fatal(err)
			}
			if len(retained[index]) == 0 {
				t.Fatal("test lost the configuration retained by a previously created client")
			}
			for _, op := range operations {
				if err := hooks.Authorize(ctx, op); !errors.Is(err, ErrDenied) {
					t.Fatalf("changed credential allowed next provider effect: %s", op.Action)
				}
			}
			if _, err := s.engine.ResolveCredential(WithPrincipal(ctx, admin), embedded.CredentialUse{ActorID: 1, CredentialID: credentials[index].ID, Version: credentials[index].Version, Purpose: "synthetic.runtime.provider", ResourceKind: "workflow", ResourceID: workflowID, WorkflowID: workflowID, RunID: runID}); err == nil {
				t.Fatal("old approved credential version could be resolved after disable or rotation")
			}
			if !tc.secondary {
				if err := s.authorizeChallenge(ctx, admin, wp, credentials[index].ID, []string{"_acme-challenge.allowed.example.com"}); !errors.Is(err, ErrDenied) {
					t.Fatal("direct DNS challenge guard did not recheck credential state")
				}
			}
		})
	}
}

func TestIdentityCredentialMetadataWriteCannotDeleteDisableOrRotate(t *testing.T) {
	s := newIdentityTestService(t)
	ctx := context.Background()
	admin := Principal{UserID: 1, Admin: true}
	member := Principal{UserID: 2}
	credential, err := s.CreateCredential(ctx, admin, embedded.CredentialInput{Name: "synthetic-metadata", Provider: "cloudflare", Config: json.RawMessage(`{"apiToken":"synthetic-original-material"}`)})
	if err != nil {
		t.Fatal(err)
	}
	for _, action := range []string{"read", "write"} {
		if _, err := s.CreateGrant(ctx, admin, Grant{UserID: 2, Kind: "credential", ResourceID: credential.ID, Action: action}); err != nil {
			t.Fatal(err)
		}
	}
	if _, err := s.NativeSave(ctx, member, "access", credential.ID, "", map[string]any{"name": "synthetic-renamed", "unified_version": credential.Version}); err != nil {
		t.Fatal("metadata write permission could not update the display name")
	}
	for _, patch := range []map[string]any{
		{"unified_enabled": false, "unified_version": credential.Version},
		{"unified_enabled": true, "unified_version": credential.Version},
		{"config": map[string]any{"apiToken": "synthetic-unauthorized-replacement"}, "unified_version": credential.Version},
	} {
		if _, err := s.NativeSave(ctx, member, "access", credential.ID, "", patch); !errors.Is(err, ErrDenied) {
			t.Fatal("metadata writer changed credential lifecycle or material without rotate")
		}
	}
	for _, deleted := range []any{"2026-10-18 00:00:00.000Z", "", nil} {
		if _, err := s.NativeSave(ctx, member, "access", credential.ID, "", map[string]any{"deleted": deleted}); !errors.Is(err, embedded.ErrInvalid) {
			t.Fatal("metadata patch reached the soft-delete or restore field")
		}
	}
	if err := s.NativeDelete(ctx, member, "access", credential.ID); !errors.Is(err, ErrDenied) {
		t.Fatal("metadata writer deleted a credential without delete permission")
	}
	info, err := s.engine.CredentialInfo(ctx, credential.ID)
	if err != nil || !info.Enabled || info.Version != credential.Version || info.Name != "synthetic-renamed" {
		t.Fatal("rejected metadata updates changed credential state")
	}
	material, err := s.engine.ResolveCredential(WithPrincipal(ctx, admin), embedded.CredentialUse{ActorID: 1, CredentialID: credential.ID, Version: credential.Version, Purpose: "credential.export", ResourceKind: "credential", ResourceID: credential.ID})
	if err != nil || material["apiToken"] != "synthetic-original-material" {
		t.Fatal("rejected metadata updates replaced credential material")
	}
	if _, err := s.CreateGrant(ctx, admin, Grant{UserID: 2, Kind: "credential", ResourceID: credential.ID, Action: "rotate"}); err != nil {
		t.Fatal(err)
	}
	if _, err := s.NativeSave(ctx, member, "access", credential.ID, "", map[string]any{"unified_enabled": false, "unified_version": credential.Version}); err != nil {
		t.Fatal("explicit rotation permission could not disable credential")
	}
	if _, err := s.NativeSave(ctx, member, "access", credential.ID, "", map[string]any{"unified_enabled": true, "unified_version": credential.Version, "config": map[string]any{"apiToken": "synthetic-authorized-replacement"}}); err != nil {
		t.Fatal("explicit rotation permission could not replace credential material")
	}
	info, err = s.engine.CredentialInfo(ctx, credential.ID)
	if err != nil || !info.Enabled || info.Version != credential.Version+1 {
		t.Fatal("authorized credential rotation did not advance version")
	}
}

func TestIdentityNativeSaveRejectsDeletedFieldAcrossCollections(t *testing.T) {
	s := newIdentityTestService(t)
	ctx := context.Background()
	admin := Principal{UserID: 1, Admin: true}
	for _, collection := range []string{"access", "certificate", "workflow", "settings", "workflow_run", "workflow_output", "workflow_logs"} {
		t.Run(collection, func(t *testing.T) {
			for _, deleted := range []any{"2026-10-18 00:00:00.000Z", "", nil} {
				if _, err := s.NativeSave(ctx, admin, collection, "synthetic-id", "", map[string]any{"deleted": deleted}); !errors.Is(err, embedded.ErrInvalid) {
					t.Fatal("generic save accepted lifecycle deletion or restoration field")
				}
			}
		})
	}
}
