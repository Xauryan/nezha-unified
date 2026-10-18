package unified

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/nezhahq/nezha/model"
)

func TestLegacyDDNSSealsMaterialPreservesBlankEditAndChecksCurrentAuthority(t *testing.T) {
	s, admin, _, _ := deploymentFixture(t)
	ctx := context.Background()
	if err := s.db.AutoMigrate(&model.DDNSProfile{}); err != nil {
		t.Fatal(err)
	}
	on := true
	p, err := s.SaveLegacyDDNS(ctx, admin, model.DDNSProfile{Common: model.Common{UserID: admin.UserID}, Name: "synthetic", Provider: model.ProviderCloudflare, Domains: []string{"legacy.example.test"}, AccessSecret: "synthetic-legacy-provider-secret", MaxRetries: 3, EnableIPv4: &on})
	if err != nil {
		t.Fatal(err)
	}
	if p.CredentialID == "" || p.CredentialVersion != 1 || p.ConfigVersion != 1 || p.AccessSecret != "" {
		t.Fatal("profile did not move to a versioned reference")
	}
	var saved model.DDNSProfile
	if err = s.db.First(&saved, p.ID).Error; err != nil {
		t.Fatal(err)
	}
	encoded, _ := json.Marshal(saved)
	if strings.Contains(string(encoded), "synthetic-legacy-provider-secret") || saved.AccessSecret != "" {
		t.Fatal("legacy material escaped encrypted storage")
	}
	resolved, err := s.LegacyDDNSProfile(ctx, admin.UserID, p.ID, p.ConfigVersion)
	if err != nil || resolved.AccessSecret != "synthetic-legacy-provider-secret" {
		t.Fatal("authorized compatibility execution failed")
	}
	p.Name = "synthetic edited"
	edited, err := s.SaveLegacyDDNS(ctx, admin, p)
	if err != nil {
		t.Fatal(err)
	}
	if _, err = s.LegacyDDNSProfile(ctx, admin.UserID, p.ID, p.ConfigVersion); err == nil {
		t.Fatal("worker retained obsolete profile authority")
	}
	current, err := s.LegacyDDNSProfile(ctx, admin.UserID, p.ID, edited.ConfigVersion)
	if err != nil || current.AccessSecret != "synthetic-legacy-provider-secret" {
		t.Fatal("blank edit lost saved material")
	}
	if _, err = s.SaveLegacyDDNS(ctx, admin, p); err == nil {
		t.Fatal("stale profile overwrite accepted")
	}
	edited.Provider = model.ProviderTencentCloud
	if _, err = s.SaveLegacyDDNS(ctx, admin, edited); err == nil {
		t.Fatal("provider switch reused incompatible credential material")
	}
	if _, err = s.SaveLegacyDDNS(ctx, Principal{UserID: 1, Admin: true, PAT: true}, edited); err == nil {
		t.Fatal("legacy PAT gained unified credential power")
	}
	if err = s.db.Create(&model.User{Common: model.Common{ID: 2}, Username: "synthetic-member", Role: model.RoleMember}).Error; err != nil {
		t.Fatal(err)
	}
	if _, err = s.LegacyDDNSProfile(ctx, 2, p.ID, current.ConfigVersion); err != nil {
		t.Fatal("real admin-owned profile lost compatible sharing")
	}
	if err = s.db.Model(&model.User{}).Where("id = ?", 1).Update("role", model.RoleMember).Error; err != nil {
		t.Fatal(err)
	}
	if _, err = s.LegacyDDNSProfile(ctx, 2, p.ID, current.ConfigVersion); err == nil {
		t.Fatal("worker ignored current profile-owner demotion")
	}
}

func TestLegacyDDNSMigrationPreservesReferencesAndFencesIncompleteOwners(t *testing.T) {
	s, _, _, _ := deploymentFixture(t)
	ctx := context.Background()
	if err := s.db.AutoMigrate(&model.DDNSProfile{}); err != nil {
		t.Fatal(err)
	}
	profiles := []model.DDNSProfile{
		{Common: model.Common{ID: 10, UserID: 1}, Name: "synthetic valid", Provider: model.ProviderCloudflare, AccessSecret: "synthetic-migration-secret", MaxRetries: 3},
		{Common: model.Common{ID: 11, UserID: 0}, Name: "synthetic unowned", Provider: model.ProviderWebHook, WebhookURL: "https://synthetic.invalid/private-token", MaxRetries: 3},
		{Common: model.Common{ID: 12, UserID: 1}, Name: "synthetic incomplete", Provider: model.ProviderCloudflare, MaxRetries: 3},
	}
	for i := range profiles {
		if err := s.db.Create(&profiles[i]).Error; err != nil {
			t.Fatal(err)
		}
	}
	if err := s.MigrateLegacyDDNS(ctx); err != nil {
		t.Fatal(err)
	}
	if err := s.db.Order("id").Find(&profiles).Error; err != nil {
		t.Fatal(err)
	}
	if profiles[0].ID != 10 || profiles[0].CredentialState != "ready" || profiles[1].CredentialState != "owner_required" || profiles[2].CredentialState != "missing" {
		t.Fatal("migration invented execution authority")
	}
	ids := []string{profiles[0].CredentialID, profiles[1].CredentialID, profiles[2].CredentialID}
	for _, p := range profiles {
		if p.AccessSecret != "" || p.WebhookURL != "" || p.CredentialID == "" {
			t.Fatal("migration retained plaintext or lost encrypted reference")
		}
	}
	if err := s.MigrateLegacyDDNS(ctx); err != nil {
		t.Fatal(err)
	}
	if err := s.db.Order("id").Find(&profiles).Error; err != nil {
		t.Fatal(err)
	}
	for i, p := range profiles {
		if p.CredentialID != ids[i] {
			t.Fatal("migration was not idempotent")
		}
	}
	if _, err := s.LegacyDDNSProfile(ctx, 1, 11, 1); err == nil {
		t.Fatal("unowned profile became executable")
	}
	if _, err := s.LegacyDDNSProfile(ctx, 1, 12, 1); err == nil {
		t.Fatal("incomplete profile became executable")
	}
}

func TestLegacyDDNSHostEffectsUseCurrentBindingAndOverrideScope(t *testing.T) {
	s, admin, _, _ := deploymentFixture(t)
	ctx := context.Background()
	if err := s.db.AutoMigrate(&model.DDNSProfile{}); err != nil {
		t.Fatal(err)
	}
	if err := s.db.Create(&model.User{Common: model.Common{ID: 2}, Username: "synthetic-host-owner", Role: model.RoleMember}).Error; err != nil {
		t.Fatal(err)
	}
	on := true
	profile, err := s.SaveLegacyDDNS(ctx, admin, model.DDNSProfile{
		Name: "synthetic host scoped", Provider: model.ProviderCloudflare,
		Domains: []string{"DEFAULT.Example.Test."}, AccessSecret: "synthetic-host-scoped-secret", MaxRetries: 1, EnableIPv4: &on,
	})
	if err != nil {
		t.Fatal(err)
	}
	profileIDs, _ := json.Marshal([]uint64{profile.ID})
	overrides, _ := json.Marshal(map[uint64][]string{profile.ID: {"OVERRIDE.Example.Test."}})
	resetHost := func(t *testing.T) {
		t.Helper()
		if err := s.db.Model(&model.Server{}).Where("id = ?", 10).Updates(map[string]any{
			"user_id": 2, "EnableDDNS": true, "ddns_profiles_raw": string(profileIDs), "override_ddns_domains_raw": string(overrides),
		}).Error; err != nil {
			t.Fatal(err)
		}
	}
	allowed := func(name string) bool {
		resolved, err := s.LegacyDDNSHostProfile(ctx, 2, 10, profile.ID, profile.ConfigVersion, []string{name})
		return err == nil && resolved.AccessSecret == "synthetic-host-scoped-secret"
	}
	mutateHost := func(t *testing.T, field string, value any) {
		t.Helper()
		if err := s.db.Model(&model.Server{}).Where("id = ?", 10).Update(field, value).Error; err != nil {
			t.Fatal(err)
		}
	}
	t.Run("current override replaces defaults and normalizes names", func(t *testing.T) {
		resetHost(t)
		if !allowed("override.example.test") || !allowed("Override.Example.Test.") {
			t.Fatal("current host override was rejected")
		}
		for _, name := range []string{"default.example.test", "other.example.test", "child.override.example.test", "override.example.test.invalid"} {
			if allowed(name) {
				t.Fatal("record outside current host override was authorized")
			}
		}
		if _, err := s.LegacyDDNSHostProfile(ctx, 2, 10, profile.ID, profile.ConfigVersion, []string{"override.example.test", "outside.example.test"}); err == nil {
			t.Fatal("batch with an unauthorized record was accepted")
		}
	})
	t.Run("updated override fences already resolved worker", func(t *testing.T) {
		resetHost(t)
		current, _ := json.Marshal(map[uint64][]string{profile.ID: {"CURRENT.example.test."}})
		mutateHost(t, "override_ddns_domains_raw", string(current))
		if allowed("override.example.test") || !allowed("current.example.test") {
			t.Fatal("worker did not use the current saved override")
		}
	})
	for _, raw := range []string{"{}", "", "null"} {
		t.Run("no override uses profile defaults "+raw, func(t *testing.T) {
			resetHost(t)
			mutateHost(t, "override_ddns_domains_raw", raw)
			if !allowed("default.example.test") || allowed("override.example.test") {
				t.Fatal("default domain fallback changed compatibility scope")
			}
		})
	}
	t.Run("empty profile override uses defaults", func(t *testing.T) {
		resetHost(t)
		empty, _ := json.Marshal(map[uint64][]string{profile.ID: {}})
		mutateHost(t, "override_ddns_domains_raw", string(empty))
		if !allowed("default.example.test") {
			t.Fatal("empty override did not preserve profile fallback")
		}
	})
	for _, tc := range []struct {
		name, field string
		value       any
	}{
		{"disabled host", "EnableDDNS", false},
		{"unbound profile", "ddns_profiles_raw", "[]"},
		{"transferred host", "user_id", uint64(3)},
		{"invalid profile binding", "ddns_profiles_raw", "invalid"},
		{"invalid override", "override_ddns_domains_raw", "invalid"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			resetHost(t)
			mutateHost(t, tc.field, tc.value)
			if allowed("override.example.test") {
				t.Fatal("worker retained revoked host authority")
			}
		})
	}
	for _, enabled := range []bool{true, false} {
		t.Run(fmt.Sprintf("unified binding owns name enabled=%t", enabled), func(t *testing.T) {
			resetHost(t)
			binding := DDNSBinding{Name: "Override.Example.Test.", RecordType: "A", Enabled: enabled}
			if err := s.db.Create(&binding).Error; err != nil {
				t.Fatal(err)
			}
			t.Cleanup(func() { s.db.Delete(&binding) })
			if allowed("override.example.test") {
				t.Fatal("legacy worker bypassed unified single writer binding")
			}
		})
	}
	t.Run("host identity is required", func(t *testing.T) {
		resetHost(t)
		for _, hostID := range []uint64{0, 999} {
			if _, err := s.LegacyDDNSHostProfile(ctx, 2, hostID, profile.ID, profile.ConfigVersion, []string{"override.example.test"}); err == nil {
				t.Fatal("missing host identity was accepted")
			}
		}
	})
	t.Run("deleted host", func(t *testing.T) {
		resetHost(t)
		if err := s.db.Delete(&model.Server{}, 10).Error; err != nil {
			t.Fatal(err)
		}
		if allowed("override.example.test") {
			t.Fatal("deleted host retained provider execution authority")
		}
	})
}
