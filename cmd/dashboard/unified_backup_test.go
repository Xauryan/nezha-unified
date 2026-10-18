package main

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/certimate-go/certimate/pkg/embedded"
)

func TestBackupKeyRequiresExactPrivateRegularFile(t *testing.T) {
	root := t.TempDir()
	name := filepath.Join(root, "key")
	if err := os.WriteFile(name, make([]byte, 32), 0600); err != nil {
		t.Fatal(err)
	}
	key, err := readBackupKey(name)
	if err != nil || len(key) != 32 {
		t.Fatalf("private synthetic key rejected: %v", err)
	}
	clear(key)
	if err := os.Chmod(name, 0644); err != nil {
		t.Fatal(err)
	}
	if _, err := readBackupKey(name); err == nil {
		t.Fatal("readable-by-others key accepted")
	}
	if err := os.Chmod(name, 0600); err != nil {
		t.Fatal(err)
	}
	link := filepath.Join(root, "link")
	if err := os.Symlink(name, link); err != nil {
		t.Fatal(err)
	}
	if _, err := readBackupKey(link); err == nil {
		t.Fatal("symlink key accepted")
	}
	if err := os.WriteFile(name, make([]byte, 33), 0600); err != nil {
		t.Fatal(err)
	}
	if _, err := readBackupKey(name); err == nil {
		t.Fatal("wrong key size accepted")
	}
}

func TestBackupCommandRejectsAmbiguousActionsBeforeOpeningStores(t *testing.T) {
	for _, p := range []DashboardCliParam{{BackupCreate: "a", BackupRestore: "b"}, {BackupCreate: "a", StorageOnly: true}, {BackupKeyFile: "missing"}} {
		if handled, err := runBackupCommand(context.Background(), p); !handled || err == nil {
			t.Fatal("ambiguous backup command accepted")
		}
	}
}

func TestUnavailableRuntimeFailsBeforeAnyDataInitialization(t *testing.T) {
	if embedded.RuntimeAvailable() {
		t.Skip("full runtime build")
	}
	previous := dashboardCliParam
	t.Cleanup(func() { dashboardCliParam = previous })
	root := filepath.Join(t.TempDir(), "uncreated")
	dashboardCliParam = DashboardCliParam{DatabaseLocation: filepath.Join(root, "sqlite.db"), ConfigFile: filepath.Join(root, "config.yaml")}
	if err := runDashboard(); err == nil {
		t.Fatal("storage build silently started executors")
	}
	if _, err := os.Stat(root); !os.IsNotExist(err) {
		t.Fatal("unavailable runtime changed data before rejection")
	}
}
