package embedded

import (
	"context"
	"encoding/json"
	"testing"
)

func TestRecoveryValidationAuthenticatesEveryCredential(t *testing.T) {
	e, err := Open(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	defer e.Close()
	for i := 0; i < 2; i++ {
		_, err := e.CreateCredential(context.Background(), "synthetic", CredentialInput{Name: "synthetic recovery test", Provider: "synthetic", Config: json.RawMessage(`{"secret":"synthetic-only"}`)})
		if err != nil {
			t.Fatal(err)
		}
	}
	records, err := e.app.FindRecordsByFilter("access", "unified_envelope != ''", "id", 0, 0)
	if err != nil || len(records) != 2 {
		t.Fatalf("records: %d, %v", len(records), err)
	}
	// Swapping an envelope into a later record must fail even though the first
	// record and the recovery key are valid.
	records[1].Set("unified_envelope", records[0].GetString("unified_envelope"))
	if err := e.app.Save(records[1]); err != nil {
		t.Fatal(err)
	}
	if err := e.verifyRecoveryKey(); err == nil {
		t.Fatal("recovery validation skipped a later credential")
	}
}
