package embedcrypto

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
	"testing"

	"github.com/pocketbase/pocketbase/core"
	_ "github.com/pocketbase/pocketbase/migrations"
)

func testApp(t *testing.T) core.App {
	t.Helper()
	app := core.NewBaseApp(core.BaseAppConfig{DataDir: t.TempDir()})
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		t.Fatal(err)
	}
	block, _ := aes.NewCipher(key)
	aead, _ := cipher.NewGCM(block)
	Configure(app, &Hooks{
		Seal: func(aad string, value []byte) string {
			nonce := make([]byte, aead.NonceSize())
			if _, err := rand.Read(nonce); err != nil {
				t.Fatal(err)
			}
			return "v1:" + base64.StdEncoding.EncodeToString(aead.Seal(nonce, nonce, value, []byte(aad)))
		},
		Open: func(aad, value string) ([]byte, error) {
			blob, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(value, "v1:"))
			if err != nil || len(blob) < aead.NonceSize()+aead.Overhead() {
				return nil, errors.New("invalid test envelope")
			}
			return aead.Open(nil, blob[:aead.NonceSize()], blob[aead.NonceSize():], []byte(aad))
		},
	})
	return app
}

func TestProtectedJSONBindsDatabaseRecordAndField(t *testing.T) {
	app := testApp(t)
	collection := core.NewBaseCollection("workflow")
	collection.Fields.Add(&core.JSONField{Name: "graphDraft"}, &core.JSONField{Name: "graphContent"})
	record := core.NewRecord(collection)
	const secret = "synthetic-inline-private-key"
	value := map[string]any{"privateKey": secret, "nodes": []any{map[string]any{"certificateId": "synthetic-ref"}}}
	if err := SetJSON(app, record, "graphDraft", value); err != nil {
		t.Fatal(err)
	}
	if len(record.Id) != 15 {
		t.Fatal("ciphertext not bound to a stable record identity")
	}
	raw, _ := json.Marshal(record.Get("graphDraft"))
	if strings.Contains(string(raw), secret) || strings.Contains(string(raw), "synthetic-ref") {
		t.Fatal("plaintext remained in persisted graph")
	}
	var recovered map[string]any
	if err := ReadJSON(app, record, "graphDraft", &recovered); err != nil || recovered["privateKey"] != secret {
		t.Fatalf("graph roundtrip failed: %v", err)
	}
	record.Set("graphContent", record.Get("graphDraft"))
	if err := ReadJSON(app, record, "graphContent", &recovered); err == nil {
		t.Fatal("cross-field ciphertext substitution accepted")
	}
	other := core.NewRecord(collection)
	_ = EnsureID(other)
	other.Set("graphDraft", record.Get("graphDraft"))
	if err := ReadJSON(app, other, "graphDraft", &recovered); err == nil {
		t.Fatal("cross-record ciphertext substitution accepted")
	}
	if err := ReadJSON(testApp(t), record, "graphDraft", &recovered); err == nil {
		t.Fatal("cross-database ciphertext substitution accepted")
	}
}

func TestEmbeddedMissingHooksFailsClosed(t *testing.T) {
	app := core.NewBaseApp(core.BaseAppConfig{DataDir: t.TempDir()})
	app.Store().Set("nezha_embedded", true)
	if _, err := SealJSON(app, "synthetic", map[string]any{"token": "synthetic-secret"}); err == nil {
		t.Fatal("embedded write accepted without encryption")
	}
	if _, err := ResolveAccess(app, context.Background(), "synthetic"); err == nil {
		t.Fatal("credential read accepted without authorization")
	}
	var value any
	if err := OpenJSON(app, "synthetic", map[string]any{"token": "synthetic-secret"}, &value); err == nil {
		t.Fatal("embedded plaintext read accepted")
	}
}

func TestSanitizeProviderDetailsAndStandaloneCompatibility(t *testing.T) {
	app := testApp(t)
	const raw = "request failed with synthetic-secret-token"
	message, data := SafeLog(app, raw, map[string]any{"error": raw})
	if strings.Contains(message, "synthetic-secret") || len(data) != 0 || strings.Contains(SafeError(app, raw), "synthetic-secret") {
		t.Fatal("provider details persisted in logs")
	}
	if SafeError(app, "") != "" {
		t.Fatal("successful runs received an error")
	}
	standalone := core.NewBaseApp(core.BaseAppConfig{DataDir: t.TempDir()})
	input := map[string]any{"token": "synthetic-token"}
	stored, err := SealJSON(standalone, "test", input)
	if err != nil {
		t.Fatal(err)
	}
	var output map[string]any
	if err := OpenJSON(standalone, "test", stored, &output); err != nil || output["token"] != input["token"] {
		t.Fatalf("standalone compatibility failed: %v", err)
	}
}

func TestMigrateProtectsLegacyRowsAndIsIdempotent(t *testing.T) {
	app := testApp(t)
	if err := app.Bootstrap(); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = app.ResetBootstrapState() })
	const secret = "synthetic-legacy-secret"
	for name, fields := range map[string][]string{
		"workflow": {"graphDraft", "graphContent"}, "workflow_run": {"graph"},
		"workflow_output": {"nodeConfig", "outputs"}, "settings": {"content"},
		"access": {"config"}, "certificate": {}, "acme_accounts": {}, "workflow_logs": {"data"},
	} {
		collection := core.NewBaseCollection(name)
		for _, field := range fields {
			collection.Fields.Add(&core.JSONField{Name: field})
		}
		for _, field := range []string{"privateKey", "unified_private_key", "unified_envelope", "error", "message"} {
			collection.Fields.Add(&core.TextField{Name: field})
		}
		if err := app.Save(collection); err != nil {
			t.Fatal(err)
		}
		record := core.NewRecord(collection)
		for _, field := range fields {
			record.Set(field, map[string]any{"secret": secret})
		}
		if name == "certificate" || name == "acme_accounts" {
			record.Set("privateKey", secret)
		}
		if name == "workflow_run" {
			record.Set("error", secret)
		}
		if name == "workflow_logs" {
			record.Set("message", secret)
		}
		if err := app.Save(record); err != nil {
			t.Fatal(err)
		}
	}
	// Transactional native writes must inherit the database's codec.
	if err := app.RunInTransaction(func(tx core.App) error {
		records, err := tx.FindAllRecords("settings")
		if err != nil {
			return err
		}
		if err := SetJSON(tx, records[0], "content", map[string]any{"secret": secret}); err != nil {
			return err
		}
		return tx.Save(records[0])
	}); err != nil {
		t.Fatal(err)
	}
	for attempt := 0; attempt < 2; attempt++ {
		if err := Migrate(app); err != nil {
			t.Fatal(err)
		}
		for _, name := range []string{"workflow", "workflow_run", "workflow_output", "settings", "access", "certificate", "acme_accounts", "workflow_logs"} {
			records, err := app.FindAllRecords(name)
			if err != nil || len(records) != 1 {
				t.Fatalf("migration lost records in %s: %v", name, err)
			}
			raw, err := json.Marshal(records[0])
			if err != nil || strings.Contains(string(raw), secret) {
				t.Fatalf("legacy plaintext remains in %s", name)
			}
			if name == "settings" {
				var value map[string]string
				if err := ReadJSON(app, records[0], "content", &value); err != nil || value["secret"] != secret {
					t.Fatalf("legacy config changed: %v", err)
				}
			}
			if name == "acme_accounts" {
				value, err := Open(app, "acme_account:"+records[0].Id, records[0].GetString("privateKey"))
				if err != nil || string(value) != secret {
					t.Fatalf("legacy ACME key changed: %v", err)
				}
			}
		}
	}
}
