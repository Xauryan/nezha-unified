// Package unifiedbackup implements a coordinated, authenticated backup format.
// Archive primitives receive a consistent snapshot through Quiesce after the
// host stops writers. Dashboard adapters open stores only offline, without
// executors. Restore never starts an executor.
package unifiedbackup

import (
	"archive/tar"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"slices"
	"strings"
	"time"
)

const (
	FormatVersion = 1
	// RecoveryMarker prevents an ordinary start of a restored store until the
	// host explicitly reconciles operation journals and remote side effects.
	RecoveryMarker       = ".unified-restore-pending.json"
	maxManifest          = 16 << 20
	maxFiles             = 100000
	maxFileSize    int64 = 256 << 30
	maxTotalSize   int64 = 1 << 40
)

// Snapshot paths are slash-separated relative paths below Root. Root must
// contain both closed/checkpointed databases, all file assets and any operation
// journals. The host hashes its stable resource references into ReferenceDigest.
// Versions and Schemas contain public version identifiers, never configuration.
type Snapshot struct {
	Root            string
	NezhaDB         string
	CertimateDB     string
	MasterKey       string
	ConfigFile      string
	Versions        map[string]string
	Schemas         map[string]string
	ReferenceDigest string
	Boundary        string
}

// Hooks must prevent new requests, enqueue and scheduled side effects, drain
// in-flight mutations, then close both database connections before returning.
// Quiesce must provide resume even after a partial failure if cleanup is needed.
// Create always calls a returned resume, including on cancellation and failure.
type Hooks struct {
	Quiesce func(context.Context) (snapshot Snapshot, resume func() error, err error)
}

type File struct {
	Path   string `json:"path"`
	Size   int64  `json:"size"`
	SHA256 string `json:"sha256"`
}

type Manifest struct {
	Format           int               `json:"format"`
	ID               string            `json:"id"`
	CreatedAt        time.Time         `json:"created_at"`
	Boundary         string            `json:"boundary"`
	Versions         map[string]string `json:"versions"`
	Schemas          map[string]string `json:"schemas"`
	ReferenceDigest  string            `json:"reference_digest"`
	NezhaDB          string            `json:"nezha_db"`
	CertimateDB      string            `json:"certimate_db"`
	MasterKey        string            `json:"master_key"`
	ConfigFile       string            `json:"config_file,omitempty"`
	WrappedMasterKey string            `json:"wrapped_master_key"`
	Files            []File            `json:"files"`
}

// RestoreHooks run offline, before publication. Validate must open both stores
// without workers, check database integrity, key decryption, schema compatibility
// and ReferenceDigest. PrepareRecovery must repair cross-store intents and mark
// in-flight remote operations unknown/reconciliation-required, without replay.
// Neither callback may leave database handles open when it returns.
type RestoreHooks struct {
	Validate        func(context.Context, string, Manifest) error
	PrepareRecovery func(context.Context, string, Manifest) error
}

// Create publishes a backup only after hashing every source file twice while
// writers are quiesced. recoveryKey is exactly 32 random bytes, managed separately
// from the backup; it is not a password and must not be passed on a command line.
// Archive data and its manifest are encrypted; the master key is additionally
// wrapped using a distinct authenticated context. Existing output is preserved.
func Create(ctx context.Context, archivePath string, recoveryKey []byte, hooks Hooks) (manifest Manifest, resultErr error) {
	if len(recoveryKey) != 32 {
		return manifest, errors.New("backup recovery key must contain exactly 32 bytes")
	}
	if hooks.Quiesce == nil {
		return manifest, errors.New("coordinated quiesce hook is required")
	}
	if err := ctx.Err(); err != nil {
		return manifest, err
	}
	snapshot, resume, err := hooks.Quiesce(ctx)
	if resume != nil {
		defer func() { resultErr = errors.Join(resultErr, resume()) }()
	}
	if err != nil {
		return manifest, err
	}
	if resume == nil {
		return manifest, errors.New("quiesce must provide a resume hook")
	}
	if err := validateSnapshot(snapshot); err != nil {
		return manifest, err
	}
	rootPath, err := filepath.Abs(snapshot.Root)
	if err != nil {
		return manifest, err
	}
	outputPath, err := filepath.Abs(archivePath)
	if err != nil {
		return manifest, err
	}
	if isWithin(rootPath, outputPath) {
		return manifest, errors.New("backup output must be outside the source tree")
	}
	if err := rejectSymlinks(rootPath); err != nil {
		return manifest, err
	}
	if err := rejectSymlinks(filepath.Dir(outputPath)); err != nil {
		return manifest, err
	}
	root, err := os.OpenRoot(rootPath)
	if err != nil {
		return manifest, errors.New("cannot open backup source")
	}
	defer root.Close()
	master, err := readRegular(root, snapshot.MasterKey, 32)
	if err != nil || len(master) != 32 {
		return manifest, errors.New("matching 32-byte master key is required")
	}
	defer clear(master)
	if subtle.ConstantTimeCompare(master, recoveryKey) == 1 {
		return manifest, errors.New("backup recovery key must be independent of the live master key")
	}
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		return manifest, err
	}
	defer clear(key)
	id := make([]byte, 16)
	if _, err := rand.Read(id); err != nil {
		return manifest, err
	}
	manifest = Manifest{Format: FormatVersion, ID: hex.EncodeToString(id), CreatedAt: time.Now().UTC(), Boundary: snapshot.Boundary,
		Versions: snapshot.Versions, Schemas: snapshot.Schemas, ReferenceDigest: snapshot.ReferenceDigest,
		NezhaDB: snapshot.NezhaDB, CertimateDB: snapshot.CertimateDB, MasterKey: snapshot.MasterKey, ConfigFile: snapshot.ConfigFile}
	wrapped, err := seal(key, master, []byte(manifest.ID+":master-key"))
	if err != nil {
		return Manifest{}, err
	}
	manifest.WrappedMasterKey = base64.StdEncoding.EncodeToString(wrapped)
	manifest.Files, err = inventory(ctx, root, snapshot.MasterKey)
	if err != nil {
		return Manifest{}, err
	}
	if err := validateManifest(manifest); err != nil {
		return Manifest{}, err
	}
	manifestJSON, err := json.Marshal(manifest)
	if err != nil || len(manifestJSON) > maxManifest {
		return Manifest{}, errors.New("backup manifest exceeds limits")
	}
	tmp, err := os.CreateTemp(filepath.Dir(outputPath), ".unified-backup-*")
	if err != nil {
		return Manifest{}, errors.New("cannot create backup output")
	}
	defer func() { _ = tmp.Close(); _ = os.Remove(tmp.Name()) }()
	stream, err := newEncryptWriter(tmp, recoveryKey, key, manifest.ID)
	if err != nil {
		return Manifest{}, err
	}
	tw := tar.NewWriter(stream)
	if err := addBytes(tw, "manifest.json", manifestJSON); err != nil {
		return Manifest{}, err
	}
	for _, item := range manifest.Files {
		if err := ctx.Err(); err != nil {
			return Manifest{}, err
		}
		f, err := openRegular(root, item.Path)
		if err != nil {
			return Manifest{}, errors.New("snapshot changed while backing up")
		}
		h := sha256.New()
		err = tw.WriteHeader(&tar.Header{Name: "data/" + item.Path, Mode: 0600, Size: item.Size, Typeflag: tar.TypeReg})
		if err == nil {
			var n int64
			n, err = io.Copy(tw, io.TeeReader(&contextReader{ctx, f}, h))
			if err == nil && (n != item.Size || hex.EncodeToString(h.Sum(nil)) != item.SHA256) {
				err = errors.New("snapshot changed while backing up")
			}
		}
		err = errors.Join(err, f.Close())
		if err != nil {
			return Manifest{}, err
		}
	}
	if err := tw.Close(); err != nil {
		return Manifest{}, err
	}
	if err := stream.Close(); err != nil {
		return Manifest{}, err
	}
	if err := tmp.Sync(); err != nil {
		return Manifest{}, err
	}
	if err := tmp.Close(); err != nil {
		return Manifest{}, err
	}
	// Hard-link publication is atomic and refuses an existing path, including a
	// symlink. Both files are deliberately on the same filesystem.
	if err := os.Link(tmp.Name(), outputPath); err != nil {
		return Manifest{}, errors.New("cannot publish backup; destination must not exist")
	}
	if err := syncDirectory(filepath.Dir(outputPath)); err != nil {
		return manifest, err
	}
	return manifest, nil
}

// Restore validates and decrypts into a new private directory. destination must
// not exist. On failure only the newly reserved directory is removed. The
// RecoveryMarker remains until CompleteRecovery explicitly allows execution.
func Restore(ctx context.Context, archivePath, destination string, recoveryKey []byte, hooks RestoreHooks) (manifest Manifest, resultErr error) {
	if len(recoveryKey) != 32 {
		return manifest, errors.New("backup recovery key must contain exactly 32 bytes")
	}
	if hooks.Validate == nil || hooks.PrepareRecovery == nil {
		return manifest, errors.New("offline validation and recovery preparation hooks are required")
	}
	if err := ctx.Err(); err != nil {
		return manifest, err
	}
	destination, err := filepath.Abs(destination)
	if err != nil {
		return manifest, err
	}
	if err := rejectSymlinks(filepath.Dir(destination)); err != nil {
		return manifest, err
	}
	if err := rejectSymlinks(archivePath); err != nil {
		return manifest, err
	}
	archive, err := os.Open(archivePath)
	if err != nil {
		return manifest, errors.New("cannot open backup archive")
	}
	defer archive.Close()
	if info, err := archive.Stat(); err != nil || !info.Mode().IsRegular() {
		return manifest, errors.New("backup archive must be a regular file")
	}
	stream, key, id, err := newDecryptReader(archive, recoveryKey)
	if err != nil {
		return manifest, err
	}
	defer clear(key)
	tr := tar.NewReader(&contextReader{ctx, stream})
	header, err := tr.Next()
	if err != nil || header.Name != "manifest.json" || header.Typeflag != tar.TypeReg || header.Size > maxManifest || header.Size < 1 {
		return manifest, errors.New("invalid backup manifest")
	}
	data, err := io.ReadAll(tr)
	if err != nil {
		return manifest, errors.New("cannot authenticate backup manifest")
	}
	if err := json.Unmarshal(data, &manifest); err != nil {
		return Manifest{}, errors.New("invalid backup manifest")
	}
	if err := validateManifest(manifest); err != nil {
		return Manifest{}, err
	}
	if manifest.ID != id {
		return Manifest{}, errors.New("backup identity mismatch")
	}
	wrapped, err := base64.StdEncoding.DecodeString(manifest.WrappedMasterKey)
	if err != nil {
		return Manifest{}, errors.New("invalid wrapped master key")
	}
	master, err := unseal(key, wrapped, []byte(id+":master-key"))
	if err != nil || len(master) != 32 {
		return Manifest{}, errors.New("cannot authenticate backup master key")
	}
	defer clear(master)
	// Reserve the destination atomically; no restore can replace existing data.
	if err := os.Mkdir(destination, 0700); err != nil {
		return Manifest{}, errors.New("restore destination must not exist")
	}
	defer func() {
		if resultErr != nil {
			_ = os.RemoveAll(destination)
		}
	}()
	root, err := os.OpenRoot(destination)
	if err != nil {
		return Manifest{}, err
	}
	defer root.Close()
	if err := writeExclusive(root, RecoveryMarker, data); err != nil {
		return Manifest{}, err
	}
	for _, item := range manifest.Files {
		header, err := tr.Next()
		if err != nil || header.Typeflag != tar.TypeReg || header.Name != "data/"+item.Path || header.Size != item.Size {
			return Manifest{}, errors.New("backup entries do not match manifest")
		}
		if err := root.MkdirAll(filepath.Dir(item.Path), 0700); err != nil {
			return Manifest{}, errors.New("cannot restore directory")
		}
		f, err := root.OpenFile(item.Path, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0600)
		if err != nil {
			return Manifest{}, errors.New("cannot restore file")
		}
		h := sha256.New()
		n, copyErr := io.Copy(f, io.TeeReader(tr, h))
		err = errors.Join(copyErr, f.Sync(), f.Close())
		if err != nil {
			return Manifest{}, errors.New("cannot restore file data")
		}
		if n != item.Size || hex.EncodeToString(h.Sum(nil)) != item.SHA256 {
			return Manifest{}, errors.New("backup file hash mismatch")
		}
	}
	if _, err := tr.Next(); err != io.EOF {
		return Manifest{}, errors.New("unexpected backup entry")
	}
	// tar EOF alone is insufficient: drain and authenticate the terminal record.
	if n, err := io.Copy(io.Discard, stream); err != nil || n != 0 {
		return Manifest{}, errors.New("backup truncated or contains trailing data")
	}
	if err := root.MkdirAll(filepath.Dir(manifest.MasterKey), 0700); err != nil {
		return Manifest{}, err
	}
	if err := writeExclusive(root, manifest.MasterKey, master); err != nil {
		return Manifest{}, err
	}
	if err := hooks.Validate(ctx, destination, manifest); err != nil {
		return Manifest{}, errors.New("restored database validation failed")
	}
	if err := hooks.PrepareRecovery(ctx, destination, manifest); err != nil {
		return Manifest{}, errors.New("restored operation recovery preparation failed")
	}
	if err := ctx.Err(); err != nil {
		return Manifest{}, err
	}
	if err := syncDirectory(destination); err != nil {
		return Manifest{}, err
	}
	return manifest, syncDirectory(filepath.Dir(destination))
}

// Rehearse performs the exact restore, integrity and recovery path in a temporary
// destination, then removes only that rehearsal. Callbacks must remain offline.
func Rehearse(ctx context.Context, archivePath, tempParent string, recoveryKey []byte, hooks RestoreHooks) (Manifest, error) {
	if err := rejectSymlinks(tempParent); err != nil {
		return Manifest{}, err
	}
	container, err := os.MkdirTemp(tempParent, ".unified-rehearsal-*")
	if err != nil {
		return Manifest{}, err
	}
	defer os.RemoveAll(container)
	return Restore(ctx, archivePath, filepath.Join(container, "restored"), recoveryKey, hooks)
}

// CheckReady must run before opening any scheduler, worker or mutating API.
func CheckReady(rootPath string) error {
	if err := rejectSymlinks(rootPath); err != nil {
		return err
	}
	root, err := os.OpenRoot(rootPath)
	if err != nil {
		return err
	}
	defer root.Close()
	_, err = root.Lstat(RecoveryMarker)
	if errors.Is(err, os.ErrNotExist) {
		return nil
	}
	if err != nil {
		return errors.New("cannot check restore recovery state")
	}
	return errors.New("restored store requires explicit operation reconciliation before execution")
}

// CompleteRecovery removes the startup fence only after a host-supplied audit
// confirms remote outcomes and repaired resource references. A missing marker
// fails closed instead of silently approving the wrong destination.
func CompleteRecovery(ctx context.Context, rootPath string, reconcile func(context.Context, string, Manifest) error) error {
	if reconcile == nil {
		return errors.New("explicit reconciliation callback is required")
	}
	if err := rejectSymlinks(rootPath); err != nil {
		return err
	}
	root, err := os.OpenRoot(rootPath)
	if err != nil {
		return err
	}
	defer root.Close()
	data, err := readRegular(root, RecoveryMarker, maxManifest)
	if err != nil {
		return errors.New("restore recovery marker is required")
	}
	var manifest Manifest
	if err := json.Unmarshal(data, &manifest); err != nil {
		return errors.New("invalid recovery marker")
	}
	if err := validateManifest(manifest); err != nil {
		return err
	}
	if err := ctx.Err(); err != nil {
		return err
	}
	if err := reconcile(ctx, rootPath, manifest); err != nil {
		return errors.New("external operation reconciliation is incomplete")
	}
	if err := ctx.Err(); err != nil {
		return err
	}
	if err := root.Remove(RecoveryMarker); err != nil {
		return err
	}
	return syncDirectory(rootPath)
}

func validateSnapshot(s Snapshot) error {
	if s.Root == "" {
		return errors.New("snapshot root is required")
	}
	for _, p := range []string{s.NezhaDB, s.CertimateDB, s.MasterKey} {
		if !safePath(p) {
			return errors.New("snapshot paths must be safe relative paths")
		}
	}
	if s.NezhaDB == s.CertimateDB || s.MasterKey == s.NezhaDB || s.MasterKey == s.CertimateDB {
		return errors.New("snapshot paths must be distinct")
	}
	if s.ConfigFile != "" && (!safePath(s.ConfigFile) || s.ConfigFile == s.NezhaDB || s.ConfigFile == s.CertimateDB || s.ConfigFile == s.MasterKey) {
		return errors.New("snapshot configuration must have a distinct safe relative path")
	}
	if s.Boundary == "" || len(s.Versions) == 0 || len(s.Schemas) == 0 || !validDigest(s.ReferenceDigest) {
		return errors.New("consistent boundary, module versions, schema versions and reference digest are required")
	}
	return nil
}

func validateManifest(m Manifest) error {
	if m.Format != FormatVersion || len(m.ID) != 32 || m.CreatedAt.IsZero() || m.WrappedMasterKey == "" {
		return errors.New("unsupported or incomplete backup manifest")
	}
	if _, err := hex.DecodeString(m.ID); err != nil {
		return errors.New("invalid backup identity")
	}
	if err := validateSnapshot(Snapshot{Root: ".", NezhaDB: m.NezhaDB, CertimateDB: m.CertimateDB, MasterKey: m.MasterKey, ConfigFile: m.ConfigFile, Boundary: m.Boundary, Versions: m.Versions, Schemas: m.Schemas, ReferenceDigest: m.ReferenceDigest}); err != nil {
		return err
	}
	if len(m.Files) < 2 || len(m.Files) > maxFiles {
		return errors.New("invalid backup file count")
	}
	var total int64
	seen := make(map[string]bool, len(m.Files))
	for _, f := range m.Files {
		if !safePath(f.Path) || f.Path == m.MasterKey || f.Path == RecoveryMarker || seen[f.Path] || f.Size < 0 || f.Size > maxFileSize || !validDigest(f.SHA256) {
			return errors.New("invalid backup file inventory")
		}
		seen[f.Path] = true
		total += f.Size
		if total > maxTotalSize {
			return errors.New("backup exceeds total size limit")
		}
	}
	if !seen[m.NezhaDB] || !seen[m.CertimateDB] {
		return errors.New("both application databases are required")
	}
	if m.ConfigFile != "" && !seen[m.ConfigFile] {
		return errors.New("declared application configuration is absent from backup")
	}
	return nil
}

func inventory(ctx context.Context, root *os.Root, masterPath string) ([]File, error) {
	var items []File
	err := fs.WalkDir(root.FS(), ".", func(name string, entry fs.DirEntry, err error) error {
		if err != nil {
			return errors.New("cannot enumerate snapshot")
		}
		if err := ctx.Err(); err != nil {
			return err
		}
		if entry.Type()&fs.ModeSymlink != 0 {
			return errors.New("snapshot may not contain symlinks")
		}
		if entry.IsDir() {
			return nil
		}
		if name == RecoveryMarker {
			return errors.New("cannot back up a store pending recovery")
		}
		if name == masterPath {
			return nil
		}
		if !safePath(name) || !entry.Type().IsRegular() {
			return errors.New("snapshot contains an unsupported file")
		}
		f, err := openRegular(root, name)
		if err != nil {
			return err
		}
		h := sha256.New()
		n, copyErr := io.Copy(h, io.LimitReader(&contextReader{ctx, f}, maxFileSize+1))
		err = errors.Join(copyErr, f.Close())
		if err != nil {
			return errors.New("cannot hash snapshot file")
		}
		if n > maxFileSize {
			return errors.New("snapshot file exceeds size limit")
		}
		items = append(items, File{Path: name, Size: n, SHA256: hex.EncodeToString(h.Sum(nil))})
		if len(items) > maxFiles {
			return errors.New("snapshot exceeds file count limit")
		}
		return nil
	})
	slices.SortFunc(items, func(a, b File) int { return strings.Compare(a.Path, b.Path) })
	return items, err
}

func safePath(name string) bool {
	return name != "." && fs.ValidPath(name) && !strings.ContainsAny(name, "\\:\x00") && path.Clean(name) == name
}
func validDigest(value string) bool {
	decoded, err := hex.DecodeString(value)
	return err == nil && len(decoded) == sha256.Size
}
func isWithin(rootPath, name string) bool {
	rel, err := filepath.Rel(rootPath, name)
	return err == nil && rel != ".." && !strings.HasPrefix(rel, ".."+string(filepath.Separator))
}

func rejectSymlinks(name string) error {
	abs, err := filepath.Abs(name)
	if err != nil {
		return err
	}
	for current := abs; ; current = filepath.Dir(current) {
		info, err := os.Lstat(current)
		if err != nil {
			return errors.New("backup path is unavailable")
		}
		if info.Mode()&os.ModeSymlink != 0 {
			return errors.New("backup paths may not contain symlinks")
		}
		if filepath.Dir(current) == current {
			break
		}
	}
	return nil
}
func openRegular(root *os.Root, name string) (*os.File, error) {
	before, err := root.Lstat(name)
	if err != nil || !before.Mode().IsRegular() {
		return nil, errors.New("snapshot file must be regular")
	}
	for parent := path.Dir(name); parent != "."; parent = path.Dir(parent) {
		info, err := root.Lstat(parent)
		if err != nil || !info.IsDir() {
			return nil, errors.New("snapshot parent must be a directory")
		}
	}
	f, err := root.Open(name)
	if err != nil {
		return nil, errors.New("cannot open snapshot file")
	}
	after, err := f.Stat()
	if err != nil || !after.Mode().IsRegular() || !os.SameFile(before, after) {
		_ = f.Close()
		return nil, errors.New("snapshot file changed")
	}
	return f, nil
}
func readRegular(root *os.Root, name string, limit int64) ([]byte, error) {
	f, err := openRegular(root, name)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	data, err := io.ReadAll(io.LimitReader(f, limit+1))
	if err != nil || int64(len(data)) > limit {
		return nil, errors.New("invalid snapshot file size")
	}
	return data, nil
}
func writeExclusive(root *os.Root, name string, data []byte) error {
	f, err := root.OpenFile(name, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0600)
	if err != nil {
		return errors.New("cannot create restored file")
	}
	_, writeErr := f.Write(data)
	return errors.Join(writeErr, f.Sync(), f.Close())
}
func addBytes(tw *tar.Writer, name string, data []byte) error {
	if err := tw.WriteHeader(&tar.Header{Name: name, Mode: 0600, Size: int64(len(data)), Typeflag: tar.TypeReg}); err != nil {
		return err
	}
	_, err := tw.Write(data)
	return err
}
func syncDirectory(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return err
	}
	return errors.Join(f.Sync(), f.Close())
}
func seal(key, plaintext, aad []byte) ([]byte, error) {
	aead, err := newAEAD(key)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, aead.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return nil, err
	}
	return aead.Seal(nonce, nonce, plaintext, aad), nil
}
func unseal(key, ciphertext, aad []byte) ([]byte, error) {
	aead, err := newAEAD(key)
	if err != nil {
		return nil, err
	}
	if len(ciphertext) < aead.NonceSize()+aead.Overhead() {
		return nil, errors.New("invalid encrypted material")
	}
	return aead.Open(nil, ciphertext[:aead.NonceSize()], ciphertext[aead.NonceSize():], aad)
}
func newAEAD(key []byte) (cipher.AEAD, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	return cipher.NewGCM(block)
}

type contextReader struct {
	ctx context.Context
	r   io.Reader
}

func (r *contextReader) Read(p []byte) (int, error) {
	if err := r.ctx.Err(); err != nil {
		return 0, err
	}
	return r.r.Read(p)
}

// ReferenceDigest computes a stable digest from canonical public references.
// Callers should include owner, grant/version, operation ID and activation state;
// never include credential config or certificate/key contents.
func ReferenceDigest(refs any) (string, error) {
	data, err := json.Marshal(refs)
	if err != nil {
		return "", fmt.Errorf("cannot encode stable references")
	}
	digest := sha256.Sum256(data)
	return hex.EncodeToString(digest[:]), nil
}
