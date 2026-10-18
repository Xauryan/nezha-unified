package unified

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/certimate-go/certimate/pkg/embedded"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func collectionKind(collection string) string {
	return map[string]string{"access": "credential", "certificate": "certificate", "workflow": "workflow", "workflow_run": "workflow", "workflow_output": "workflow", "workflow_logs": "workflow", "settings": "settings"}[collection]
}

func (s *Service) resourceScope(ctx context.Context, p Principal, kind string) (embedded.NativeScope, error) {
	p, err := s.CurrentPrincipal(ctx, p)
	if err != nil {
		return embedded.NativeScope{}, err
	}
	if p.Admin && p.TokenID == 0 && p.ServiceID == 0 {
		return embedded.NativeScope{All: true}, nil
	}
	ids := []string{}
	if p.TokenID != 0 {
		var token ScopedToken
		if err = s.db.WithContext(ctx).First(&token, p.TokenID).Error; err != nil {
			return embedded.NativeScope{}, err
		}
		for _, scope := range token.Scopes {
			if scope.Kind == kind && scope.Action == "read" {
				ids = append(ids, scope.ResourceID)
			}
		}
	} else if p.ServiceID != 0 {
		var service ServiceIdentity
		if err = s.db.WithContext(ctx).First(&service, p.ServiceID).Error; err != nil {
			return embedded.NativeScope{}, err
		}
		for _, scope := range service.Scopes {
			if scope.Kind == kind && scope.Action == "read" {
				ids = append(ids, scope.ResourceID)
			}
		}
	} else {
		if err = s.db.WithContext(ctx).Model(&Grant{}).Where("user_id = ? AND kind = ? AND action = ?", p.UserID, kind, "read").Pluck("resource_id", &ids).Error; err != nil {
			return embedded.NativeScope{}, err
		}
		var owned []string
		if err = s.db.WithContext(ctx).Model(&ResourceBinding{}).Where("owner_id = ? AND kind = ? AND active = ?", p.UserID, kind, true).Pluck("resource_id", &owned).Error; err != nil {
			return embedded.NativeScope{}, err
		}
		ids = append(ids, owned...)
	}
	ids = slices.DeleteFunc(ids, func(id string) bool { return !s.Can(ctx, p, kind, id, "read") })
	slices.Sort(ids)
	ids = slices.Compact(ids)
	scope := embedded.NativeScope{IDs: ids}
	// Certificate visibility is an explicit workflow policy capability, separate
	// from being able to inspect the workflow or its logs.
	if kind == "certificate" && p.TokenID == 0 && p.ServiceID == 0 {
		var policies []WorkflowPolicy
		if err = s.db.WithContext(ctx).Where("enabled = ? AND certificate_read = ?", true, true).Find(&policies).Error; err != nil {
			return scope, err
		}
		for _, wp := range policies {
			if s.Can(ctx, p, "workflow", wp.WorkflowID, "read") {
				scope.WorkflowIDs = append(scope.WorkflowIDs, wp.WorkflowID)
			}
		}
	}
	return scope, nil
}
func (s *Service) NativeList(ctx context.Context, p Principal, collection string, opts embedded.NativeOptions) (embedded.NativePage, error) {
	kind := collectionKind(collection)
	if kind == "" {
		return embedded.NativePage{}, embedded.ErrInvalid
	}
	if kind == "settings" {
		p, err := s.CurrentPrincipal(ctx, p)
		if err != nil {
			return embedded.NativePage{}, err
		}
		if requireAdmin(p) != nil {
			return embedded.NativePage{}, ErrDenied
		}
		return s.engine.NativeList(ctx, collection, embedded.NativeScope{All: true}, opts)
	}
	scope, err := s.resourceScope(ctx, p, kind)
	if err != nil {
		return embedded.NativePage{}, err
	}
	if kind == "workflow" && collection != "workflow" {
		scope.WorkflowIDs = scope.IDs
		scope.IDs = nil
	}
	return s.engine.NativeList(ctx, collection, scope, opts)
}
func (s *Service) nativeCan(ctx context.Context, p Principal, collection, id, action string) bool {
	kind, resourceID, err := s.engine.NativeResource(ctx, collection, id)
	if err != nil {
		return false
	}
	if kind == "settings" {
		current, err := s.CurrentPrincipal(ctx, p)
		return err == nil && requireAdmin(current) == nil
	}
	if s.Can(ctx, p, kind, resourceID, action) {
		return true
	}
	if kind == "certificate" && action == "read" && p.TokenID == 0 && p.ServiceID == 0 {
		record, err := s.engine.NativeGet(ctx, collection, id)
		if err != nil {
			return false
		}
		wf, _ := record["workflowRef"].(string)
		var wp WorkflowPolicy
		return wf != "" && s.db.WithContext(ctx).First(&wp, "workflow_id = ? AND enabled = ? AND certificate_read = ?", wf, true, true).Error == nil && s.Can(ctx, p, "workflow", wf, "read")
	}
	return false
}
func (s *Service) NativeGet(ctx context.Context, p Principal, collection, id string) (map[string]any, error) {
	if !s.nativeCan(ctx, p, collection, id, "read") {
		return nil, ErrDenied
	}
	return s.engine.NativeGet(ctx, collection, id)
}
func operationID() (string, error) {
	var raw [24]byte
	if _, err := rand.Read(raw[:]); err != nil {
		return "", err
	}
	return hex.EncodeToString(raw[:]), nil
}

// NativeSave is the cross-store commit coordinator. Every newly created record
// remains inactive until the host owner/policy transaction has committed.
func (s *Service) NativeSave(ctx context.Context, p Principal, collection, id, op string, input map[string]any) (map[string]any, error) {
	p, err := s.CurrentPrincipal(ctx, p)
	if err != nil {
		return nil, err
	}
	kind := collectionKind(collection)
	if kind == "" || collection == "workflow_run" || collection == "workflow_output" || collection == "workflow_logs" {
		return nil, embedded.ErrInvalid
	}
	for key := range input {
		if key == "deleted" {
			return nil, embedded.ErrInvalid
		}
		if strings.HasPrefix(key, "unified_") && key != "unified_version" && key != "unified_enabled" {
			return nil, embedded.ErrInvalid
		}
	}
	if kind == "settings" {
		if requireAdmin(p) != nil {
			return nil, ErrDenied
		}
	} else if id == "" {
		if !s.Can(ctx, p, kind, "new", "create") {
			return nil, ErrDenied
		}
	} else if !s.nativeCan(ctx, p, collection, id, "write") {
		return nil, ErrDenied
	}
	if collection == "access" && id != "" {
		_, configChange := input["config"]
		_, enabledChange := input["unified_enabled"]
		if (configChange || enabledChange) && !s.Can(ctx, p, "credential", id, "rotate") {
			return nil, ErrDenied
		}
		if _, has := input["provider"]; has && requireAdmin(p) != nil {
			return nil, ErrDenied
		}
	}
	s.gate.Lock()
	defer s.gate.Unlock()
	p, err = s.CurrentPrincipal(ctx, p)
	if err != nil {
		return nil, err
	}
	if kind == "settings" {
		if requireAdmin(p) != nil {
			return nil, ErrDenied
		}
	} else if id == "" {
		if !s.Can(ctx, p, kind, "new", "create") {
			return nil, ErrDenied
		}
	} else if !s.nativeCan(ctx, p, collection, id, "write") {
		return nil, ErrDenied
	}
	if collection == "access" && id != "" {
		_, configChange := input["config"]
		_, enabledChange := input["unified_enabled"]
		if (configChange || enabledChange) && !s.Can(ctx, p, "credential", id, "rotate") {
			return nil, ErrDenied
		}
	}
	ctx = WithPrincipal(ctx, p)
	if id != "" || collection == "settings" {
		out, err := s.engine.NativeSave(ctx, fmt.Sprint(p.UserID), collection, id, input)
		if err != nil {
			return nil, err
		}
		if collection == "workflow" {
			if err = s.updateWorkflowApproval(ctx, p, id); err != nil {
				return nil, err
			}
		}
		return out, nil
	}
	if op == "" {
		op, err = operationID()
		if err != nil {
			return nil, err
		}
	}
	if len(op) > 64 || strings.ContainsAny(op, "/\\ \r\n") {
		return nil, embedded.ErrInvalid
	}
	raw, err := json.Marshal(input)
	if err != nil {
		return nil, embedded.ErrInvalid
	}
	sum := sha256.Sum256(raw)
	digest := hex.EncodeToString(sum[:])
	intent := ResourceIntent{OperationID: op, ActorID: p.UserID, TokenID: p.TokenID, ServiceID: p.ServiceID, IdentityVersion: p.IdentityVersion, Kind: kind, Collection: collection, Digest: digest, State: "prepared"}
	inserted := s.db.WithContext(ctx).Clauses(clause.OnConflict{DoNothing: true}).Create(&intent)
	if inserted.Error != nil {
		return nil, inserted.Error
	}
	if inserted.RowsAffected == 0 {
		if err = s.db.WithContext(ctx).First(&intent, "operation_id = ?", op).Error; err != nil {
			return nil, err
		}
		if intent.ActorID != p.UserID || intent.TokenID != p.TokenID || intent.ServiceID != p.ServiceID || intent.IdentityVersion != p.IdentityVersion || intent.Collection != collection || intent.Digest != digest {
			return nil, embedded.ErrInvalid
		}
	}
	if intent.State == "active" {
		if !s.nativeCan(ctx, p, collection, intent.ResourceID, "read") {
			return nil, ErrDenied
		}
		return s.engine.NativeGet(ctx, collection, intent.ResourceID)
	}
	record, err := s.engine.FindNativeOperation(ctx, collection, op)
	if errors.Is(err, embedded.ErrNotFound) {
		body := make(map[string]any, len(input)+2)
		for k, v := range input {
			body[k] = v
		}
		body["unified_operation_id"] = op
		body["unified_active"] = false
		record, err = s.engine.NativeSave(ctx, fmt.Sprint(p.UserID), collection, "", body)
	}
	if err != nil {
		return nil, err
	}
	id, _ = record["id"].(string)
	if id == "" {
		return nil, embedded.ErrInvalid
	}
	intent.ResourceID = id
	if err = s.commitNativeOwnership(ctx, p, intent); err != nil {
		return nil, err
	}
	return s.engine.NativeGet(ctx, collection, id)
}
func (s *Service) commitNativeOwnership(ctx context.Context, p Principal, intent ResourceIntent) error {
	binding := ResourceBinding{Kind: intent.Kind, ResourceID: intent.ResourceID, OwnerID: intent.ActorID, Version: 1, Active: true, OperationID: intent.OperationID}
	if err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Clauses(clause.OnConflict{DoNothing: true}).Create(&binding).Error; err != nil {
			return err
		}
		var saved ResourceBinding
		if err := tx.First(&saved, "kind = ? AND resource_id = ?", intent.Kind, intent.ResourceID).Error; err != nil {
			return err
		}
		if !saved.Active || saved.OwnerID != intent.ActorID || saved.OperationID != intent.OperationID {
			return ErrDenied
		}
		if err := tx.Model(&ResourceIntent{}).Where("operation_id = ?", intent.OperationID).Updates(map[string]any{"resource_id": intent.ResourceID, "state": "linked"}).Error; err != nil {
			return err
		}
		return tx.Create(&Audit{ActorID: intent.ActorID, Action: "resource.link", Resource: intent.Kind + ":" + intent.ResourceID}).Error
	}); err != nil {
		return err
	}
	if intent.Kind == "workflow" {
		if err := s.updateWorkflowApproval(ctx, p, intent.ResourceID); err != nil {
			return err
		}
	}
	if err := s.engine.ActivateNative(ctx, intent.Collection, intent.ResourceID, intent.OperationID); err != nil {
		return err
	}
	return s.db.WithContext(ctx).Model(&ResourceIntent{}).Where("operation_id = ?", intent.OperationID).Update("state", "active").Error
}
func (s *Service) RecoverIntents(ctx context.Context) error {
	rows := []ResourceIntent{}
	if err := s.db.WithContext(ctx).Where("state IN ?", []string{"prepared", "linked"}).Order("created_at").Find(&rows).Error; err != nil {
		return err
	}
	for _, intent := range rows {
		record, err := s.engine.FindNativeOperation(ctx, intent.Collection, intent.OperationID)
		if errors.Is(err, embedded.ErrNotFound) {
			continue
		}
		if err != nil {
			return err
		}
		p, err := s.CurrentPrincipal(ctx, Principal{UserID: intent.ActorID, TokenID: intent.TokenID, ServiceID: intent.ServiceID, IdentityVersion: intent.IdentityVersion})
		if err != nil {
			if err = s.db.WithContext(ctx).Model(&ResourceIntent{}).Where("operation_id = ?", intent.OperationID).Update("state", "reconciliation_required").Error; err != nil {
				return err
			}
			continue
		}
		intent.ResourceID, _ = record["id"].(string)
		if !s.Can(ctx, p, intent.Kind, "new", "create") {
			continue
		}
		if err = s.commitNativeOwnership(ctx, p, intent); err != nil {
			return err
		}
	}
	return nil
}
func (s *Service) NativeDelete(ctx context.Context, p Principal, collection, id string) error {
	if !s.nativeCan(ctx, p, collection, id, "delete") {
		return ErrDenied
	}
	if collection == "workflow_output" || collection == "workflow_logs" {
		return embedded.ErrInvalid
	}
	s.gate.Lock()
	defer s.gate.Unlock()
	if !s.nativeCan(ctx, p, collection, id, "delete") {
		return ErrDenied
	}
	if collection == "workflow_run" {
		return s.engine.NativeDelete(WithPrincipal(ctx, p), fmt.Sprint(p.UserID), collection, id)
	}
	// Revocation commits first: partial deletion must never leave an executable
	// resource whose source ownership was already removed.
	kind := collectionKind(collection)
	if err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&ResourceBinding{}).Where("kind = ? AND resource_id = ?", kind, id).Updates(map[string]any{"active": false, "version": gorm.Expr("version + 1")}).Error; err != nil {
			return err
		}
		if kind == "workflow" {
			return tx.Model(&WorkflowPolicy{}).Where("workflow_id = ?", id).Updates(map[string]any{"enabled": false, "version": gorm.Expr("version + 1")}).Error
		}
		return nil
	}); err != nil {
		return err
	}
	return s.engine.NativeDelete(WithPrincipal(ctx, p), fmt.Sprint(p.UserID), collection, id)
}
