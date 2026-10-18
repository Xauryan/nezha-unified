package unifiedbackup

import (
	"archive/tar"
	"bytes"
	"context"
	"crypto/rand"
	"encoding/json"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

type fixture struct {
	snapshot Snapshot
	archive  string
	key      []byte
	resumed  int
}

func newFixture(t *testing.T) *fixture {
	t.Helper()
	root := filepath.Join(t.TempDir(), "source")
	if err := os.MkdirAll(filepath.Join(root, "certimate", "storage"), 0700); err != nil {
		t.Fatal(err)
	}
	master := make([]byte, 32)
	if _, err := rand.Read(master); err != nil {
		t.Fatal(err)
	}
	for name, data := range map[string][]byte{
		"sqlite.db": []byte("synthetic-Nezha-database-and-grants"), "certimate/data.db": []byte("synthetic-Certimate-database-and-refs"),
		"certimate/unified.key": master, "certimate/storage/asset.pem": bytes.Repeat([]byte("synthetic-private-material"), 8192),
		"operation-journal.json": []byte(`{"synthetic-operation":"awaiting_reconciliation"}`),
	} {
		if err := os.WriteFile(filepath.Join(root, name), data, 0600); err != nil {
			t.Fatal(err)
		}
	}
	digest, _ := ReferenceDigest(map[string]string{"credential:synthetic": "owner:1"})
	return &fixture{snapshot: Snapshot{Root: root, NezhaDB: "sqlite.db", CertimateDB: "certimate/data.db", MasterKey: "certimate/unified.key", Versions: map[string]string{"nezha": "v2.3.18", "certimate": "v0.4.35"}, Schemas: map[string]string{"nezha": "1", "certimate": "1"}, ReferenceDigest: digest, Boundary: "closed-stores-synthetic"}, archive: filepath.Join(t.TempDir(), "backup.nzb"), key: bytes.Repeat([]byte{0x42}, 32)}
}
func (f *fixture) hooks() Hooks {
	return Hooks{Quiesce: func(context.Context) (Snapshot, func() error, error) {
		return f.snapshot, func() error { f.resumed++; return nil }, nil
	}}
}
func (f *fixture) create(t *testing.T) Manifest {
	t.Helper()
	m, err := Create(context.Background(), f.archive, f.key, f.hooks())
	if err != nil {
		t.Fatal(err)
	}
	return m
}
func restoreHooks() RestoreHooks {
	return RestoreHooks{Validate: func(context.Context, string, Manifest) error { return nil }, PrepareRecovery: func(context.Context, string, Manifest) error { return nil }}
}

func TestRoundTripEncryptedBackupReferencesAndRecoveryFence(t *testing.T) {
	f := newFixture(t)
	m := f.create(t)
	if f.resumed != 1 {
		t.Fatal("quiesced source was not resumed")
	}
	archive, err := os.ReadFile(f.archive)
	if err != nil {
		t.Fatal(err)
	}
	master, _ := os.ReadFile(filepath.Join(f.snapshot.Root, f.snapshot.MasterKey))
	for _, sensitive := range [][]byte{master, []byte("synthetic-Nezha-database"), []byte("synthetic-Certimate-database"), []byte("synthetic-private-material"), []byte(m.ReferenceDigest)} {
		if bytes.Contains(archive, sensitive) {
			t.Fatal("archive leaked unencrypted material")
		}
	}
	dest := filepath.Join(t.TempDir(), "restored")
	validated, prepared := false, false
	hooks := restoreHooks()
	hooks.Validate = func(_ context.Context, root string, got Manifest) error {
		if got.ID != m.ID || got.ReferenceDigest != m.ReferenceDigest {
			t.Fatal("stable reference/backup identity changed")
		}
		if CheckReady(root) == nil {
			t.Fatal("restore validation ran without execution fence")
		}
		validated = true
		return nil
	}
	hooks.PrepareRecovery = func(_ context.Context, root string, _ Manifest) error {
		if !validated {
			t.Fatal("recovery preceded validation")
		}
		prepared = true
		return nil
	}
	if _, err := Restore(context.Background(), f.archive, dest, f.key, hooks); err != nil {
		t.Fatal(err)
	}
	if !prepared {
		t.Fatal("recovery preparation missing")
	}
	for _, file := range append(m.Files, File{Path: m.MasterKey}) {
		want, _ := os.ReadFile(filepath.Join(f.snapshot.Root, file.Path))
		got, err := os.ReadFile(filepath.Join(dest, file.Path))
		if err != nil || !bytes.Equal(want, got) {
			t.Fatalf("restored data mismatch: %s", file.Path)
		}
		info, _ := os.Stat(filepath.Join(dest, file.Path))
		if info.Mode().Perm() != 0600 {
			t.Fatal("restored file is not private")
		}
	}
	if CheckReady(dest) == nil {
		t.Fatal("restored source may execute before reconciliation")
	}
	if err := CompleteRecovery(context.Background(), dest, func(context.Context, string, Manifest) error { return errors.New("synthetic remote outcome unknown") }); err == nil {
		t.Fatal("unresolved operation accepted")
	}
	if CheckReady(dest) == nil {
		t.Fatal("failed reconciliation removed fence")
	}
	if err := CompleteRecovery(context.Background(), dest, func(context.Context, string, Manifest) error { return nil }); err != nil {
		t.Fatal(err)
	}
	if err := CheckReady(dest); err != nil {
		t.Fatal(err)
	}
	if err := CompleteRecovery(context.Background(), dest, func(context.Context, string, Manifest) error { return nil }); err == nil {
		t.Fatal("missing recovery marker accepted")
	}
}

func TestBackupRefusesMissingKeyDBAndUnsafeSource(t *testing.T) {
	for _, kind := range []string{"missing-key", "invalid-key", "missing-nezha", "missing-certimate", "symlink-file", "symlink-directory", "pending-recovery", "output-in-source", "missing-version", "missing-reference"} {
		t.Run(kind, func(t *testing.T) {
			f := newFixture(t)
			switch kind {
			case "missing-key":
				_ = os.Remove(filepath.Join(f.snapshot.Root, f.snapshot.MasterKey))
			case "invalid-key":
				_ = os.WriteFile(filepath.Join(f.snapshot.Root, f.snapshot.MasterKey), []byte("synthetic invalid"), 0600)
			case "missing-nezha":
				_ = os.Remove(filepath.Join(f.snapshot.Root, f.snapshot.NezhaDB))
			case "missing-certimate":
				_ = os.Remove(filepath.Join(f.snapshot.Root, f.snapshot.CertimateDB))
			case "symlink-file":
				_ = os.Symlink(filepath.Join(f.snapshot.Root, f.snapshot.MasterKey), filepath.Join(f.snapshot.Root, "alias"))
			case "symlink-directory":
				_ = os.Symlink(filepath.Join(f.snapshot.Root, "certimate"), filepath.Join(f.snapshot.Root, "alias"))
			case "pending-recovery":
				_ = os.WriteFile(filepath.Join(f.snapshot.Root, RecoveryMarker), []byte("{}"), 0600)
			case "output-in-source":
				f.archive = filepath.Join(f.snapshot.Root, "backup.nzb")
			case "missing-version":
				f.snapshot.Versions = nil
			case "missing-reference":
				f.snapshot.ReferenceDigest = ""
			}
			if _, err := Create(context.Background(), f.archive, f.key, f.hooks()); err == nil {
				t.Fatal("unsafe snapshot accepted")
			}
			if f.resumed != 1 {
				t.Fatal("failure did not resume")
			}
			if _, err := os.Stat(f.archive); !errors.Is(err, os.ErrNotExist) {
				t.Fatal("failed backup published output")
			}
		})
	}
}

func TestQuiesceAndResumeAreMandatoryAndFailuresPreserveOutput(t *testing.T) {
	f := newFixture(t)
	if _, err := Create(context.Background(), f.archive, f.key, Hooks{}); err == nil {
		t.Fatal("missing quiesce accepted")
	}
	if _, err := Create(context.Background(), f.archive, f.key, Hooks{Quiesce: func(context.Context) (Snapshot, func() error, error) { return f.snapshot, nil, nil }}); err == nil {
		t.Fatal("missing resume accepted")
	}
	cleanup := false
	_, err := Create(context.Background(), f.archive, f.key, Hooks{Quiesce: func(context.Context) (Snapshot, func() error, error) {
		return Snapshot{}, func() error { cleanup = true; return nil }, errors.New("quiesce failed")
	}})
	if err == nil || !cleanup {
		t.Fatal("partial quiesce cleanup failed")
	}
	f.create(t)
	before, _ := os.ReadFile(f.archive)
	if _, err := Create(context.Background(), f.archive, f.key, f.hooks()); err == nil {
		t.Fatal("existing backup replaced")
	}
	after, _ := os.ReadFile(f.archive)
	if !bytes.Equal(before, after) {
		t.Fatal("existing backup modified")
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := Create(ctx, f.archive, f.key, f.hooks()); !errors.Is(err, context.Canceled) {
		t.Fatal("canceled backup accepted")
	}
}

func TestBackupRequiresIndependentRecoveryKey(t *testing.T) {
	f := newFixture(t)
	f.key, _ = os.ReadFile(filepath.Join(f.snapshot.Root, f.snapshot.MasterKey))
	if _, err := Create(context.Background(), f.archive, f.key, f.hooks()); err == nil {
		t.Fatal("live master key accepted as backup wrapping key")
	}
	if f.resumed != 1 {
		t.Fatal("recovery key validation failed to resume snapshot")
	}
}

func TestConfigurationPathIsBoundToInventory(t *testing.T) {
	for _, name := range []string{"config.yaml", "../outside", "missing.yaml", "sqlite.db", "certimate/unified.key"} {
		t.Run(name, func(t *testing.T) {
			f := newFixture(t)
			f.snapshot.ConfigFile = name
			if err := os.WriteFile(filepath.Join(f.snapshot.Root, "config.yaml"), []byte("site_name: synthetic"), 0600); err != nil {
				t.Fatal(err)
			}
			manifest, err := Create(context.Background(), f.archive, f.key, f.hooks())
			if name != "config.yaml" {
				if err == nil {
					t.Fatal("invalid configuration reference accepted")
				}
				return
			}
			if err != nil || manifest.ConfigFile != name {
				t.Fatal("configuration reference missing from backup")
			}
			restored, err := Restore(context.Background(), f.archive, filepath.Join(t.TempDir(), "restored"), f.key, restoreHooks())
			if err != nil || restored.ConfigFile != name {
				t.Fatal("configuration reference missing after restore")
			}
		})
	}
}

func TestRestoreRejectsTamperTruncationWrongKeyAndPreservesExistingData(t *testing.T) {
	f := newFixture(t)
	f.create(t)
	original, _ := os.ReadFile(f.archive)
	for _, kind := range []string{"header", "middle", "last", "truncate-record", "truncate-final", "trailing", "wrong-key"} {
		t.Run(kind, func(t *testing.T) {
			data := bytes.Clone(original)
			key := bytes.Clone(f.key)
			switch kind {
			case "header":
				data[len(streamMagic)] ^= 1
			case "middle":
				data[len(data)/2] ^= 1
			case "last":
				data[len(data)-1] ^= 1
			case "truncate-record":
				data = data[:len(data)/2]
			case "truncate-final":
				data = data[:len(data)-21]
			case "trailing":
				data = append(data, 1)
			case "wrong-key":
				key[0] ^= 1
			}
			archive := filepath.Join(t.TempDir(), "bad.nzb")
			_ = os.WriteFile(archive, data, 0600)
			dest := filepath.Join(t.TempDir(), "restore")
			if _, err := Restore(context.Background(), archive, dest, key, restoreHooks()); err == nil {
				t.Fatal("corrupt archive accepted")
			}
			if _, err := os.Stat(dest); !errors.Is(err, os.ErrNotExist) {
				t.Fatal("failed restore left data")
			}
		})
	}
	dest := t.TempDir()
	sentinel := filepath.Join(dest, "preserve")
	_ = os.WriteFile(sentinel, []byte("existing runtime data"), 0600)
	if _, err := Restore(context.Background(), f.archive, dest, f.key, restoreHooks()); err == nil {
		t.Fatal("existing destination accepted")
	}
	if b, _ := os.ReadFile(sentinel); string(b) != "existing runtime data" {
		t.Fatal("existing runtime data modified")
	}
}

func TestRestoreRejectsAuthenticatedTraversalSymlinkAndHashMismatch(t *testing.T) {
	f := newFixture(t)
	f.create(t)
	for _, kind := range []string{"traversal", "absolute", "backslash", "duplicate", "master-raw", "hash", "symlink", "missing-db", "extra-entry", "key-wrap", "unsupported-format"} {
		t.Run(kind, func(t *testing.T) {
			bad := rewriteArchive(t, f, func(m *Manifest, headers []*tar.Header, bodies [][]byte) ([]*tar.Header, [][]byte) {
				switch kind {
				case "traversal":
					m.Files[0].Path = "../outside"
				case "absolute":
					m.Files[0].Path = "/tmp/outside"
				case "backslash":
					m.Files[0].Path = "..\\outside"
				case "duplicate":
					m.Files[1].Path = m.Files[0].Path
				case "master-raw":
					m.Files[0].Path = m.MasterKey
				case "hash":
					m.Files[0].SHA256 = strings.Repeat("0", 64)
				case "symlink":
					headers[0].Typeflag = tar.TypeSymlink
					headers[0].Linkname = "../../outside"
					headers[0].Size = 0
					bodies[0] = nil
				case "missing-db":
					m.NezhaDB = "absent.db"
				case "extra-entry":
					headers = append(headers, &tar.Header{Name: "data/extra", Typeflag: tar.TypeReg, Size: 1, Mode: 0600})
					bodies = append(bodies, []byte{1})
				case "key-wrap":
					m.WrappedMasterKey = "invalid"
				case "unsupported-format":
					m.Format++
				}
				return headers, bodies
			})
			dest := filepath.Join(t.TempDir(), "restore")
			if _, err := Restore(context.Background(), bad, dest, f.key, restoreHooks()); err == nil {
				t.Fatal("invalid authenticated archive accepted")
			}
			if _, err := os.Stat(dest); !errors.Is(err, os.ErrNotExist) {
				t.Fatal("failed restore retained destination")
			}
		})
	}
}

func TestRestoreHooksAndOfflineRehearsal(t *testing.T) {
	f := newFixture(t)
	f.create(t)
	for _, kind := range []string{"missing", "validation", "recovery"} {
		t.Run(kind, func(t *testing.T) {
			h := restoreHooks()
			switch kind {
			case "missing":
				h.Validate = nil
			case "validation":
				h.Validate = func(context.Context, string, Manifest) error { return errors.New("synthetic secret must not escape") }
			case "recovery":
				h.PrepareRecovery = func(context.Context, string, Manifest) error { return errors.New("synthetic secret must not escape") }
			}
			dest := filepath.Join(t.TempDir(), "restore")
			_, err := Restore(context.Background(), f.archive, dest, f.key, h)
			if err == nil {
				t.Fatal("failed hook accepted")
			}
			if strings.Contains(err.Error(), "synthetic secret") {
				t.Fatal("hook error leaked confidential context")
			}
			if _, err := os.Stat(dest); !errors.Is(err, os.ErrNotExist) {
				t.Fatal("failed hook retained destination")
			}
		})
	}
	parent := t.TempDir()
	count := 0
	h := restoreHooks()
	h.Validate = func(context.Context, string, Manifest) error { count++; return nil }
	if _, err := Rehearse(context.Background(), f.archive, parent, f.key, h); err != nil {
		t.Fatal(err)
	}
	entries, _ := os.ReadDir(parent)
	if len(entries) != 0 || count != 1 {
		t.Fatal("rehearsal failed to validate or clean up")
	}
}

func TestSourceAndDestinationSymlinkRootsRejected(t *testing.T) {
	f := newFixture(t)
	f.create(t)
	link := filepath.Join(t.TempDir(), "linked")
	if err := os.Symlink(t.TempDir(), link); err != nil {
		t.Fatal(err)
	}
	if _, err := Restore(context.Background(), f.archive, filepath.Join(link, "restored"), f.key, restoreHooks()); err == nil {
		t.Fatal("symlink parent accepted")
	}
	sourceLink := filepath.Join(t.TempDir(), "source")
	_ = os.Symlink(f.snapshot.Root, sourceLink)
	f.snapshot.Root = sourceLink
	f.archive = filepath.Join(t.TempDir(), "backup")
	if _, err := Create(context.Background(), f.archive, f.key, f.hooks()); err == nil {
		t.Fatal("symlink snapshot root accepted")
	}
}

func rewriteArchive(t *testing.T, f *fixture, mutate func(*Manifest, []*tar.Header, [][]byte) ([]*tar.Header, [][]byte)) string {
	t.Helper()
	raw, err := os.ReadFile(f.archive)
	if err != nil {
		t.Fatal(err)
	}
	r, key, id, err := newDecryptReader(bytes.NewReader(raw), f.key)
	if err != nil {
		t.Fatal(err)
	}
	defer clear(key)
	tr := tar.NewReader(r)
	_, err = tr.Next()
	if err != nil {
		t.Fatal(err)
	}
	data, _ := io.ReadAll(tr)
	var m Manifest
	if err := json.Unmarshal(data, &m); err != nil {
		t.Fatal(err)
	}
	var headers []*tar.Header
	var bodies [][]byte
	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatal(err)
		}
		body, err := io.ReadAll(tr)
		if err != nil {
			t.Fatal(err)
		}
		headers = append(headers, header)
		bodies = append(bodies, body)
	}
	headers, bodies = mutate(&m, headers, bodies)
	var out bytes.Buffer
	w, err := newEncryptWriter(&out, f.key, key, id)
	if err != nil {
		t.Fatal(err)
	}
	tw := tar.NewWriter(w)
	data, _ = json.Marshal(m)
	if err := addBytes(tw, "manifest.json", data); err != nil {
		t.Fatal(err)
	}
	for i, h := range headers {
		if err := tw.WriteHeader(h); err != nil {
			t.Fatal(err)
		}
		if _, err := tw.Write(bodies[i]); err != nil {
			t.Fatal(err)
		}
	}
	if err := tw.Close(); err != nil {
		t.Fatal(err)
	}
	if err := w.Close(); err != nil {
		t.Fatal(err)
	}
	dest := filepath.Join(t.TempDir(), "rewritten.nzb")
	if err := os.WriteFile(dest, out.Bytes(), 0600); err != nil {
		t.Fatal(err)
	}
	return dest
}
