package unified

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"github.com/certimate-go/certimate/pkg/embedded"
	"github.com/nezhahq/nezha/service/unified/policy"
)

func TestIdentityAdministratorTokenCannotBypassRuntimePolicy(t *testing.T) {
	s := newIdentityTestService(t)
	ctx := context.Background()
	admin := Principal{UserID: 1, Admin: true}
	const workflowID = "synthetic-runtime-workflow"
	wp := WorkflowPolicy{WorkflowID: workflowID, Version: 1, Enabled: true, RunAsUserID: 1, GraphDigest: "synthetic-approved-graph", AllowInline: true, Names: []policy.NameScope{{Root: "allowed.example.com", IncludeRoot: true, Descendants: true}}}
	if err := s.db.Create(&wp).Error; err != nil {
		t.Fatal(err)
	}
	if err := s.db.Create(&ResourceBinding{Kind: "workflow", ResourceID: workflowID, OwnerID: 1, Version: 1, Active: true}).Error; err != nil {
		t.Fatal(err)
	}
	issued := issueIdentityToken(t, s, admin, []Scope{{Kind: "workflow", ResourceID: workflowID, Action: "execute"}})
	p, err := s.AuthenticateScopedToken(ctx, issued.Secret)
	if err != nil {
		t.Fatal(err)
	}
	scopedCtx := WithPrincipal(ctx, p)
	for _, action := range []string{"certificate.issue", "certificate.monitor"} {
		op := embedded.RuntimeOperation{WorkflowID: workflowID, GraphDigest: wp.GraphDigest, Action: action, Domains: []string{"outside.example.net"}}
		if err := s.authorizeRuntime(scopedCtx, op); !errors.Is(err, ErrDenied) {
			t.Fatalf("administrator scoped token bypassed name policy for %s", action)
		}
		op.Domains = []string{"www.allowed.example.com"}
		if err := s.authorizeRuntime(scopedCtx, op); err != nil {
			t.Fatalf("permitted workflow names denied for %s", action)
		}
	}
	importOp := embedded.RuntimeOperation{WorkflowID: workflowID, GraphDigest: wp.GraphDigest, Action: "certificate.import", Domains: []string{"allowed.example.com"}}
	if err := s.authorizeRuntime(scopedCtx, importOp); !errors.Is(err, ErrDenied) {
		t.Fatal("administrator token bypassed explicit certificate import policy")
	}
	inlineOp := embedded.RuntimeOperation{WorkflowID: workflowID, GraphDigest: wp.GraphDigest, Action: "deployment.execute"}
	if err := s.authorizeRuntime(scopedCtx, inlineOp); !errors.Is(err, ErrDenied) {
		t.Fatal("administrator token bypassed uncredentialed inline execution restriction")
	}
	if err := s.authorizeRuntime(WithPrincipal(ctx, admin), inlineOp); err != nil {
		t.Fatal("unrestricted administrator lost approved inline execution authority")
	}
}

func TestIdentityInactiveBindingAndManagementMetadataRemainClosed(t *testing.T) {
	s := newIdentityTestService(t)
	ctx := context.Background()
	admin := Principal{UserID: 1, Admin: true}
	credential, err := s.engine.CreateCredential(ctx, "1", embedded.CredentialInput{Name: "synthetic-inactive", Provider: "cloudflare", Config: json.RawMessage(`{"apiToken":"synthetic-inactive-secret"}`)})
	if err != nil {
		t.Fatal(err)
	}
	scopes := []Scope{{Kind: "credential", ResourceID: credential.ID, Action: "read"}, {Kind: "credential", ResourceID: credential.ID, Action: "use"}}
	for _, scope := range scopes {
		if err := s.db.Create(&Grant{UserID: 2, Kind: scope.Kind, ResourceID: scope.ResourceID, Action: scope.Action}).Error; err != nil {
			t.Fatal(err)
		}
	}
	binding := ResourceBinding{Kind: "credential", ResourceID: credential.ID, OwnerID: 2, Active: true, Version: 1}
	if err := s.db.Create(&binding).Error; err != nil {
		t.Fatal(err)
	}
	issued := issueIdentityToken(t, s, admin, scopes)
	token, err := s.AuthenticateScopedToken(ctx, issued.Secret)
	if err != nil {
		t.Fatal(err)
	}
	identity, err := s.CreateServiceIdentity(ctx, admin, ServiceIdentityInput{Name: "synthetic-reader", OwnerID: 2, Scopes: scopes})
	if err != nil {
		t.Fatal(err)
	}
	service, err := s.ServicePrincipal(ctx, identity.ID)
	if err != nil {
		t.Fatal(err)
	}
	for _, p := range []Principal{token, service} {
		if _, err := s.Ownership(ctx, p); !errors.Is(err, ErrDenied) {
			t.Fatal("restricted identity enumerated ownership management metadata")
		}
		if _, err := s.ListGrants(ctx, p); !errors.Is(err, ErrDenied) {
			t.Fatal("restricted identity enumerated grants")
		}
	}
	if err := s.db.Model(&binding).Update("active", false).Error; err != nil {
		t.Fatal(err)
	}
	for _, p := range []Principal{admin, token, service, {UserID: 2}} {
		if s.Can(ctx, p, "credential", credential.ID, "use") {
			t.Fatal("inactive binding remained usable while engine record survived")
		}
	}
	for _, p := range []Principal{token, service, {UserID: 2}} {
		if s.Can(ctx, p, "credential", credential.ID, "read") {
			t.Fatal("inactive source remained visible to restricted identities")
		}
	}
	if !s.Can(ctx, admin, "credential", credential.ID, "read") || !s.Can(ctx, admin, "credential", credential.ID, "delete") {
		t.Fatal("unrestricted administrator lost reconciliation access")
	}
}

func TestIdentityRecoveryKeepsRevokedOriginRestrictions(t *testing.T) {
	for _, mode := range []string{"token", "service"} {
		t.Run(mode, func(t *testing.T) {
			s := newIdentityTestService(t)
			ctx := context.Background()
			admin := Principal{UserID: 1, Admin: true}
			scopes := []Scope{{Kind: "credential", ResourceID: "new", Action: "create"}}
			var p Principal
			var revoke func() error
			if mode == "token" {
				issued := issueIdentityToken(t, s, admin, scopes)
				var err error
				p, err = s.AuthenticateScopedToken(ctx, issued.Secret)
				if err != nil {
					t.Fatal(err)
				}
				revoke = func() error { return s.RevokeScopedToken(ctx, admin, issued.Token.ID) }
			} else {
				identity, err := s.CreateServiceIdentity(ctx, admin, ServiceIdentityInput{Name: "synthetic-recovery", OwnerID: 2, Scopes: scopes})
				if err != nil {
					t.Fatal(err)
				}
				p, err = s.ServicePrincipal(ctx, identity.ID)
				if err != nil {
					t.Fatal(err)
				}
				revoke = func() error {
					active := false
					_, err := s.UpdateServiceIdentity(ctx, admin, identity.ID, ServiceIdentityUpdate{Version: identity.Version, Active: &active})
					return err
				}
			}
			const operationID = "synthetic-interrupted-creation"
			record, err := s.engine.NativeSave(ctx, "1", "access", "", map[string]any{"name": "synthetic-pending", "provider": "cloudflare", "config": map[string]any{"apiToken": "synthetic-pending-secret"}, "unified_operation_id": operationID, "unified_active": false})
			if err != nil {
				t.Fatal(err)
			}
			id, _ := record["id"].(string)
			intent := ResourceIntent{OperationID: operationID, ActorID: p.UserID, TokenID: p.TokenID, ServiceID: p.ServiceID, IdentityVersion: p.IdentityVersion, Kind: "credential", Collection: "access", ResourceID: id, State: "prepared"}
			if err := s.db.Create(&intent).Error; err != nil {
				t.Fatal(err)
			}
			if err := revoke(); err != nil {
				t.Fatal(err)
			}
			if err := s.RecoverIntents(ctx); err != nil {
				t.Fatal(err)
			}
			if err := s.db.First(&intent, "operation_id = ?", operationID).Error; err != nil {
				t.Fatal(err)
			}
			if intent.State != "reconciliation_required" {
				t.Fatal("recovery discarded originating identity restrictions")
			}
			if _, err := s.engine.NativeGet(ctx, "access", id); !errors.Is(err, embedded.ErrNotFound) {
				t.Fatal("recovery activated a resource created by a revoked identity")
			}
		})
	}
}

func TestIdentityPersistedRunRechecksTokenAndServiceVersions(t *testing.T) {
	for _, mode := range []string{"token", "service"} {
		t.Run(mode, func(t *testing.T) {
			s := newIdentityTestService(t)
			ctx := context.Background()
			admin := Principal{UserID: 1, Admin: true}
			const workflowID = "synthetic-durable-workflow"
			wp := WorkflowPolicy{WorkflowID: workflowID, Version: 1, Enabled: true, RunAsUserID: 1, GraphDigest: "synthetic-approved-graph"}
			if err := s.db.Create(&wp).Error; err != nil {
				t.Fatal(err)
			}
			if err := s.db.Create(&ResourceBinding{Kind: "workflow", ResourceID: workflowID, OwnerID: 1, Version: 1, Active: true}).Error; err != nil {
				t.Fatal(err)
			}
			scopes := []Scope{{Kind: "workflow", ResourceID: workflowID, Action: "execute"}}
			var p Principal
			var revoke func() error
			if mode == "token" {
				issued := issueIdentityToken(t, s, admin, scopes)
				var err error
				p, err = s.AuthenticateScopedToken(ctx, issued.Secret)
				if err != nil {
					t.Fatal(err)
				}
				revoke = func() error { return s.RevokeScopedToken(ctx, admin, issued.Token.ID) }
			} else {
				identity, err := s.CreateServiceIdentity(ctx, admin, ServiceIdentityInput{Name: "synthetic-durable", OwnerID: 2, Scopes: scopes})
				if err != nil {
					t.Fatal(err)
				}
				p, err = s.ServicePrincipal(ctx, identity.ID)
				if err != nil {
					t.Fatal(err)
				}
				revoke = func() error {
					owner := uint64(3)
					_, err := s.UpdateServiceIdentity(ctx, admin, identity.ID, ServiceIdentityUpdate{Version: identity.Version, OwnerID: &owner})
					return err
				}
			}
			const runID = "synthetic-accepted-run"
			if err := s.bindRun(WithPrincipal(ctx, p), workflowID, runID); err != nil {
				t.Fatal(err)
			}
			op := embedded.RuntimeOperation{WorkflowID: workflowID, RunID: runID, GraphDigest: wp.GraphDigest, Action: "node.execute"}
			if err := s.authorizeRuntime(ctx, op); err != nil {
				t.Fatal("accepted run could not resume with persisted identity")
			}
			if err := revoke(); err != nil {
				t.Fatal(err)
			}
			if err := s.authorizeRuntime(ctx, op); !errors.Is(err, ErrDenied) {
				t.Fatal("persisted run survived origin token revoke or service transfer")
			}
		})
	}
}

func TestIdentityFormalServiceDNSUsesExplicitNamePolicy(t *testing.T) {
	s := newIdentityTestService(t)
	ctx := context.Background()
	admin := Principal{UserID: 1, Admin: true}
	credential, err := s.engine.CreateCredential(ctx, "1", embedded.CredentialInput{Name: "synthetic-service-dns", Provider: "cloudflare", Config: json.RawMessage(`{"apiToken":"synthetic-service-dns-secret"}`)})
	if err != nil {
		t.Fatal(err)
	}
	account := DNSAccount{Name: "synthetic", OwnerID: 1, Provider: "cloudflare", CredentialID: credential.ID, CredentialVersion: credential.Version, Enabled: true}
	if err := s.db.Create(&account).Error; err != nil {
		t.Fatal(err)
	}
	zone := DNSZone{AccountID: account.ID, ProviderID: "synthetic-zone", Name: "example.com", Active: true}
	if err := s.db.Create(&zone).Error; err != nil {
		t.Fatal(err)
	}
	scopes := []Scope{{Kind: "credential", ResourceID: credential.ID, Action: "use"}, {Kind: "dns_zone", ResourceID: fmt.Sprint(zone.ID), Action: "create"}, {Kind: "dns_zone", ResourceID: fmt.Sprint(zone.ID), Action: "delete"}}
	identity, err := s.CreateServiceIdentity(ctx, admin, ServiceIdentityInput{Name: "synthetic-service-dns", OwnerID: 2, Scopes: scopes})
	if err != nil {
		t.Fatal(err)
	}
	p, err := s.ServicePrincipal(ctx, identity.ID)
	if err != nil {
		t.Fatal(err)
	}
	serviceGrant, err := s.CreateDNSPolicy(ctx, admin, DNSPolicy{ServiceID: identity.ID, ZoneID: zone.ID, NameScope: policy.NameScope{Root: "_acme-challenge.allowed.example.com", IncludeRoot: true}, Types: []string{"TXT"}, Actions: []string{"create", "delete"}})
	if err != nil {
		t.Fatal(err)
	}
	ownerGrant, err := s.CreateDNSPolicy(ctx, admin, DNSPolicy{UserID: 2, ZoneID: zone.ID, NameScope: policy.NameScope{Root: "_acme-challenge.other.example.com", IncludeRoot: true}, Types: []string{"TXT"}, Actions: []string{"create", "delete"}})
	if err != nil {
		t.Fatal(err)
	}
	wp := WorkflowPolicy{Credentials: map[string]int{credential.ID: credential.Version}}
	allowed := []string{"_acme-challenge.allowed.example.com"}
	if err := s.authorizeChallenge(ctx, p, wp, credential.ID, allowed); err != nil {
		t.Fatal("formal service could not use its own explicit DNS policy")
	}
	if err := s.authorizeChallenge(ctx, p, wp, credential.ID, []string{"_acme-challenge.other.example.com"}); !errors.Is(err, ErrDenied) {
		t.Fatal("formal service borrowed owner's unrelated DNS name policy")
	}
	if err := s.db.Delete(&ownerGrant).Error; err != nil {
		t.Fatal(err)
	}
	if err := s.authorizeChallenge(ctx, p, wp, credential.ID, allowed); err != nil {
		t.Fatal("owner user-grant revocation removed independent service authority")
	}
	if err := s.db.Delete(&serviceGrant).Error; err != nil {
		t.Fatal(err)
	}
	if err := s.authorizeChallenge(ctx, p, wp, credential.ID, allowed); !errors.Is(err, ErrDenied) {
		t.Fatal("revoked service DNS name policy retained challenge authority")
	}
}
