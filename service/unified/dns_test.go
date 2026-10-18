package unified

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/certimate-go/certimate/pkg/embedded"
	"github.com/nezhahq/nezha/model"
	"github.com/nezhahq/nezha/service/unified/dnsprovider"
	"github.com/nezhahq/nezha/service/unified/policy"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"path/filepath"
	"sync"
	"testing"
	"time"
)

type fakeDNS struct {
	mu      sync.Mutex
	records []dnsprovider.Record
	writes  int
	unknown bool
}

func (f *fakeDNS) Capabilities() dnsprovider.Capabilities {
	return dnsprovider.Capabilities{Types: []string{"A", "AAAA", "TXT"}, MinTTL: 60}
}
func (f *fakeDNS) Zones(context.Context) ([]dnsprovider.Zone, error) {
	return []dnsprovider.Zone{{ID: "remote-zone", Name: "example.com", Mode: "full"}}, nil
}
func (f *fakeDNS) Records(context.Context, dnsprovider.Zone) ([]dnsprovider.Record, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	return append([]dnsprovider.Record(nil), f.records...), nil
}
func (f *fakeDNS) Get(_ context.Context, _ dnsprovider.Zone, id string) (dnsprovider.Record, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	for _, r := range f.records {
		if r.ID == id {
			return r, nil
		}
	}
	return dnsprovider.Record{}, errors.New("not found")
}
func (f *fakeDNS) Create(_ context.Context, _ dnsprovider.Zone, r dnsprovider.Record) (dnsprovider.Record, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.writes++
	r.ID = "created"
	if f.unknown {
		return r, &dnsprovider.APIError{Code: "timeout", Unknown: true}
	}
	f.records = append(f.records, r)
	return r, nil
}
func (f *fakeDNS) Update(_ context.Context, _ dnsprovider.Zone, r dnsprovider.Record) (dnsprovider.Record, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.writes++
	if f.unknown {
		return r, &dnsprovider.APIError{Code: "timeout", Unknown: true}
	}
	for i, old := range f.records {
		if old.ID == r.ID {
			f.records[i] = r
		}
	}
	return r, nil
}
func (f *fakeDNS) Delete(context.Context, dnsprovider.Zone, string) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.writes++
	return nil
}
func ddnsFixture(t *testing.T) (*Service, *fakeDNS, Principal, Principal, DNSZone, DNSPolicy, *model.Server) {
	t.Helper()
	ctx := context.Background()
	dir := t.TempDir()
	e, err := embedded.Open(filepath.Join(dir, "pb"))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { e.Close() })
	db, err := gorm.Open(sqlite.Open(filepath.Join(dir, "host.db")), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	sql, _ := db.DB()
	sql.SetMaxOpenConns(1)
	t.Cleanup(func() { sql.Close() })
	if err := db.AutoMigrate(&model.User{}, &model.Server{}, &model.DDNSProfile{}); err != nil {
		t.Fatal(err)
	}
	a := model.User{Username: "admin", Role: model.RoleAdmin}
	m := model.User{Username: "member", Role: model.RoleMember}
	db.Create(&a)
	db.Create(&m)
	admin, member := Principal{UserID: a.ID, Admin: true}, Principal{UserID: m.ID}
	host := model.Server{Common: model.Common{UserID: m.ID}, Name: "test-host", UUID: "synthetic-host"}
	db.Create(&host)
	s, err := New(db, e)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(s.StopDDNS)
	f := &fakeDNS{records: []dnsprovider.Record{{ID: "in", Name: "www.delegated.example.com", Type: "A", Value: "192.0.2.1", TTL: 600}, {ID: "out", Name: "elsewhere.example.com", Type: "A", Value: "192.0.2.2", TTL: 600}}}
	s.dnsFactory = func(string, map[string]any) (dnsprovider.Provider, error) { return f, nil }
	cred, err := s.CreateCredential(ctx, admin, embedded.CredentialInput{Name: "test", Provider: "cloudflare", Config: json.RawMessage(`{"apiToken":"synthetic"}`)})
	if err != nil {
		t.Fatal(err)
	}
	a1, err := s.CreateDNSAccount(ctx, admin, DNSAccount{Name: "test", Provider: "cloudflare", CredentialID: cred.ID, CredentialVersion: 1})
	if err != nil {
		t.Fatal(err)
	}
	zones, err := s.RefreshDNSZones(ctx, admin, a1.ID)
	if err != nil {
		t.Fatal(err)
	}
	g, err := s.CreateDNSPolicy(ctx, admin, DNSPolicy{UserID: member.UserID, ZoneID: zones[0].ID, NameScope: policy.NameScope{Root: "delegated.example.com", Descendants: true}, Types: []string{"A", "AAAA"}, Actions: []string{"read", "update", "create", "delete", "ddns"}})
	if err != nil {
		t.Fatal(err)
	}
	s.ddns.source = func(context.Context, uint64) (DDNSHostSnapshot, error) {
		return DDNSHostSnapshot{OwnerID: m.ID, Online: true, ObservedAt: time.Now(), IPv4: "192.0.2.9"}, nil
	}
	return s, f, admin, member, zones[0], g, &host
}
func TestDNSAuthorizationPaginationRenameAndUnknownReplay(t *testing.T) {
	s, f, admin, member, z, g, _ := ddnsFixture(t)
	ctx := context.Background()
	page, err := s.DNSRecords(ctx, member, z.ID, 1, 1)
	if err != nil || page.Total != 1 || page.Items[0].ID != "in" {
		t.Fatalf("filtered page %#v %v", page, err)
	}
	rename := DNSChange{OperationKey: "rename-outside-scope", Action: "update", Record: f.records[0]}
	rename.Record.Name = "outside.example.com"
	if _, err := s.ChangeDNS(ctx, member, z.ID, rename); !errors.Is(err, ErrDenied) || f.writes != 0 {
		t.Fatal("rename escaped scope")
	}
	replace := DNSChange{OperationKey: "replace-original-record", Action: "update", Record: f.records[0]}
	replace.Record.ID = "out"
	if _, err := s.ChangeDNS(ctx, member, z.ID, replace); !errors.Is(err, ErrDenied) || f.writes != 0 {
		t.Fatal("ID substitution")
	}
	f.unknown = true
	change := DNSChange{OperationKey: "create-ambiguous-record", Action: "create", Record: dnsprovider.Record{Name: "new.delegated.example.com", Type: "A", Value: "192.0.2.3", TTL: 600}}
	op, err := s.ChangeDNS(ctx, member, z.ID, change)
	if err != nil || op.State != "unknown" || f.writes != 1 {
		t.Fatalf("unknown %#v %v", op, err)
	}
	if _, err := s.ChangeDNS(ctx, member, z.ID, change); err != nil || f.writes != 1 {
		t.Fatal("ambiguous write replayed")
	}
	if err := s.RevokeDNSPolicy(ctx, admin, g.ID); err != nil {
		t.Fatal(err)
	}
	if _, err := s.DNSRecords(ctx, member, z.ID, 1, 20); !errors.Is(err, ErrDenied) {
		t.Fatal("revoked grant usable")
	}
}
