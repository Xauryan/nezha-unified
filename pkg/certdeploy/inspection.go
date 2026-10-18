package certdeploy

import (
	"context"
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"github.com/nezhahq/nezha/pkg/agentproto"
	"os"
	"path/filepath"
	"strings"
)

// Inspect reads one exact operation and the current certificate/key pair. It
// never creates files, repeats scripts, applies a bundle, reloads or uses a CA.
func (e *Executor) Inspect(ctx context.Context, d agentproto.Deployment) (agentproto.Receipt, error) {
	e.mu.Lock()
	defer e.mu.Unlock()
	return e.inspectLocked(ctx, d)
}
func (e *Executor) inspectLocked(ctx context.Context, d agentproto.Deployment) (agentproto.Receipt, error) {
	r := agentproto.Receipt{OperationID: d.OperationID, TargetID: d.Target.ID, TargetVersion: d.Target.Version, Sequence: d.Sequence, Fingerprint: d.Fingerprint}
	if ctx.Err() != nil {
		return r, ctx.Err()
	}
	if !e.allowed() || !safeID.MatchString(d.OperationID) || !safeID.MatchString(d.Target.ID) || !filepath.IsAbs(d.Target.BundlePath) || filepath.Clean(d.Target.BundlePath) == string(filepath.Separator) {
		return r, ErrInvalid
	}
	for path := filepath.Clean(d.Target.BundlePath); ; path = filepath.Dir(path) {
		info, err := os.Lstat(path)
		if err == nil && info.Mode()&os.ModeSymlink != 0 {
			return r, ErrInvalid
		}
		if err != nil && !os.IsNotExist(err) {
			return r, err
		}
		if path == filepath.Dir(path) {
			break
		}
	}
	j, err := loadJournal(d.Target.BundlePath, d.OperationID)
	if os.IsNotExist(err) {
		r.Inspected = true
		r.Phase = "not_found"
		return r, nil
	}
	if err != nil {
		return r, err
	}
	if j.Receipt.TargetID != d.Target.ID || j.Receipt.TargetVersion != d.Target.Version || j.Receipt.Sequence != d.Sequence || j.Receipt.Fingerprint != d.Fingerprint {
		return r, ErrInvalid
	}
	r = j.Receipt
	r.Inspected = true
	link, err := current(d.Target.BundlePath)
	if err != nil {
		return r, err
	}
	r.CurrentOperationID = strings.TrimPrefix(link, "releases/")
	if link != "" {
		certificate, e1 := os.ReadFile(filepath.Join(d.Target.BundlePath, link, "certificate.pem"))
		key, e2 := os.ReadFile(filepath.Join(d.Target.BundlePath, link, "private.key"))
		if e1 == nil && e2 == nil {
			pair, e3 := tls.X509KeyPair(certificate, key)
			if e3 == nil && len(pair.Certificate) > 0 {
				leaf, e4 := x509.ParseCertificate(pair.Certificate[0])
				if e4 == nil {
					sum := sha256.Sum256(leaf.Raw)
					r.CurrentFingerprint = hex.EncodeToString(sum[:])
					r.CurrentPairValid = true
				}
			}
		}
	}
	if r.Phase == "verified" && r.CurrentOperationID == r.OperationID && (!r.CurrentPairValid || r.CurrentFingerprint != r.Fingerprint) {
		r.Phase = "unknown"
		r.Verified = false
		r.CertificateObserved = false
		r.Code = "installed_pair"
	}
	if (r.Phase == "staging" || r.Phase == "validating") && r.CurrentOperationID != r.OperationID {
		r.Phase = "not_applied"
		r.Verified = false
	}
	return r, nil
}

// Retire records the operator's explicit acceptance of an uncertain outcome.
// It is not a rollback or success assertion and never changes service files.
func (e *Executor) Retire(ctx context.Context, d agentproto.Deployment) (agentproto.Receipt, error) {
	e.mu.Lock()
	defer e.mu.Unlock()
	r, err := e.inspectLocked(ctx, d)
	if err != nil {
		return r, err
	}
	if r.Phase == "verified" || r.Phase == "rolled_back" || r.Phase == "failed" || r.Phase == "retired" {
		return r, nil
	}
	var j journal
	if r.Phase == "not_found" {
		if err := safeRoot(d.Target.BundlePath); err != nil {
			return r, err
		}
		j.Receipt = r
	} else {
		j, err = loadJournal(d.Target.BundlePath, d.OperationID)
		if err != nil {
			return r, err
		}
	}
	r.Phase = "retired"
	r.Code = "operator_retired_unknown"
	r.Verified = false
	r.RolledBack = false
	r.CertificateObserved = false
	j.Receipt = r
	return r, writeJournal(d.Target.BundlePath, j)
}
