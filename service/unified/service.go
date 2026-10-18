// Package unified coordinates Nezha-owned authorization and Certimate-owned
// business resources. PocketBase tables are never modified through GORM.
package unified

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/certimate-go/certimate/pkg/embedded"
	"gorm.io/gorm"
)

var ErrDenied = errors.New("permission denied")

type Principal struct {
	UserID uint64
	Admin  bool
	PAT    bool
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
	db     *gorm.DB
	engine *embedded.Engine
}

func New(db *gorm.DB, engine *embedded.Engine) (*Service, error) {
	if db == nil || engine == nil {
		return nil, errors.New("unified storage is required")
	}
	if err := db.AutoMigrate(&Grant{}, &Audit{}); err != nil {
		return nil, err
	}
	return &Service{db: db, engine: engine}, nil
}

func validPrincipal(p Principal) bool { return p.UserID > 0 && !p.PAT }
func requireAdmin(p Principal) error {
	if !validPrincipal(p) || !p.Admin {
		return ErrDenied
	}
	return nil
}

func (s *Service) Can(ctx context.Context, p Principal, kind, id, action string) bool {
	if !validPrincipal(p) {
		return false
	}
	if p.Admin {
		return true
	}
	var count int64
	err := s.db.WithContext(ctx).Model(&Grant{}).Where("user_id = ? AND kind = ? AND resource_id = ? AND action = ?", p.UserID, kind, id, action).Count(&count).Error
	return err == nil && count > 0
}

func (s *Service) List(ctx context.Context, p Principal, kind string, page, perPage int) (embedded.Page, error) {
	if !validPrincipal(p) {
		return embedded.Page{}, ErrDenied
	}
	var ids []string
	if !p.Admin {
		if err := s.db.WithContext(ctx).Model(&Grant{}).Where("user_id = ? AND kind = ? AND action = ?", p.UserID, kind, "read").Pluck("resource_id", &ids).Error; err != nil {
			return embedded.Page{}, err
		}
	}
	return s.engine.List(ctx, kind, ids, p.Admin, page, perPage)
}

func (s *Service) CreateCredential(ctx context.Context, p Principal, in embedded.CredentialInput) (embedded.Resource, error) {
	if err := requireAdmin(p); err != nil {
		return embedded.Resource{}, err
	}
	return s.engine.CreateCredential(ctx, fmt.Sprint(p.UserID), in)
}

func (s *Service) ImportCertificate(ctx context.Context, p Principal, in embedded.CertificateInput) (embedded.Resource, error) {
	if err := requireAdmin(p); err != nil {
		return embedded.Resource{}, err
	}
	return s.engine.ImportCertificate(ctx, fmt.Sprint(p.UserID), in)
}

func (s *Service) CertificatePEM(ctx context.Context, p Principal, id string) (string, error) {
	if !s.Can(ctx, p, "certificate", id, "download") {
		return "", ErrDenied
	}
	return s.engine.CertificatePEM(ctx, id)
}

func (s *Service) ListGrants(ctx context.Context, p Principal) ([]Grant, error) {
	if !validPrincipal(p) {
		return nil, ErrDenied
	}
	grants := []Grant{}
	q := s.db.WithContext(ctx).Order("id ASC")
	if !p.Admin {
		q = q.Where("user_id = ?", p.UserID)
	}
	err := q.Find(&grants).Error
	return grants, err
}

func (s *Service) CreateGrant(ctx context.Context, p Principal, g Grant) (Grant, error) {
	if err := requireAdmin(p); err != nil {
		return Grant{}, err
	}
	if g.UserID == 0 || !(g.Action == "read" || g.Kind == "certificate" && g.Action == "download") || !s.engine.Exists(ctx, g.Kind, g.ResourceID) {
		return Grant{}, embedded.ErrInvalid
	}
	g.ID = 0
	g.CreatedAt = time.Time{}
	err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&g).Error; err != nil {
			return err
		}
		return tx.Create(&Audit{ActorID: p.UserID, Action: "grant.create", Resource: fmt.Sprintf("%d:%s:%s:%s", g.UserID, g.Kind, g.ResourceID, g.Action)}).Error
	})
	return g, err
}

func (s *Service) RevokeGrant(ctx context.Context, p Principal, id uint64) error {
	if err := requireAdmin(p); err != nil {
		return err
	}
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
