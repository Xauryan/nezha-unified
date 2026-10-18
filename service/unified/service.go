// Package unified coordinates Nezha-owned authorization and Certimate-owned
// business resources. PocketBase tables are never modified through GORM.
package unified

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/certimate-go/certimate/pkg/embedded"
	"github.com/nezhahq/nezha/service/unified/dnsprovider"
	"gorm.io/gorm"
)

var ErrDenied = errors.New("permission denied")

type Principal struct {
	UserID          uint64
	Admin           bool
	PAT             bool
	TokenID         uint64
	ServiceID       uint64
	IdentityVersion uint64
}

type Grant struct {
	ID         uint64    `json:"id" gorm:"primaryKey"`
	UserID     uint64    `json:"user_id" gorm:"uniqueIndex:unified_grant"`
	Kind       string    `json:"kind" gorm:"uniqueIndex:unified_grant;size:32"`
	ResourceID string    `json:"resource_id" gorm:"uniqueIndex:unified_grant;size:64"`
	Action     string    `json:"action" gorm:"uniqueIndex:unified_grant;size:32"`
	CreatedAt  time.Time `json:"created_at"`
}

func (Grant) TableName() string { return "unified_grants" }

type Audit struct {
	ID        uint64 `gorm:"primaryKey"`
	ActorID   uint64
	Action    string
	Resource  string
	CreatedAt time.Time
}

func (Audit) TableName() string { return "unified_audit" }

type Service struct {
	db          *gorm.DB
	engine      *embedded.Engine
	gate        sync.RWMutex
	recordLocks [64]sync.Mutex
	dnsFactory  func(string, map[string]any) (dnsprovider.Provider, error)
	ddns        *DDNSWorker
	deployments *deploymentRuntime
}

func New(db *gorm.DB, engine *embedded.Engine) (*Service, error) {
	if db == nil || engine == nil {
		return nil, errors.New("unified storage is required")
	}
	if err := db.AutoMigrate(&Grant{}, &Audit{}, &DNSAccount{}, &DNSZone{}, &DNSPolicy{}, &DNSOperation{}, &DDNSBinding{}, &DDNSDevice{}, &DDNSHostReport{}, &DeploymentTarget{}, &Deployment{}, &ResourceBinding{}, &WorkflowPolicy{}, &RunBinding{}, &ResourceIntent{}, &ScopedToken{}, &ServiceIdentity{}); err != nil {
		return nil, err
	}
	s := &Service{db: db, engine: engine, dnsFactory: dnsprovider.New, deployments: &deploymentRuntime{wake: make(chan struct{}, 1)}}
	s.ddns = newDDNSWorker(s)
	engine.SetCredentialAuthorizer(s.authorizeCredential)
	return s, nil
}

func validPrincipal(p Principal) bool { return p.UserID > 0 && !p.PAT }
func requireAdmin(p Principal) error {
	if !validPrincipal(p) || !p.Admin || p.TokenID != 0 || p.ServiceID != 0 {
		return ErrDenied
	}
	return nil
}

func (s *Service) Can(ctx context.Context, p Principal, kind, id, action string) bool {
	p, err := s.CurrentPrincipal(ctx, p)
	if err != nil || !s.RestrictPrincipal(ctx, p, kind, id, action) {
		return false
	}
	var source ResourceBinding
	lookup := s.db.WithContext(ctx).Where("kind = ? AND resource_id = ?", kind, id).First(&source).Error
	if lookup != nil && !errors.Is(lookup, gorm.ErrRecordNotFound) {
		return false
	}
	if lookup == nil && !source.Active && !(requireAdmin(p) == nil && (action == "read" || action == "delete")) {
		return false
	}
	if p.Admin || p.ServiceID != 0 {
		return true
	}
	var count int64
	err = s.db.WithContext(ctx).Model(&Grant{}).Where("user_id = ? AND kind = ? AND resource_id = ? AND action = ?", p.UserID, kind, id, action).Count(&count).Error
	if err == nil && count > 0 {
		return true
	}
	// Ownership grants metadata and editing; secret use, export and execution
	// remain explicit independently revocable permissions.
	if action != "read" && action != "write" && action != "delete" {
		return false
	}
	err = s.db.WithContext(ctx).Model(&ResourceBinding{}).Where("owner_id = ? AND kind = ? AND resource_id = ? AND active = ?", p.UserID, kind, id, true).Count(&count).Error
	return err == nil && count > 0
}

func (s *Service) List(ctx context.Context, p Principal, kind string, page, perPage int) (embedded.Page, error) {
	scope, err := s.resourceScope(ctx, p, kind)
	if err != nil {
		return embedded.Page{}, err
	}
	return s.engine.List(ctx, kind, scope.IDs, scope.All, page, perPage)
}

func (s *Service) CreateCredential(ctx context.Context, p Principal, in embedded.CredentialInput) (embedded.Resource, error) {
	p, err := s.CurrentPrincipal(ctx, p)
	if err != nil {
		return embedded.Resource{}, err
	}
	if err := requireAdmin(p); err != nil {
		return embedded.Resource{}, err
	}
	var config map[string]any
	if json.Unmarshal(in.Config, &config) != nil {
		return embedded.Resource{}, embedded.ErrInvalid
	}
	record, err := s.NativeSave(ctx, p, "access", "", "", map[string]any{"name": in.Name, "provider": in.Provider, "config": config})
	if err != nil {
		return embedded.Resource{}, err
	}
	id, _ := record["id"].(string)
	return s.engine.CredentialInfo(ctx, id)
}

func (s *Service) ImportCertificate(ctx context.Context, p Principal, in embedded.CertificateInput) (embedded.Resource, error) {
	p, err := s.CurrentPrincipal(ctx, p)
	if err != nil {
		return embedded.Resource{}, err
	}
	if err := requireAdmin(p); err != nil {
		return embedded.Resource{}, err
	}
	resource, err := s.engine.ImportCertificate(ctx, fmt.Sprint(p.UserID), in)
	if err != nil {
		return resource, err
	}
	return resource, s.bindImported(ctx, p, "certificate", resource.ID)
}

func (s *Service) CertificatePEM(ctx context.Context, p Principal, id string) (string, error) {
	if !s.Can(ctx, p, "certificate", id, "download") {
		return "", ErrDenied
	}
	return s.engine.CertificatePEM(ctx, id)
}

func (s *Service) ListGrants(ctx context.Context, p Principal) ([]Grant, error) {
	p, err := s.CurrentPrincipal(ctx, p)
	if err != nil || p.TokenID != 0 || p.ServiceID != 0 {
		return nil, ErrDenied
	}
	grants := []Grant{}
	q := s.db.WithContext(ctx).Order("id ASC")
	if !p.Admin {
		q = q.Where("user_id = ?", p.UserID)
	}
	err = q.Find(&grants).Error
	return grants, err
}

func (s *Service) CreateGrant(ctx context.Context, p Principal, g Grant) (Grant, error) {
	p, err := s.CurrentPrincipal(ctx, p)
	if err != nil {
		return Grant{}, err
	}
	if err := requireAdmin(p); err != nil {
		return Grant{}, err
	}
	if g.UserID == 0 || !validGrantAction(g.Kind, g.Action) || !s.resourceExists(ctx, g.Kind, g.ResourceID) {
		return Grant{}, embedded.ErrInvalid
	}
	g.ID = 0
	g.CreatedAt = time.Time{}
	err = s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&g).Error; err != nil {
			return err
		}
		return tx.Create(&Audit{ActorID: p.UserID, Action: "grant.create", Resource: fmt.Sprintf("%d:%s:%s:%s", g.UserID, g.Kind, g.ResourceID, g.Action)}).Error
	})
	return g, err
}

func (s *Service) RevokeGrant(ctx context.Context, p Principal, id uint64) error {
	p, err := s.CurrentPrincipal(ctx, p)
	if err != nil {
		return err
	}
	if err := requireAdmin(p); err != nil {
		return err
	}
	s.gate.Lock()
	defer s.gate.Unlock()
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var g Grant
		if err := tx.First(&g, id).Error; err != nil {
			return err
		}
		if err := tx.Delete(&g).Error; err != nil {
			return err
		}
		return tx.Create(&Audit{ActorID: p.UserID, Action: "grant.revoke", Resource: fmt.Sprintf("%d:%s:%s:%s", g.UserID, g.Kind, g.ResourceID, g.Action)}).Error
	})
}

func (s *Service) Close() error { return s.engine.Close() }
