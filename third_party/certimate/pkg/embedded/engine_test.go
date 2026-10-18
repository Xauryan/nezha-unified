package embedded

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"math/big"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/pocketbase/pocketbase/core"
)

func testEngine(t *testing.T) *Engine {
	t.Helper()
	e, err := Open(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = e.Close() })
	return e
}

func TestEmbeddedLifecycleAndNoSecondIdentity(t *testing.T) {
	dir := t.TempDir()
	e, err := Open(dir)
	if err != nil {
		t.Fatal(err)
	}
	for kind := range collections {
		page, err := e.List(context.Background(), kind, nil, true, 1, 20)
		if err != nil || page.Total != 0 || len(page.Items) != 0 {
			t.Fatalf("fresh %s inventory: %#v %v", kind, page, err)
		}
	}
	users, err := e.app.FindAllRecords(core.CollectionNameSuperusers)
	if err != nil || len(users) != 0 {
		t.Fatalf("unexpected superusers: %d, %v", len(users), err)
	}
	if err := e.Close(); err != nil {
		t.Fatal(err)
	}
	if err := e.Close(); err != nil {
		t.Fatal(err)
	}
	if _, err := e.List(context.Background(), "credential", nil, true, 1, 20); err == nil {
		t.Fatal("closed engine usable")
	}
	e, err = Open(dir)
	if err != nil {
		t.Fatal(err)
	}
	_ = e.Close()
	if err := os.Remove(filepath.Join(dir, "unified.key")); err != nil {
		t.Fatal(err)
	}
	if _, err := Open(dir); err == nil {
		t.Fatal("missing recovery key accepted")
	}
}

func TestRecoveryRejectsWrongKeyAndPreservesEncryptedData(t *testing.T) {
	dir := t.TempDir()
	e, err := Open(dir)
	if err != nil {
		t.Fatal(err)
	}
	resource, err := e.CreateCredential(context.Background(), "1", CredentialInput{Name: "recovery", Provider: "cloudflare", Config: json.RawMessage(`{"apiToken":"synthetic"}`)})
	if err != nil {
		t.Fatal(err)
	}
	keyPath := filepath.Join(dir, "unified.key")
	key, err := os.ReadFile(keyPath)
	if err != nil {
		t.Fatal(err)
	}
	if err := e.Close(); err != nil {
		t.Fatal(err)
	}
	wrong := make([]byte, 32)
	if _, err := rand.Read(wrong); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(keyPath, wrong, 0600); err != nil {
		t.Fatal(err)
	}
	if reopened, err := Open(dir); err == nil {
		_ = reopened.Close()
		t.Fatal("wrong recovery key accepted")
	}
	if err := os.WriteFile(keyPath, key, 0600); err != nil {
		t.Fatal(err)
	}
	e, err = Open(dir)
	if err != nil {
		t.Fatal(err)
	}
	defer e.Close()
	page, err := e.List(context.Background(), "credential", nil, true, 1, 20)
	if err != nil || page.Total != 1 || page.Items[0].ID != resource.ID {
		t.Fatal("restored data did not survive reopen")
	}
}

func TestCredentialEncryptionAndAuthorizedPagination(t *testing.T) {
	e := testEngine(t)
	ctx := context.Background()
	secret := "synthetic-secret-never-return"
	var ids []string
	for _, name := range []string{"one", "two", "three"} {
		r, err := e.CreateCredential(ctx, "1", CredentialInput{Name: name, Provider: "cloudflare", Config: json.RawMessage(`{"apiToken":"` + secret + `"}`)})
		if err != nil {
			t.Fatal(err)
		}
		ids = append(ids, r.ID)
	}
	page, err := e.List(ctx, "credential", ids[1:], false, 2, 1)
	if err != nil || page.Total != 2 || len(page.Items) != 1 {
		t.Fatalf("pagination leaks/loses resources: %#v %v", page, err)
	}
	data, _ := json.Marshal(page)
	if strings.Contains(string(data), secret) || strings.Contains(string(data), "envelope") {
		t.Fatal("metadata leaked secret")
	}
	r, err := e.app.FindRecordById("access", ids[0])
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(r.GetString("config"), secret) {
		t.Fatal("plaintext stored in config")
	}
	envelope, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(r.GetString("unified_envelope"), "v1:"))
	if err != nil {
		t.Fatal(err)
	}
	n := e.aead.NonceSize()
	plain, err := e.aead.Open(nil, envelope[:n], envelope[n:], []byte("credential:"+ids[0]+":1"))
	if err != nil || !strings.Contains(string(plain), secret) {
		t.Fatalf("cannot recover: %v", err)
	}
	if _, err := e.aead.Open(nil, envelope[:n], envelope[n:], []byte("credential:"+ids[1]+":1")); err == nil {
		t.Fatal("envelope can be transplanted")
	}
	page, err = e.List(ctx, "credential", nil, false, 1, 20)
	if err != nil || page.Total != 0 || len(page.Items) != 0 {
		t.Fatal("empty grant list leaked")
	}
	if _, err = e.List(ctx, "credential", nil, true, 1, 101); err == nil {
		t.Fatal("unbounded page")
	}
}

func testCertificate(t *testing.T) CertificateInput {
	t.Helper()
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	template := &x509.Certificate{SerialNumber: big.NewInt(1), Subject: pkix.Name{CommonName: "display-label.example"}, DNSNames: []string{"real.example.com", "www.real.example.com"}, NotBefore: time.Now().Add(-time.Hour), NotAfter: time.Now().Add(24 * time.Hour), KeyUsage: x509.KeyUsageDigitalSignature}
	der, err := x509.CreateCertificate(rand.Reader, template, template, &key.PublicKey, key)
	if err != nil {
		t.Fatal(err)
	}
	k, err := x509.MarshalPKCS8PrivateKey(key)
	if err != nil {
		t.Fatal(err)
	}
	return CertificateInput{Certificate: string(pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: der})), PrivateKey: string(pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: k}))}
}

func TestCertificateImportUsesSANAndProtectsKey(t *testing.T) {
	e := testEngine(t)
	ctx := context.Background()
	in := testCertificate(t)
	publicPEM := in.Certificate
	in.Certificate += in.PrivateKey // A common combined PEM bundle must not leak.
	r, err := e.ImportCertificate(ctx, "1", in)
	if err != nil {
		t.Fatal(err)
	}
	if r.Names != "real.example.com;www.real.example.com" {
		t.Fatalf("labels treated as SANs: %#v", r)
	}
	stored, err := e.app.FindRecordById("certificate", r.ID)
	if err != nil {
		t.Fatal(err)
	}
	if stored.GetString("privateKey") != "" || !strings.HasPrefix(stored.GetString("unified_private_key"), "v1:") {
		t.Fatal("unencrypted private key")
	}
	pub, err := e.CertificatePEM(ctx, r.ID)
	if err != nil || pub != publicPEM || strings.Contains(pub, "PRIVATE KEY") {
		t.Fatal("incorrect certificate download")
	}
	wrong := testCertificate(t)
	in.PrivateKey = wrong.PrivateKey
	if _, err = e.ImportCertificate(ctx, "1", in); err == nil {
		t.Fatal("mismatched key accepted")
	}
	logs, err := e.app.FindAllRecords("unified_audit")
	if err != nil || len(logs) != 1 {
		t.Fatal("audit not committed with import")
	}
}
