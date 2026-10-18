package unified

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/nezhahq/nezha/pkg/agentproto"
	"github.com/nezhahq/nezha/service/agentcall"
	"gorm.io/gorm"
	"time"
)

func (s *Service) RetireDeployment(ctx context.Context, p Principal, id uint64, acknowledge bool, fingerprint string) (Deployment, error) {
	current, err := s.CurrentPrincipal(ctx, p)
	if err != nil {
		return Deployment{}, err
	}
	if requireAdmin(current) != nil || !acknowledge {
		return Deployment{}, ErrDenied
	}
	return s.inspectDeployment(ctx, current, id, true, fingerprint)
}

// Manual recovery authorizes the CURRENT operator, independently of the
// original job's now-revoked user, run, token, cancellation or restore state.
func (s *Service) inspectDeployment(ctx context.Context, p Principal, id uint64, retire bool, fingerprint string) (Deployment, error) {
	d, err := s.Deployment(ctx, p, id)
	if err != nil {
		return d, err
	}
	auth := func(ctx context.Context) error {
		current, err := s.CurrentPrincipal(ctx, p)
		if err != nil {
			return err
		}
		if !s.Can(ctx, current, "deployment_target", fmt.Sprint(d.TargetID), "read") || (requireAdmin(current) != nil && !s.Can(ctx, current, "deployment_target", fmt.Sprint(d.TargetID), "reconcile") && !s.Can(ctx, current, "deployment_target", fmt.Sprint(d.TargetID), "execute")) {
			return ErrDenied
		}
		if retire {
			current, err := s.CurrentPrincipal(ctx, p)
			if err != nil || requireAdmin(current) != nil {
				return ErrDenied
			}
		}
		return nil
	}
	if err = auth(ctx); err != nil {
		return d, err
	}
	if s.deployments == nil {
		return d, agentcall.ErrOffline
	}
	runtime := s.deployments
	runtime.locks[id%uint64(len(runtime.locks))].Lock()
	defer runtime.locks[id%uint64(len(runtime.locks))].Unlock()
	if err = s.db.WithContext(ctx).First(&d, id).Error; err != nil {
		return d, err
	}
	if retire && (fingerprint != d.Fingerprint || d.State != "unknown" && d.State != "reconciliation_required") {
		return d, ErrTargetConflict
	}
	var target DeploymentTarget
	if json.Unmarshal([]byte(d.TargetJSON), &target) != nil {
		return d, ErrTargetConflict
	}
	if target.Kind == "certimate" {
		if !retire {
			return d, ErrTargetConflict
		}
		if err = auth(ctx); err != nil {
			return d, err
		}
		err = s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
			result := tx.Model(&Deployment{}).Where("id = ? AND state = ? AND fingerprint = ?", id, d.State, fingerprint).Updates(map[string]any{"state": "retired", "code": "operator_retired_unknown", "lease_until": time.Time{}, "inspected_at": time.Now()})
			if result.Error != nil {
				return result.Error
			}
			if result.RowsAffected != 1 {
				return ErrTargetConflict
			}
			return tx.Create(&Audit{ActorID: p.UserID, Action: "deployment.retire", Resource: fmt.Sprint(id)}).Error
		})
		if err != nil {
			return d, err
		}
		return s.Deployment(ctx, p, id)
	}
	if runtime.calls == nil {
		return d, agentcall.ErrOffline
	}
	method := "inspect"
	if retire {
		method = "retire"
	}
	payload, _ := json.Marshal(agentproto.Deployment{OperationID: d.OperationID, Sequence: d.ID, Target: target.agentTarget(), Fingerprint: d.Fingerprint})
	raw, err := runtime.calls.Call(ctx, target.ServerID, agentproto.TaskCall, "certificate_inspect", agentproto.Request{Protocol: agentproto.Protocol, OperationID: d.OperationID, Method: method, Payload: payload}, auth, nil)
	if err != nil {
		return d, err
	}
	var receipt agentproto.Receipt
	if json.Unmarshal(raw, &receipt) != nil || !validDeploymentReceipt(&receipt, d) || !receipt.Inspected {
		return d, agentcall.ErrProtocol
	}
	state := "reconciliation_required"
	switch receipt.Phase {
	case "verified":
		if receipt.Verified {
			state = "verified"
		}
	case "failed":
		state = "failed"
	case "rolled_back":
		if receipt.RolledBack {
			state = "rolled_back"
		}
	case "not_applied":
		state = "cancelled"
	case "retired":
		state = "retired"
	}
	if err = auth(ctx); err != nil {
		return d, err
	}
	err = s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		query := tx.Model(&Deployment{}).Where("id = ?", d.ID)
		if retire {
			query = query.Where("state = ? AND fingerprint = ?", d.State, fingerprint)
		}
		result := query.Updates(map[string]any{"state": state, "code": safeDeploymentCode(receipt.Code), "receipt": mustDeploymentJSON(receipt), "lease_until": time.Time{}, "inspected_at": time.Now()})
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected != 1 {
			return ErrTargetConflict
		}
		if state == "verified" && receipt.CurrentOperationID == d.OperationID && receipt.CurrentPairValid && receipt.CurrentFingerprint == d.Fingerprint {
			if err := tx.Model(&DeploymentTarget{}).Where("id = ? AND version = ? AND applied_version < ?", d.TargetID, d.TargetVersion, d.ID).Updates(map[string]any{"applied_version": d.ID, "applied_fingerprint": d.Fingerprint}).Error; err != nil {
				return err
			}
		}
		return tx.Create(&Audit{ActorID: p.UserID, Action: "deployment." + method, Resource: fmt.Sprint(d.ID)}).Error
	})
	if err != nil {
		return d, err
	}
	return s.Deployment(ctx, p, id)
}
func safeDeploymentCode(code string) string {
	switch code {
	case "", "stage", "validation", "authorization", "apply", "rollback", "rollback_reload", "rollback_health", "first_install_no_restore", "operator_retired_unknown", "installed_pair", "permissions", "ownership", "cancelled":
		return code
	default:
		return "remote_result"
	}
}

func validDeploymentReceipt(receipt *agentproto.Receipt, d Deployment) bool {
	if receipt.OperationID != d.OperationID || receipt.TargetID != fmt.Sprint(d.TargetID) || receipt.TargetVersion != d.TargetVersion || receipt.Sequence != d.ID || receipt.Fingerprint != d.Fingerprint {
		return false
	}
	if receipt.CurrentOperationID != "" && !deploymentKey.MatchString(receipt.CurrentOperationID) {
		return false
	}
	if receipt.CurrentFingerprint != "" {
		if len(receipt.CurrentFingerprint) != 64 {
			return false
		}
		for _, c := range receipt.CurrentFingerprint {
			if !(c >= '0' && c <= '9' || c >= 'a' && c <= 'f') {
				return false
			}
		}
	}
	switch receipt.Phase {
	case "staging", "validating", "applying", "unknown", "verified", "failed", "rolled_back", "not_found", "not_applied", "retired", "applied":
	default:
		return false
	}
	receipt.Code = safeDeploymentCode(receipt.Code)
	return true
}
