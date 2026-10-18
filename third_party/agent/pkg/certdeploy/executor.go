// Package certdeploy implements the Agent's journalled certificate transaction.
// The bundle directory belongs to this executor; services consume current/*.
package certdeploy

import (
	"context"
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
	"sync"
	"time"

	"github.com/nezhahq/agent/pkg/agentproto"
)

var ErrInvalid = errors.New("invalid certificate deployment")
var safeID = regexp.MustCompile(`^[A-Za-z0-9_-]{1,128}$`)

type Executor struct {
	mu      sync.Mutex
	Allowed func() bool
	Run     func(context.Context, []string) error
}
type journal struct {
	Receipt  agentproto.Receipt `json:"receipt"`
	Digest   string             `json:"digest"`
	Previous string             `json:"previous"`
}

func Validate(d agentproto.Deployment) error {
	t := d.Target
	if !safeID.MatchString(t.ID) || !safeID.MatchString(d.OperationID) || t.Version == 0 || d.Sequence == 0 || !filepath.IsAbs(t.BundlePath) || filepath.Clean(t.BundlePath) == string(filepath.Separator) || len(t.AllowedNames) == 0 || len(t.ReloadCommand) == 0 || len(t.HealthCommand) == 0 || len(t.CheckCommand) == 0 {
		return ErrInvalid
	}
	for _, argv := range [][]string{t.CheckCommand, t.ReloadCommand, t.HealthCommand} {
		if len(argv) > 64 || !filepath.IsAbs(argv[0]) {
			return ErrInvalid
		}
		for _, v := range argv {
			if len(v) > 4096 || strings.ContainsRune(v, 0) {
				return ErrInvalid
			}
		}
	}
	if (t.TLSAddress == "") != (t.TLSServerName == "") {
		return ErrInvalid
	}
	if t.FileMode != 0 && t.FileMode != 0600 && t.FileMode != 0640 {
		return ErrInvalid
	}
	if t.OwnerUID != nil && *t.OwnerUID < 0 || t.OwnerGID != nil && *t.OwnerGID < 0 {
		return ErrInvalid
	}
	return ValidateCertificate(d.CertificatePEM, d.PrivateKeyPEM, d.Fingerprint, t.AllowedNames)
}

// ValidateCertificate checks the actual key, chain, validity and every SAN before
// an Agent or cloud target can use certificate material.
func ValidateCertificate(certificatePEM, privateKeyPEM, fingerprint string, allowedNames []string) error {
	if len(certificatePEM) > 131072 || len(privateKeyPEM) > 32768 {
		return ErrInvalid
	}
	pair, err := tls.X509KeyPair([]byte(certificatePEM), []byte(privateKeyPEM))
	if err != nil || len(pair.Certificate) == 0 {
		return ErrInvalid
	}
	leaf, err := x509.ParseCertificate(pair.Certificate[0])
	if err != nil || time.Now().Before(leaf.NotBefore) || time.Now().After(leaf.NotAfter) || len(leaf.DNSNames) == 0 {
		return ErrInvalid
	}
	if len(leaf.IPAddresses) > 0 || len(leaf.EmailAddresses) > 0 || len(leaf.URIs) > 0 {
		return ErrInvalid
	}
	hash := sha256.Sum256(leaf.Raw)
	if fingerprint != hex.EncodeToString(hash[:]) {
		return ErrInvalid
	}
	// Every actual SAN must be explicitly permitted. A literal wildcard does not
	// silently confer permission to upload any other wildcard or base domain.
	for _, name := range leaf.DNSNames {
		if !slices.Contains(allowedNames, strings.ToLower(strings.TrimSuffix(name, "."))) {
			return ErrInvalid
		}
	}
	for i := 0; i+1 < len(pair.Certificate); i++ {
		cert, err := x509.ParseCertificate(pair.Certificate[i])
		if err != nil {
			return ErrInvalid
		}
		issuer, err := x509.ParseCertificate(pair.Certificate[i+1])
		if err != nil || cert.CheckSignatureFrom(issuer) != nil {
			return ErrInvalid
		}
	}
	return nil
}
func (e *Executor) allowed() bool { return e.Allowed != nil && e.Allowed() }
func (e *Executor) run(ctx context.Context, argv []string, dir string) error {
	if !e.allowed() {
		return errors.New("disabled")
	}
	expanded := make([]string, len(argv))
	for i, s := range argv {
		expanded[i] = strings.NewReplacer("{directory}", dir, "{certificate}", filepath.Join(dir, "certificate.pem"), "{key}", filepath.Join(dir, "private.key")).Replace(s)
	}
	bounded, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	if e.Run != nil {
		return e.Run(bounded, expanded)
	}
	// Never record stdout/stderr: commands may print key or configuration values.
	return exec.CommandContext(bounded, expanded[0], expanded[1:]...).Run()
}
func atomicFile(path string, data []byte, mode os.FileMode) error {
	f, err := os.CreateTemp(filepath.Dir(path), ".write-")
	if err != nil {
		return err
	}
	tmp := f.Name()
	defer os.Remove(tmp)
	if err = f.Chmod(mode); err == nil {
		_, err = f.Write(data)
	}
	if err == nil {
		err = f.Sync()
	}
	closeErr := f.Close()
	if err != nil {
		return err
	}
	if closeErr != nil {
		return closeErr
	}
	if err = os.Rename(tmp, path); err != nil {
		return err
	}
	return syncDir(filepath.Dir(path))
}
func syncDir(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return f.Sync()
}
func writeJournal(root string, j journal) error {
	b, err := json.Marshal(j)
	if err != nil {
		return err
	}
	return atomicFile(filepath.Join(root, "receipts", j.Receipt.OperationID+".json"), b, 0600)
}
func loadJournal(root, id string) (journal, error) {
	var j journal
	b, err := os.ReadFile(filepath.Join(root, "receipts", id+".json"))
	if err == nil {
		err = json.Unmarshal(b, &j)
	}
	return j, err
}
func safeRoot(root string) error {
	if !filepath.IsAbs(root) || filepath.Clean(root) == string(filepath.Separator) {
		return ErrInvalid
	}
	// Reject symlink ancestors; the executor never follows a caller-controlled
	// redirect when creating journal, release, certificate or private-key files.
	for path := filepath.Clean(root); ; path = filepath.Dir(path) {
		info, err := os.Lstat(path)
		if err == nil && info.Mode()&os.ModeSymlink != 0 {
			return ErrInvalid
		}
		if err != nil && !os.IsNotExist(err) {
			return err
		}
		if path == filepath.Dir(path) {
			break
		}
	}
	if err := os.MkdirAll(root, 0700); err != nil {
		return err
	}
	info, err := os.Stat(root)
	if err != nil {
		return err
	}
	if info.Mode().Perm()&0022 != 0 {
		return ErrInvalid
	}
	for _, part := range []string{"receipts", "releases"} {
		p := filepath.Join(root, part)
		info, err := os.Lstat(p)
		if err == nil && (!info.IsDir() || info.Mode()&os.ModeSymlink != 0) {
			return ErrInvalid
		}
		if err = os.MkdirAll(p, 0700); err != nil {
			return err
		}
	}
	return nil
}
func current(root string) (string, error) {
	p := filepath.Join(root, "current")
	v, err := os.Readlink(p)
	if os.IsNotExist(err) {
		return "", nil
	}
	if err != nil {
		return "", err
	}
	if filepath.IsAbs(v) || !strings.HasPrefix(v, "releases/") || filepath.Clean(v) != v {
		return "", ErrInvalid
	}
	return v, nil
}
func switchCurrent(root, release string) error {
	if release == "" {
		err := os.Remove(filepath.Join(root, "current"))
		if err != nil && !os.IsNotExist(err) {
			return err
		}
		return syncDir(root)
	}
	tmp := filepath.Join(root, ".current-next")
	if err := os.Remove(tmp); err != nil && !os.IsNotExist(err) {
		return err
	}
	if err := os.Symlink(release, tmp); err != nil {
		return err
	}
	if err := os.Rename(tmp, filepath.Join(root, "current")); err != nil {
		return err
	}
	return syncDir(root)
}
func (e *Executor) Execute(ctx context.Context, d agentproto.Deployment, progress func(agentproto.Receipt), permit func(context.Context) error) (agentproto.Receipt, error) {
	e.mu.Lock()
	defer e.mu.Unlock()
	r := agentproto.Receipt{OperationID: d.OperationID, TargetID: d.Target.ID, TargetVersion: d.Target.Version, Sequence: d.Sequence, Fingerprint: d.Fingerprint, Phase: "staging"}
	if !e.allowed() {
		return r, errors.New("disabled")
	}
	if err := Validate(d); err != nil {
		return r, err
	}
	root := d.Target.BundlePath
	if err := safeRoot(root); err != nil {
		return r, err
	}
	body, _ := json.Marshal(d)
	sum := sha256.Sum256(body)
	digest := hex.EncodeToString(sum[:])
	if old, err := loadJournal(root, d.OperationID); err == nil {
		if old.Digest != digest {
			return r, errors.New("operation conflict")
		}
		return old.Receipt, phaseError(old.Receipt.Phase)
	} else if !os.IsNotExist(err) {
		return r, err
	}
	entries, err := os.ReadDir(filepath.Join(root, "receipts"))
	if err != nil {
		return r, err
	}
	for _, entry := range entries {
		if !strings.HasSuffix(entry.Name(), ".json") {
			continue
		}
		old, err := loadJournal(root, strings.TrimSuffix(entry.Name(), ".json"))
		if err != nil {
			return r, err
		}
		if old.Receipt.TargetID != d.Target.ID {
			return r, errors.New("bundle belongs to another target")
		}
		if old.Receipt.Sequence >= d.Sequence || old.Receipt.TargetVersion > d.Target.Version {
			return r, errors.New("stale deployment")
		}
		if old.Receipt.Phase == "applying" || old.Receipt.Phase == "unknown" {
			return r, errors.New("unresolved prior deployment")
		}
	}
	previous, err := current(root)
	if err != nil {
		return r, err
	}
	j := journal{Receipt: r, Digest: digest, Previous: previous}
	if err = writeJournal(root, j); err != nil {
		return r, err
	}
	phase := func(state string) error {
		j.Receipt.Phase = state
		if err := writeJournal(root, j); err != nil {
			return err
		}
		if progress != nil {
			progress(j.Receipt)
		}
		return nil
	}
	failed := func(code string, cause error) (agentproto.Receipt, error) {
		j.Receipt.Code = code
		_ = phase("failed")
		return j.Receipt, cause
	}
	release := filepath.Join("releases", d.OperationID)
	dir := filepath.Join(root, release)
	if err = os.Mkdir(dir, 0700); err != nil {
		return failed("stage", err)
	}
	if err = atomicFile(filepath.Join(dir, "certificate.pem"), []byte(d.CertificatePEM), 0600); err != nil {
		return failed("stage", err)
	}
	if err = atomicFile(filepath.Join(dir, "private.key"), []byte(d.PrivateKeyPEM), 0600); err != nil {
		return failed("stage", err)
	}
	mode := os.FileMode(d.Target.FileMode)
	if mode == 0 {
		mode = 0600
	}
	uid, gid := -1, -1
	if d.Target.OwnerUID != nil {
		uid = *d.Target.OwnerUID
	}
	if d.Target.OwnerGID != nil {
		gid = *d.Target.OwnerGID
	}
	for _, name := range []string{"certificate.pem", "private.key"} {
		file := filepath.Join(dir, name)
		if uid >= 0 || gid >= 0 {
			if err = os.Chown(file, uid, gid); err != nil {
				return failed("ownership", err)
			}
		}
		if err = os.Chmod(file, mode); err != nil {
			return failed("permissions", err)
		}
	}
	if uid >= 0 || gid >= 0 {
		// Allow traversal without directory listing or modification. Receipt
		// journals remain private and keys keep their explicit owner and mode.
		for _, path := range []string{root, filepath.Join(root, "releases"), dir} {
			if err = os.Chmod(path, 0711); err != nil {
				return failed("permissions", err)
			}
		}
	}
	if err = phase("validating"); err != nil {
		return j.Receipt, err
	}
	if err = e.run(ctx, d.Target.CheckCommand, dir); err != nil {
		return failed("validation", err)
	}
	if ctx.Err() != nil || !e.allowed() {
		return failed("cancelled", errors.New("cancelled before apply"))
	}
	// A staged bundle needs a fresh Dashboard authorization acknowledgement.
	// Losing the session or revoking the run prevents the atomic commit.
	if permit == nil {
		return failed("authorization", errors.New("apply permission required"))
	}
	if err = permit(ctx); err != nil {
		return failed("authorization", err)
	}
	if err = phase("applying"); err != nil {
		return j.Receipt, err
	}
	rollback := func(cause error) (agentproto.Receipt, error) {
		recovery, cancel := context.WithTimeout(context.Background(), 45*time.Second)
		defer cancel()
		if err := switchCurrent(root, previous); err != nil {
			j.Receipt.Code = "rollback"
			_ = phase("unknown")
			return j.Receipt, cause
		}
		if previous != "" {
			if err := e.run(recovery, d.Target.ReloadCommand, filepath.Join(root, "current")); err != nil {
				j.Receipt.Code = "rollback_reload"
				_ = phase("unknown")
				return j.Receipt, cause
			}
			if err := e.run(recovery, d.Target.HealthCommand, filepath.Join(root, "current")); err != nil {
				j.Receipt.Code = "rollback_health"
				_ = phase("unknown")
				return j.Receipt, cause
			}
		} else {
			// Removing a file cannot unload a certificate already held by a
			// running service. There is no previous bundle to reload or verify.
			j.Receipt.Code = "first_install_no_restore"
			_ = phase("unknown")
			return j.Receipt, cause
		}
		j.Receipt.RolledBack = true
		j.Receipt.Code = "apply"
		_ = phase("rolled_back")
		return j.Receipt, cause
	}
	if err = switchCurrent(root, release); err != nil {
		return rollback(err)
	}
	if err = e.run(ctx, d.Target.ReloadCommand, filepath.Join(root, "current")); err != nil {
		return rollback(err)
	}
	if err = e.verify(ctx, d); err != nil {
		return rollback(err)
	}
	j.Receipt.Verified = true
	j.Receipt.CertificateObserved = d.Target.TLSAddress != ""
	if err = phase("verified"); err != nil {
		return j.Receipt, err
	}
	return j.Receipt, nil
}
func phaseError(phase string) error {
	if phase == "verified" {
		return nil
	}
	return fmt.Errorf("deployment receipt: %s", phase)
}
func (e *Executor) verify(ctx context.Context, d agentproto.Deployment) error {
	if err := e.run(ctx, d.Target.HealthCommand, filepath.Join(d.Target.BundlePath, "current")); err != nil {
		return err
	}
	return e.verifyTLS(ctx, d)
}
func (e *Executor) verifyTLS(ctx context.Context, d agentproto.Deployment) error {
	if d.Target.TLSAddress != "" {
		dialer := tls.Dialer{NetDialer: &net.Dialer{Timeout: 10 * time.Second}, Config: &tls.Config{ServerName: d.Target.TLSServerName, MinVersion: tls.VersionTLS12}}
		conn, err := dialer.DialContext(ctx, "tcp", d.Target.TLSAddress)
		if err != nil {
			return err
		}
		defer conn.Close()
		certs := conn.(*tls.Conn).ConnectionState().PeerCertificates
		if len(certs) == 0 {
			return ErrInvalid
		}
		sum := sha256.Sum256(certs[0].Raw)
		if hex.EncodeToString(sum[:]) != d.Fingerprint {
			return errors.New("served certificate mismatch")
		}
	}
	return nil
}

// Reconcile never repeats scripts. An interrupted transaction needs both a
// matching installed pair and a freshly observed TLS leaf to become verified.
func (e *Executor) Reconcile(ctx context.Context, d agentproto.Deployment) (agentproto.Receipt, error) {
	e.mu.Lock()
	defer e.mu.Unlock()
	r, err := e.inspectLocked(ctx, d)
	if err != nil {
		return r, err
	}
	if r.Phase == "verified" || r.Phase == "rolled_back" || r.Phase == "failed" || r.Phase == "retired" {
		return r, phaseError(r.Phase)
	}
	if d.Target.TLSAddress == "" || !r.CurrentPairValid || r.CurrentOperationID != d.OperationID || r.CurrentFingerprint != d.Fingerprint {
		return r, phaseError("unknown")
	}
	if err = e.verifyTLS(ctx, d); err != nil {
		return r, err
	}
	j, err := loadJournal(d.Target.BundlePath, d.OperationID)
	if err != nil {
		return r, err
	}
	r.Phase = "verified"
	r.Verified = true
	r.CertificateObserved = true
	r.Code = ""
	j.Receipt = r
	return r, writeJournal(d.Target.BundlePath, j)
}
