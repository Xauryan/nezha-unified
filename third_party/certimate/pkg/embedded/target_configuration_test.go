package embedded

import (
	"context"
	"encoding/json"
	"errors"
	"strings"
	"testing"
)

func TestTargetConfigurationEnvelopeBindsIdentityVersionAndDatabase(t *testing.T) {
	e := testEngine(t)
	ctx := context.Background()
	config := map[string]any{"path": "/synthetic/cert.pem", "password": "synthetic-target-password"}
	envelope, err := e.SealTargetConfiguration(ctx, "target-1", 7, config)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(envelope, "synthetic-target-password") {
		t.Fatal("target config stored plaintext")
	}
	decoded, err := e.OpenTargetConfiguration(ctx, "target-1", 7, envelope)
	if err != nil || decoded["password"] != config["password"] {
		t.Fatalf("target config did not round trip: %v", err)
	}
	for _, tc := range []struct {
		id      string
		version uint64
	}{{"target-2", 7}, {"target-1", 8}, {"target-1", 0}, {"", 7}} {
		if _, err := e.OpenTargetConfiguration(ctx, tc.id, tc.version, envelope); err == nil {
			t.Fatalf("target config accepted identity/version %q/%d", tc.id, tc.version)
		}
	}
	other := testEngine(t)
	if _, err := other.OpenTargetConfiguration(ctx, "target-1", 7, envelope); err == nil {
		t.Fatal("foreign database decrypted target config")
	}
	if _, err := e.OpenTargetConfiguration(ctx, "target-1", 7, envelope+"corrupt"); err == nil {
		t.Fatal("tampered target config authenticated")
	}
	canceled, cancel := context.WithCancel(ctx)
	cancel()
	if _, err := e.OpenTargetConfiguration(canceled, "target-1", 7, envelope); !errors.Is(err, context.Canceled) {
		t.Fatalf("canceled decryption returned %v", err)
	}
}

func TestNativeSaveCannotSoftDeleteResources(t *testing.T) {
	e := testEngine(t)
	ctx := context.Background()
	credential, err := e.CreateCredential(ctx, "synthetic", CredentialInput{Name: "synthetic", Provider: "synthetic", Config: json.RawMessage(`{"token":"synthetic-token"}`)})
	if err != nil {
		t.Fatal(err)
	}
	certificate, err := e.ImportCertificate(ctx, "synthetic", testCertificate(t))
	if err != nil {
		t.Fatal(err)
	}
	for _, tc := range []struct{ collection, id string }{{"access", credential.ID}, {"certificate", certificate.ID}} {
		if _, err := e.NativeSave(ctx, "synthetic", tc.collection, tc.id, map[string]any{"deleted": "2026-10-18 00:00:00Z"}); !errors.Is(err, ErrInvalid) {
			t.Fatalf("soft-delete PATCH accepted for %s: %v", tc.collection, err)
		}
		if _, err := e.NativeGet(ctx, tc.collection, tc.id); err != nil {
			t.Fatalf("rejected PATCH changed %s: %v", tc.collection, err)
		}
	}
}

func TestCertificateImportRemainsInactiveUntilOwnershipActivation(t *testing.T) {
	e := testEngine(t)
	ctx := context.Background()
	in := testCertificate(t)
	in.OperationID = "synthetic-import-operation"
	r, err := e.ImportCertificate(ctx, "synthetic", in)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := e.CertificatePEM(ctx, r.ID); !errors.Is(err, ErrNotFound) {
		t.Fatal("provisional certificate public material was visible")
	}
	if _, err := e.CertificateMaterial(ctx, r.ID); !errors.Is(err, ErrNotFound) {
		t.Fatal("provisional certificate private material was visible")
	}
	page, err := e.List(ctx, "certificate", nil, true, 1, 20)
	if err != nil || page.Total != 0 {
		t.Fatalf("provisional certificate counted: %d %v", page.Total, err)
	}
	recovery, err := e.FindNativeOperation(ctx, "certificate", in.OperationID)
	if err != nil || recovery["id"] != r.ID {
		t.Fatalf("provisional import could not be recovered: %v", err)
	}
	if _, err := e.ImportCertificate(ctx, "synthetic", in); err == nil {
		t.Fatal("duplicate import operation created a second certificate")
	}
	if err := e.ActivateNative(ctx, "certificate", r.ID, in.OperationID); err != nil {
		t.Fatal(err)
	}
	if _, err := e.CertificateMaterial(ctx, r.ID); err != nil {
		t.Fatal(err)
	}
	page, err = e.List(ctx, "certificate", nil, true, 1, 20)
	if err != nil || page.Total != 1 {
		t.Fatalf("activation count incorrect: %d %v", page.Total, err)
	}
}
