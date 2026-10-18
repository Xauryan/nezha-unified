package unified

import (
	"context"
	"encoding/json"
	"path/filepath"
	"testing"

	"github.com/certimate-go/certimate/pkg/embedded"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestResourceGrantsRevokeAndAdminPATRestrictions(t *testing.T) {
	dir := t.TempDir()
	engine, err := embedded.Open(filepath.Join(dir, "certimate"))
	if err != nil {
		t.Fatal(err)
	}
	defer engine.Close()
	db, err := gorm.Open(sqlite.Open(filepath.Join(dir, "nezha.db")), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
	s, err := New(db, engine)
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()
	admin := Principal{UserID: 1, Admin: true}
	member := Principal{UserID: 2}
	in := embedded.CredentialInput{Name: "test", Provider: "cloudflare", Config: json.RawMessage(`{"apiToken":"synthetic"}`)}
	r, err := s.CreateCredential(ctx, admin, in)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := s.CreateCredential(ctx, member, in); err == nil {
		t.Fatal("member created a credential")
	}
	page, err := s.List(ctx, member, "credential", 1, 20)
	if err != nil || page.Total != 0 {
		t.Fatal("member saw ungranted resource")
	}
	g, err := s.CreateGrant(ctx, admin, Grant{UserID: 2, Kind: "credential", ResourceID: r.ID, Action: "read"})
	if err != nil {
		t.Fatal(err)
	}
	page, err = s.List(ctx, member, "credential", 1, 20)
	if err != nil || page.Total != 1 {
		t.Fatalf("grant not visible: %v", err)
	}
	if s.Can(ctx, member, "credential", r.ID, "use") {
		t.Fatal("metadata implies use")
	}
	if s.Can(ctx, Principal{UserID: 3}, "credential", r.ID, "read") {
		t.Fatal("grant shared between users")
	}
	if _, err = s.CreateGrant(ctx, member, Grant{UserID: 2, Kind: "credential", ResourceID: r.ID, Action: "read"}); err == nil {
		t.Fatal("member granted itself permission")
	}
	if err = s.RevokeGrant(ctx, admin, g.ID); err != nil {
		t.Fatal(err)
	}
	page, err = s.List(ctx, member, "credential", 1, 20)
	if err != nil || page.Total != 0 {
		t.Fatal("revocation not immediate")
	}
	pat := Principal{UserID: 1, Admin: true, PAT: true}
	if s.Can(ctx, pat, "credential", r.ID, "read") {
		t.Fatal("admin PAT bypassed restriction")
	}
	if _, err = s.List(ctx, pat, "credential", 1, 20); err == nil {
		t.Fatal("PAT listed resources")
	}
	if _, err = s.CreateCredential(ctx, pat, in); err == nil {
		t.Fatal("PAT created resource")
	}
	if _, err = s.CreateGrant(ctx, admin, Grant{UserID: 2, Kind: "credential", ResourceID: "missing", Action: "read"}); err == nil {
		t.Fatal("grant created for nonexistent record")
	}
}
