package unified

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"slices"
	"strconv"
	"strings"

	"github.com/certimate-go/certimate/pkg/embedded"
	"github.com/nezhahq/nezha/model"
	"github.com/nezhahq/nezha/service/unified/policy"
	"gorm.io/gorm"
)

func clearLegacyDDNSMaterial(p *model.DDNSProfile) {
	p.AccessID, p.AccessSecret, p.WebhookURL, p.WebhookRequestBody, p.WebhookHeaders = "", "", "", "", ""
}
func legacyDDNSMaterial(p model.DDNSProfile) map[string]any {
	return map[string]any{"access_id": p.AccessID, "access_secret": p.AccessSecret, "webhook_url": p.WebhookURL, "webhook_request_body": p.WebhookRequestBody, "webhook_headers": p.WebhookHeaders}
}
func applyLegacyDDNSMaterial(p *model.DDNSProfile, config map[string]any) {
	p.AccessID, _ = config["access_id"].(string)
	p.AccessSecret, _ = config["access_secret"].(string)
	p.WebhookURL, _ = config["webhook_url"].(string)
	p.WebhookRequestBody, _ = config["webhook_request_body"].(string)
	p.WebhookHeaders, _ = config["webhook_headers"].(string)
}

func (s *Service) authorizeLegacyCredential(ctx context.Context, use embedded.CredentialUse) error {
	p, ok := principalFromContext(ctx)
	if !ok || p.UserID != use.ActorID || use.Purpose != "legacy.ddns.execute" && use.Purpose != "legacy.ddns.configure" {
		return ErrDenied
	}
	current, err := s.CurrentPrincipal(ctx, p)
	if err != nil || current.TokenID != 0 || current.ServiceID != 0 {
		return ErrDenied
	}
	var profile model.DDNSProfile
	if err = s.db.WithContext(ctx).First(&profile, "id = ?", use.ResourceID).Error; err != nil || profile.CredentialID != use.CredentialID || profile.CredentialVersion != use.Version || fmt.Sprint(profile.ConfigVersion) != use.OperationID || profile.UserID == 0 {
		return ErrDenied
	}
	if use.Purpose == "legacy.ddns.execute" && profile.CredentialState != "ready" {
		return ErrDenied
	}
	if use.Purpose == "legacy.ddns.configure" {
		if profile.UserID == current.UserID || requireAdmin(current) == nil {
			return nil
		}
		return ErrDenied
	}
	if profile.UserID == current.UserID {
		return nil
	}
	var owner model.User
	if s.db.WithContext(ctx).First(&owner, profile.UserID).Error == nil && owner.Role.IsAdmin() {
		return nil
	}
	return ErrDenied
}

// LegacyDDNSProfile resolves an exact saved profile for each legacy worker effect.
// The returned material is server-only and is never passed to a metadata response.
func (s *Service) LegacyDDNSProfile(ctx context.Context, actor, id, version uint64) (model.DDNSProfile, error) {
	var p model.DDNSProfile
	if err := s.db.WithContext(ctx).First(&p, id).Error; err != nil {
		return p, err
	}
	if p.ConfigVersion != version || p.CredentialID == "" {
		return model.DDNSProfile{}, ErrDenied
	}
	config, err := s.engine.ResolveCredential(WithPrincipal(ctx, Principal{UserID: actor}), embedded.CredentialUse{ActorID: actor, CredentialID: p.CredentialID, Version: p.CredentialVersion, Purpose: "legacy.ddns.execute", ResourceKind: "legacy_ddns", ResourceID: fmt.Sprint(p.ID), OperationID: fmt.Sprint(version)})
	if err != nil {
		return model.DDNSProfile{}, ErrDenied
	}
	applyLegacyDDNSMaterial(&p, config)
	return p, nil
}

// LegacyDDNSHostProfile checks the current saved host before every legacy
// provider effect. Nonempty host overrides replace the profile's default names.
func (s *Service) LegacyDDNSHostProfile(ctx context.Context, actor, hostID, id, version uint64, names []string) (model.DDNSProfile, error) {
	if actor == 0 || hostID == 0 || len(names) == 0 {
		return model.DDNSProfile{}, ErrDenied
	}
	s.gate.RLock()
	defer s.gate.RUnlock()
	var host model.Server
	if err := s.db.WithContext(ctx).First(&host, hostID).Error; err != nil || host.UserID != actor || !host.EnableDDNS {
		return model.DDNSProfile{}, ErrDenied
	}
	// Server.AfterFind preserves compatibility by swallowing invalid legacy JSON;
	// authorization must reject it rather than fall back to a wider domain set.
	if json.Unmarshal([]byte(host.DDNSProfilesRaw), &host.DDNSProfiles) != nil || !slices.Contains(host.DDNSProfiles, id) {
		return model.DDNSProfile{}, ErrDenied
	}
	if host.OverrideDDNSDomainsRaw != "" && json.Unmarshal([]byte(host.OverrideDDNSDomainsRaw), &host.OverrideDDNSDomains) != nil {
		return model.DDNSProfile{}, ErrDenied
	}
	var profile model.DDNSProfile
	if err := s.db.WithContext(ctx).First(&profile, id).Error; err != nil || profile.ConfigVersion != version {
		return model.DDNSProfile{}, ErrDenied
	}
	domains := profile.Domains
	if override := host.OverrideDDNSDomains[id]; len(override) > 0 {
		domains = override
	}
	for _, name := range names {
		if !s.legacyDDNSRecordAllowed(ctx, domains, name) {
			return model.DDNSProfile{}, ErrDenied
		}
	}
	return s.LegacyDDNSProfile(ctx, actor, id, version)
}

func (s *Service) LegacyDDNSRecordAllowed(ctx context.Context, profile model.DDNSProfile, name string) bool {
	return s.legacyDDNSRecordAllowed(ctx, profile.Domains, name)
}

func (s *Service) legacyDDNSRecordAllowed(ctx context.Context, domains []string, name string) bool {
	name, err := policy.NormalizeName(name)
	if err != nil {
		return false
	}
	allowed := false
	for _, domain := range domains {
		normalized, err := policy.NormalizeName(domain)
		if err == nil && normalized == name {
			allowed = true
			break
		}
	}
	if !allowed {
		return false
	}
	var count int64
	return s.db.WithContext(ctx).Model(&DDNSBinding{}).Where("LOWER(RTRIM(name, '.')) = ?", name).Count(&count).Error == nil && count == 0
}

// SaveLegacyDDNS retains the original entry point while moving all provider and
// webhook material to the same encrypted, versioned credential authority.
func (s *Service) SaveLegacyDDNS(ctx context.Context, p Principal, profile model.DDNSProfile) (model.DDNSProfile, error) {
	p, err := s.CurrentPrincipal(ctx, p)
	if err != nil || p.TokenID != 0 || p.ServiceID != 0 {
		return model.DDNSProfile{}, ErrDenied
	}
	s.gate.Lock()
	defer s.gate.Unlock()
	var prior model.DDNSProfile
	if profile.ID != 0 {
		if err = s.db.WithContext(ctx).First(&prior, profile.ID).Error; err != nil {
			return model.DDNSProfile{}, err
		}
		if prior.UserID != p.UserID && requireAdmin(p) != nil {
			return model.DDNSProfile{}, ErrDenied
		}
		if profile.ConfigVersion != prior.ConfigVersion {
			return model.DDNSProfile{}, ErrTargetConflict
		}
		profile.UserID = prior.UserID
	} else {
		profile.UserID = p.UserID
	}
	if len(profile.Name) == 0 || len(profile.Name) > 100 || !slices.Contains(model.ProviderList[:], profile.Provider) || len(profile.Domains) > 100 || profile.MaxRetries < 1 || profile.MaxRetries > 10 {
		return model.DDNSProfile{}, embedded.ErrInvalid
	}
	if profile.UserID == 0 {
		return model.DDNSProfile{}, ErrDenied
	}
	material := legacyDDNSMaterial(profile)
	originalMaterial := map[string]any{}
	if prior.ID != 0 && profile.Provider == prior.Provider {
		existing := legacyDDNSMaterial(prior)
		if prior.CredentialID != "" {
			existing, err = s.engine.ResolveCredential(WithPrincipal(ctx, p), embedded.CredentialUse{ActorID: p.UserID, CredentialID: prior.CredentialID, Version: prior.CredentialVersion, Purpose: "legacy.ddns.configure", ResourceKind: "legacy_ddns", ResourceID: fmt.Sprint(prior.ID), OperationID: fmt.Sprint(prior.ConfigVersion)})
			if err != nil {
				return model.DDNSProfile{}, ErrDenied
			}
		}
		originalMaterial = existing
		for key, value := range material {
			if value == "" {
				material[key] = existing[key]
			}
		}
	}
	applyLegacyDDNSMaterial(&profile, material)
	if profile.Provider != model.ProviderDummy && profile.Provider != model.ProviderWebHook && profile.AccessSecret == "" || profile.Provider == model.ProviderWebHook && profile.WebhookURL == "" || profile.Provider == model.ProviderTencentCloud && profile.AccessID == "" {
		return model.DDNSProfile{}, embedded.ErrInvalid
	}
	clearLegacyDDNSMaterial(&profile)
	profile.ConfigVersion = prior.ConfigVersion + 1
	profile.CredentialID = ""
	profile.CredentialVersion = 0
	profile.CredentialState = "ready"
	if prior.CredentialID != "" && profile.Provider == prior.Provider && reflect.DeepEqual(material, originalMaterial) {
		profile.CredentialID, profile.CredentialVersion = prior.CredentialID, prior.CredentialVersion
	} else if profile.Provider != model.ProviderDummy {
		raw, _ := json.Marshal(material)
		credential, e := s.engine.CreateCredential(ctx, strconv.FormatUint(p.UserID, 10), embedded.CredentialInput{Name: legacyCredentialName(profile.Name), Provider: "nezha-legacy-ddns:" + profile.Provider, Config: raw})
		if e != nil {
			return model.DDNSProfile{}, e
		}
		profile.CredentialID, profile.CredentialVersion = credential.ID, credential.Version
	}
	err = s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if e := tx.Save(&profile).Error; e != nil {
			return e
		}
		if profile.CredentialID != "" && profile.CredentialID != prior.CredentialID {
			if e := tx.Create(&ResourceBinding{Kind: "credential", ResourceID: profile.CredentialID, OwnerID: profile.UserID, Version: 1, Active: true}).Error; e != nil {
				return e
			}
		}
		return tx.Create(&Audit{ActorID: p.UserID, Action: "legacy_ddns.save", Resource: fmt.Sprint(profile.ID)}).Error
	})
	return profile, err
}

// Migration only moves local secret storage; it does not start a provider call.
func (s *Service) MigrateLegacyDDNS(ctx context.Context) error {
	if !s.db.Migrator().HasTable(&model.DDNSProfile{}) {
		return nil
	}
	var profiles []model.DDNSProfile
	if err := s.db.WithContext(ctx).Where("credential_id = '' OR credential_id IS NULL").Find(&profiles).Error; err != nil {
		return err
	}
	for _, p := range profiles {
		if p.Provider == model.ProviderDummy {
			continue
		}
		material := legacyDDNSMaterial(p)
		raw, _ := json.Marshal(material)
		r, err := s.engine.CreateCredential(ctx, "legacy-migration", embedded.CredentialInput{Name: legacyCredentialName(p.Name), Provider: "nezha-legacy-ddns:" + p.Provider, Config: raw})
		if err != nil {
			return err
		}
		p.CredentialID, p.CredentialVersion, p.CredentialState = r.ID, r.Version, "ready"
		var owner model.User
		ownerExists := p.UserID != 0 && s.db.WithContext(ctx).First(&owner, p.UserID).Error == nil
		if !ownerExists {
			p.CredentialState = "owner_required"
		} else if p.Provider == model.ProviderWebHook && p.WebhookURL == "" || p.Provider != model.ProviderWebHook && p.AccessSecret == "" {
			p.CredentialState = "missing"
		}
		clearLegacyDDNSMaterial(&p)
		p.ConfigVersion++
		if err = s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
			if err := tx.Save(&p).Error; err != nil {
				return err
			}
			if ownerExists {
				return tx.Create(&ResourceBinding{Kind: "credential", ResourceID: r.ID, OwnerID: p.UserID, Version: 1, Active: true}).Error
			}
			return nil
		}); err != nil {
			return err
		}
	}
	return nil
}

func legacyCredentialName(name string) string {
	name = "DDNS: " + name
	if len(name) > 120 {
		name = strings.ToValidUTF8(name[:120], "")
	}
	return name
}
