package unified

import (
	"context"
	"encoding/json"
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
	ID                    uint64         `json:"id" gorm:"primaryKey"`
	Name                  string         `json:"name"`
	Kind                  string         `json:"kind"`
	Provider              string         `json:"provider,omitempty"`
	CredentialID          string         `json:"credential_id,omitempty"`
	CredentialVersion     int            `json:"credential_version,omitempty"`
	ConfigurationEnvelope string         `json:"-"`
	Config                map[string]any `json:"config,omitempty" gorm:"-"`
	Configured            bool           `json:"configured" gorm:"-"`
	ServerID              uint64         `json:"server_id" gorm:"index"`
	AllowedNames          []string       `json:"allowed_names" gorm:"serializer:json"`
	BundlePath            string         `json:"bundle_path"`
	CheckCommand          []string       `json:"check_command" gorm:"serializer:json"`
	ReloadCommand         []string       `json:"reload_command" gorm:"serializer:json"`
	HealthCommand         []string       `json:"health_command" gorm:"serializer:json"`
	TLSAddress            string         `json:"tls_address,omitempty"`
	TLSServerName         string         `json:"tls_server_name,omitempty"`
	FileMode              uint32         `json:"file_mode,omitempty"`
	OwnerUID              *int           `json:"owner_uid,omitempty"`
	OwnerGID              *int           `json:"owner_gid,omitempty"`
	Version               uint64         `json:"version"`
	Enabled               bool           `json:"enabled"`
	DesiredVersion        uint64         `json:"desired_version"`
	AppliedVersion        uint64         `json:"applied_version"`
	AppliedFingerprint    string         `json:"applied_fingerprint,omitempty"`
	CreatedAt             time.Time      `json:"created_at"`
	UpdatedAt             time.Time      `json:"updated_at"`
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
	t.Configured = t.ConfigurationEnvelope != ""
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
			t.Configured = t.ConfigurationEnvelope != ""
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
	if (t.Kind != "nezha" && t.Kind != "certimate") || strings.TrimSpace(t.Name) == "" || len(t.Name) > 120 || len(t.AllowedNames) == 0 || len(t.AllowedNames) > 100 {
		return DeploymentTarget{}, embedded.ErrInvalid
	}
	for i, name := range t.AllowedNames {
		t.AllowedNames[i], err = policy.NormalizeName(name)
		if err != nil {
			return DeploymentTarget{}, embedded.ErrInvalid
		}
	}
	if t.Kind == "nezha" {
		if err = s.validateAgentTarget(ctx, t); err != nil {
			return DeploymentTarget{}, err
		}
		t.Provider, t.CredentialID, t.CredentialVersion, t.ConfigurationEnvelope, t.Config = "", "", 0, "", nil
		if t.FileMode == 0 {
			t.FileMode = 0600
		}
	} else {
		if t.Provider == "" || len(t.Provider) > 100 || t.Provider == "nezha-agent" || t.CredentialVersion < 1 || !s.Can(ctx, p, "credential", t.CredentialID, "use") {
			return DeploymentTarget{}, embedded.ErrInvalid
		}
		info, infoErr := s.engine.CredentialInfo(ctx, t.CredentialID)
		if infoErr != nil || !info.Enabled || info.Version != t.CredentialVersion {
			return DeploymentTarget{}, embedded.ErrInvalid
		}
		if t.Config != nil {
			if encoded, e := json.Marshal(t.Config); e != nil || len(encoded) > 2<<20 {
				return DeploymentTarget{}, embedded.ErrInvalid
			}
		}
		t.ServerID, t.BundlePath, t.TLSAddress, t.TLSServerName = 0, "", "", ""
		t.CheckCommand, t.ReloadCommand, t.HealthCommand, t.OwnerUID, t.OwnerGID = nil, nil, nil, nil, nil
		t.FileMode = 0
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
			if t.Kind == "certimate" {
				envelope, e := s.engine.SealTargetConfiguration(ctx, fmt.Sprint(t.ID), t.Version, t.Config)
				if e != nil {
					return e
				}
				t.ConfigurationEnvelope = envelope
				if e = tx.Model(&DeploymentTarget{}).Where("id = ?", t.ID).Update("configuration_envelope", envelope).Error; e != nil {
					return e
				}
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
			if t.Kind == "certimate" {
				config := t.Config
				if config == nil {
					if prior.Kind != t.Kind || prior.Provider != t.Provider || prior.CredentialID != t.CredentialID {
						return embedded.ErrInvalid
					}
					var e error
					config, e = s.engine.OpenTargetConfiguration(ctx, fmt.Sprint(prior.ID), prior.Version, prior.ConfigurationEnvelope)
					if e != nil {
						return e
					}
				}
				envelope, e := s.engine.SealTargetConfiguration(ctx, fmt.Sprint(t.ID), t.Version, config)
				if e != nil {
					return e
				}
				t.ConfigurationEnvelope = envelope
			}
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
	t.Configured = t.ConfigurationEnvelope != ""
	t.Config = nil
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
		if err := tx.Model(&Deployment{}).Where("target_id = ? AND state NOT IN ?", id, []string{"verified", "applied", "failed", "rolled_back", "cancelled", "superseded", "retired"}).Count(&active).Error; err != nil {
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
	if t.Kind != "nezha" || s.deployments == nil || s.deployments.calls == nil {
		return agentproto.Capabilities{}, nil
	}
	return s.deployments.calls.Capabilities(ctx, t.ServerID, func(ctx context.Context) error {
		if !s.Can(ctx, p, "deployment_target", fmt.Sprint(id), "read") {
			return ErrDenied
		}
		return nil
	})
}

func (s *Service) validateAgentTarget(ctx context.Context, t DeploymentTarget) error {
	if t.ServerID == 0 || !filepath.IsAbs(t.BundlePath) || filepath.Clean(t.BundlePath) == "/" {
		return embedded.ErrInvalid
	}
	for _, argv := range [][]string{t.CheckCommand, t.ReloadCommand, t.HealthCommand} {
		if len(argv) == 0 || len(argv) > 64 || !filepath.IsAbs(argv[0]) {
			return embedded.ErrInvalid
		}
		for _, arg := range argv {
			if len(arg) > 4096 || strings.ContainsRune(arg, 0) {
				return embedded.ErrInvalid
			}
		}
	}
	if (t.TLSAddress == "") != (t.TLSServerName == "") || (t.FileMode != 0 && t.FileMode != 0600 && t.FileMode != 0640) || t.OwnerUID != nil && *t.OwnerUID < 0 || t.OwnerGID != nil && *t.OwnerGID < 0 {
		return embedded.ErrInvalid
	}
	var count int64
	if err := s.db.WithContext(ctx).Model(&model.Server{}).Where("id = ?", t.ServerID).Count(&count).Error; err != nil {
		return err
	}
	if count != 1 {
		return embedded.ErrInvalid
	}
	return nil
}
