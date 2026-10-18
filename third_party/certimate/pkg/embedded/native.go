package embedded

// Native services use no PocketBase HTTP, authentication or public record
// serialization. Host-authorized scopes apply before count and pagination.
import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/certimate-go/certimate/internal/domain"
	"github.com/certimate-go/certimate/internal/embedcrypto"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/cron"
	"github.com/pocketbase/pocketbase/tools/search"
	"slices"
	"strings"
	"time"
)

type NativeOptions struct {
	Page, PerPage                int
	Filter, Sort, Fields, Expand string
}
type NativeScope struct {
	All              bool
	IDs, WorkflowIDs []string
}
type NativePage struct {
	Page       int              `json:"page"`
	PerPage    int              `json:"perPage"`
	TotalItems int              `json:"totalItems"`
	TotalPages int              `json:"totalPages"`
	Items      []map[string]any `json:"items"`
}

var nativeReadFields = map[string][]string{
	"access":          {"id", "name", "provider", "reserve", "created", "updated", "deleted", "unified_enabled", "unified_version"},
	"certificate":     {"id", "source", "serialNumber", "subjectName", "subjectAltNames", "issuerName", "issuerOrg", "keyAlgorithm", "validationPolicy", "validityNotBefore", "validityNotAfter", "validityInterval", "isRenewed", "isRevoked", "workflowRef", "workflowRunRef", "created", "updated", "deleted"},
	"workflow":        {"id", "name", "description", "trigger", "triggerCron", "enabled", "hasDraft", "hasContent", "lastRunRef", "lastRunStatus", "lastRunTime", "created", "updated"},
	"workflow_run":    {"id", "workflowRef", "status", "trigger", "startedAt", "endedAt", "error", "created", "updated"},
	"workflow_output": {"id", "runRef", "nodeId", "nodeType", "nodeName", "succeeded", "startedAt", "endedAt", "error", "created", "updated"},
	"workflow_logs":   {"id", "runRef", "nodeId", "nodeName", "level", "message", "timestamp", "created", "updated"},
	"settings":        {"id", "name", "created", "updated"},
}

func nativeManaged(collection string) bool {
	return collection == "access" || collection == "certificate" || collection == "workflow"
}
func nativeActive(r *core.Record) bool {
	return r.GetString("unified_operation_id") == "" || r.GetBool("unified_active")
}
func nativeScopeExpression(collection string, scope NativeScope) dbx.Expression {
	if scope.All {
		return dbx.NewExp("1=1")
	}
	var expressions []dbx.Expression
	if len(scope.IDs) > 0 {
		expressions = append(expressions, dbx.In("id", slicesToAny(scope.IDs)...))
	}
	if len(scope.WorkflowIDs) > 0 {
		switch collection {
		case "workflow":
			expressions = append(expressions, dbx.In("id", slicesToAny(scope.WorkflowIDs)...))
		case "workflow_run", "certificate":
			expressions = append(expressions, dbx.In("workflowRef", slicesToAny(scope.WorkflowIDs)...))
		case "workflow_logs", "workflow_output":
			params := dbx.Params{}
			placeholders := make([]string, len(scope.WorkflowIDs))
			for i, id := range scope.WorkflowIDs {
				key := fmt.Sprintf("native_workflow_%d", i)
				placeholders[i] = "{:" + key + "}"
				params[key] = id
			}
			expressions = append(expressions, dbx.NewExp("[[runRef]] IN (SELECT [[id]] FROM [[workflow_run]] WHERE [[workflowRef]] IN ("+strings.Join(placeholders, ",")+"))", params))
		}
	}
	if len(expressions) == 0 {
		return dbx.NewExp("1=0")
	}
	return dbx.Or(expressions...)
}
func (e *Engine) NativeList(ctx context.Context, collection string, scope NativeScope, opts NativeOptions) (NativePage, error) {
	e.mu.RLock()
	defer e.mu.RUnlock()
	result := NativePage{Page: opts.Page, PerPage: opts.PerPage, Items: []map[string]any{}}
	if result.Page == 0 {
		result.Page = 1
	}
	if result.PerPage == 0 {
		result.PerPage = 20
	}
	if err := e.check(ctx); err != nil {
		return result, err
	}
	fields, ok := nativeReadFields[collection]
	if !ok || result.Page < 1 || result.Page > 100000 || result.PerPage < 1 || result.PerPage > 200 || len(opts.Filter) > 16384 || len(scope.IDs)+len(scope.WorkflowIDs) > 10000 {
		return result, ErrInvalid
	}
	q := e.app.RecordQuery(collection).WithContext(ctx).AndWhere(nativeScopeExpression(collection, scope))
	if collection == "access" || collection == "certificate" {
		q.AndWhere(dbx.HashExp{"deleted": ""})
	}
	if nativeManaged(collection) {
		q.AndWhere(dbx.Or(dbx.HashExp{"unified_operation_id": ""}, dbx.HashExp{"unified_active": true}))
	}
	filterFields := slices.Clone(fields)
	if collection == "workflow_output" {
		filterFields = append(filterFields, "outputs")
	}
	resolver := search.NewSimpleFieldResolver(filterFields...)
	if opts.Filter != "" {
		expr, err := search.FilterData(opts.Filter).BuildExpr(resolver)
		if err != nil {
			return result, ErrInvalid
		}
		q.AndWhere(expr)
	}
	count := *q
	if err := count.Select("COUNT(*)").Row(&result.TotalItems); err != nil {
		return result, err
	}
	if opts.Sort == "" {
		opts.Sort = "-created"
	}
	for _, field := range search.ParseSortFromString(opts.Sort) {
		if strings.HasPrefix(field.Name, "@") {
			return result, ErrInvalid
		}
		expr, err := field.BuildExpr(resolver)
		if err != nil {
			return result, ErrInvalid
		}
		q.AndOrderBy(expr)
	}
	q.AndOrderBy("id ASC").Limit(int64(result.PerPage)).Offset(int64((result.Page - 1) * result.PerPage))
	var records []*core.Record
	if err := q.All(&records); err != nil {
		return result, err
	}
	for _, r := range records {
		item, err := nativeProjection(e.app, r, false)
		if err != nil {
			return result, err
		}
		result.Items = append(result.Items, item)
	}
	result.TotalPages = (result.TotalItems + result.PerPage - 1) / result.PerPage
	return result, nil
}
func (e *Engine) NativeGet(ctx context.Context, collection, id string) (map[string]any, error) {
	e.mu.RLock()
	defer e.mu.RUnlock()
	if err := e.check(ctx); err != nil {
		return nil, err
	}
	r, err := e.nativeRecord(collection, id)
	if err != nil {
		return nil, err
	}
	return nativeProjection(e.app, r, true)
}
func (e *Engine) nativeRecord(collection, id string) (*core.Record, error) {
	if _, ok := nativeReadFields[collection]; !ok {
		return nil, ErrInvalid
	}
	r, err := e.app.FindRecordById(collection, id)
	if err != nil || !r.GetDateTime("deleted").IsZero() || (nativeManaged(collection) && !nativeActive(r)) {
		return nil, ErrNotFound
	}
	return r, nil
}
func (e *Engine) NativeResource(ctx context.Context, collection, id string) (string, string, error) {
	e.mu.RLock()
	defer e.mu.RUnlock()
	if err := e.check(ctx); err != nil {
		return "", "", err
	}
	r, err := e.nativeRecord(collection, id)
	if err != nil {
		return "", "", err
	}
	switch collection {
	case "access":
		return "credential", id, nil
	case "certificate":
		return "certificate", id, nil
	case "workflow":
		return "workflow", id, nil
	case "settings":
		return "settings", id, nil
	case "workflow_run":
		workflowID := r.GetString("workflowRef")
		if _, err := e.nativeRecord("workflow", workflowID); err != nil {
			return "", "", err
		}
		return "workflow", workflowID, nil
	case "workflow_output", "workflow_logs":
		run, err := e.nativeRecord("workflow_run", r.GetString("runRef"))
		if err != nil {
			return "", "", err
		}
		workflowID := run.GetString("workflowRef")
		if _, err := e.nativeRecord("workflow", workflowID); err != nil {
			return "", "", err
		}
		return "workflow", workflowID, nil
	}
	return "", "", ErrInvalid
}
func nativeProjection(app core.App, r *core.Record, detail bool) (map[string]any, error) {
	collection := r.Collection().Name
	out := map[string]any{"collectionName": collection, "collectionId": r.Collection().Id}
	for _, field := range nativeReadFields[collection] {
		out[field] = r.Get(field)
	}
	if collection == "access" {
		out["configConfigured"] = r.GetString("unified_envelope") != ""
	}
	if collection == "certificate" && detail {
		out["certificate"] = r.GetString("certificate")
		out["privateKey"] = ""
	}
	var fields []string
	switch collection {
	case "workflow":
		if detail {
			fields = []string{"graphDraft", "graphContent"}
		}
	case "workflow_run":
		if detail {
			fields = []string{"graph"}
		}
	case "workflow_output":
		fields = []string{"outputs"}
		if detail {
			fields = append(fields, "nodeConfig")
		}
	case "settings":
		fields = []string{"content"}
	}
	for _, field := range fields {
		var value any
		if err := embedcrypto.ReadJSON(app, r, field, &value); err != nil {
			return nil, err
		}
		out[field] = nativeRedact(value)
	}
	if collection == "workflow_output" && out["outputs"] == nil {
		out["outputs"] = []any{}
	}
	// These expansions remain inside the same workflow authorization root.
	// Certificate -> workflow expansion requires a separate host decision.
	if collection == "workflow" && r.GetString("lastRunRef") != "" {
		if run, err := app.FindRecordById("workflow_run", r.GetString("lastRunRef")); err == nil && run.GetString("workflowRef") == r.Id {
			out["expand"] = map[string]any{"lastRunRef": map[string]any{"id": run.Id, "status": run.GetString("status"), "trigger": run.GetString("trigger"), "startedAt": run.Get("startedAt"), "endedAt": run.Get("endedAt"), "error": run.GetString("error")}}
		}
	}
	if collection == "workflow_run" && r.GetString("workflowRef") != "" {
		if workflow, err := app.FindRecordById("workflow", r.GetString("workflowRef")); err == nil && nativeActive(workflow) {
			out["expand"] = map[string]any{"workflowRef": map[string]any{"id": workflow.Id, "name": workflow.GetString("name"), "description": workflow.GetString("description")}}
		}
	}
	return out, nil
}

func nativeValidGraph(graph *domain.WorkflowGraph) bool {
	if graph == nil || len(graph.Nodes) < 2 {
		return false
	}
	count := 0
	seen := map[string]bool{}
	var walk func([]*domain.WorkflowNode, int) bool
	walk = func(nodes []*domain.WorkflowNode, depth int) bool {
		if depth > 64 {
			return false
		}
		for _, node := range nodes {
			count++
			if count > 4096 || node == nil || node.Id == "" || seen[node.Id] {
				return false
			}
			seen[node.Id] = true
			if !slices.Contains([]string{"start", "end", "condition", "branchBlock", "tryCatch", "tryBlock", "catchBlock", "delay", "bizApply", "bizUpload", "bizMonitor", "bizDeploy", "bizNotify"}, string(node.Type)) {
				return false
			}
			if !walk(node.Blocks, depth+1) {
				return false
			}
		}
		return true
	}
	return walk(graph.Nodes, 0) && graph.Verify() == nil
}
func nativeSecretField(key string) bool {
	key = strings.ToLower(strings.ReplaceAll(strings.ReplaceAll(key, "_", ""), "-", ""))
	return strings.Contains(key, "privatekey") || strings.Contains(key, "password") || strings.Contains(key, "secret") || strings.Contains(key, "token") || key == "keycontent" || key == "clientkey" || key == "authorization" || key == "cookie" || key == "apikey" || key == "accesskey" || key == "eabhmackey" || strings.HasPrefix(key, "unified")
}

func nativeSecretEntry(value map[string]any) bool {
	for _, field := range []string{"name", "key"} {
		if name, ok := value[field].(string); ok && nativeSecretField(name) {
			return true
		}
	}
	return false
}
func nativeRedact(value any) any {
	switch v := value.(type) {
	case map[string]any:
		out := map[string]any{}
		for key, item := range v {
			if nativeSecretField(key) || key == "value" && nativeSecretEntry(v) {
				out[key] = ""
			} else {
				out[key] = nativeRedact(item)
			}
		}
		return out
	case []any:
		out := make([]any, len(v))
		for i, item := range v {
			out[i] = nativeRedact(item)
		}
		return out
	default:
		return value
	}
}

// Secret leaves are merged by node identity so reordering never crosses secrets.
func nativeMergeRedacted(value, previous any) any {
	switch v := value.(type) {
	case map[string]any:
		old, _ := previous.(map[string]any)
		out := map[string]any{}
		for key, item := range v {
			if str, ok := item.(string); (nativeSecretField(key) || key == "value" && nativeSecretEntry(v)) && ok && str == "" {
				if prior, exists := old[key]; exists {
					out[key] = prior
					continue
				}
			}
			out[key] = nativeMergeRedacted(item, old[key])
		}
		return out
	case []any:
		old, _ := previous.([]any)
		byID := map[string]any{}
		for _, item := range old {
			if m, ok := item.(map[string]any); ok {
				if id, ok := m["id"].(string); ok && id != "" {
					byID[id] = item
				}
			}
		}
		out := make([]any, len(v))
		for i, item := range v {
			var prior any
			if m, ok := item.(map[string]any); ok {
				if id, ok := m["id"].(string); ok {
					prior = byID[id]
				}
			}
			out[i] = nativeMergeRedacted(item, prior)
		}
		return out
	default:
		return value
	}
}

var nativeWriteFields = map[string][]string{"access": {"name", "provider", "reserve", "config", "unified_enabled"}, "workflow": {"name", "description", "trigger", "triggerCron", "enabled", "graphDraft", "graphContent", "hasDraft", "hasContent"}, "settings": {"name", "content"}}

func (e *Engine) NativeSave(ctx context.Context, actor, collection, id string, input map[string]any) (map[string]any, error) {
	e.mu.RLock()
	defer e.mu.RUnlock()
	if err := e.check(ctx); err != nil {
		return nil, err
	}
	fields, ok := nativeWriteFields[collection]
	if !ok || actor == "" {
		return nil, ErrInvalid
	}
	if _, deleting := input["deleted"]; deleting {
		return nil, ErrInvalid
	}
	encoded, err := json.Marshal(input)
	if err != nil || len(encoded) > 2<<20 {
		return nil, ErrInvalid
	}
	var result *core.Record
	err = e.app.RunInTransaction(func(tx core.App) error {
		var r *core.Record
		if id == "" {
			c, err := tx.FindCollectionByNameOrId(collection)
			if err != nil {
				return err
			}
			r = core.NewRecord(c)
			if err := embedcrypto.EnsureID(r); err != nil {
				return err
			}
			if collection == "workflow" {
				for _, field := range []string{"graphDraft", "graphContent"} {
					if err := embedcrypto.SetJSON(tx, r, field, nil); err != nil {
						return err
					}
				}
			}
		} else {
			var err error
			r, err = tx.FindRecordById(collection, id)
			if err != nil || !r.GetDateTime("deleted").IsZero() || (nativeManaged(collection) && !nativeActive(r)) {
				return ErrNotFound
			}
		}
		if collection == "access" && id != "" {
			if provider, exists := input["provider"]; exists && provider != r.GetString("provider") {
				return errors.New("provider is immutable; create a new credential")
			}
		}
		for _, field := range fields {
			value, exists := input[field]
			if !exists || field == "config" {
				continue
			}
			if field == "graphDraft" || field == "graphContent" || field == "content" {
				var prior any
				if id != "" {
					if err := embedcrypto.ReadJSON(tx, r, field, &prior); err != nil {
						return err
					}
				}
				merged := nativeMergeRedacted(value, prior)
				if field == "graphContent" && merged != nil {
					blob, err := json.Marshal(merged)
					var graph domain.WorkflowGraph
					if err != nil || json.Unmarshal(blob, &graph) != nil || !nativeValidGraph(&graph) {
						return ErrInvalid
					}
				}
				if err := embedcrypto.SetJSON(tx, r, field, merged); err != nil {
					return err
				}
			} else {
				r.Set(field, value)
			}
		}
		if collection == "settings" && !slices.Contains([]string{"emails", "notifyTemplate", "scriptTemplate", "sslProvider", "persistence"}, r.GetString("name")) {
			return ErrInvalid
		}
		if collection == "workflow" && r.GetBool("enabled") && r.GetString("trigger") == "scheduled" {
			if _, err := cron.NewSchedule(r.GetString("triggerCron")); err != nil {
				return ErrInvalid
			}
		}
		if collection == "access" {
			if id == "" {
				r.Set("unified_enabled", true)
				r.Set("unified_version", 1)
			}
			if old, exists := input["unified_version"]; exists && id != "" {
				expected := 0
				switch v := old.(type) {
				case float64:
					if v != float64(int(v)) {
						return ErrInvalid
					}
					expected = int(v)
				case int:
					expected = v
				default:
					return ErrInvalid
				}
				if expected != r.GetInt("unified_version") {
					return ErrInvalid
				}
			}
			config, supplied := input["config"].(map[string]any)
			if supplied && len(config) > 0 {
				if id != "" {
					if _, exists := input["unified_version"]; !exists {
						return errors.New("credential version is required when replacing material")
					}
					priorBytes, err := e.unseal(fmt.Sprintf("credential:%s:%d", r.Id, r.GetInt("unified_version")), r.GetString("unified_envelope"))
					if err != nil {
						return err
					}
					var prior map[string]any
					if json.Unmarshal(priorBytes, &prior) != nil {
						return ErrInvalid
					}
					config = nativeMergeRedacted(config, prior).(map[string]any)
					r.Set("unified_version", r.GetInt("unified_version")+1)
				}
				data, err := json.Marshal(config)
				if err != nil || len(data) > 65536 {
					return ErrInvalid
				}
				r.Set("unified_envelope", e.seal(fmt.Sprintf("credential:%s:%d", r.Id, r.GetInt("unified_version")), data))
			} else if id == "" {
				return ErrInvalid
			}
			r.Set("config", map[string]any{})
		}
		if id == "" && nativeManaged(collection) {
			if op, ok := input["unified_operation_id"].(string); ok && op != "" {
				r.Set("unified_operation_id", op)
				r.Set("unified_active", false)
			}
		}
		if err := tx.Save(r); err != nil {
			return err
		}
		if err := audit(tx, actor, "native."+collection+".save", collection+":"+r.Id); err != nil {
			return err
		}
		result = r
		return nil
	})
	if err != nil {
		return nil, err
	}
	if nativeActive(result) {
		if err := e.reconcileNativeWrite(ctx, collection, result, false); err != nil {
			return nil, err
		}
	}
	return nativeProjection(e.app, result, true)
}
func (e *Engine) NativeDelete(ctx context.Context, actor, collection, id string) error {
	e.mu.RLock()
	defer e.mu.RUnlock()
	if err := e.check(ctx); err != nil {
		return err
	}
	if actor == "" || !slices.Contains([]string{"access", "certificate", "workflow", "workflow_run"}, collection) {
		return ErrInvalid
	}
	var deleted *core.Record
	err := e.app.RunInTransaction(func(tx core.App) error {
		r, err := tx.FindRecordById(collection, id)
		if err != nil || (nativeManaged(collection) && !nativeActive(r)) {
			return ErrNotFound
		}
		if collection == "workflow" {
			var active int
			if err := tx.RecordQuery("workflow_run").Select("COUNT(*)").AndWhere(dbx.HashExp{"workflowRef": id}).AndWhere(dbx.In("status", "pending", "processing")).Row(&active); err != nil {
				return err
			}
			if active > 0 {
				return errors.New("cancel active runs before deleting the workflow")
			}
		}
		if collection == "workflow_run" {
			if !slices.Contains([]string{"succeeded", "failed", "canceled"}, r.GetString("status")) {
				return errors.New("only terminal workflow runs can be removed")
			}
			if workflow, err := tx.FindRecordById("workflow", r.GetString("workflowRef")); err == nil && workflow.GetString("lastRunRef") == id {
				workflow.Set("lastRunRef", "")
				workflow.Set("lastRunStatus", "")
				workflow.Set("lastRunTime", "")
				if err := tx.Save(workflow); err != nil {
					return err
				}
			}
			// Explicit deletion keeps the lifecycle contract visible even if an
			// upstream migration changes relation cascade flags in the future.
			for _, childCollection := range []string{"workflow_output", "workflow_logs"} {
				var children []*core.Record
				if err := tx.RecordQuery(childCollection).AndWhere(dbx.HashExp{"runRef": id}).All(&children); err != nil {
					return err
				}
				for _, child := range children {
					if err := tx.Delete(child); err != nil {
						return err
					}
				}
			}
		}
		if collection == "access" || collection == "certificate" {
			r.Set("deleted", time.Now().UTC())
			err = tx.Save(r)
		} else {
			err = tx.Delete(r)
		}
		if err != nil {
			return err
		}
		deleted = r
		return audit(tx, actor, "native."+collection+".delete", collection+":"+id)
	})
	if err != nil {
		return err
	}
	return e.reconcileNativeWrite(ctx, collection, deleted, true)
}
func (e *Engine) ActivateNative(ctx context.Context, collection, id, operationID string) error {
	e.mu.RLock()
	defer e.mu.RUnlock()
	if err := e.check(ctx); err != nil {
		return err
	}
	if !nativeManaged(collection) || operationID == "" {
		return ErrInvalid
	}
	var activated *core.Record
	err := e.app.RunInTransaction(func(tx core.App) error {
		r, err := tx.FindRecordById(collection, id)
		if err != nil {
			return ErrNotFound
		}
		if r.GetString("unified_operation_id") != operationID {
			return ErrInvalid
		}
		r.Set("unified_active", true)
		if err := tx.Save(r); err != nil {
			return err
		}
		activated = r
		return nil
	})
	if err != nil {
		return err
	}
	return e.reconcileNativeWrite(ctx, collection, activated, false)
}
func (e *Engine) FindNativeOperation(ctx context.Context, collection, operationID string) (map[string]any, error) {
	e.mu.RLock()
	defer e.mu.RUnlock()
	if err := e.check(ctx); err != nil {
		return nil, err
	}
	if !nativeManaged(collection) || operationID == "" {
		return nil, ErrInvalid
	}
	records, err := e.app.FindRecordsByFilter(collection, "unified_operation_id={:op}", "", 1, 0, dbx.Params{"op": operationID})
	if err != nil {
		return nil, err
	}
	if len(records) == 0 {
		return nil, ErrNotFound
	}
	out, err := nativeProjection(e.app, records[0], collection == "workflow")
	if err != nil {
		return nil, err
	}
	out["unified_active"] = records[0].GetBool("unified_active")
	return out, nil
}

// NativeCloneCredential is a separately authorized server operation. The source
// material never appears in a read response, and the new record remains inactive
// until the host has durably bound its ownership and exact resource identity.
func (e *Engine) NativeCloneCredential(ctx context.Context, use CredentialUse, actor, name, operationID string) (map[string]any, error) {
	if actor == "" || strings.TrimSpace(name) == "" || operationID == "" || use.Purpose != "credential.clone" || use.ResourceKind != "credential" || use.ResourceID != use.CredentialID {
		return nil, ErrInvalid
	}
	config, err := e.ResolveCredential(ctx, use)
	if err != nil {
		return nil, err
	}
	e.mu.RLock()
	if err := e.check(ctx); err != nil {
		e.mu.RUnlock()
		return nil, err
	}
	r, err := e.nativeRecord("access", use.CredentialID)
	if err != nil {
		e.mu.RUnlock()
		return nil, err
	}
	if r.GetInt("unified_version") != use.Version || !r.GetBool("unified_enabled") {
		e.mu.RUnlock()
		return nil, ErrInvalid
	}
	provider, reserve := r.GetString("provider"), r.GetString("reserve")
	err = audit(e.app, actor, "credential.clone.request", "credential:"+use.CredentialID)
	e.mu.RUnlock()
	if err != nil {
		return nil, err
	}
	return e.NativeSave(ctx, actor, "access", "", map[string]any{"name": name, "provider": provider, "reserve": reserve, "config": config, "unified_operation_id": operationID})
}
