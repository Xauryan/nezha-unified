package unified

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/certimate-go/certimate/pkg/embedded"
	"github.com/nezhahq/nezha/model"
	"github.com/nezhahq/nezha/service/unified/policy"
	"gorm.io/gorm"
)

// ResourceBinding is the sole owner authority for cross-module business objects.
// Certimate stores only the stable reference and activation/operation marker.
type ResourceBinding struct {
	ID          uint64    `json:"id" gorm:"primaryKey"`
	Kind        string    `json:"kind" gorm:"uniqueIndex:resource_owner;size:32"`
	ResourceID  string    `json:"resource_id" gorm:"uniqueIndex:resource_owner;size:64"`
	OwnerID     uint64    `json:"owner_id" gorm:"index"`
	Active      bool      `json:"active"`
	Version     uint64    `json:"version"`
	OperationID string    `json:"operation_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (ResourceBinding) TableName() string { return "unified_resource_bindings" }

type WorkflowPolicy struct {
	WorkflowID          string             `json:"workflow_id" gorm:"primaryKey;size:64"`
	Version             uint64             `json:"version"`
	RunAsUserID         uint64             `json:"run_as_user_id"`
	ServiceID           uint64             `json:"service_id"`
	IdentityVersion     uint64             `json:"identity_version"`
	Enabled             bool               `json:"enabled"`
	AllowInline         bool               `json:"allow_inline"`
	AllowImport         bool               `json:"allow_import"`
	CertificateRead     bool               `json:"certificate_read"`
	CertificateDownload bool               `json:"certificate_download"`
	GraphDigest         string             `json:"graph_digest"`
	Names               []policy.NameScope `json:"names" gorm:"serializer:json"`
	Credentials         map[string]int     `json:"credentials" gorm:"serializer:json"`
	Targets             map[string]uint64  `json:"targets" gorm:"serializer:json"`
	UpdatedAt           time.Time          `json:"updated_at"`
}

func (WorkflowPolicy) TableName() string { return "unified_workflow_policies" }

type RunBinding struct {
	RunID           string    `json:"run_id" gorm:"primaryKey;size:64"`
	WorkflowID      string    `json:"workflow_id" gorm:"index"`
	PolicyVersion   uint64    `json:"policy_version"`
	OwnerVersion    uint64    `json:"owner_version"`
	ActorID         uint64    `json:"actor_id"`
	TokenID         uint64    `json:"token_id"`
	ServiceID       uint64    `json:"service_id"`
	IdentityVersion uint64    `json:"identity_version"`
	GraphDigest     string    `json:"graph_digest"`
	Revoked         bool      `json:"revoked"`
	CreatedAt       time.Time `json:"created_at"`
}

func (RunBinding) TableName() string { return "unified_run_bindings" }

type ResourceIntent struct {
	OperationID     string    `json:"operation_id" gorm:"primaryKey;size:64"`
	ActorID         uint64    `json:"actor_id"`
	TokenID         uint64    `json:"token_id"`
	ServiceID       uint64    `json:"service_id"`
	IdentityVersion uint64    `json:"identity_version"`
	Kind            string    `json:"kind"`
	Collection      string    `json:"collection"`
	ResourceID      string    `json:"resource_id"`
	Digest          string    `json:"-"`
	State           string    `json:"state"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func (ResourceIntent) TableName() string { return "unified_resource_intents" }

type principalKey struct{}

func WithPrincipal(ctx context.Context, p Principal) context.Context {
	return context.WithValue(ctx, principalKey{}, p)
}
func principalFromContext(ctx context.Context) (Principal, bool) {
	p, ok := ctx.Value(principalKey{}).(Principal)
	return p, ok
}

// CurrentPrincipal deliberately reads the live account and identity versions.
// Persisted asynchronous requests do not retain administrator authority in memory.
func (s *Service) CurrentPrincipal(ctx context.Context, p Principal) (Principal, error) {
	if !validPrincipal(p) {
		return Principal{}, ErrDenied
	}
	var u model.User
	if err := s.db.WithContext(ctx).First(&u, p.UserID).Error; err != nil {
		return Principal{}, ErrDenied
	}
	p.Admin = u.Role.IsAdmin() && p.ServiceID == 0
	if err := s.ValidateIdentity(ctx, p); err != nil {
		return Principal{}, ErrDenied
	}
	return p, nil
}
func validGrantAction(kind, action string) bool {
	allowed := map[string]string{
		"credential":        "read create write delete use export rotate clone",
		"certificate":       "read create write delete use download download_key revoke",
		"workflow":          "read create write delete execute cancel retry events",
		"deployment_target": "read create write delete execute use",
		"deployment_task":   "read cancel retry",
		"dns_account":       "read create write delete", "dns_zone": "read create write delete execute", "ddns_binding": "read create write delete execute", "settings": "read write", "system": "read",
	}
	return strings.Contains(" "+allowed[kind]+" ", " "+action+" ") && action != ""
}
func (s *Service) resourceExists(ctx context.Context, kind, id string) bool {
	if id == "new" {
		return kind == "workflow" || kind == "credential" || kind == "certificate" || kind == "deployment_target" || kind == "ddns_binding"
	}
	if kind == "credential" || kind == "certificate" || kind == "workflow" {
		return s.engine.Exists(ctx, kind, id)
	}
	var count int64
	table := map[string]string{"deployment_target": "unified_deployment_targets", "deployment_task": "unified_deployments", "dns_account": "unified_dns_accounts", "dns_zone": "unified_dns_zones", "ddns_binding": "unified_ddns_bindings"}[kind]
	return table != "" && s.db.WithContext(ctx).Table(table).Where("id = ?", id).Count(&count).Error == nil && count == 1
}
func (s *Service) Ownership(ctx context.Context, p Principal) ([]ResourceBinding, error) {
	p, err := s.CurrentPrincipal(ctx, p)
	if err != nil || p.TokenID != 0 || p.ServiceID != 0 {
		return nil, ErrDenied
	}
	rows := []ResourceBinding{}
	q := s.db.WithContext(ctx).Order("kind, resource_id")
	if requireAdmin(p) != nil {
		q = q.Where("owner_id = ?", p.UserID)
	}
	err = q.Find(&rows).Error
	return rows, err
}
func (s *Service) TransferOwnership(ctx context.Context, p Principal, kind, id string, owner, expected uint64) (ResourceBinding, error) {
	p, err := s.CurrentPrincipal(ctx, p)
	if err != nil {
		return ResourceBinding{}, err
	}
	if err := requireAdmin(p); err != nil {
		return ResourceBinding{}, err
	}
	s.gate.Lock()
	defer s.gate.Unlock()
	var user model.User
	if s.db.WithContext(ctx).First(&user, owner).Error != nil {
		return ResourceBinding{}, embedded.ErrInvalid
	}
	var binding ResourceBinding
	err = s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("kind = ? AND resource_id = ?", kind, id).First(&binding).Error; err != nil {
			return err
		}
		if expected != binding.Version {
			return embedded.ErrInvalid
		}
		change := tx.Model(&ResourceBinding{}).Where("id = ? AND version = ?", binding.ID, expected).Updates(map[string]any{"owner_id": owner, "version": expected + 1})
		if change.Error != nil {
			return change.Error
		}
		if change.RowsAffected != 1 {
			return embedded.ErrInvalid
		}
		if kind == "workflow" {
			if err := tx.Model(&WorkflowPolicy{}).Where("workflow_id = ?", id).Updates(map[string]any{"enabled": false, "version": gorm.Expr("version + 1")}).Error; err != nil {
				return err
			}
		}
		return tx.Create(&Audit{ActorID: p.UserID, Action: "ownership.transfer", Resource: fmt.Sprintf("%s:%s:%d", kind, id, owner)}).Error
	})
	binding.OwnerID = owner
	binding.Version = expected + 1
	return binding, err
}
func (s *Service) WorkflowPolicies(ctx context.Context, p Principal) ([]WorkflowPolicy, error) {
	if !validPrincipal(p) {
		return nil, ErrDenied
	}
	rows := []WorkflowPolicy{}
	if err := s.db.WithContext(ctx).Order("workflow_id").Find(&rows).Error; err != nil {
		return nil, err
	}
	out := make([]WorkflowPolicy, 0, len(rows))
	for _, row := range rows {
		if s.Can(ctx, p, "workflow", row.WorkflowID, "read") {
			out = append(out, row)
		}
	}
	return out, nil
}
func (s *Service) SaveWorkflowPolicy(ctx context.Context, p Principal, id string, expected uint64, in WorkflowPolicy) (WorkflowPolicy, error) {
	p, err := s.CurrentPrincipal(ctx, p)
	if err != nil {
		return WorkflowPolicy{}, err
	}
	if err := requireAdmin(p); err != nil {
		return WorkflowPolicy{}, err
	}
	s.gate.Lock()
	defer s.gate.Unlock()
	if !s.engine.Exists(ctx, "workflow", id) {
		return WorkflowPolicy{}, embedded.ErrNotFound
	}
	actor, err := s.CurrentPrincipal(ctx, Principal{UserID: in.RunAsUserID, ServiceID: in.ServiceID, IdentityVersion: in.IdentityVersion})
	if err != nil {
		return WorkflowPolicy{}, err
	}
	if in.Enabled && !s.Can(ctx, actor, "workflow", id, "execute") {
		return WorkflowPolicy{}, ErrDenied
	}
	for i := range in.Names {
		root, err := policy.NormalizeName(in.Names[i].Root)
		if err != nil || strings.Contains(root, "*") {
			return WorkflowPolicy{}, embedded.ErrInvalid
		}
		in.Names[i].Root = root
	}
	for credential, version := range in.Credentials {
		info, err := s.engine.CredentialInfo(ctx, credential)
		if err != nil || info.Version != version || !info.Enabled {
			return WorkflowPolicy{}, embedded.ErrInvalid
		}
	}
	for target, version := range in.Targets {
		if version == 0 || !s.resourceExists(ctx, "deployment_target", target) {
			return WorkflowPolicy{}, embedded.ErrInvalid
		}
	}
	digest, err := s.engine.WorkflowGraphDigest(ctx, id)
	if err != nil {
		return WorkflowPolicy{}, err
	}
	in.GraphDigest = digest
	in.WorkflowID = id
	in.Version = expected + 1
	err = s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var old WorkflowPolicy
		err := tx.First(&old, "workflow_id = ?", id).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if expected != 0 {
				return embedded.ErrInvalid
			}
			if err = tx.Create(&in).Error; err != nil {
				return err
			}
		} else if err != nil {
			return err
		} else {
			if old.Version != expected {
				return embedded.ErrInvalid
			}
			if err = tx.Save(&in).Error; err != nil {
				return err
			}
		}
		return tx.Create(&Audit{ActorID: p.UserID, Action: "workflow.policy", Resource: fmt.Sprintf("%s:%d", id, in.Version)}).Error
	})
	return in, err
}
