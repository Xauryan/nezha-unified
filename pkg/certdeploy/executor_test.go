package certdeploy

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"github.com/nezhahq/nezha/pkg/agentproto"
	"math/big"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func fixture(t *testing.T, root, id string, sequence uint64) agentproto.Deployment {
	t.Helper()
	key, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	tmpl := x509.Certificate{SerialNumber: big.NewInt(int64(sequence)), DNSNames: []string{"synthetic.example.test"}, NotBefore: time.Now().Add(-time.Hour), NotAfter: time.Now().Add(time.Hour), KeyUsage: x509.KeyUsageDigitalSignature}
	der, err := x509.CreateCertificate(rand.Reader, &tmpl, &tmpl, &key.PublicKey, key)
	if err != nil {
		t.Fatal(err)
	}
	keyDER, _ := x509.MarshalPKCS8PrivateKey(key)
	hash := sha256.Sum256(der)
	return agentproto.Deployment{OperationID: id, Sequence: sequence, Fingerprint: hex.EncodeToString(hash[:]), CertificatePEM: string(pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: der})), PrivateKeyPEM: string(pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: keyDER})), Target: agentproto.Target{ID: "target", Version: 1, BundlePath: root, AllowedNames: []string{"synthetic.example.test"}, CheckCommand: []string{"/synthetic/check", "{certificate}"}, ReloadCommand: []string{"/synthetic/reload"}, HealthCommand: []string{"/synthetic/health"}}}
}
func allowApply(context.Context) error { return nil }
func TestAtomicBundleReplayVersionAndJournal(t *testing.T) {
	root := filepath.Join(t.TempDir(), "bundle")
	d := fixture(t, root, "first", 1)
	calls := 0
	e := Executor{Allowed: func() bool { return true }, Run: func(context.Context, []string) error { calls++; return nil }}
	receipt, err := e.Execute(context.Background(), d, nil, allowApply)
	if err != nil || !receipt.Verified {
		t.Fatalf("deploy: %+v %v", receipt, err)
	}
	current, _ := os.Readlink(filepath.Join(root, "current"))
	if current != "releases/first" {
		t.Fatal("bundle not atomic current")
	}
	key, _ := os.Stat(filepath.Join(root, "current/private.key"))
	if key.Mode().Perm() != 0600 {
		t.Fatal("key permissions")
	}
	before := calls
	_, err = e.Execute(context.Background(), d, nil, allowApply)
	if err != nil || calls != before {
		t.Fatal("replay reran commands")
	}
	raw, _ := os.ReadFile(filepath.Join(root, "receipts/first.json"))
	if strings.Contains(string(raw), "PRIVATE KEY") || strings.Contains(string(raw), d.PrivateKeyPEM) {
		t.Fatal("journal contains key")
	}
	newer := fixture(t, root, "second", 2)
	if _, err = e.Execute(context.Background(), newer, nil, allowApply); err != nil {
		t.Fatal(err)
	}
	stale := fixture(t, root, "third", 1)
	if _, err = e.Execute(context.Background(), stale, nil, allowApply); err == nil {
		t.Fatal("stale version replaced current")
	}
}
func TestBeforeApplyPermissionAndRollback(t *testing.T) {
	root := filepath.Join(t.TempDir(), "bundle")
	first := fixture(t, root, "first", 1)
	failHealth := false
	e := Executor{Allowed: func() bool { return true }, Run: func(_ context.Context, args []string) error {
		if failHealth && strings.HasSuffix(args[0], "health") {
			return errors.New("synthetic health fail")
		}
		return nil
	}}
	if _, err := e.Execute(context.Background(), first, nil, allowApply); err != nil {
		t.Fatal(err)
	}
	next := fixture(t, root, "denied", 2)
	_, err := e.Execute(context.Background(), next, nil, func(context.Context) error { return errors.New("revoked") })
	if err == nil {
		t.Fatal("revoked commit allowed")
	}
	current, _ := os.Readlink(filepath.Join(root, "current"))
	if current != "releases/first" {
		t.Fatal("revoked commit changed active bundle")
	}
	next = fixture(t, root, "third", 3)
	failHealth = true
	r, err := e.Execute(context.Background(), next, nil, allowApply)
	if err == nil || !r.RolledBack {
		t.Fatalf("rollback: %+v %v", r, err)
	}
	current, _ = os.Readlink(filepath.Join(root, "current"))
	if current != "releases/first" {
		t.Fatal("rollback did not restore previous bundle")
	}
}
func TestUnknownReconciliationDoesNotRepeatReload(t *testing.T) {
	root := filepath.Join(t.TempDir(), "bundle")
	d := fixture(t, root, "first", 1)
	calls := []string{}
	e := Executor{Allowed: func() bool { return true }, Run: func(_ context.Context, args []string) error { calls = append(calls, args[0]); return nil }}
	if _, err := e.Execute(context.Background(), d, nil, allowApply); err != nil {
		t.Fatal(err)
	}
	j, _ := loadJournal(root, d.OperationID)
	j.Receipt.Phase = "applying"
	j.Receipt.Verified = false
	if err := writeJournal(root, j); err != nil {
		t.Fatal(err)
	}
	calls = nil
	d.CertificatePEM = ""
	d.PrivateKeyPEM = ""
	r, err := e.Reconcile(context.Background(), d)
	if err != nil || !r.Verified {
		t.Fatal("reconcile", err)
	}
	if len(calls) != 1 || calls[0] != "/synthetic/health" {
		t.Fatal("reconciliation repeated effect", calls)
	}
}
func TestNamesMismatchAndSymlinkRejected(t *testing.T) {
	root := filepath.Join(t.TempDir(), "bundle")
	d := fixture(t, root, "first", 1)
	d.Target.AllowedNames = []string{"other.example.test"}
	if Validate(d) == nil {
		t.Fatal("unapproved actual SAN accepted")
	}
	parent := t.TempDir()
	if err := os.Symlink(parent, root); err != nil {
		t.Fatal(err)
	}
	if safeRoot(root) == nil {
		t.Fatal("symlink bundle root accepted")
	}
}

func TestExplicitOwnerModeAndReconcilePreserveTraversal(t *testing.T) {
	root := filepath.Join(t.TempDir(), "bundle")
	d := fixture(t, root, "owned", 1)
	uid, gid := os.Getuid(), os.Getgid()
	d.Target.OwnerUID = &uid
	d.Target.OwnerGID = &gid
	d.Target.FileMode = 0640
	e := Executor{Allowed: func() bool { return true }, Run: func(context.Context, []string) error { return nil }}
	if _, err := e.Execute(context.Background(), d, nil, allowApply); err != nil {
		t.Fatal(err)
	}
	info, _ := os.Stat(filepath.Join(root, "current/private.key"))
	if info.Mode().Perm() != 0640 {
		t.Fatal("explicit private key mode ignored")
	}
	d.CertificatePEM = ""
	d.PrivateKeyPEM = ""
	if _, err := e.Reconcile(context.Background(), d); err != nil {
		t.Fatal(err)
	}
	info, _ = os.Stat(root)
	if info.Mode().Perm() != 0711 {
		t.Fatal("reconciliation broke service traversal")
	}
	bad := fixture(t, filepath.Join(t.TempDir(), "bad"), "bad", 1)
	bad.Target.FileMode = 0644
	if Validate(bad) == nil {
		t.Fatal("world-readable key accepted")
	}
}
