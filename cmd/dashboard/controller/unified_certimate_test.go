package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/certimate-go/certimate/pkg/embedded"
	"github.com/gin-gonic/gin"
	"github.com/nezhahq/nezha/model"
	"github.com/nezhahq/nezha/service/singleton"
	"github.com/nezhahq/nezha/service/unified"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type nativeHTTPFixture struct {
	service   *unified.Service
	engine    *embedded.Engine
	db        *gorm.DB
	router    *gin.Engine
	principal unified.Principal
}

func newNativeHTTPFixture(t *testing.T) *nativeHTTPFixture {
	t.Helper()
	dir := t.TempDir()
	engine, err := embedded.Open(filepath.Join(dir, "certimate"))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = engine.Close() })
	db, err := gorm.Open(sqlite.Open(filepath.Join(dir, "nezha.db")), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = sqlDB.Close() })
	if err = db.AutoMigrate(&model.User{}); err != nil {
		t.Fatal(err)
	}
	for _, user := range []model.User{
		{Common: model.Common{ID: 1}, Username: "synthetic-native-admin", Role: model.RoleAdmin},
		{Common: model.Common{ID: 2}, Username: "synthetic-native-member", Role: model.RoleMember},
		{Common: model.Common{ID: 3}, Username: "synthetic-native-other", Role: model.RoleMember},
	} {
		if err = db.Create(&user).Error; err != nil {
			t.Fatal(err)
		}
	}
	s, err := unified.New(db, engine)
	if err != nil {
		t.Fatal(err)
	}
	prior := singleton.Unified
	singleton.Unified = s
	t.Cleanup(func() { singleton.Unified = prior })
	f := &nativeHTTPFixture{service: s, engine: engine, db: db, principal: unified.Principal{UserID: 1, Admin: true}}
	r := gin.New()
	r.Use(func(c *gin.Context) {
		c.Header("Cache-Control", "no-store")
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, 8<<20)
		if f.principal.UserID != 0 {
			role := model.RoleMember
			if f.principal.Admin {
				role = model.RoleAdmin
			}
			c.Set(model.CtxKeyAuthorizedUser, &model.User{Common: model.Common{ID: f.principal.UserID}, Role: role})
			if f.principal.PAT {
				c.Set(model.CtxKeyAPIToken, true)
			}
			if f.principal.TokenID != 0 || f.principal.ServiceID != 0 {
				c.Set(unifiedScopedPrincipalCtxKey, f.principal)
			}
		}
	})
	group := r.Group("/api/v1/unified")
	bindUnifiedCertimateRoutes(group)
	bindUnifiedOwnershipRoutes(group)
	f.router = r
	return f
}

func (f *nativeHTTPFixture) request(t *testing.T, method, path string, body any, want int, operationID ...string) json.RawMessage {
	t.Helper()
	var encoded []byte
	if body != nil {
		var err error
		encoded, err = json.Marshal(body)
		if err != nil {
			t.Fatal(err)
		}
	}
	req := httptest.NewRequest(method, "/api/v1/unified"+path, bytes.NewReader(encoded))
	req.Header.Set("Content-Type", "application/json")
	if len(operationID) > 0 {
		req.Header.Set("Idempotency-Key", operationID[0])
	}
	w := httptest.NewRecorder()
	f.router.ServeHTTP(w, req)
	if w.Code != want {
		t.Fatalf("%s %s status %d, want %d (response %s)", method, path, w.Code, want, w.Body.String())
	}
	if w.Header().Get("Cache-Control") != "no-store" {
		t.Fatal("native response missing no-store")
	}
	var result struct {
		Success bool            `json:"success"`
		Data    json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal(w.Body.Bytes(), &result); err != nil {
		t.Fatal(err)
	}
	if result.Success != (want == http.StatusOK) {
		t.Fatalf("response success=%v status=%d", result.Success, w.Code)
	}
	return result.Data
}

func nativeHTTPWorkflowBody(name string) map[string]any {
	return map[string]any{"name": name, "trigger": "manual", "enabled": false, "hasContent": true, "graphContent": map[string]any{"nodes": []any{
		map[string]any{"id": "start", "type": "start", "data": map[string]any{"config": map[string]any{"trigger": "manual"}}},
		map[string]any{"id": "end", "type": "end", "data": map[string]any{}},
	}}}
}

func (f *nativeHTTPFixture) workflow(t *testing.T, name string) string {
	t.Helper()
	row, err := f.engine.NativeSave(context.Background(), "1", "workflow", "", nativeHTTPWorkflowBody(name))
	if err != nil {
		t.Fatal(err)
	}
	id, ok := row["id"].(string)
	if !ok || id == "" {
		t.Fatal("workflow has no id")
	}
	if err := f.db.Create(&unified.ResourceBinding{Kind: "workflow", ResourceID: id, OwnerID: 1, Active: true, Version: 1}).Error; err != nil {
		t.Fatal(err)
	}
	return id
}

func (f *nativeHTTPFixture) grant(t *testing.T, id, action string) {
	t.Helper()
	_, err := f.service.CreateGrant(context.Background(), unified.Principal{UserID: 1, Admin: true}, unified.Grant{UserID: 2, Kind: "workflow", ResourceID: id, Action: action})
	if err != nil {
		t.Fatal(err)
	}
}

func TestUnifiedNativeHTTPAuthorizesBeforePaginationAndDetail(t *testing.T) {
	f := newNativeHTTPFixture(t)
	hidden := f.workflow(t, "a-hidden")
	visible := f.workflow(t, "b-visible")
	_ = f.workflow(t, "c-hidden")
	f.grant(t, visible, "read")
	f.principal = unified.Principal{UserID: 2}
	var page embedded.NativePage
	filter := url.QueryEscape("name~'hidden' || id='" + visible + "'")
	data := f.request(t, http.MethodGet, "/certimate/resources/workflow?page=1&perPage=1&sort=name&filter="+filter, nil, http.StatusOK)
	if err := json.Unmarshal(data, &page); err != nil {
		t.Fatal(err)
	}
	if page.TotalItems != 1 || page.TotalPages != 1 || len(page.Items) != 1 || page.Items[0]["id"] != visible {
		t.Fatalf("scope/count/pagination mismatch: %s", data)
	}
	data = f.request(t, http.MethodGet, "/certimate/resources/workflow?page=2&perPage=1", nil, http.StatusOK)
	if err := json.Unmarshal(data, &page); err != nil {
		t.Fatal(err)
	}
	if page.TotalItems != 1 || len(page.Items) != 0 {
		t.Fatalf("hidden rows affected second page: %s", data)
	}
	f.request(t, http.MethodGet, "/certimate/resources/workflow/"+visible, nil, http.StatusOK)
	f.request(t, http.MethodGet, "/certimate/resources/workflow/"+hidden, nil, http.StatusForbidden)
	f.request(t, http.MethodGet, "/certimate/resources/workflow?perPage=201", nil, http.StatusBadRequest)
	f.request(t, http.MethodGet, "/certimate/resources/workflow?page=0", nil, http.StatusBadRequest)

	issued, err := f.service.CreateScopedToken(context.Background(), unified.Principal{UserID: 1, Admin: true}, unified.ScopedTokenInput{Name: "synthetic-scoped-native", ExpiresAt: time.Now().Add(time.Hour), Scopes: []unified.Scope{{Kind: "workflow", ResourceID: visible, Action: "read"}}})
	if err != nil {
		t.Fatal(err)
	}
	f.principal, err = f.service.AuthenticateScopedToken(context.Background(), issued.Secret)
	if err != nil {
		t.Fatal(err)
	}
	data = f.request(t, http.MethodGet, "/certimate/resources/workflow?perPage=1", nil, http.StatusOK)
	if err = json.Unmarshal(data, &page); err != nil {
		t.Fatal(err)
	}
	if page.TotalItems != 1 || len(page.Items) != 1 || page.Items[0]["id"] != visible {
		t.Fatalf("administrator token scope bypass: %s", data)
	}
	f.request(t, http.MethodGet, "/certimate/resources/workflow/"+hidden, nil, http.StatusForbidden)
	f.request(t, http.MethodPatch, "/certimate/resources/workflow/"+visible, map[string]any{"name": "forbidden-write"}, http.StatusForbidden)
}

func TestUnifiedNativeHTTPDirectSideEffectsRequireAuthority(t *testing.T) {
	f := newNativeHTTPFixture(t)
	id := f.workflow(t, "private-workflow")
	for _, principal := range []unified.Principal{{}, {UserID: 2}, {UserID: 1, Admin: true, PAT: true}} {
		f.principal = principal
		for _, request := range []struct {
			method, path string
			body         any
		}{
			{http.MethodGet, "/certimate/resources/workflow/" + id, nil},
			{http.MethodPatch, "/certimate/resources/workflow/" + id, map[string]any{"enabled": true}},
			{http.MethodDelete, "/certimate/resources/workflow/" + id, nil},
			{http.MethodPost, "/certimate/workflows/" + id + "/runs", map[string]any{"trigger": "manual"}},
			{http.MethodPost, "/certimate/workflows/" + id + "/runs/synthetic-run/cancel", nil},
			{http.MethodPost, "/certimate/certificates/synthetic-cert/download", map[string]any{"fileFormat": "PEM"}},
			{http.MethodPost, "/certimate/certificates/synthetic-cert/revoke", nil},
			{http.MethodPost, "/certimate/notifications/test", map[string]any{"provider": "email", "accessId": "synthetic-credential"}},
			{http.MethodPost, "/certimate/credentials/synthetic-credential/duplicate", map[string]any{"name": "synthetic-clone"}},
			{http.MethodPost, "/certimate/credentials/synthetic-credential/export", nil},
			{http.MethodGet, "/certimate/system/crons", nil},
			{http.MethodGet, "/certimate/system/logs", nil},
			{http.MethodPost, "/ownership/transfer", map[string]any{"kind": "workflow", "resource_id": id, "owner_id": 2, "expected_version": 1}},
			{http.MethodPut, "/workflow-policies/" + id, map[string]any{"expected_version": 0, "run_as_user_id": 1, "enabled": true}},
		} {
			f.request(t, request.method, request.path, request.body, http.StatusForbidden)
		}
	}
}

func TestUnifiedNativeHTTPBatchReportsIndependentAuthorization(t *testing.T) {
	f := newNativeHTTPFixture(t)
	allowed := f.workflow(t, "allowed-original")
	denied := f.workflow(t, "denied-original")
	f.grant(t, allowed, "read")
	f.grant(t, allowed, "write")
	f.principal = unified.Principal{UserID: 2}
	data := f.request(t, http.MethodPost, "/certimate/batch", map[string]any{"operations": []any{
		map[string]any{"collection": "workflow", "method": "PATCH", "id": allowed, "body": map[string]any{"description": "permitted-change"}},
		map[string]any{"collection": "workflow", "method": "PATCH", "id": denied, "body": map[string]any{"description": "forbidden-change"}},
	}}, http.StatusOK)
	var outcomes []struct {
		Status int `json:"status"`
	}
	if err := json.Unmarshal(data, &outcomes); err != nil {
		t.Fatal(err)
	}
	if len(outcomes) != 2 || outcomes[0].Status != 200 || outcomes[1].Status != 403 {
		t.Fatalf("batch statuses: %s", data)
	}
	allowedRow, err := f.engine.NativeGet(context.Background(), "workflow", allowed)
	if err != nil || allowedRow["description"] != "permitted-change" {
		t.Fatal("authorized batch operation was not saved")
	}
	deniedRow, err := f.engine.NativeGet(context.Background(), "workflow", denied)
	if err != nil || deniedRow["description"] == "forbidden-change" {
		t.Fatal("unauthorized batch operation was saved")
	}
	f.request(t, http.MethodPost, "/certimate/batch", map[string]any{"operations": []any{map[string]any{"collection": "workflow", "method": "PUT", "id": allowed, "body": map[string]any{}}}}, http.StatusBadRequest)
}

func TestUnifiedNativeHTTPCreateIdempotency(t *testing.T) {
	f := newNativeHTTPFixture(t)
	body := nativeHTTPWorkflowBody("idempotent-workflow")
	first := f.request(t, http.MethodPost, "/certimate/resources/workflow", body, http.StatusOK, "synthetic-workflow-create")
	second := f.request(t, http.MethodPost, "/certimate/resources/workflow", body, http.StatusOK, "synthetic-workflow-create")
	var a, b map[string]any
	if err := json.Unmarshal(first, &a); err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal(second, &b); err != nil {
		t.Fatal(err)
	}
	if a["id"] != b["id"] {
		t.Fatal("same operation created duplicate resource")
	}
	f.request(t, http.MethodPost, "/certimate/resources/workflow", nativeHTTPWorkflowBody("different"), http.StatusBadRequest, "synthetic-workflow-create")
	f.request(t, http.MethodPost, "/certimate/resources/workflow", body, http.StatusBadRequest, "invalid operation")
}

func TestUnifiedNativeHTTPCredentialReadCloneAndExportAreSeparate(t *testing.T) {
	f := newNativeHTTPFixture(t)
	const syntheticSecret = "synthetic-native-http-secret"
	data := f.request(t, http.MethodPost, "/certimate/resources/access", map[string]any{
		"name": "Synthetic DNS credential", "provider": "cloudflare", "config": map[string]any{"apiToken": syntheticSecret},
	}, http.StatusOK, "synthetic-credential-create")
	var record map[string]any
	if err := json.Unmarshal(data, &record); err != nil {
		t.Fatal(err)
	}
	id, ok := record["id"].(string)
	if !ok || id == "" {
		t.Fatal("missing credential id")
	}
	if strings.Contains(string(data), syntheticSecret) || record["config"] != nil {
		t.Fatal("create returned secret material")
	}
	_, err := f.service.CreateGrant(context.Background(), unified.Principal{UserID: 1, Admin: true}, unified.Grant{UserID: 2, Kind: "credential", ResourceID: id, Action: "read"})
	if err != nil {
		t.Fatal(err)
	}
	f.principal = unified.Principal{UserID: 2}
	data = f.request(t, http.MethodGet, "/certimate/resources/access/"+id+"?fields=*", nil, http.StatusOK)
	if strings.Contains(string(data), syntheticSecret) {
		t.Fatal("metadata read returned secret material")
	}
	f.request(t, http.MethodPost, "/certimate/credentials/"+id+"/export", nil, http.StatusForbidden)
	f.request(t, http.MethodPost, "/certimate/credentials/"+id+"/duplicate", map[string]any{"name": "Unauthorized clone"}, http.StatusForbidden)
	f.principal = unified.Principal{UserID: 1, Admin: true}
	data = f.request(t, http.MethodPost, "/certimate/credentials/"+id+"/export", nil, http.StatusOK)
	var exported struct {
		Config map[string]any `json:"config"`
	}
	if err := json.Unmarshal(data, &exported); err != nil {
		t.Fatal(err)
	}
	if exported.Config["apiToken"] != syntheticSecret {
		t.Fatal("authorized explicit export did not return expected material")
	}
	data = f.request(t, http.MethodPost, "/certimate/credentials/"+id+"/duplicate", map[string]any{"name": "Authorized clone"}, http.StatusOK, "synthetic-credential-clone")
	if strings.Contains(string(data), syntheticSecret) {
		t.Fatal("clone returned secret material")
	}
	var cloned map[string]any
	if err := json.Unmarshal(data, &cloned); err != nil {
		t.Fatal(err)
	}
	if cloned["id"] == id || cloned["configConfigured"] != true {
		t.Fatal("server-side clone did not create distinct configured credential")
	}
	issued, err := f.service.CreateScopedToken(context.Background(), f.principal, unified.ScopedTokenInput{
		Name: "synthetic-credential-read-only", ExpiresAt: time.Now().Add(time.Hour),
		Scopes: []unified.Scope{{Kind: "credential", ResourceID: id, Action: "read"}},
	})
	if err != nil {
		t.Fatal(err)
	}
	f.principal, err = f.service.AuthenticateScopedToken(context.Background(), issued.Secret)
	if err != nil {
		t.Fatal(err)
	}
	f.request(t, http.MethodGet, "/certimate/resources/access/"+id, nil, http.StatusOK)
	f.request(t, http.MethodPost, "/certimate/credentials/"+id+"/export", nil, http.StatusForbidden)
	f.request(t, http.MethodPost, "/certimate/credentials/"+id+"/duplicate", map[string]any{"name": "Scoped clone"}, http.StatusForbidden)
}
