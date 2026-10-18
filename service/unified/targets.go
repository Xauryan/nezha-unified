package unified

import (
	"context"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/certimate-go/certimate/pkg/embedded"
	"github.com/nezhahq/nezha/model"
	"github.com/nezhahq/nezha/pkg/agentproto"
	"github.com/nezhahq/nezha/service/unified/policy"
	"gorm.io/gorm"
)

type DeploymentTarget struct {
	ID                 uint64    `json:"id" gorm:"primaryKey"`
	Name               string    `json:"name"`
	Kind               string    `json:"kind"`
	ServerID           uint64    `json:"server_id" gorm:"index"`
	AllowedNames       []string  `json:"allowed_names" gorm:"serializer:json"`
	BundlePath         string    `json:"bundle_path"`
	CheckCommand       []string  `json:"check_command" gorm:"serializer:json"`
	ReloadCommand      []string  `json:"reload_command" gorm:"serializer:json"`
	HealthCommand      []string  `json:"health_command" gorm:"serializer:json"`
	TLSAddress         string    `json:"tls_address,omitempty"`
	TLSServerName      string    `json:"tls_server_name,omitempty"`
	FileMode           uint32    `json:"file_mode,omitempty"`
	OwnerUID           *int      `json:"owner_uid,omitempty"`
	OwnerGID           *int      `json:"owner_gid,omitempty"`
	Version            uint64    `json:"version"`
	Enabled            bool      `json:"enabled"`
	DesiredVersion     uint64    `json:"desired_version"`
	AppliedVersion     uint64    `json:"applied_version"`
	AppliedFingerprint string    `json:"applied_fingerprint,omitempty"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

func (DeploymentTarget) TableName() string { return "unified_deployment_targets" }
func (t DeploymentTarget) agentTarget() agentproto.Target {
	return agentproto.Target{ID: strconv.FormatUint(t.ID, 10), Version: t.Version, BundlePath: t.BundlePath, AllowedNames: t.AllowedNames, CheckCommand: t.CheckCommand, ReloadCommand: t.ReloadCommand, HealthCommand: t.HealthCommand, TLSAddress: t.TLSAddress, TLSServerName: t.TLSServerName, FileMode: t.FileMode, OwnerUID: t.OwnerUID, OwnerGID: t.OwnerGID}
}
func (s *Service) Target(ctx context.Context, p Principal, id uint64) (DeploymentTarget, error) {
	if !s.Can(ctx, p, "deployment_target", fmt.Sprint(id), "read") {
		return DeploymentTarget{}, ErrDenied
	}
	var t DeploymentTarget
	err := s.db.WithContext(ctx).First(&t, id).Error
	return t, err
}
func (s *Service) Targets(ctx context.Context, p Principal) ([]DeploymentTarget, error) {
	scope, err := s.resourceScope(ctx, p, "deployment_target")
	if err != nil {
		return nil, err
	}
	var rows []DeploymentTarget
	q := s.db.WithContext(ctx).Order("id ASC")
	if !scope.All {
		q = q.Where("id IN ?", scope.IDs)
	}
	if err = q.Find(&rows).Error; err != nil {
		return nil, err
	}
	out := []DeploymentTarget{}
	for _, t := range rows {
		if s.Can(ctx, p, "deployment_target", fmt.Sprint(t.ID), "read") {
			out = append(out, t)
		}
	}
	return out, nil
}
func (s *Service) SaveTarget(ctx context.Context, p Principal, id uint64, t DeploymentTarget) (DeploymentTarget, error) {
	p, err := s.CurrentPrincipal(ctx, p)
	if err != nil {
		return t, err
	}
	if err = requireAdmin(p); err != nil {
		return t, err
	}
	if t.Kind != "nezha" || strings.TrimSpace(t.Name) == "" || len(t.Name) > 120 || t.ServerID == 0 || !filepath.IsAbs(t.BundlePath) || filepath.Clean(t.BundlePath) == "/" || len(t.AllowedNames) == 0 || len(t.AllowedNames) > 100 {
		return t, embedded.ErrInvalid
	}
	for i, name := range t.AllowedNames {
		t.AllowedNames[i], err = policy.NormalizeName(name)
		if err != nil {
			return t, embedded.ErrInvalid
		}
	}
	for _, argv := range [][]string{t.CheckCommand, t.ReloadCommand, t.HealthCommand} {
		if len(argv) == 0 || len(argv) > 64 || !filepath.IsAbs(argv[0]) {
			return t, embedded.ErrInvalid
		}
		for _, arg := range argv {
			if len(arg) > 4096 || strings.ContainsRune(arg, 0) {
				return t, embedded.ErrInvalid
			}
		}
	}
	if (t.TLSAddress == "") != (t.TLSServerName == "") {
		return t, embedded.ErrInvalid
	}
	if t.FileMode == 0 {
		t.FileMode = 0600
	}
	if t.FileMode != 0600 && t.FileMode != 0640 {
		return t, embedded.ErrInvalid
	}
	if t.OwnerUID != nil && *t.OwnerUID < 0 || t.OwnerGID != nil && *t.OwnerGID < 0 {
		return t, embedded.ErrInvalid
	}
	var count int64
	if err = s.db.WithContext(ctx).Model(&model.Server{}).Where("id = ?", t.ServerID).Count(&count).Error; err != nil || count != 1 {
		return t, embedded.ErrInvalid
	}
	t.ID = id
	t.Name = strings.TrimSpace(t.Name)
	err = s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if id == 0 {
			t.Version = 1
			t.DesiredVersion = 0
			t.AppliedVersion = 0
			t.AppliedFingerprint = ""
			t.CreatedAt = time.Time{}
			t.UpdatedAt = time.Time{}
			if err := tx.Create(&t).Error; err != nil {
				return err
			}
			if err := tx.Create(&ResourceBinding{Kind: "deployment_target", ResourceID: fmt.Sprint(t.ID), OwnerID: p.UserID, Version: 1, Active: true}).Error; err != nil {
				return err
			}
		} else {
			var prior DeploymentTarget
			if err := tx.First(&prior, id).Error; err != nil {
				return err
			}
			if t.Version != prior.Version {
				return ErrTargetConflict
			}
			var active int64
			if err := tx.Model(&Deployment{}).Where("target_id = ? AND state IN ?", id, []string{"dispatching", "staging", "validating", "applying", "unknown", "reconciling", "reconciliation_required"}).Count(&active).Error; err != nil {
				return err
			}
			if active > 0 {
				return ErrTargetConflict
			}
			t.Version++
			t.DesiredVersion = prior.DesiredVersion
			t.AppliedVersion = prior.AppliedVersion
			t.AppliedFingerprint = prior.AppliedFingerprint
			t.CreatedAt = prior.CreatedAt
			res := tx.Model(&DeploymentTarget{}).Where("id = ? AND version = ?", id, prior.Version).Select("*").Updates(&t)
			if res.Error != nil {
				return res.Error
			}
			if res.RowsAffected != 1 {
				return ErrTargetConflict
			}
		}
		return tx.Create(&Audit{ActorID: p.UserID, Action: "target.save", Resource: fmt.Sprint(t.ID)}).Error
	})
	return t, err
}
func (s *Service) DeleteTarget(ctx context.Context, p Principal, id uint64) error {
	p, err := s.CurrentPrincipal(ctx, p)
	if err != nil {
		return err
	}
	if err = requireAdmin(p); err != nil {
		return err
	}
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var active int64
		if err := tx.Model(&Deployment{}).Where("target_id = ? AND state NOT IN ?", id, []string{"verified", "failed", "rolled_back", "cancelled", "superseded"}).Count(&active).Error; err != nil {
			return err
		}
		if active > 0 {
			return ErrTargetConflict
		}
		if err := tx.Delete(&DeploymentTarget{}, id).Error; err != nil {
			return err
		}
		if err := tx.Model(&ResourceBinding{}).Where("kind = ? AND resource_id = ?", "deployment_target", fmt.Sprint(id)).Updates(map[string]any{"active": false, "version": gorm.Expr("version + 1")}).Error; err != nil {
			return err
		}
		return tx.Create(&Audit{ActorID: p.UserID, Action: "target.delete", Resource: fmt.Sprint(id)}).Error
	})
}
func (s *Service) TargetCapabilities(ctx context.Context, p Principal, id uint64) (agentproto.Capabilities, error) {
	t, err := s.Target(ctx, p, id)
	if err != nil {
		return agentproto.Capabilities{}, err
	}
	if s.deployments == nil || s.deployments.calls == nil {
		return agentproto.Capabilities{}, nil
	}
	return s.deployments.calls.Capabilities(ctx, t.ServerID, func(ctx context.Context) error {
		if !s.Can(ctx, p, "deployment_target", fmt.Sprint(id), "read") {
			return ErrDenied
		}
		return nil
	})
}
