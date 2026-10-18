package embedded

import (
	"context"
	"encoding/json"
	"errors"
	"strings"
	"testing"

	"github.com/certimate-go/certimate/internal/embedcrypto"
	"github.com/pocketbase/pocketbase/core"
)

func TestNativeCredentialsActivationProjectionAndRotation(t *testing.T) {
	e := testEngine(t)
	ctx := context.Background()
	in := map[string]any{"name": "native fixture", "provider": "cloudflare", "config": map[string]any{"token": "synthetic-native-secret"}, "unified_operation_id": "synthetic-operation-1"}
	created, err := e.NativeSave(ctx, "actor:1", "access", "", in)
	if err != nil {
		t.Fatal(err)
	}
	id := created["id"].(string)
	if _, exists := created["config"]; exists {
		t.Fatal("credential metadata includes config")
	}
	if _, err := e.NativeGet(ctx, "access", id); !errors.Is(err, ErrNotFound) {
		t.Fatalf("inactive record visible: %v", err)
	}
	page, err := e.NativeList(ctx, "access", NativeScope{All: true}, NativeOptions{})
	if err != nil || page.TotalItems != 0 {
		t.Fatalf("inactive record counted: %d %v", page.TotalItems, err)
	}
	found, err := e.FindNativeOperation(ctx, "access", "synthetic-operation-1")
	if err != nil || found["id"] != id || found["unified_active"] != false {
		t.Fatalf("operation recovery failed: %v", err)
	}
	if err := e.ActivateNative(ctx, "access", id, "wrong-operation"); !errors.Is(err, ErrInvalid) {
		t.Fatalf("wrong activation identity accepted: %v", err)
	}
	if err := e.ActivateNative(ctx, "access", id, "synthetic-operation-1"); err != nil {
		t.Fatal(err)
	}
	if _, err := e.NativeSave(ctx, "actor:1", "access", id, map[string]any{"name": "renamed"}); err != nil {
		t.Fatal(err)
	}
	r, err := e.app.FindRecordById("access", id)
	if err != nil || r.GetInt("unified_version") != 1 || r.GetString("config") == "synthetic-native-secret" {
		t.Fatalf("metadata changed credential: %v", err)
	}
	e.SetCredentialAuthorizer(func(context.Context, CredentialUse) error { return nil })
	use := CredentialUse{ActorID: 1, CredentialID: id, Version: 1, Purpose: "test", ResourceKind: "credential", ResourceID: id}
	config, err := e.ResolveCredential(ctx, use)
	if err != nil || config["token"] != "synthetic-native-secret" {
		t.Fatalf("metadata update lost secret: %v", err)
	}
	if _, err := e.NativeSave(ctx, "actor:1", "access", id, map[string]any{"config": map[string]any{"token": "replacement"}}); err == nil {
		t.Fatal("rotation without version accepted")
	}
	if _, err := e.NativeSave(ctx, "actor:1", "access", id, map[string]any{"unified_version": 1, "config": map[string]any{"token": "replacement"}}); err != nil {
		t.Fatal(err)
	}
	if _, err := e.NativeSave(ctx, "actor:1", "access", id, map[string]any{"unified_version": 1, "config": map[string]any{"token": "stale"}}); err == nil {
		t.Fatal("stale rotation accepted")
	}
	if _, err := e.ResolveCredential(ctx, use); err == nil {
		t.Fatal("old credential version usable")
	}
	use.Version = 2
	config, err = e.ResolveCredential(ctx, use)
	if err != nil || config["token"] != "replacement" {
		t.Fatalf("rotation not retained: %v", err)
	}
	page, err = e.NativeList(ctx, "access", NativeScope{IDs: []string{id}}, NativeOptions{Page: 1, PerPage: 1})
	if err != nil || page.TotalItems != 1 || len(page.Items) != 1 {
		t.Fatalf("authorized count incorrect: %d %v", page.TotalItems, err)
	}
	page, err = e.NativeList(ctx, "access", NativeScope{}, NativeOptions{Page: 1, PerPage: 1})
	if err != nil || page.TotalItems != 0 {
		t.Fatalf("empty scope leaked count: %d %v", page.TotalItems, err)
	}
	for _, filter := range []string{"config.token='replacement'", "unified_envelope!=''", "@request.auth.id!=''"} {
		if _, err := e.NativeList(ctx, "access", NativeScope{All: true}, NativeOptions{Filter: filter}); !errors.Is(err, ErrInvalid) {
			t.Fatalf("secret filter accepted: %v", err)
		}
	}
	if _, err := e.NativeCloneCredential(ctx, use, "actor:1", "clone", "clone-operation"); !errors.Is(err, ErrInvalid) {
		t.Fatal("non-clone purpose accepted")
	}
	use.Purpose = "credential.clone"
	e.SetCredentialAuthorizer(func(context.Context, CredentialUse) error { return errors.New("denied") })
	if _, err := e.NativeCloneCredential(ctx, use, "actor:1", "clone", "clone-operation"); err == nil {
		t.Fatal("unauthorized clone accepted")
	}
	e.SetCredentialAuthorizer(func(_ context.Context, requested CredentialUse) error {
		if requested.Purpose != "credential.clone" || requested.CredentialID != id {
			return errors.New("denied")
		}
		return nil
	})
	clone, err := e.NativeCloneCredential(ctx, use, "actor:1", "clone", "clone-operation")
	if err != nil {
		t.Fatal(err)
	}
	cloneID := clone["id"].(string)
	if _, err := e.NativeGet(ctx, "access", cloneID); !errors.Is(err, ErrNotFound) {
		t.Fatal("clone active before ownership binding")
	}
	if _, exists := clone["config"]; exists {
		t.Fatal("clone response exposed material")
	}
	if err := e.ActivateNative(ctx, "access", cloneID, "clone-operation"); err != nil {
		t.Fatal(err)
	}
	clonedRecord, err := e.app.FindRecordById("access", cloneID)
	if err != nil {
		t.Fatal(err)
	}
	plain, err := e.unseal("credential:"+cloneID+":1", clonedRecord.GetString("unified_envelope"))
	if err != nil || !strings.Contains(string(plain), "replacement") {
		t.Fatalf("clone did not copy active material: %v", err)
	}
}

func nativeFixtureGraph() map[string]any {
	return map[string]any{"nodes": []any{
		map[string]any{"id": "start", "type": "start", "data": map[string]any{}},
		map[string]any{"id": "upload", "type": "bizUpload", "data": map[string]any{"config": map[string]any{"privateKey": "synthetic-workflow-key", "certificate": "public-fixture"}}},
		map[string]any{"id": "end", "type": "end", "data": map[string]any{}},
	}}
}

func TestNativeWorkflowEncryptedRoundtripAndInvalidPublication(t *testing.T) {
	e := testEngine(t)
	ctx := context.Background()
	created, err := e.NativeSave(ctx, "actor:1", "workflow", "", map[string]any{"name": "native workflow", "trigger": "manual", "graphContent": nativeFixtureGraph(), "hasContent": true})
	if err != nil {
		t.Fatal(err)
	}
	id := created["id"].(string)
	data, _ := json.Marshal(created)
	if strings.Contains(string(data), "synthetic-workflow-key") {
		t.Fatal("workflow metadata exposed inline key")
	}
	r, err := e.app.FindRecordById("workflow", id)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(r.GetString("graphContent"), "synthetic-workflow-key") {
		t.Fatal("workflow key persisted plaintext")
	}
	if _, err := e.NativeSave(ctx, "actor:1", "workflow", id, map[string]any{"graphContent": created["graphContent"]}); err != nil {
		t.Fatal(err)
	}
	r, _ = e.app.FindRecordById("workflow", id)
	var graph map[string]any
	if err := embedcrypto.ReadJSON(e.app, r, "graphContent", &graph); err != nil {
		t.Fatal(err)
	}
	encoded, _ := json.Marshal(graph)
	if !strings.Contains(string(encoded), "synthetic-workflow-key") {
		t.Fatal("redacted edit lost key")
	}
	for _, bad := range []any{map[string]any{"nodes": []any{nil, nil}}, map[string]any{"nodes": []any{map[string]any{"id": "dup", "type": "start"}, map[string]any{"id": "dup", "type": "end"}}}} {
		if _, err := e.NativeSave(ctx, "actor:1", "workflow", id, map[string]any{"name": "must rollback", "graphContent": bad}); !errors.Is(err, ErrInvalid) {
			t.Fatalf("invalid graph accepted: %v", err)
		}
	}
	if _, err := e.NativeSave(ctx, "actor:1", "workflow", id, map[string]any{"name": "must rollback", "enabled": true, "trigger": "scheduled", "triggerCron": "invalid"}); !errors.Is(err, ErrInvalid) {
		t.Fatalf("invalid schedule accepted: %v", err)
	}
	r, _ = e.app.FindRecordById("workflow", id)
	if r.GetString("name") != "native workflow" {
		t.Fatal("invalid publication partially committed")
	}
}

func TestNativeWorkflowChildScopeBeforePagination(t *testing.T) {
	e := testEngine(t)
	ctx := context.Background()
	workflows := make([]string, 2)
	for i := range workflows {
		workflow, err := e.NativeSave(ctx, "actor:1", "workflow", "", map[string]any{"name": "fixture", "trigger": "manual"})
		if err != nil {
			t.Fatal(err)
		}
		workflows[i] = workflow["id"].(string)
		c, err := e.app.FindCollectionByNameOrId("workflow_run")
		if err != nil {
			t.Fatal(err)
		}
		r := core.NewRecord(c)
		r.Set("workflowRef", workflows[i])
		r.Set("status", "succeeded")
		r.Set("trigger", "manual")
		if err := embedcrypto.SetJSON(e.app, r, "graph", nil); err != nil {
			t.Fatal(err)
		}
		if err := e.app.Save(r); err != nil {
			t.Fatal(err)
		}
	}
	page, err := e.NativeList(ctx, "workflow_run", NativeScope{WorkflowIDs: []string{workflows[1]}}, NativeOptions{Page: 1, PerPage: 1})
	if err != nil || page.TotalItems != 1 || len(page.Items) != 1 {
		t.Fatalf("child scope count: %d %v", page.TotalItems, err)
	}
	id := page.Items[0]["id"].(string)
	kind, resource, err := e.NativeResource(ctx, "workflow_run", id)
	if err != nil || kind != "workflow" || resource != workflows[1] {
		t.Fatalf("wrong child root: %q %q %v", kind, resource, err)
	}
	if _, err := e.NativeGet(ctx, "_superusers", id); !errors.Is(err, ErrInvalid) {
		t.Fatal("internal collection accessible")
	}
}

func TestNativeRunDeletePreservesCertificatesAndRejectsActiveRuns(t *testing.T) {
	e := testEngine(t)
	ctx := context.Background()
	workflow, err := e.NativeSave(ctx, "actor:1", "workflow", "", map[string]any{"name": "run history fixture", "trigger": "manual"})
	if err != nil {
		t.Fatal(err)
	}
	wid := workflow["id"].(string)
	c, err := e.app.FindCollectionByNameOrId("workflow_run")
	if err != nil {
		t.Fatal(err)
	}
	run := core.NewRecord(c)
	run.Set("workflowRef", wid)
	run.Set("status", "processing")
	run.Set("trigger", "manual")
	if err := embedcrypto.SetJSON(e.app, run, "graph", nil); err != nil {
		t.Fatal(err)
	}
	if err := e.app.Save(run); err != nil {
		t.Fatal(err)
	}
	if err := e.NativeDelete(ctx, "actor:1", "workflow_run", run.Id); err == nil {
		t.Fatal("active run removed")
	}
	if err := e.NativeDelete(ctx, "actor:1", "workflow", wid); err == nil {
		t.Fatal("workflow with active run removed")
	}
	for _, collection := range []string{"workflow_output", "workflow_logs"} {
		c, err := e.app.FindCollectionByNameOrId(collection)
		if err != nil {
			t.Fatal(err)
		}
		r := core.NewRecord(c)
		r.Set("runRef", run.Id)
		r.Set("workflowRef", wid)
		r.Set("nodeId", "node")
		r.Set("timestamp", int64(1000))
		r.Set("message", "safe test event")
		if collection == "workflow_output" {
			if err := embedcrypto.SetJSON(e.app, r, "nodeConfig", map[string]any{}); err != nil {
				t.Fatal(err)
			}
			if err := embedcrypto.SetJSON(e.app, r, "outputs", []any{map[string]any{"name": "keyContent", "value": "synthetic-output-key", "type": "ref", "valueType": "string"}}); err != nil {
				t.Fatal(err)
			}
		}
		if err := e.app.Save(r); err != nil {
			t.Fatal(err)
		}
	}
	page, err := e.NativeList(ctx, "workflow_output", NativeScope{WorkflowIDs: []string{wid}}, NativeOptions{})
	if err != nil {
		t.Fatal(err)
	}
	outputJSON, _ := json.Marshal(page)
	if strings.Contains(string(outputJSON), "synthetic-output-key") {
		t.Fatal("named secret output leaked")
	}
	cert, err := e.ImportCertificate(ctx, "actor:1", testCertificate(t))
	if err != nil {
		t.Fatal(err)
	}
	certificate, err := e.app.FindRecordById("certificate", cert.ID)
	if err != nil {
		t.Fatal(err)
	}
	certificate.Set("workflowRef", wid)
	certificate.Set("workflowRunRef", run.Id)
	if err := e.app.Save(certificate); err != nil {
		t.Fatal(err)
	}
	run.Set("status", "succeeded")
	if err := e.app.Save(run); err != nil {
		t.Fatal(err)
	}
	wr, _ := e.app.FindRecordById("workflow", wid)
	wr.Set("lastRunRef", run.Id)
	wr.Set("lastRunStatus", "succeeded")
	if err := e.app.Save(wr); err != nil {
		t.Fatal(err)
	}
	if err := e.NativeDelete(ctx, "actor:1", "workflow_run", run.Id); err != nil {
		t.Fatal(err)
	}
	for _, collection := range []string{"workflow_output", "workflow_logs"} {
		records, err := e.app.FindAllRecords(collection)
		if err != nil || len(records) != 0 {
			t.Fatalf("run children remain in %s: %v", collection, err)
		}
	}
	wr, err = e.app.FindRecordById("workflow", wid)
	if err != nil || wr.GetString("lastRunRef") != "" || wr.GetString("lastRunStatus") != "" {
		t.Fatalf("workflow last run not cleared: %v", err)
	}
	if _, err := e.NativeGet(ctx, "certificate", cert.ID); err != nil {
		t.Fatalf("history removal deleted certificate: %v", err)
	}
}
