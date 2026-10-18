package unifiedbackup

import (
	"context"
	"database/sql"
	"encoding/json"
	"os"
	"path/filepath"
	"sigs.k8s.io/yaml"
	"testing"
)

func dashboardFixture(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	if err := os.Mkdir(filepath.Join(root, "certimate"), 0700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(root, "certimate", "unified.key"), make([]byte, 32), 0600); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(root, "config.yaml"), []byte("synthetic_test: true\n"), 0600); err != nil {
		t.Fatal(err)
	}
	makeDB := func(name string, queries []string) {
		db, err := sql.Open("sqlite3", filepath.Join(root, name))
		if err != nil {
			t.Fatal(err)
		}
		defer db.Close()
		for _, q := range queries {
			if _, err := db.Exec(q); err != nil {
				t.Fatal(err)
			}
		}
	}
	makeDB("sqlite.db", []string{
		"CREATE TABLE unified_resource_bindings (kind TEXT,resource_id TEXT,owner_id INTEGER,active BOOLEAN,version INTEGER,operation_id TEXT)",
		"INSERT INTO unified_resource_bindings VALUES ('workflow','workflow1',1,1,1,'create1')",
		"CREATE TABLE unified_run_bindings (run_id TEXT,workflow_id TEXT,policy_version INTEGER,owner_version INTEGER,revoked BOOLEAN)",
		"INSERT INTO unified_run_bindings VALUES ('run1','workflow1',1,1,0)",
		"CREATE TABLE unified_dns_operations (id INTEGER,state TEXT,error TEXT)",
		"INSERT INTO unified_dns_operations VALUES (1,'applying',''),(2,'succeeded','')",
		"CREATE TABLE unified_deployments (id INTEGER,state TEXT,operation_id TEXT,target_id INTEGER,target_version INTEGER,certificate_id TEXT,certificate_version TEXT,workflow_run_id TEXT,actor_id INTEGER)",
		"INSERT INTO unified_deployments (id,state) VALUES (1,'queued'),(2,'verified')",
		"CREATE TABLE unified_workflow_policies (enabled BOOLEAN,workflow_id TEXT,version INTEGER,run_as_user_id INTEGER,service_id INTEGER,identity_version INTEGER,graph_digest TEXT,names TEXT,credentials TEXT,targets TEXT)", "INSERT INTO unified_workflow_policies (enabled) VALUES (1)",
		"CREATE TABLE unified_ddns_bindings (enabled BOOLEAN,network_enabled BOOLEAN,id INTEGER,zone_id INTEGER,record_id TEXT,owner_id INTEGER,version INTEGER,host_id INTEGER)", "INSERT INTO unified_ddns_bindings (enabled,network_enabled) VALUES (1,1)",
		"CREATE TABLE unified_deployment_targets (enabled BOOLEAN,id INTEGER,server_id INTEGER,version INTEGER,desired_version INTEGER,applied_version INTEGER)", "INSERT INTO unified_deployment_targets (enabled) VALUES (1)",
		"CREATE TABLE servers (enable_ddns BOOLEAN)", "INSERT INTO servers VALUES (1)",
	})
	makeDB("certimate/data.db", []string{
		"CREATE TABLE access (id TEXT)", "CREATE TABLE certificate (id TEXT)",
		"CREATE TABLE workflow (id TEXT,enabled BOOLEAN)", "INSERT INTO workflow VALUES ('workflow1',1)",
		"CREATE TABLE workflow_run (id TEXT,status TEXT)", "INSERT INTO workflow_run VALUES ('run1','processing'),('run2','succeeded')",
	})
	return root
}

func TestDashboardLockBlocksBackupAndAllowsRelease(t *testing.T) {
	root := dashboardFixture(t)
	unlock, err := LockDataRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if other, err := LockDataRoot(root); err == nil {
		other()
		t.Fatal("live Dashboard lock accepted")
	}
	if err := unlock(); err != nil {
		t.Fatal(err)
	}
	unlock, err = LockDataRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	defer unlock()
	if _, err := DashboardSnapshot(context.Background(), root, "sqlite.db", "test"); err != nil {
		t.Fatal(err)
	}
}

func TestDashboardRestoreQuarantinesAndRequiresExactEvidence(t *testing.T) {
	ctx := context.Background()
	root := dashboardFixture(t)
	key := make([]byte, 32)
	key[0] = 1 // Synthetic recovery key distinct from the fixture store key.
	archive := filepath.Join(t.TempDir(), "backup.nzu")
	manifest, err := Create(ctx, archive, key, Hooks{Quiesce: func(ctx context.Context) (Snapshot, func() error, error) {
		unlock, err := LockDataRoot(root)
		if err != nil {
			return Snapshot{}, nil, err
		}
		s, err := DashboardSnapshot(ctx, root, "sqlite.db", "test")
		return s, unlock, err
	}})
	if err != nil {
		t.Fatal(err)
	}
	restored := filepath.Join(t.TempDir(), "restored")
	_, err = Restore(ctx, archive, restored, key, RestoreHooks{Validate: func(ctx context.Context, root string, m Manifest) error {
		return ValidateDashboard(ctx, root, m, "test")
	}, PrepareRecovery: PrepareDashboardRecovery})
	if err != nil {
		t.Fatal(err)
	}
	if CheckReady(restored) == nil {
		t.Fatal("restore startup fence missing")
	}
	report, err := ReadRecoveryReport(restored)
	if err != nil {
		t.Fatal(err)
	}
	if len(report.Operations) != 3 {
		t.Fatalf("wanted 3 ambiguous operations, got %d", len(report.Operations))
	}
	db, err := offlineDB(restored, "sqlite.db", false)
	if err != nil {
		t.Fatal(err)
	}
	var state string
	if err := db.QueryRow("SELECT state FROM unified_deployments WHERE id=1").Scan(&state); err != nil || state != "reconciliation_required" {
		t.Fatalf("deployment replay remained enabled: %s %v", state, err)
	}
	var enabled bool
	if err := db.QueryRow("SELECT enabled FROM unified_ddns_bindings").Scan(&enabled); err != nil || enabled {
		t.Fatalf("DDNS was not quarantined: %v", err)
	}
	db.Close()
	evidence := RecoveryEvidence{ArchiveID: manifest.ID, ReferenceDigest: manifest.ReferenceDigest, BoundaryEvidence: "Synthetic fixture: inspected isolated simulated provider journal"}
	reconcile := func(ctx context.Context, root string, m Manifest) error {
		return ReconcileDashboard(ctx, root, m, evidence)
	}
	if err := CompleteRecovery(ctx, restored, reconcile); err == nil {
		t.Fatal("missing operation evidence accepted")
	}
	for _, item := range report.Operations {
		evidence.Operations = append(evidence.Operations, OperationEvidence{Kind: item.Kind, ID: item.ID, Outcome: "retired_without_replay", Evidence: "Synthetic operation retired; no network executor available"})
	}
	if err := CompleteRecovery(ctx, restored, reconcile); err != nil {
		t.Fatal(err)
	}
	if err := CheckReady(restored); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(restored, RecoveryAudit)); err != nil {
		t.Fatal(err)
	}
	if err := CompleteRecovery(ctx, restored, reconcile); err == nil {
		t.Fatal("reconciliation accepted without fence")
	}
	// A new backup of a recovered installation must retain the earlier audit
	// without colliding with this restore's new report/audit names.
	secondArchive := filepath.Join(t.TempDir(), "second.nzu")
	second, err := Create(ctx, secondArchive, key, Hooks{Quiesce: func(ctx context.Context) (Snapshot, func() error, error) {
		unlock, err := LockDataRoot(restored)
		if err != nil {
			return Snapshot{}, nil, err
		}
		snapshot, err := DashboardSnapshot(ctx, restored, "sqlite.db", "test")
		return snapshot, unlock, err
	}})
	if err != nil {
		t.Fatal(err)
	}
	secondRoot := filepath.Join(t.TempDir(), "second-restored")
	_, err = Restore(ctx, secondArchive, secondRoot, key, RestoreHooks{Validate: func(ctx context.Context, root string, m Manifest) error {
		return ValidateDashboard(ctx, root, m, "test")
	}, PrepareRecovery: PrepareDashboardRecovery})
	if err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(secondRoot, ".unified-recovery-history", second.ID, RecoveryAudit)); err != nil {
		t.Fatal("prior reconciliation evidence was not preserved")
	}
}

func TestDashboardSnapshotRejectsDanglingReference(t *testing.T) {
	root := dashboardFixture(t)
	db, err := offlineDB(root, "certimate/data.db", true)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := db.Exec("DELETE FROM workflow"); err != nil {
		t.Fatal(err)
	}
	db.Close()
	if _, err := DashboardSnapshot(context.Background(), root, "sqlite.db", "test"); err == nil {
		t.Fatal("cross-store missing resource accepted")
	}
}

func TestDashboardVersionAndReferenceMismatchRejected(t *testing.T) {
	root := dashboardFixture(t)
	s, err := DashboardSnapshot(context.Background(), root, "sqlite.db", "test")
	if err != nil {
		t.Fatal(err)
	}
	m := Manifest{Versions: s.Versions, Schemas: s.Schemas, ReferenceDigest: s.ReferenceDigest, NezhaDB: s.NezhaDB, CertimateDB: s.CertimateDB}
	if err := ValidateDashboard(context.Background(), root, m, "other-version"); err == nil {
		t.Fatal("incompatible build accepted")
	}
	m.ReferenceDigest = "different"
	if err := ValidateDashboard(context.Background(), root, m, "test"); err == nil {
		t.Fatal("reference mismatch accepted")
	}
	raw, _ := json.Marshal(m)
	if len(raw) == 0 {
		t.Fatal("empty manifest")
	}
}

func TestDashboardTSDBMustBeIncludedAndRestoredPathIsDisabled(t *testing.T) {
	root := dashboardFixture(t)
	inside := filepath.Join(root, "metrics")
	if err := os.Mkdir(inside, 0700); err != nil {
		t.Fatal(err)
	}
	config := filepath.Join(root, "config.yaml")
	writeConfig := func(path string) {
		raw, err := yaml.Marshal(map[string]any{"tsdb": map[string]any{"data_path": path}, "synthetic_test": true})
		if err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(config, raw, 0600); err != nil {
			t.Fatal(err)
		}
	}
	writeConfig(t.TempDir())
	if err := ValidateDashboardConfiguration(root, "config.yaml"); err == nil {
		t.Fatal("external TSDB was silently excluded")
	}
	writeConfig(inside)
	if err := ValidateDashboardConfiguration(root, "config.yaml"); err != nil {
		t.Fatal(err)
	}
	if err := PrepareDashboardRecovery(context.Background(), root, Manifest{ID: "synthetic", NezhaDB: "sqlite.db", CertimateDB: "certimate/data.db", ConfigFile: "config.yaml"}); err != nil {
		t.Fatal(err)
	}
	raw, err := os.ReadFile(config)
	if err != nil {
		t.Fatal(err)
	}
	var restored map[string]any
	if err := yaml.Unmarshal(raw, &restored); err != nil {
		t.Fatal(err)
	}
	if restored["tsdb"].(map[string]any)["data_path"] != "" {
		t.Fatal("restored metrics would write into the old store")
	}
	if _, err := os.Stat(inside); err != nil {
		t.Fatal("retained metrics directory was removed")
	}
}
