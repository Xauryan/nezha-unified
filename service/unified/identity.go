package unified

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
	"unicode"

	"github.com/certimate-go/certimate/pkg/embedded"
	"github.com/nezhahq/nezha/model"
	"gorm.io/gorm"
)

// Scope is an exact resource/action pair. Wildcards are deliberately unsupported;
// callers must enumerate each authorized resource, even for administrator tokens.
type Scope struct {
	Kind       string `json:"kind"`
	ResourceID string `json:"resource_id"`
	Action     string `json:"action"`
}

const ScopedTokenPrefix = "nzu_"

type ScopedToken struct {
	ID          uint64     `json:"id" gorm:"primaryKey"`
	UserID      uint64     `json:"user_id" gorm:"index;not null"`
	Name        string     `json:"name" gorm:"size:128;not null"`
	PublicID    string     `json:"-" gorm:"uniqueIndex;size:32;not null"`
	SecretHash  []byte     `json:"-" gorm:"not null"`
	UserVersion uint64     `json:"-" gorm:"not null"`
	Scopes      []Scope    `json:"scopes" gorm:"serializer:json;type:text;not null"`
	CreatedAt   time.Time  `json:"created_at"`
	ExpiresAt   time.Time  `json:"expires_at" gorm:"not null"`
	RevokedAt   *time.Time `json:"revoked_at,omitempty"`
}

func (ScopedToken) TableName() string { return "unified_scoped_tokens" }

type ScopedTokenInput struct {
	Name      string    `json:"name"`
	Scopes    []Scope   `json:"scopes"`
	ExpiresAt time.Time `json:"expires_at"`
}

// IssuedScopedToken is the only response that contains the bearer secret. The
// secret is never persisted and is not reconstructed by list or detail methods.
type IssuedScopedToken struct {
	Token  ScopedToken `json:"token"`
	Secret string      `json:"secret"`
}

type ServiceIdentity struct {
	ID        uint64    `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"size:128;not null"`
	OwnerID   uint64    `json:"owner_id" gorm:"index;not null"`
	Active    bool      `json:"active" gorm:"not null"`
	Version   uint64    `json:"version" gorm:"not null"`
	Scopes    []Scope   `json:"scopes" gorm:"serializer:json;type:text;not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (ServiceIdentity) TableName() string { return "unified_service_identities" }

type ServiceIdentityInput struct {
	Name    string  `json:"name"`
	OwnerID uint64  `json:"owner_id"`
	Active  *bool   `json:"active,omitempty"`
	Scopes  []Scope `json:"scopes"`
}

// ServiceIdentityUpdate is a partial update. Version is required so concurrent
// transfer, deactivation and policy edits cannot overwrite one another.
type ServiceIdentityUpdate struct {
	Version uint64   `json:"version"`
	Name    *string  `json:"name,omitempty"`
	OwnerID *uint64  `json:"owner_id,omitempty"`
	Active  *bool    `json:"active,omitempty"`
	Scopes  *[]Scope `json:"scopes,omitempty"`
}

func validateScopes(scopes []Scope) error {
	if len(scopes) == 0 || len(scopes) > 128 {
		return embedded.ErrInvalid
	}
	kinds := map[string]bool{"credential": true, "certificate": true, "workflow": true, "workflow_run": true, "deployment_target": true, "deployment_task": true, "dns_account": true, "dns_zone": true, "dns_record": true, "dns_policy": true, "ddns_binding": true, "dns_operation": true, "settings": true, "system": true}
	actions := map[string]bool{"read": true, "create": true, "write": true, "delete": true, "use": true, "download": true, "download_key": true, "export": true, "rotate": true, "revoke": true, "clone": true, "execute": true, "cancel": true, "retry": true, "events": true, "deploy": true}
	seen := make(map[Scope]bool, len(scopes))
	for _, scope := range scopes {
		if !kinds[scope.Kind] || !actions[scope.Action] || scope.ResourceID == "" || len(scope.ResourceID) > 128 || strings.ContainsAny(scope.ResourceID, "*\x00") || strings.IndexFunc(scope.ResourceID, unicode.IsSpace) >= 0 || strings.IndexFunc(scope.ResourceID, unicode.IsControl) >= 0 || seen[scope] {
			return embedded.ErrInvalid
		}
		if scope.ResourceID == "new" && scope.Action != "create" {
			return embedded.ErrInvalid
		}
		seen[scope] = true
	}
	return nil
}

func validIdentityName(name string) bool {
	return strings.TrimSpace(name) == name && name != "" && len(name) <= 128 && strings.IndexFunc(name, unicode.IsControl) == -1
}

func identityScopeAllows(scopes []Scope, kind, id, action string) bool {
	for _, scope := range scopes {
		if scope.Kind == kind && scope.ResourceID == id && scope.Action == action {
			return true
		}
	}
	return false
}

func scopeJSON(scopes []Scope) string {
	encoded, _ := json.Marshal(scopes)
	return string(encoded)
}

// identitySession rejects both kinds of restricted identity before considering
// administrator status. Token management can never mint a less restricted token.
func (s *Service) identitySession(ctx context.Context, p Principal, admin bool) (model.User, error) {
	var user model.User
	if !validPrincipal(p) || p.TokenID != 0 || p.ServiceID != 0 || s.db.WithContext(ctx).First(&user, p.UserID).Error != nil || admin && !user.Role.IsAdmin() {
		return model.User{}, ErrDenied
	}
	return user, nil
}

func (s *Service) ListScopedTokens(ctx context.Context, p Principal) ([]ScopedToken, error) {
	if _, err := s.identitySession(ctx, p, false); err != nil {
		return nil, err
	}
	rows := []ScopedToken{}
	err := s.db.WithContext(ctx).Omit("secret_hash", "public_id", "user_version").Where("user_id = ?", p.UserID).Order("id DESC").Find(&rows).Error
	return rows, err
}

func (s *Service) CreateScopedToken(ctx context.Context, p Principal, in ScopedTokenInput) (IssuedScopedToken, error) {
	user, err := s.identitySession(ctx, p, false)
	if err != nil {
		return IssuedScopedToken{}, err
	}
	now := time.Now().UTC()
	if !validIdentityName(in.Name) || validateScopes(in.Scopes) != nil || !in.ExpiresAt.After(now) || in.ExpiresAt.After(now.Add(366*24*time.Hour)) {
		return IssuedScopedToken{}, embedded.ErrInvalid
	}
	p.Admin = user.Role.IsAdmin()
	for _, scope := range in.Scopes {
		if !s.Can(ctx, p, scope.Kind, scope.ResourceID, scope.Action) {
			return IssuedScopedToken{}, ErrDenied
		}
	}
	idBytes := make([]byte, 16)
	secretBytes := make([]byte, 32)
	if _, err := rand.Read(idBytes); err != nil {
		return IssuedScopedToken{}, err
	}
	if _, err := rand.Read(secretBytes); err != nil {
		return IssuedScopedToken{}, err
	}
	publicID := base64.RawURLEncoding.EncodeToString(idBytes)
	secretPart := base64.RawURLEncoding.EncodeToString(secretBytes)
	digest := sha256.Sum256([]byte(secretPart))
	row := ScopedToken{UserID: user.ID, UserVersion: user.TokenVersion, Name: in.Name, PublicID: publicID, SecretHash: digest[:], Scopes: append([]Scope(nil), in.Scopes...), CreatedAt: now, ExpiresAt: in.ExpiresAt.UTC()}
	err = s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&row).Error; err != nil {
			return err
		}
		return tx.Create(&Audit{ActorID: user.ID, Action: "token.create", Resource: fmt.Sprint(row.ID)}).Error
	})
	if err != nil {
		return IssuedScopedToken{}, err
	}
	row.SecretHash, row.PublicID, row.UserVersion = nil, "", 0
	return IssuedScopedToken{Token: row, Secret: ScopedTokenPrefix + publicID + "." + secretPart}, nil
}

func (s *Service) RevokeScopedToken(ctx context.Context, p Principal, id uint64) error {
	if _, err := s.identitySession(ctx, p, false); err != nil {
		return err
	}
	if id == 0 {
		return embedded.ErrInvalid
	}
	s.gate.Lock()
	defer s.gate.Unlock()
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var token ScopedToken
		if err := tx.Where("id = ? AND user_id = ?", id, p.UserID).First(&token).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return embedded.ErrNotFound
			}
			return err
		}
		if token.RevokedAt != nil {
			return nil
		}
		now := time.Now().UTC()
		if err := tx.Model(&token).Update("revoked_at", now).Error; err != nil {
			return err
		}
		return tx.Create(&Audit{ActorID: p.UserID, Action: "token.revoke", Resource: fmt.Sprint(id)}).Error
	})
}

// AuthenticateScopedToken indexes by the nonsensitive random public identifier,
// then compares fixed length SHA-256 hashes in constant time. The bearer secret
// never appears in a SQL query, an audit record, or a returned error.
func (s *Service) AuthenticateScopedToken(ctx context.Context, raw string) (Principal, error) {
	if !strings.HasPrefix(raw, ScopedTokenPrefix) || len(raw) != len(ScopedTokenPrefix)+22+1+43 {
		return Principal{}, ErrDenied
	}
	publicID, secret, ok := strings.Cut(strings.TrimPrefix(raw, ScopedTokenPrefix), ".")
	if !ok || len(publicID) != 22 || len(secret) != 43 {
		return Principal{}, ErrDenied
	}
	if value, err := base64.RawURLEncoding.DecodeString(publicID); err != nil || len(value) != 16 {
		return Principal{}, ErrDenied
	}
	if value, err := base64.RawURLEncoding.DecodeString(secret); err != nil || len(value) != 32 {
		return Principal{}, ErrDenied
	}
	var row ScopedToken
	err := s.db.WithContext(ctx).Where("public_id = ?", publicID).First(&row).Error
	digest := sha256.Sum256([]byte(secret))
	expected := make([]byte, sha256.Size)
	if len(row.SecretHash) == sha256.Size {
		copy(expected, row.SecretHash)
	}
	matched := subtle.ConstantTimeCompare(expected, digest[:])
	if err != nil || len(row.SecretHash) != sha256.Size || matched != 1 {
		return Principal{}, ErrDenied
	}
	p := Principal{UserID: row.UserID, TokenID: row.ID}
	if err := s.ValidateIdentity(ctx, p); err != nil {
		return Principal{}, ErrDenied
	}
	var user model.User
	if err := s.db.WithContext(ctx).First(&user, p.UserID).Error; err != nil {
		return Principal{}, ErrDenied
	}
	p.Admin = user.Role.IsAdmin()
	return p, nil
}

// ValidateIdentity is re-run at authorization and execution boundaries. A cached
// principal cannot outlive revocation, expiry, owner deletion or policy version.
func (s *Service) ValidateIdentity(ctx context.Context, p Principal) error {
	_, _, err := s.identityRestrictions(ctx, p)
	return err
}

func (s *Service) identityRestrictions(ctx context.Context, p Principal) ([]Scope, []Scope, error) {
	if !validPrincipal(p) {
		return nil, nil, ErrDenied
	}
	var user model.User
	if s.db.WithContext(ctx).First(&user, p.UserID).Error != nil {
		return nil, nil, ErrDenied
	}
	var tokenScopes, serviceScopes []Scope
	if p.TokenID != 0 {
		var token ScopedToken
		if s.db.WithContext(ctx).First(&token, p.TokenID).Error != nil || token.UserID != p.UserID || token.UserVersion != user.TokenVersion || token.RevokedAt != nil || !token.ExpiresAt.After(time.Now()) || validateScopes(token.Scopes) != nil {
			return nil, nil, ErrDenied
		}
		tokenScopes = token.Scopes
	}
	if p.ServiceID != 0 {
		var identity ServiceIdentity
		if s.db.WithContext(ctx).First(&identity, p.ServiceID).Error != nil || !identity.Active || identity.OwnerID != p.UserID || p.IdentityVersion == 0 || identity.Version != p.IdentityVersion || validateScopes(identity.Scopes) != nil {
			return nil, nil, ErrDenied
		}
		serviceScopes = identity.Scopes
	}
	return tokenScopes, serviceScopes, nil
}

// RestrictPrincipal must run before any administrator or owner shortcut. It
// grants no ordinary user privilege itself: it only narrows a current principal.
func (s *Service) RestrictPrincipal(ctx context.Context, p Principal, kind, id, action string) bool {
	tokenScopes, serviceScopes, err := s.identityRestrictions(ctx, p)
	if err != nil {
		return false
	}
	if p.TokenID != 0 && !identityScopeAllows(tokenScopes, kind, id, action) {
		return false
	}
	if p.ServiceID != 0 && !identityScopeAllows(serviceScopes, kind, id, action) {
		return false
	}
	return true
}

func (s *Service) ListServiceIdentities(ctx context.Context, p Principal) ([]ServiceIdentity, error) {
	user, err := s.identitySession(ctx, p, false)
	if err != nil {
		return nil, err
	}
	rows := []ServiceIdentity{}
	q := s.db.WithContext(ctx).Order("id ASC")
	if !user.Role.IsAdmin() {
		q = q.Where("owner_id = ?", user.ID)
	}
	err = q.Find(&rows).Error
	return rows, err
}

func (s *Service) CreateServiceIdentity(ctx context.Context, p Principal, in ServiceIdentityInput) (ServiceIdentity, error) {
	if _, err := s.identitySession(ctx, p, true); err != nil {
		return ServiceIdentity{}, err
	}
	if !validIdentityName(in.Name) || in.OwnerID == 0 || validateScopes(in.Scopes) != nil {
		return ServiceIdentity{}, embedded.ErrInvalid
	}
	var owner model.User
	if s.db.WithContext(ctx).First(&owner, in.OwnerID).Error != nil {
		return ServiceIdentity{}, embedded.ErrInvalid
	}
	row := ServiceIdentity{Name: in.Name, OwnerID: in.OwnerID, Active: in.Active == nil || *in.Active, Version: 1, Scopes: append([]Scope(nil), in.Scopes...)}
	err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&row).Error; err != nil {
			return err
		}
		return tx.Create(&Audit{ActorID: p.UserID, Action: "identity.create", Resource: fmt.Sprintf("%d:owner:%d:version:%d", row.ID, row.OwnerID, row.Version)}).Error
	})
	return row, err
}

func (s *Service) UpdateServiceIdentity(ctx context.Context, p Principal, id uint64, in ServiceIdentityUpdate) (ServiceIdentity, error) {
	if _, err := s.identitySession(ctx, p, true); err != nil {
		return ServiceIdentity{}, err
	}
	if id == 0 || in.Version == 0 || in.Name == nil && in.OwnerID == nil && in.Active == nil && in.Scopes == nil {
		return ServiceIdentity{}, embedded.ErrInvalid
	}
	s.gate.Lock()
	defer s.gate.Unlock()
	var row ServiceIdentity
	err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&row, id).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return embedded.ErrNotFound
			}
			return err
		}
		if row.Version != in.Version {
			return embedded.ErrInvalid
		}
		previousOwner := row.OwnerID
		if in.Name != nil {
			if !validIdentityName(*in.Name) {
				return embedded.ErrInvalid
			}
			row.Name = *in.Name
		}
		if in.OwnerID != nil {
			var owner model.User
			if *in.OwnerID == 0 || tx.First(&owner, *in.OwnerID).Error != nil {
				return embedded.ErrInvalid
			}
			row.OwnerID = *in.OwnerID
		}
		if in.Scopes != nil {
			if validateScopes(*in.Scopes) != nil {
				return embedded.ErrInvalid
			}
			row.Scopes = append([]Scope(nil), (*in.Scopes)...)
		}
		if in.Active != nil {
			row.Active = *in.Active
		}
		row.Version++
		result := tx.Model(&ServiceIdentity{}).Where("id = ? AND version = ?", id, in.Version).Updates(map[string]any{"name": row.Name, "owner_id": row.OwnerID, "active": row.Active, "version": row.Version, "scopes": scopeJSON(row.Scopes), "updated_at": time.Now().UTC()})
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected != 1 {
			return embedded.ErrInvalid
		}
		return tx.Create(&Audit{ActorID: p.UserID, Action: "identity.update", Resource: fmt.Sprintf("%d:owner:%d>%d:version:%d", id, previousOwner, row.OwnerID, row.Version)}).Error
	})
	if err != nil {
		return ServiceIdentity{}, err
	}
	err = s.db.WithContext(ctx).First(&row, id).Error
	return row, err
}

// ServicePrincipal captures a versioned identity for a new run. Persist all
// fields; do not call this again to silently upgrade an already accepted run.
func (s *Service) ServicePrincipal(ctx context.Context, id uint64) (Principal, error) {
	var row ServiceIdentity
	if id == 0 || s.db.WithContext(ctx).First(&row, id).Error != nil {
		return Principal{}, ErrDenied
	}
	p := Principal{UserID: row.OwnerID, ServiceID: row.ID, IdentityVersion: row.Version}
	if err := s.ValidateIdentity(ctx, p); err != nil {
		return Principal{}, err
	}
	return p, nil
}
