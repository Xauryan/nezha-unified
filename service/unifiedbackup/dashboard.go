package unifiedbackup

import (
	"bytes"
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"slices"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"sigs.k8s.io/yaml"
)

const RecoveryReport = ".unified-recovery-report.json"
const RecoveryAudit = ".unified-recovery-audit.json"

// DashboardVersions is tied to the retained sources in upstream.lock.json.
// Dashboard version is supplied by the build's singleton.Version.
func DashboardVersions(version string) map[string]string {
	return map[string]string{"dashboard": version, "nezha": "v2.3.18", "certimate": "v0.4.35", "unified_backup": "1"}
}

func offlineDB(root, relative string, writable bool) (*sql.DB, error) {
	if !safePath(relative) {
		return nil, errors.New("invalid database path")
	}
	name := filepath.Join(root, filepath.FromSlash(relative))
	if err := rejectSymlinks(name); err != nil {
		return nil, err
	}
	mode := "ro"
	if writable {
		mode = "rw"
	}
	u := url.URL{Scheme: "file", Path: name, RawQuery: "mode=" + mode + "&_busy_timeout=1000&_foreign_keys=on"}
	db, err := sql.Open("sqlite3", u.String())
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(1)
	return db, nil
}

func checkIntegrity(ctx context.Context, db *sql.DB) error {
	var result string
	if err := db.QueryRowContext(ctx, "PRAGMA integrity_check").Scan(&result); err != nil || result != "ok" {
		return errors.New("database integrity check failed")
	}
	rows, err := db.QueryContext(ctx, "PRAGMA foreign_key_check")
	if err != nil {
		return errors.New("foreign key check failed")
	}
	defer rows.Close()
	if rows.Next() {
		return errors.New("database foreign key check failed")
	}
	return rows.Err()
}

func tableExists(ctx context.Context, db *sql.DB, table string) bool {
	var count int
	return db.QueryRowContext(ctx, "SELECT count(*) FROM sqlite_master WHERE type='table' AND name=?", table).Scan(&count) == nil && count == 1
}
func tableDigest(ctx context.Context, db *sql.DB, query string, args ...any) (string, error) {
	rows, err := db.QueryContext(ctx, query, args...)
	if err != nil {
		return "", err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return "", err
	}
	hash := sha256.New()
	enc := json.NewEncoder(hash)
	for rows.Next() {
		values := make([]any, len(columns))
		pointers := make([]any, len(columns))
		for i := range values {
			pointers[i] = &values[i]
		}
		if err := rows.Scan(pointers...); err != nil {
			return "", err
		}
		if err := enc.Encode(values); err != nil {
			return "", err
		}
	}
	if err := rows.Err(); err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}

// DashboardSnapshot requires the caller to hold LockDataRoot; it never starts
// services, upgrades schemas or contacts a provider. SQLite journals are kept.
func DashboardSnapshot(ctx context.Context, root, nezhaDB, version string) (Snapshot, error) {
	if err := CheckReady(root); err != nil {
		return Snapshot{}, err
	}
	s := Snapshot{Root: root, NezhaDB: nezhaDB, CertimateDB: "certimate/data.db", MasterKey: "certimate/unified.key", Versions: DashboardVersions(version), Schemas: map[string]string{}, Boundary: "offline-exclusive-lock:" + time.Now().UTC().Format(time.RFC3339Nano)}
	digest, schemas, err := dashboardIntegrity(ctx, root, s.NezhaDB, s.CertimateDB)
	if err != nil {
		return Snapshot{}, err
	}
	s.ReferenceDigest = digest
	s.Schemas = schemas
	return s, nil
}

func dashboardIntegrity(ctx context.Context, root, nezhaPath, certimatePath string) (string, map[string]string, error) {
	nezha, err := offlineDB(root, nezhaPath, false)
	if err != nil {
		return "", nil, err
	}
	defer nezha.Close()
	cert, err := offlineDB(root, certimatePath, false)
	if err != nil {
		return "", nil, err
	}
	defer cert.Close()
	schemas := map[string]string{}
	for name, db := range map[string]*sql.DB{"nezha": nezha, "certimate": cert} {
		if err := checkIntegrity(ctx, db); err != nil {
			return "", nil, err
		}
		schema, err := tableDigest(ctx, db, "SELECT type,name,tbl_name,sql FROM sqlite_master WHERE name NOT LIKE 'sqlite_%' ORDER BY type,name")
		if err != nil {
			return "", nil, err
		}
		schemas[name] = schema
	}
	if tableExists(ctx, cert, "_migrations") {
		digest, err := tableDigest(ctx, cert, "SELECT * FROM _migrations ORDER BY 1")
		if err != nil {
			return "", nil, err
		}
		schemas["certimate_migrations"] = digest
	}
	// Hash only public stable references. No secrets, ciphertext, graph contents,
	// names, provider responses or private keys enter a manifest.
	digests := map[string]string{}
	queries := map[string]string{
		"users":                      "SELECT id,role,token_version FROM users ORDER BY id",
		"unified_resource_bindings":  "SELECT kind,resource_id,owner_id,active,version,operation_id FROM unified_resource_bindings ORDER BY kind,resource_id",
		"unified_run_bindings":       "SELECT run_id,workflow_id,policy_version,owner_version,revoked FROM unified_run_bindings ORDER BY run_id",
		"unified_dns_accounts":       "SELECT id,owner_id,credential_id,credential_version FROM unified_dns_accounts ORDER BY id",
		"unified_dns_zones":          "SELECT id,account_id,provider_id FROM unified_dns_zones ORDER BY id",
		"unified_resource_intents":   "SELECT operation_id,kind,resource_id,state FROM unified_resource_intents ORDER BY operation_id",
		"unified_grants":             "SELECT user_id,kind,resource_id,action FROM unified_grants ORDER BY user_id,kind,resource_id,action",
		"unified_scoped_tokens":      "SELECT id,user_id,user_version,scopes,expires_at,revoked_at FROM unified_scoped_tokens ORDER BY id",
		"unified_service_identities": "SELECT id,owner_id,active,version,scopes FROM unified_service_identities ORDER BY id",
		"unified_workflow_policies":  "SELECT workflow_id,version,run_as_user_id,service_id,identity_version,enabled,graph_digest,names,credentials,targets FROM unified_workflow_policies ORDER BY workflow_id",
		"unified_dns_policies":       "SELECT * FROM unified_dns_policies ORDER BY id",
		"unified_ddns_bindings":      "SELECT id,zone_id,record_id,owner_id,version,host_id,enabled,network_enabled FROM unified_ddns_bindings ORDER BY id",
		"unified_deployment_targets": "SELECT id,server_id,version,enabled,desired_version,applied_version FROM unified_deployment_targets ORDER BY id",
		"unified_deployments":        "SELECT id,operation_id,target_id,target_version,certificate_id,certificate_version,workflow_run_id,actor_id,state FROM unified_deployments ORDER BY id",
	}
	for table, query := range queries {
		if tableExists(ctx, nezha, table) {
			digest, err := tableDigest(ctx, nezha, query)
			if err != nil {
				return "", nil, err
			}
			digests[table] = digest
		}
	}
	if !tableExists(ctx, nezha, "unified_resource_bindings") {
		return "", nil, errors.New("unified ownership schema is required")
	}
	rows, err := nezha.QueryContext(ctx, "SELECT kind,resource_id,owner_id,active FROM unified_resource_bindings ORDER BY kind,resource_id")
	if err != nil {
		return "", nil, err
	}
	for rows.Next() {
		var kind, id string
		var owner uint64
		var active bool
		if err := rows.Scan(&kind, &id, &owner, &active); err != nil {
			rows.Close()
			return "", nil, err
		}
		if !active {
			continue
		}
		if owner == 0 {
			rows.Close()
			return "", nil, errors.New("active resource has no owner")
		}
		table := map[string]string{"credential": "access", "certificate": "certificate", "workflow": "workflow"}[kind]
		if table != "" {
			var count int
			if cert.QueryRowContext(ctx, "SELECT count(*) FROM "+table+" WHERE id=?", id).Scan(&count) != nil || count != 1 {
				rows.Close()
				return "", nil, errors.New("cross-store resource reference is missing")
			}
		}
	}
	if err := rows.Close(); err != nil {
		return "", nil, err
	}
	for _, table := range []string{"access", "certificate", "workflow", "workflow_run"} {
		if !tableExists(ctx, cert, table) {
			return "", nil, errors.New("Certimate resource schema is required")
		}
		digest, err := tableDigest(ctx, cert, "SELECT id FROM "+table+" ORDER BY id")
		if err != nil {
			return "", nil, err
		}
		digests["certimate_"+table] = digest
	}
	raw, _ := json.Marshal(digests)
	digest := sha256.Sum256(raw)
	return hex.EncodeToString(digest[:]), schemas, nil
}

// ValidateDashboard checks an exact supported source version and the recorded
// schema/reference boundary before an optional storage-only crypto validation.
func ValidateDashboard(ctx context.Context, root string, m Manifest, version string) error {
	if !reflect.DeepEqual(m.Versions, DashboardVersions(version)) {
		return errors.New("backup module versions do not match this Dashboard build")
	}
	digest, schemas, err := dashboardIntegrity(ctx, root, m.NezhaDB, m.CertimateDB)
	if err != nil {
		return err
	}
	if digest != m.ReferenceDigest || !reflect.DeepEqual(schemas, m.Schemas) {
		return errors.New("restored schema or resource references differ from backup")
	}
	return nil
}

type RecoveryOperation struct {
	Kind       string `json:"kind"`
	ID         string `json:"id"`
	PriorState string `json:"prior_state"`
}
type DashboardRecoveryReport struct {
	ArchiveID       string              `json:"archive_id"`
	ReferenceDigest string              `json:"reference_digest"`
	Operations      []RecoveryOperation `json:"operations"`
}
type RecoveryEvidence struct {
	ArchiveID        string              `json:"archive_id"`
	ReferenceDigest  string              `json:"reference_digest"`
	BoundaryEvidence string              `json:"boundary_evidence"`
	Operations       []OperationEvidence `json:"operations"`
}
type OperationEvidence struct {
	Kind     string `json:"kind"`
	ID       string `json:"id"`
	Outcome  string `json:"outcome"`
	Evidence string `json:"evidence"`
}

// PrepareDashboardRecovery quarantines every restored automation and any
// unfinished side effect. Local cancellation is not evidence of remote failure.
// Evidence is collected before allowing the new Dashboard to start; no old run
// or operation is automatically retried.
func PrepareDashboardRecovery(ctx context.Context, root string, m Manifest) error {
	report := DashboardRecoveryReport{ArchiveID: m.ID, ReferenceDigest: m.ReferenceDigest, Operations: []RecoveryOperation{}}
	for _, store := range []struct {
		path string
		cert bool
	}{{m.NezhaDB, false}, {m.CertimateDB, true}} {
		db, err := offlineDB(root, store.path, true)
		if err != nil {
			return err
		}
		tx, err := db.BeginTx(ctx, nil)
		if err != nil {
			db.Close()
			return err
		}
		type mutation struct{ table, kind, id, state, where, update string }
		mutations := []mutation{
			{"unified_dns_operations", "dns", "id", "state", "state NOT IN ('succeeded','failed','cancelled')", "state='reconciliation_required',error='restore_requires_reconciliation'"},
			{"unified_deployments", "deployment", "id", "state", "state NOT IN ('verified','failed','rolled_back','cancelled','superseded')", "state='reconciliation_required'"},
			{"unified_resource_intents", "resource_intent", "operation_id", "state", "state NOT IN ('active','committed','completed','failed')", "state='reconciliation_required'"},
		}
		if store.cert {
			mutations = []mutation{{"workflow_run", "workflow_run", "id", "status", "status IN ('pending','processing')", "status='canceled'"}}
		}
		for _, item := range mutations {
			var exists int
			if err := tx.QueryRowContext(ctx, "SELECT count(*) FROM sqlite_master WHERE type='table' AND name=?", item.table).Scan(&exists); err != nil {
				tx.Rollback()
				db.Close()
				return err
			}
			if exists == 0 {
				continue
			}
			rows, err := tx.QueryContext(ctx, "SELECT CAST("+item.id+" AS TEXT),"+item.state+" FROM "+item.table+" WHERE "+item.where+" ORDER BY "+item.id)
			if err != nil {
				tx.Rollback()
				db.Close()
				return err
			}
			for rows.Next() {
				var id, state string
				if err := rows.Scan(&id, &state); err != nil {
					rows.Close()
					tx.Rollback()
					db.Close()
					return err
				}
				report.Operations = append(report.Operations, RecoveryOperation{Kind: item.kind, ID: id, PriorState: state})
			}
			err = rows.Err()
			rows.Close()
			if err != nil {
				tx.Rollback()
				db.Close()
				return err
			}
			if _, err := tx.ExecContext(ctx, "UPDATE "+item.table+" SET "+item.update+" WHERE "+item.where); err != nil {
				tx.Rollback()
				db.Close()
				return err
			}
		}
		disables := map[string]string{"unified_workflow_policies": "enabled=0", "unified_ddns_bindings": "enabled=0,network_enabled=0", "unified_deployment_targets": "enabled=0", "unified_run_bindings": "revoked=1", "servers": "enable_ddns=0"}
		if store.cert {
			disables = map[string]string{"workflow": "enabled=0"}
		}
		for table, update := range disables {
			var exists int
			if err := tx.QueryRowContext(ctx, "SELECT count(*) FROM sqlite_master WHERE type='table' AND name=?", table).Scan(&exists); err != nil {
				tx.Rollback()
				db.Close()
				return err
			}
			if exists != 0 {
				if _, err := tx.ExecContext(ctx, "UPDATE "+table+" SET "+update); err != nil {
					tx.Rollback()
					db.Close()
					return err
				}
			}
		}
		err = tx.Commit()
		closeErr := db.Close()
		if err != nil || closeErr != nil {
			return errors.Join(err, closeErr)
		}
	}
	if err := quarantineDashboardConfiguration(root, m.ConfigFile); err != nil {
		return err
	}
	slices.SortFunc(report.Operations, func(a, b RecoveryOperation) int { return strings.Compare(a.Kind+":"+a.ID, b.Kind+":"+b.ID) })
	raw, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		return err
	}
	rootDir, err := os.OpenRoot(root)
	if err != nil {
		return err
	}
	defer rootDir.Close()
	// A restored installation can itself be backed up later. Preserve its old
	// reconciliation evidence under a distinct history path before creating the
	// report for this new restore boundary.
	for _, name := range []string{RecoveryReport, RecoveryAudit} {
		if _, err := rootDir.Lstat(name); errors.Is(err, os.ErrNotExist) {
			continue
		} else if err != nil {
			return err
		}
		directory := ".unified-recovery-history/" + m.ID
		if !safePath(directory) {
			return errors.New("invalid recovery history identity")
		}
		if err := rootDir.MkdirAll(directory, 0700); err != nil {
			return err
		}
		destination := directory + "/" + name
		if _, err := rootDir.Lstat(destination); !errors.Is(err, os.ErrNotExist) {
			return errors.New("recovery history already exists")
		}
		if err := rootDir.Rename(name, destination); err != nil {
			return err
		}
	}
	return writeExclusive(rootDir, RecoveryReport, raw)
}

// ValidateDashboardConfiguration rejects a configured TSDB outside the backup
// tree. Such a store cannot be covered by the exclusive data-root boundary.
func ValidateDashboardConfiguration(rootPath, configPath string) error {
	root, err := os.OpenRoot(rootPath)
	if err != nil {
		return err
	}
	defer root.Close()
	data, err := readRegular(root, configPath, maxManifest)
	if err != nil {
		return errors.New("cannot read Dashboard configuration")
	}
	var config map[string]any
	if err := yaml.Unmarshal(data, &config); err != nil {
		return errors.New("invalid Dashboard configuration")
	}
	tsdb, _ := config["tsdb"].(map[string]any)
	location, _ := tsdb["data_path"].(string)
	if location == "" {
		return nil
	}
	abs, err := filepath.Abs(location)
	if err != nil || !isWithin(rootPath, abs) {
		return errors.New("configured TSDB must reside inside the data directory for a complete backup")
	}
	return rejectSymlinks(abs)
}

func quarantineDashboardConfiguration(rootPath, configPath string) error {
	if configPath == "" {
		return nil
	}
	root, err := os.OpenRoot(rootPath)
	if err != nil {
		return err
	}
	defer root.Close()
	data, err := readRegular(root, configPath, maxManifest)
	if err != nil {
		return errors.New("restored Dashboard configuration is unavailable")
	}
	var config map[string]any
	if err := yaml.Unmarshal(data, &config); err != nil {
		return errors.New("invalid restored Dashboard configuration")
	}
	tsdb, _ := config["tsdb"].(map[string]any)
	if tsdb == nil {
		return nil
	}
	// A relative or absolute path may still address the original installation.
	// Retained TSDB files stay in the restored tree for explicit reconfiguration.
	tsdb["data_path"] = ""
	encoded, err := yaml.Marshal(config)
	if err != nil {
		return errors.New("cannot quarantine restored TSDB configuration")
	}
	f, err := root.OpenFile(configPath, os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	_, writeErr := f.Write(encoded)
	return errors.Join(writeErr, f.Sync(), f.Close())
}

func ReadRecoveryReport(root string) (DashboardRecoveryReport, error) {
	var report DashboardRecoveryReport
	rootDir, err := os.OpenRoot(root)
	if err != nil {
		return report, err
	}
	defer rootDir.Close()
	data, err := readRegular(rootDir, RecoveryReport, maxManifest)
	if err != nil {
		return report, err
	}
	err = json.Unmarshal(data, &report)
	return report, err
}

// ReconcileDashboard stores an operator's per-operation evidence and leaves
// automation disabled. "unknown" is intentionally not an acceptable outcome.
func ReconcileDashboard(ctx context.Context, root string, m Manifest, evidence RecoveryEvidence) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	report, err := ReadRecoveryReport(root)
	if err != nil {
		return err
	}
	if evidence.ArchiveID != m.ID || report.ArchiveID != m.ID || evidence.ReferenceDigest != report.ReferenceDigest || strings.TrimSpace(evidence.BoundaryEvidence) == "" || len(evidence.BoundaryEvidence) > 4096 {
		return errors.New("evidence must identify the exact restored boundary")
	}
	if len(evidence.Operations) != len(report.Operations) {
		return errors.New("every ambiguous operation requires evidence")
	}
	seen := map[string]bool{}
	for _, item := range evidence.Operations {
		key := item.Kind + ":" + item.ID
		if seen[key] || len(item.Evidence) > 4096 || strings.TrimSpace(item.Evidence) == "" {
			return errors.New("invalid or duplicate operation evidence")
		}
		if item.Outcome != "confirmed_applied" && item.Outcome != "confirmed_not_applied" && item.Outcome != "retired_without_replay" {
			return errors.New("unsupported recovery outcome")
		}
		found := false
		for _, wanted := range report.Operations {
			if wanted.Kind == item.Kind && wanted.ID == item.ID {
				found = true
				break
			}
		}
		if !found {
			return errors.New("evidence names an unknown operation")
		}
		seen[key] = true
	}
	for _, name := range []string{m.NezhaDB, m.CertimateDB} {
		db, err := offlineDB(root, name, false)
		if err != nil {
			return err
		}
		err = checkIntegrity(ctx, db)
		db.Close()
		if err != nil {
			return err
		}
	}
	raw, err := json.MarshalIndent(evidence, "", "  ")
	if err != nil {
		return err
	}
	rootDir, err := os.OpenRoot(root)
	if err != nil {
		return err
	}
	defer rootDir.Close()
	if previous, readErr := readRegular(rootDir, RecoveryAudit, maxManifest); readErr == nil {
		if bytes.Equal(previous, raw) {
			return nil
		}
		return errors.New("different recovery evidence is already recorded")
	} else if _, statErr := rootDir.Lstat(RecoveryAudit); !errors.Is(statErr, os.ErrNotExist) {
		return errors.New("existing recovery evidence is unreadable")
	}
	if err := writeExclusive(rootDir, RecoveryAudit, raw); err != nil {
		return fmt.Errorf("cannot record recovery evidence: %w", err)
	}
	return nil
}
