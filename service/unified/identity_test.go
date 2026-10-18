package unified

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/certimate-go/certimate/pkg/embedded"
	"github.com/nezhahq/nezha/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func newIdentityTestService(t *testing.T) *Service {
	t.Helper()
	dir := t.TempDir()
	e, err := embedded.Open(filepath.Join(dir, "certimate"))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = e.Close() })
	db, err := gorm.Open(sqlite.Open(filepath.Join(dir, "identity.db")), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = sqlDB.Close() })
	if err := db.AutoMigrate(&model.User{}); err != nil {
		t.Fatal(err)
	}
	for _, user := range []model.User{
		{Common: model.Common{ID: 1}, Username: "synthetic-admin", Role: model.RoleAdmin, TokenVersion: 7},
		{Common: model.Common{ID: 2}, Username: "synthetic-member", Role: model.RoleMember, TokenVersion: 9},
		{Common: model.Common{ID: 3}, Username: "synthetic-other", Role: model.RoleMember, TokenVersion: 11},
	} {
		if err := db.Create(&user).Error; err != nil {
			t.Fatal(err)
		}
	}
	s, err := New(db, e)
	if err != nil {
		t.Fatal(err)
	}
	return s
}

func issueIdentityToken(t *testing.T, s *Service, p Principal, scopes []Scope) IssuedScopedToken {
	t.Helper()
	v, err := s.CreateScopedToken(context.Background(), p, ScopedTokenInput{Name: "synthetic-token", Scopes: scopes, ExpiresAt: time.Now().Add(time.Hour)})
	if err != nil {
		t.Fatal(err)
	}
	return v
}

func TestIdentityScopedAdministratorResourceActionAndSecretLifecycle(t *testing.T) {
	s := newIdentityTestService(t)
	ctx := context.Background()
	admin := Principal{UserID: 1, Admin: true}
	first, err := s.engine.CreateCredential(ctx, "1", embedded.CredentialInput{Name: "first", Provider: "cloudflare", Config: json.RawMessage(`{"apiToken":"synthetic-first"}`)})
	if err != nil {
		t.Fatal(err)
	}
	second, err := s.engine.CreateCredential(ctx, "1", embedded.CredentialInput{Name: "second", Provider: "cloudflare", Config: json.RawMessage(`{"apiToken":"synthetic-second"}`)})
	if err != nil {
		t.Fatal(err)
	}
	issued := issueIdentityToken(t, s, admin, []Scope{{Kind: "credential", ResourceID: first.ID, Action: "read"}})
	p, err := s.AuthenticateScopedToken(ctx, issued.Secret)
	if err != nil || p.TokenID != issued.Token.ID || !p.Admin {
		t.Fatal("scoped admin authentication failed")
	}
	if !s.Can(ctx, p, "credential", first.ID, "read") || s.Can(ctx, p, "credential", second.ID, "read") || s.Can(ctx, p, "credential", first.ID, "use") {
		t.Fatal("administrator token escaped its exact scope")
	}
	page, err := s.List(ctx, p, "credential", 1, 20)
	if err != nil || page.Total != 1 || len(page.Items) != 1 || page.Items[0].ID != first.ID {
		t.Fatal("scope filtering did not precede pagination and count")
	}
	if _, err := s.ListScopedTokens(ctx, p); !errors.Is(err, ErrDenied) {
		t.Fatal("token reached identity management")
	}
	if _, err := s.CreateScopedToken(ctx, p, ScopedTokenInput{}); !errors.Is(err, ErrDenied) {
		t.Fatal("token minted another token")
	}
	if _, err := s.CreateServiceIdentity(ctx, p, ServiceIdentityInput{}); !errors.Is(err, ErrDenied) {
		t.Fatal("token created service identity")
	}
	if err := s.RevokeScopedToken(ctx, p, issued.Token.ID); !errors.Is(err, ErrDenied) {
		t.Fatal("token reached revocation management")
	}
	if s.Can(ctx, Principal{UserID: 1, Admin: true, PAT: true}, "credential", first.ID, "read") {
		t.Fatal("legacy PAT gained unified access")
	}
	if _, err := s.AuthenticateScopedToken(ctx, "nzp_synthetic-legacy-token"); !errors.Is(err, ErrDenied) {
		t.Fatal("legacy PAT authenticated")
	}
	wrong := issued.Secret[:len(issued.Secret)-2] + "xx"
	if wrong == issued.Secret {
		wrong = issued.Secret[:len(issued.Secret)-2] + "yy"
	}
	if _, err := s.AuthenticateScopedToken(ctx, wrong); !errors.Is(err, ErrDenied) {
		t.Fatal("changed secret authenticated")
	}
	var stored ScopedToken
	if err := s.db.First(&stored, issued.Token.ID).Error; err != nil {
		t.Fatal(err)
	}
	_, secretPart, _ := strings.Cut(issued.Secret, ".")
	digest := sha256.Sum256([]byte(secretPart))
	if string(stored.SecretHash) != string(digest[:]) || string(stored.SecretHash) == secretPart {
		t.Fatal("secret was not stored solely as hash")
	}
	rows, err := s.ListScopedTokens(ctx, admin)
	if err != nil || len(rows) != 1 || len(rows[0].SecretHash) != 0 || rows[0].PublicID != "" {
		t.Fatal("token metadata leaked authentication material")
	}
	encoded, err := json.Marshal(rows)
	if err != nil || strings.Contains(string(encoded), secretPart) || strings.Contains(string(encoded), "secret_hash") || strings.Contains(string(encoded), "public_id") {
		t.Fatal("metadata included authentication material")
	}
	otherRows, err := s.ListScopedTokens(ctx, Principal{UserID: 2})
	if err != nil || len(otherRows) != 0 {
		t.Fatal("another user's tokens exposed")
	}
	if err := s.RevokeScopedToken(ctx, Principal{UserID: 2}, issued.Token.ID); !errors.Is(err, embedded.ErrNotFound) {
		t.Fatal("cross-user revocation accepted")
	}
	if err := s.RevokeScopedToken(ctx, admin, issued.Token.ID); err != nil {
		t.Fatal(err)
	}
	if err := s.RevokeScopedToken(ctx, admin, issued.Token.ID); err != nil {
		t.Fatal("revoke not idempotent")
	}
	if _, err := s.AuthenticateScopedToken(ctx, issued.Secret); !errors.Is(err, ErrDenied) || s.Can(ctx, p, "credential", first.ID, "read") {
		t.Fatal("revoked bearer or captured principal retained access")
	}
	var audits []Audit
	if err := s.db.Find(&audits).Error; err != nil {
		t.Fatal(err)
	}
	auditJSON, err := json.Marshal(audits)
	if err != nil || strings.Contains(string(auditJSON), secretPart) || len(audits) != 2 {
		t.Fatal("audit leaked secret or duplicate revoke")
	}
}

func TestIdentityTokenFollowsCurrentUserAndGrantState(t *testing.T) {
	s := newIdentityTestService(t)
	ctx := context.Background()
	scope := Scope{Kind: "certificate", ResourceID: "synthetic-certificate", Action: "download"}
	grant := Grant{UserID: 2, Kind: scope.Kind, ResourceID: scope.ResourceID, Action: scope.Action}
	if err := s.db.Create(&grant).Error; err != nil {
		t.Fatal(err)
	}
	member := Principal{UserID: 2}
	issued := issueIdentityToken(t, s, member, []Scope{scope})
	p, err := s.AuthenticateScopedToken(ctx, issued.Secret)
	if err != nil || !s.Can(ctx, p, scope.Kind, scope.ResourceID, scope.Action) {
		t.Fatal("granted member token failed")
	}
	if _, err := s.CreateScopedToken(ctx, member, ScopedTokenInput{Name: "escalation", ExpiresAt: time.Now().Add(time.Hour), Scopes: []Scope{{Kind: scope.Kind, ResourceID: scope.ResourceID, Action: "download_key"}}}); !errors.Is(err, ErrDenied) {
		t.Fatal("token issued an ungranted scope")
	}
	if err := s.db.Delete(&grant).Error; err != nil {
		t.Fatal(err)
	}
	if s.Can(ctx, p, scope.Kind, scope.ResourceID, scope.Action) {
		t.Fatal("token retained revoked grant")
	}
	if err := s.db.Model(&model.User{}).Where("id = ?", 2).Update("token_version", 10).Error; err != nil {
		t.Fatal(err)
	}
	if _, err := s.AuthenticateScopedToken(ctx, issued.Secret); !errors.Is(err, ErrDenied) || s.ValidateIdentity(ctx, p) == nil {
		t.Fatal("token survived user token-version bump")
	}
	adminToken := issueIdentityToken(t, s, Principal{UserID: 1, Admin: true}, []Scope{scope})
	adminP, err := s.AuthenticateScopedToken(ctx, adminToken.Secret)
	if err != nil {
		t.Fatal(err)
	}
	if err := s.db.Model(&model.User{}).Where("id = ?", 1).Update("role", model.RoleMember).Error; err != nil {
		t.Fatal(err)
	}
	if s.Can(ctx, adminP, scope.Kind, scope.ResourceID, scope.Action) {
		t.Fatal("cached admin survived demotion")
	}
	if err := s.db.Delete(&model.User{}, 1).Error; err != nil {
		t.Fatal(err)
	}
	if _, err := s.AuthenticateScopedToken(ctx, adminToken.Secret); !errors.Is(err, ErrDenied) {
		t.Fatal("deleted user authenticated")
	}
}

func TestIdentityServicePolicyTransferAndDeactivation(t *testing.T) {
	s := newIdentityTestService(t)
	ctx := context.Background()
	admin := Principal{UserID: 1, Admin: true}
	scope := Scope{Kind: "workflow", ResourceID: "synthetic-workflow", Action: "execute"}
	v, err := s.CreateServiceIdentity(ctx, admin, ServiceIdentityInput{Name: "synthetic-automation", OwnerID: 2, Scopes: []Scope{scope}})
	if err != nil {
		t.Fatal(err)
	}
	p, err := s.ServicePrincipal(ctx, v.ID)
	if err != nil || !s.Can(ctx, p, scope.Kind, scope.ResourceID, scope.Action) {
		t.Fatal("service explicit scope did not authorize automation")
	}
	if s.Can(ctx, p, scope.Kind, "different-workflow", scope.Action) || s.Can(ctx, p, scope.Kind, scope.ResourceID, "write") {
		t.Fatal("service escaped explicit scopes")
	}
	unversioned := p
	unversioned.IdentityVersion = 0
	if s.ValidateIdentity(ctx, unversioned) == nil {
		t.Fatal("unversioned run accepted")
	}
	if _, err := s.CreateServiceIdentity(ctx, Principal{UserID: 2, Admin: true}, ServiceIdentityInput{Name: "escalation", OwnerID: 2, Scopes: []Scope{scope}}); !errors.Is(err, ErrDenied) {
		t.Fatal("stale admin claim created service")
	}
	if _, err := s.CreateScopedToken(ctx, p, ScopedTokenInput{}); !errors.Is(err, ErrDenied) {
		t.Fatal("service minted user token")
	}
	rows, err := s.ListServiceIdentities(ctx, Principal{UserID: 3})
	if err != nil || len(rows) != 0 {
		t.Fatal("unrelated user saw identity")
	}
	rows, err = s.ListServiceIdentities(ctx, Principal{UserID: 2})
	if err != nil || len(rows) != 1 {
		t.Fatal("owner did not see identity")
	}
	owner := uint64(3)
	transferred, err := s.UpdateServiceIdentity(ctx, admin, v.ID, ServiceIdentityUpdate{Version: v.Version, OwnerID: &owner})
	if err != nil || transferred.OwnerID != 3 || transferred.Version != 2 || s.ValidateIdentity(ctx, p) == nil {
		t.Fatal("transfer did not invalidate captured run")
	}
	if _, err := s.UpdateServiceIdentity(ctx, admin, v.ID, ServiceIdentityUpdate{Version: v.Version, OwnerID: &owner}); !errors.Is(err, embedded.ErrInvalid) {
		t.Fatal("stale policy edit accepted")
	}
	owner = 2
	returned, err := s.UpdateServiceIdentity(ctx, admin, v.ID, ServiceIdentityUpdate{Version: transferred.Version, OwnerID: &owner})
	if err != nil || s.ValidateIdentity(ctx, p) == nil {
		t.Fatal("round trip ownership resurrected old run")
	}
	current, err := s.ServicePrincipal(ctx, v.ID)
	if err != nil {
		t.Fatal(err)
	}
	active := false
	disabled, err := s.UpdateServiceIdentity(ctx, admin, v.ID, ServiceIdentityUpdate{Version: returned.Version, Active: &active})
	if err != nil || disabled.Active || s.ValidateIdentity(ctx, current) == nil {
		t.Fatal("deactivation did not invalidate run")
	}
	if _, err := s.ServicePrincipal(ctx, v.ID); !errors.Is(err, ErrDenied) {
		t.Fatal("inactive service started run")
	}
	active = true
	reactivated, err := s.UpdateServiceIdentity(ctx, admin, v.ID, ServiceIdentityUpdate{Version: disabled.Version, Active: &active})
	if err != nil || s.ValidateIdentity(ctx, current) == nil {
		t.Fatal("reactivation resurrected captured run")
	}
	current, err = s.ServicePrincipal(ctx, v.ID)
	if err != nil {
		t.Fatal(err)
	}
	narrowed := []Scope{{Kind: scope.Kind, ResourceID: scope.ResourceID, Action: "read"}}
	_, err = s.UpdateServiceIdentity(ctx, admin, v.ID, ServiceIdentityUpdate{Version: reactivated.Version, Scopes: &narrowed})
	if err != nil || s.ValidateIdentity(ctx, current) == nil {
		t.Fatal("scope edit did not invalidate captured run")
	}
	current, err = s.ServicePrincipal(ctx, v.ID)
	if err != nil || s.Can(ctx, current, scope.Kind, scope.ResourceID, "execute") || !s.Can(ctx, current, scope.Kind, scope.ResourceID, "read") {
		t.Fatal("updated service scope was not enforced on new runs")
	}
	if err := s.db.Delete(&model.User{}, 2).Error; err != nil {
		t.Fatal(err)
	}
	if _, err := s.ServicePrincipal(ctx, v.ID); !errors.Is(err, ErrDenied) {
		t.Fatal("service with deleted owner started run")
	}
}

func TestIdentityRejectsInvalidScopesAndExpiredTokens(t *testing.T) {
	s := newIdentityTestService(t)
	ctx := context.Background()
	admin := Principal{UserID: 1, Admin: true}
	for _, scopes := range [][]Scope{
		nil,
		{{Kind: "credential", ResourceID: "*", Action: "read"}},
		{{Kind: "credential", ResourceID: "new", Action: "read"}},
		{{Kind: "credential", ResourceID: "one", Action: "*"}},
		{{Kind: "unknown", ResourceID: "one", Action: "read"}},
		{{Kind: "credential", ResourceID: "one\ntwo", Action: "read"}},
		{{Kind: "credential", ResourceID: "one", Action: "read"}, {Kind: "credential", ResourceID: "one", Action: "read"}},
	} {
		if _, err := s.CreateScopedToken(ctx, admin, ScopedTokenInput{Name: "invalid", Scopes: scopes, ExpiresAt: time.Now().Add(time.Hour)}); !errors.Is(err, embedded.ErrInvalid) {
			t.Fatal("invalid scope accepted")
		}
	}
	scopes := []Scope{{Kind: "certificate", ResourceID: "synthetic", Action: "read"}}
	for _, expiry := range []time.Time{{}, time.Now().Add(-time.Hour), time.Now().Add(367 * 24 * time.Hour)} {
		if _, err := s.CreateScopedToken(ctx, admin, ScopedTokenInput{Name: "invalid-expiry", Scopes: scopes, ExpiresAt: expiry}); !errors.Is(err, embedded.ErrInvalid) {
			t.Fatal("invalid expiry accepted")
		}
	}
	issued := issueIdentityToken(t, s, admin, scopes)
	p, err := s.AuthenticateScopedToken(ctx, issued.Secret)
	if err != nil {
		t.Fatal(err)
	}
	if err := s.db.Model(&ScopedToken{}).Where("id = ?", issued.Token.ID).Update("expires_at", time.Now().Add(-time.Second)).Error; err != nil {
		t.Fatal(err)
	}
	if _, err := s.AuthenticateScopedToken(ctx, issued.Secret); !errors.Is(err, ErrDenied) || s.ValidateIdentity(ctx, p) == nil {
		t.Fatal("expired token retained authority")
	}
}
