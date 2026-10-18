package unified

import (
	"context"
	"errors"
	"fmt"
	"github.com/certimate-go/certimate/pkg/embedded"
	"github.com/nezhahq/nezha/service/unified/dnsprovider"
	"github.com/nezhahq/nezha/service/unified/policy"
	"sync"
	"testing"
	"time"
)

func TestDNSFormalServiceHasIndependentExactGrants(t *testing.T) {
	s, f, admin, _, z, userGrant, _ := ddnsFixture(t)
	ctx := context.Background()
	identity, err := s.CreateServiceIdentity(ctx, admin, ServiceIdentityInput{Name: "dns-platform", OwnerID: admin.UserID, Scopes: []Scope{{Kind: "dns_zone", ResourceID: fmt.Sprint(z.ID), Action: "read"}, {Kind: "dns_zone", ResourceID: fmt.Sprint(z.ID), Action: "write"}}})
	if err != nil {
		t.Fatal(err)
	}
	p := Principal{UserID: admin.UserID, ServiceID: identity.ID, IdentityVersion: identity.Version}
	if _, err := s.DNSRecords(ctx, p, z.ID, 1, 20); !errors.Is(err, ErrDenied) {
		t.Fatal("service inherited administrator account access without DNS grant")
	}
	g, err := s.CreateDNSPolicy(ctx, admin, DNSPolicy{ServiceID: identity.ID, ZoneID: z.ID, NameScope: policy.NameScope{Root: "delegated.example.com", Descendants: true}, Types: []string{"A"}, Actions: []string{"read", "update"}})
	if err != nil {
		t.Fatal(err)
	}
	if err := s.RevokeDNSPolicy(ctx, admin, userGrant.ID); err != nil {
		t.Fatal(err)
	}
	page, err := s.DNSRecords(ctx, p, z.ID, 1, 20)
	if err != nil || page.Total != 1 {
		t.Fatalf("formal service lost independent grant or escaped subtree: %#v %v", page, err)
	}
	r := f.records[0]
	r.Value = "192.0.2.20"
	if _, err := s.ChangeDNS(ctx, p, z.ID, DNSChange{OperationKey: "platform-service-change", Action: "update", Record: r}); err != nil || f.writes != 1 {
		t.Fatalf("formal service update %v", err)
	}
	if err := s.RevokeDNSPolicy(ctx, admin, g.ID); err != nil {
		t.Fatal(err)
	}
	if _, err := s.ChangeDNS(ctx, p, z.ID, DNSChange{OperationKey: "platform-revoked-change", Action: "update", Record: r}); !errors.Is(err, ErrDenied) || f.writes != 1 {
		t.Fatal("revoked formal DNS grant still usable")
	}
}

func TestDDNSFormalHandoverAndHostTransfer(t *testing.T) {
	s, f, admin, member, z, g, host := ddnsFixture(t)
	ctx := context.Background()
	b, err := s.SaveDDNSBinding(ctx, member, 0, DDNSBinding{ZoneID: z.ID, RecordID: "in", SourceKind: "host", HostID: host.ID, Enabled: true, NetworkEnabled: true})
	if err != nil {
		t.Fatal(err)
	}
	identity, err := s.CreateServiceIdentity(ctx, admin, ServiceIdentityInput{Name: "platform-ddns", OwnerID: admin.UserID, Scopes: []Scope{{Kind: "dns_zone", ResourceID: fmt.Sprint(z.ID), Action: "execute"}, {Kind: "ddns_binding", ResourceID: fmt.Sprint(b.ID), Action: "execute"}}})
	if err != nil {
		t.Fatal(err)
	}
	_, err = s.CreateDNSPolicy(ctx, admin, DNSPolicy{ServiceID: identity.ID, ZoneID: z.ID, NameScope: policy.NameScope{Root: b.Name, IncludeRoot: true}, Types: []string{"A"}, Actions: []string{"ddns"}})
	if err != nil {
		t.Fatal(err)
	}
	b, err = s.AssignDDNSService(ctx, admin, b.ID, identity.ID, b.Version)
	if err != nil {
		t.Fatal(err)
	}
	if err := s.RevokeDNSPolicy(ctx, admin, g.ID); err != nil {
		t.Fatal(err)
	}
	b, err = s.executeDDNS(ctx, b.ID)
	if err != nil || b.State != "current" || f.writes != 1 {
		t.Fatalf("formal handover %#v %v", b, err)
	}
	if err := s.db.Model(&host).Update("user_id", admin.UserID).Error; err != nil {
		t.Fatal(err)
	}
	b, err = s.executeDDNS(ctx, b.ID)
	if err != nil || b.Error != "host_unavailable" || f.writes != 1 {
		t.Fatalf("host transfer retained old approval %#v %v", b, err)
	}
}

func TestDDNSCrashIntentRequiresReadbackOrExplicitReconciliation(t *testing.T) {
	s, f, _, member, z, _, host := ddnsFixture(t)
	ctx := context.Background()
	b, err := s.SaveDDNSBinding(ctx, member, 0, DDNSBinding{ZoneID: z.ID, RecordID: "in", SourceKind: "host", HostID: host.ID, Enabled: true, NetworkEnabled: true})
	if err != nil {
		t.Fatal(err)
	}
	intent := DNSOperation{ActorID: member.UserID, Key: "crashed-ddns-intent", ZoneID: z.ID, RecordID: "in", RecordType: "A", Name: b.Name, Action: "update", State: "applying"}
	if err := s.db.Create(&intent).Error; err != nil {
		t.Fatal(err)
	}
	b, err = s.executeDDNS(ctx, b.ID)
	if err != nil || b.State != "unknown" || f.writes != 0 {
		t.Fatalf("crashed intent replayed %#v %v", b, err)
	}
	result, err := s.ReconcileDNSOperation(ctx, member, intent.ID, "")
	if err != nil || result.Operation.State != "reconciled" || result.Record.Value != "192.0.2.1" {
		t.Fatalf("reconcile %#v %v", result, err)
	}
	b, err = s.executeDDNS(ctx, b.ID)
	if err != nil || b.State != "current" || f.writes != 1 {
		t.Fatalf("acknowledged operation did not resume %#v %v", b, err)
	}
}

func TestDNSTokenRestrictionsAndWorkerSourceRevocation(t *testing.T) {
	s, f, admin, _, z, _, host := ddnsFixture(t)
	ctx := context.Background()
	issued, err := s.CreateScopedToken(ctx, admin, ScopedTokenInput{Name: "dns-reader", ExpiresAt: time.Now().Add(time.Hour), Scopes: []Scope{{Kind: "dns_zone", ResourceID: fmt.Sprint(z.ID), Action: "read"}}})
	if err != nil {
		t.Fatal(err)
	}
	p, err := s.AuthenticateScopedToken(ctx, issued.Secret)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := s.DNSRecords(ctx, p, z.ID, 1, 10); err != nil {
		t.Fatal(err)
	}
	if _, err := s.DNSRecords(ctx, p, z.ID+1, 1, 10); !errors.Is(err, ErrDenied) {
		t.Fatal("restricted administrator crossed zone")
	}
	if _, err := s.ChangeDNS(ctx, p, z.ID, DNSChange{OperationKey: "reader-tries-to-write", Action: "update", Record: f.records[0]}); !errors.Is(err, ErrDenied) || f.writes != 0 {
		t.Fatal("restricted administrator wrote")
	}
	b, err := s.SaveDDNSBinding(ctx, admin, 0, DDNSBinding{ZoneID: z.ID, RecordID: "in", SourceKind: "host", HostID: host.ID, Enabled: true, NetworkEnabled: true})
	if err != nil {
		t.Fatal(err)
	}
	worker, err := s.CreateScopedToken(ctx, admin, ScopedTokenInput{Name: "dns-worker", ExpiresAt: time.Now().Add(time.Hour), Scopes: []Scope{{Kind: "dns_zone", ResourceID: fmt.Sprint(z.ID), Action: "execute"}, {Kind: "ddns_binding", ResourceID: fmt.Sprint(b.ID), Action: "execute"}}})
	if err != nil {
		t.Fatal(err)
	}
	s.db.Model(&b).Update("token_id", worker.Token.ID)
	if err := s.RevokeScopedToken(ctx, admin, worker.Token.ID); err != nil {
		t.Fatal(err)
	}
	b, err = s.executeDDNS(ctx, b.ID)
	if err != nil || b.State != "blocked" || f.writes != 0 {
		t.Fatalf("revoked asynchronous token %#v %v", b, err)
	}
}

func TestDDNSStabilityChangeOnlyPauseAndRevocation(t *testing.T) {
	s, f, admin, member, z, g, host := ddnsFixture(t)
	ctx := context.Background()
	b, err := s.SaveDDNSBinding(ctx, member, 0, DDNSBinding{ZoneID: z.ID, RecordID: "in", SourceKind: "host", HostID: host.ID, StabilitySeconds: 60, Enabled: true, NetworkEnabled: true})
	if err != nil {
		t.Fatal(err)
	}
	b, err = s.RunDDNS(ctx, member, b.ID)
	if err != nil || b.State != "stabilizing" || f.writes != 0 {
		t.Fatalf("stability %#v %v", b, err)
	}
	s.db.Model(&b).Update("candidate_since", time.Now().Add(-2*time.Minute))
	b, err = s.RunDDNS(ctx, member, b.ID)
	if err != nil || b.LastSuccess != "192.0.2.9" || f.writes != 1 {
		t.Fatalf("update %#v %v", b, err)
	}
	var wg sync.WaitGroup
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func() { defer wg.Done(); s.RunDDNS(ctx, member, b.ID) }()
	}
	wg.Wait()
	if f.writes != 1 {
		t.Fatal("concurrent no-op repeated write")
	}
	change := DNSChange{OperationKey: "manual-record-update", Action: "update", Record: dnsprovider.Record{ID: "in", Name: b.Name, Type: "A", TTL: 600, Value: "192.0.2.12"}}
	if _, err := s.ChangeDNS(ctx, member, z.ID, change); !errors.Is(err, ErrDDNSManaged) {
		t.Fatalf("manual protection %v", err)
	}
	change.PauseDDNS = true
	if _, err := s.ChangeDNS(ctx, member, z.ID, change); err != nil {
		t.Fatal(err)
	}
	b, _ = s.RunDDNS(ctx, member, b.ID)
	if b.State != "paused" || f.writes != 2 {
		t.Fatal("paused binding wrote")
	}
	b.Enabled = true
	b.NetworkEnabled = false
	b, err = s.SaveDDNSBinding(ctx, member, b.ID, b)
	if err != nil {
		t.Fatal(err)
	}
	b, _ = s.RunDDNS(ctx, member, b.ID)
	if b.State != "paused" || f.writes != 2 {
		t.Fatal("network disabled binding wrote")
	}
	if err := s.RevokeDNSPolicy(ctx, admin, g.ID); err != nil {
		t.Fatal(err)
	}
	if _, err := s.RunDDNS(ctx, member, b.ID); !errors.Is(err, ErrDenied) {
		t.Fatal("revoked owner executed")
	}
	rows, err := s.DDNSBindings(ctx, member)
	if err != nil || len(rows) != 0 {
		t.Fatal("revoked bindings leaked")
	}
}
func TestDDNSDeviceExactBindingAndRevocation(t *testing.T) {
	s, f, admin, member, z, _, _ := ddnsFixture(t)
	ctx := context.Background()
	b, err := s.SaveDDNSBinding(ctx, member, 0, DDNSBinding{ZoneID: z.ID, RecordID: "in", SourceKind: "device", Enabled: true, NetworkEnabled: true})
	if err != nil {
		t.Fatal(err)
	}
	d, token, err := s.CreateDDNSDevice(ctx, member, b.ID)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := s.ReportDDNSDevice(ctx, token, "2001:db8::1"); !errors.Is(err, embedded.ErrInvalid) {
		t.Fatal("wrong family accepted")
	}
	b, err = s.ReportDDNSDevice(ctx, token, "192.0.2.18")
	if err != nil || b.LastSuccess != "192.0.2.18" || f.writes != 1 {
		t.Fatalf("device %#v %v", b, err)
	}
	if err := s.RevokeDDNSDevice(ctx, admin, d.ID); err != nil {
		t.Fatal(err)
	}
	if _, err := s.ReportDDNSDevice(ctx, token, "192.0.2.19"); !errors.Is(err, ErrDenied) {
		t.Fatal("revoked device accepted")
	}
	b, _ = s.RunDDNS(ctx, member, b.ID)
	if b.Error != "device_report_required" || f.writes != 1 {
		t.Fatal("worker used revoked observation")
	}
	if _, err := s.DDNSBindings(ctx, Principal{UserID: admin.UserID, Admin: true, PAT: true}); !errors.Is(err, ErrDenied) {
		t.Fatal("admin PAT bypass")
	}
}
func TestDDNSLifecycleAndStableIPv6Selection(t *testing.T) {
	s, _, _, _, _, _, _ := ddnsFixture(t)
	ctx := context.Background()
	s.ddns.interval = time.Hour
	source := func(context.Context, uint64) (DDNSHostSnapshot, error) { return DDNSHostSnapshot{}, nil }
	if err := s.StartDDNS(ctx, source); err != nil {
		t.Fatal(err)
	}
	if err := s.StartDDNS(ctx, source); err != nil {
		t.Fatal(err)
	}
	s.NotifyDDNSHost(1)
	s.StopDDNS()
	s.StopDDNS()
	if err := s.StartDDNS(ctx, source); err != nil {
		t.Fatal(err)
	}
	s.StopDDNS()
	b := DDNSBinding{RecordType: "AAAA", Interface: "eth0", LastSuccess: "2001:db8::2"}
	h := DDNSHostSnapshot{Addresses: []DDNSAddress{{Interface: "eth0", Address: "2001:db8::1"}, {Interface: "eth0", Address: "2001:db8::2"}, {Interface: "eth0", Address: "2001:db8::3", Temporary: true}, {Interface: "eth1", Address: "2001:db8::4"}}}
	if got := selectDDNSAddress(b, h); got != b.LastSuccess {
		t.Fatalf("stable IPv6 changed %s", got)
	}
	b.Address = "2001:db8::3"
	if got := selectDDNSAddress(b, h); got != "" {
		t.Fatal("temporary selected")
	}
	b.Address = "fe80::1"
	h.Addresses = append(h.Addresses, DDNSAddress{Interface: "eth0", Address: "fe80::1"})
	if got := selectDDNSAddress(b, h); got != "" {
		t.Fatal("link-local selected")
	}
}

func TestDDNSStopDrainsAndRejectsNewExecutions(t *testing.T) {
	s, _, _, member, z, _, host := ddnsFixture(t)
	ctx := context.Background()
	b, err := s.SaveDDNSBinding(ctx, member, 0, DDNSBinding{ZoneID: z.ID, RecordID: "in", SourceKind: "host", HostID: host.ID, Enabled: true, NetworkEnabled: true})
	if err != nil {
		t.Fatal(err)
	}
	started, release := make(chan struct{}, 1), make(chan struct{})
	source := func(context.Context, uint64) (DDNSHostSnapshot, error) {
		started <- struct{}{}
		<-release
		return DDNSHostSnapshot{}, nil
	}
	if err := s.StartDDNS(ctx, source); err != nil {
		t.Fatal(err)
	}
	select {
	case <-started:
	case <-time.After(3 * time.Second):
		t.Fatal("worker did not start")
	}
	deadline, cancel := context.WithTimeout(ctx, 20*time.Millisecond)
	defer cancel()
	if err := s.StopDDNSContext(deadline); !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("uncooperative source bypassed grace: %v", err)
	}
	if _, err := s.RunDDNS(ctx, member, b.ID); !errors.Is(err, ErrDDNSUnavailable) {
		t.Fatalf("new work admitted while stopping: %v", err)
	}
	close(release)
	if err := s.StopDDNSContext(ctx); err != nil {
		t.Fatal(err)
	}
}
func TestDNSBulkExportUsesCurrentGrants(t *testing.T) {
	s, f, admin, member, z, g, _ := ddnsFixture(t)
	ctx := context.Background()
	out, err := s.ExportDNS(ctx, member, z.ID)
	if err != nil || len(out.Records) != 1 {
		t.Fatalf("export %#v %v", out, err)
	}
	rows, err := s.ImportDNS(ctx, member, z.ID, "batch-import-synthetic", []dnsprovider.Record{{Name: "new.delegated.example.com", Type: "A", Value: "192.0.2.9", TTL: 600}, {Name: "evil-delegated.example.com", Type: "A", Value: "192.0.2.10", TTL: 600}})
	if err != nil || rows[0].Operation == nil || rows[1].Error == "" || f.writes != 1 {
		t.Fatalf("batch %#v %v", rows, err)
	}
	if err := s.RevokeDNSPolicy(ctx, admin, g.ID); err != nil {
		t.Fatal(err)
	}
	if _, err := s.ExportDNS(ctx, member, z.ID); !errors.Is(err, ErrDenied) {
		t.Fatal("revoked export")
	}
}
