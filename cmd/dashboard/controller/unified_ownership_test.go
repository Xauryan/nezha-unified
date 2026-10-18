package controller

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/nezhahq/nezha/model"
	"github.com/nezhahq/nezha/service/unified"
)

func TestUnifiedOwnershipHTTPTransferRequiresVersionAndDisablesExecution(t *testing.T) {
	f := newNativeHTTPFixture(t)
	id := f.workflow(t, "transfer-workflow")
	policy := unified.WorkflowPolicy{WorkflowID: id, RunAsUserID: 1, Version: 4, Enabled: true}
	if err := f.db.Create(&policy).Error; err != nil {
		t.Fatal(err)
	}
	body := map[string]any{"kind": "workflow", "resource_id": id, "owner_id": 2, "expected_version": 1}
	data := f.request(t, http.MethodPost, "/ownership/transfer", body, http.StatusOK)
	var changed unified.ResourceBinding
	if err := json.Unmarshal(data, &changed); err != nil {
		t.Fatal(err)
	}
	if changed.OwnerID != 2 || changed.Version != 2 {
		t.Fatalf("unexpected transferred owner/version: %d/%d", changed.OwnerID, changed.Version)
	}
	if err := f.db.First(&policy, "workflow_id = ?", id).Error; err != nil {
		t.Fatal(err)
	}
	if policy.Enabled || policy.Version != 5 {
		t.Fatal("ownership transfer retained previous executable approval")
	}
	f.request(t, http.MethodPost, "/ownership/transfer", body, http.StatusBadRequest)
	f.principal = unified.Principal{UserID: 2}
	data = f.request(t, http.MethodGet, "/ownership", nil, http.StatusOK)
	var bindings []unified.ResourceBinding
	if err := json.Unmarshal(data, &bindings); err != nil {
		t.Fatal(err)
	}
	if len(bindings) != 1 || bindings[0].ResourceID != id {
		t.Fatal("new owner cannot read resource binding")
	}
	f.request(t, http.MethodGet, "/certimate/resources/workflow/"+id, nil, http.StatusOK)
	f.request(t, http.MethodPost, "/ownership/transfer", map[string]any{"kind": "workflow", "resource_id": id, "owner_id": 3, "expected_version": 2}, http.StatusForbidden)
	f.principal = unified.Principal{UserID: 3}
	data = f.request(t, http.MethodGet, "/ownership", nil, http.StatusOK)
	if err := json.Unmarshal(data, &bindings); err != nil {
		t.Fatal(err)
	}
	if len(bindings) != 0 {
		t.Fatal("unrelated user can read resource bindings")
	}
	if err := f.db.Model(&model.User{}).Where("id = ?", 1).Update("role", model.RoleMember).Error; err != nil {
		t.Fatal(err)
	}
	f.principal = unified.Principal{UserID: 1, Admin: true}
	f.request(t, http.MethodPost, "/ownership/transfer", map[string]any{"kind": "workflow", "resource_id": id, "owner_id": 3, "expected_version": 2}, http.StatusForbidden)
}

func TestUnifiedWorkflowPolicyHTTPRequiresExplicitExpectedVersion(t *testing.T) {
	f := newNativeHTTPFixture(t)
	id := f.workflow(t, "policy-workflow")
	path := "/workflow-policies/" + id
	f.request(t, http.MethodPut, path, map[string]any{"run_as_user_id": 1, "enabled": false}, http.StatusBadRequest)
	f.request(t, http.MethodPut, path, map[string]any{"workflow_id": "different", "expected_version": 0, "run_as_user_id": 1, "enabled": false}, http.StatusBadRequest)
	data := f.request(t, http.MethodPut, path, map[string]any{"expected_version": 0, "run_as_user_id": 1, "enabled": false}, http.StatusOK)
	var policy unified.WorkflowPolicy
	if err := json.Unmarshal(data, &policy); err != nil {
		t.Fatal(err)
	}
	if policy.WorkflowID != id || policy.Version != 1 || policy.Enabled {
		t.Fatal("policy identity/version contract mismatch")
	}
	f.request(t, http.MethodPut, path, map[string]any{"expected_version": 0, "run_as_user_id": 1, "enabled": true}, http.StatusBadRequest)
}
