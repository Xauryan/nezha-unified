package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/certimate-go/certimate/pkg/embedded"
	"github.com/nezhahq/nezha/service/singleton"
	"github.com/nezhahq/nezha/service/unifiedbackup"
)

func registerBackupFlags() {
	flag.StringVar(&dashboardCliParam.BackupCreate, "backup-create", "", "Create encrypted offline backup at a new archive path")
	flag.StringVar(&dashboardCliParam.BackupRestore, "backup-restore", "", "Restore an encrypted backup into a new directory")
	flag.StringVar(&dashboardCliParam.BackupDestination, "backup-destination", "", "New directory for --backup-restore")
	flag.StringVar(&dashboardCliParam.BackupKeyFile, "backup-key-file", "", "Private file containing exactly 32 random recovery-key bytes")
	flag.StringVar(&dashboardCliParam.BackupRehearse, "backup-rehearse", "", "Validate a backup by performing an isolated temporary restore")
	flag.StringVar(&dashboardCliParam.BackupInspect, "backup-inspect-recovery", "", "Read pending operation IDs from a restored data directory")
	flag.StringVar(&dashboardCliParam.BackupReconcile, "backup-reconcile", "", "Reconcile a restored data directory using operation-specific evidence")
	flag.StringVar(&dashboardCliParam.BackupEvidence, "backup-evidence", "", "JSON evidence file for --backup-reconcile; never include secrets")
}

func runBackupCommand(ctx context.Context, p DashboardCliParam) (bool, error) {
	count := 0
	for _, value := range []string{p.BackupCreate, p.BackupRestore, p.BackupRehearse, p.BackupInspect, p.BackupReconcile} {
		if value != "" {
			count++
		}
	}
	if count == 0 {
		if p.BackupDestination != "" || p.BackupKeyFile != "" || p.BackupEvidence != "" {
			return true, errors.New("backup options require a backup action")
		}
		return false, nil
	}
	if count != 1 || p.StorageOnly {
		return true, errors.New("choose exactly one offline backup action")
	}
	if p.BackupInspect != "" || p.BackupReconcile != "" {
		root := p.BackupInspect
		if root == "" {
			root = p.BackupReconcile
		}
		if _, err := os.Stat(root); err != nil {
			return true, errors.New("restored data directory is unavailable")
		}
		unlock, err := unifiedbackup.LockDataRoot(root)
		if err != nil {
			return true, err
		}
		defer unlock()
		if p.BackupInspect != "" {
			report, err := unifiedbackup.ReadRecoveryReport(root)
			if err != nil {
				return true, err
			}
			return true, json.NewEncoder(os.Stdout).Encode(report)
		}
		evidenceFile, err := os.Open(p.BackupEvidence)
		if err != nil {
			return true, errors.New("operation evidence file is required")
		}
		defer evidenceFile.Close()
		var evidence unifiedbackup.RecoveryEvidence
		decoder := json.NewDecoder(io.LimitReader(evidenceFile, 16<<20))
		decoder.DisallowUnknownFields()
		if err := decoder.Decode(&evidence); err != nil {
			return true, errors.New("invalid recovery evidence JSON")
		}
		var trailing any
		if decoder.Decode(&trailing) != io.EOF {
			return true, errors.New("unexpected trailing recovery evidence")
		}
		err = unifiedbackup.CompleteRecovery(ctx, root, func(ctx context.Context, root string, m unifiedbackup.Manifest) error {
			return unifiedbackup.ReconcileDashboard(ctx, root, m, evidence)
		})
		if err == nil {
			fmt.Println("Recovery boundary reconciled. Restored workflows, DDNS and deployment targets remain disabled; review each explicitly. Reconfigure TSDB to the restored metrics directory before enabling it.")
		}
		return true, err
	}
	key, err := readBackupKey(p.BackupKeyFile)
	if err != nil {
		return true, err
	}
	defer clear(key)
	hooks := unifiedbackup.RestoreHooks{
		Validate: func(ctx context.Context, root string, m unifiedbackup.Manifest) error {
			if err := unifiedbackup.ValidateDashboard(ctx, root, m, singleton.Version); err != nil {
				return err
			}
			// Open performs migrations/key validation but never starts HTTP, a worker,
			// a scheduler or any provider. Only the newly restored copy is opened.
			engine, err := embedded.Open(filepath.Join(root, filepath.Dir(m.CertimateDB)))
			if err != nil {
				return errors.New("restored certificate storage/key validation failed")
			}
			return engine.Close()
		},
		PrepareRecovery: unifiedbackup.PrepareDashboardRecovery,
	}
	var manifest unifiedbackup.Manifest
	switch {
	case p.BackupCreate != "":
		root, err := filepath.Abs(filepath.Dir(p.DatabaseLocation))
		if err != nil {
			return true, err
		}
		keyPath, err := filepath.EvalSymlinks(p.BackupKeyFile)
		if err != nil {
			return true, err
		}
		keyPath, err = filepath.Abs(keyPath)
		if err != nil {
			return true, err
		}
		keyRelative, err := filepath.Rel(root, keyPath)
		if err != nil || keyRelative != ".." && !strings.HasPrefix(keyRelative, ".."+string(filepath.Separator)) {
			return true, errors.New("backup recovery key must be stored outside the Dashboard data directory")
		}
		if _, err := os.Stat(root); err != nil {
			return true, errors.New("Dashboard data directory is unavailable")
		}
		config, err := filepath.Abs(p.ConfigFile)
		if err != nil {
			return true, err
		}
		relative, err := filepath.Rel(root, config)
		if err != nil || relative == ".." || strings.HasPrefix(relative, ".."+string(filepath.Separator)) {
			return true, errors.New("configuration must reside inside the data directory for a complete backup")
		}
		if _, err := os.Stat(config); err != nil {
			return true, errors.New("Dashboard configuration is unavailable")
		}
		manifest, err = unifiedbackup.Create(ctx, p.BackupCreate, key, unifiedbackup.Hooks{Quiesce: func(ctx context.Context) (unifiedbackup.Snapshot, func() error, error) {
			unlock, err := unifiedbackup.LockDataRoot(root)
			if err != nil {
				return unifiedbackup.Snapshot{}, nil, err
			}
			if err := unifiedbackup.ValidateDashboardConfiguration(root, filepath.ToSlash(relative)); err != nil {
				return unifiedbackup.Snapshot{}, unlock, err
			}
			snapshot, err := unifiedbackup.DashboardSnapshot(ctx, root, filepath.Base(p.DatabaseLocation), singleton.Version)
			snapshot.ConfigFile = filepath.ToSlash(relative)
			return snapshot, unlock, err
		}})
		if err != nil {
			return true, err
		}
	case p.BackupRestore != "":
		if p.BackupDestination == "" {
			return true, errors.New("--backup-destination must name a new directory")
		}
		manifest, err = unifiedbackup.Restore(ctx, p.BackupRestore, p.BackupDestination, key, hooks)
		if err != nil {
			return true, err
		}
	case p.BackupRehearse != "":
		manifest, err = unifiedbackup.Rehearse(ctx, p.BackupRehearse, os.TempDir(), key, hooks)
		if err != nil {
			return true, err
		}
	}
	// IDs and counts are public metadata. Do not print wrapped keys, manifests,
	// configuration, database content or storage errors containing secret fields.
	fmt.Printf("Backup operation complete: archive_id=%s files=%d\n", manifest.ID, len(manifest.Files))
	if p.BackupRestore != "" {
		fmt.Println("Restored data is fenced. Inspect the recovery report and supply operation-specific evidence before starting Dashboard.")
	}
	return true, nil
}

func readBackupKey(name string) ([]byte, error) {
	if name == "" {
		return nil, errors.New("--backup-key-file is required")
	}
	before, err := os.Lstat(name)
	if err != nil || !before.Mode().IsRegular() || before.Mode().Perm()&0077 != 0 {
		return nil, errors.New("backup key must be a private regular file (mode 0600)")
	}
	file, err := os.Open(name)
	if err != nil {
		return nil, errors.New("cannot open backup key file")
	}
	defer file.Close()
	after, err := file.Stat()
	if err != nil || !os.SameFile(before, after) {
		return nil, errors.New("backup key file changed")
	}
	data, err := io.ReadAll(io.LimitReader(file, 33))
	if err != nil || len(data) != 32 {
		clear(data)
		return nil, errors.New("backup key must contain exactly 32 random bytes")
	}
	return data, nil
}
