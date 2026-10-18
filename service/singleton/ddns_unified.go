package singleton

import (
	"context"
	"errors"
	"strings"

	"github.com/libdns/cloudflare"
	"github.com/libdns/he"
	"github.com/libdns/libdns"
	tencentcloud "github.com/nezhahq/libdns-tencentcloud"
	"github.com/nezhahq/nezha/model"
	"github.com/nezhahq/nezha/pkg/ddns/webhook"
)

type legacyDDNSSetter struct{ actor, hostID, profileID, version uint64 }

func (p *legacyDDNSSetter) provider(ctx context.Context, zone string, recs []libdns.Record) (libdns.RecordSetter, error) {
	if Unified == nil {
		return nil, errors.New("unified credential runtime unavailable")
	}
	names := make([]string, 0, len(recs))
	for _, rec := range recs {
		name := rec.RR().Name
		if name == "@" || name == "" {
			name = strings.TrimSuffix(zone, ".")
		} else {
			name = strings.TrimSuffix(name, ".") + "." + strings.TrimSuffix(zone, ".")
		}
		names = append(names, name)
	}
	profile, err := Unified.LegacyDDNSHostProfile(ctx, p.actor, p.hostID, p.profileID, p.version, names)
	if err != nil {
		return nil, errors.New("DDNS host or record authorization denied")
	}
	switch profile.Provider {
	case model.ProviderWebHook:
		return &webhook.Provider{DDNSProfile: &profile}, nil
	case model.ProviderCloudflare:
		return &cloudflare.Provider{APIToken: profile.AccessSecret}, nil
	case model.ProviderTencentCloud:
		return &tencentcloud.Provider{SecretId: profile.AccessID, SecretKey: profile.AccessSecret}, nil
	case model.ProviderHE:
		return &he.Provider{APIKey: profile.AccessSecret}, nil
	}
	return nil, errors.New("DDNS provider unavailable")
}
func (p *legacyDDNSSetter) SetRecords(ctx context.Context, zone string, recs []libdns.Record) ([]libdns.Record, error) {
	provider, err := p.provider(ctx, zone, recs)
	if err != nil {
		return nil, err
	}
	out, err := provider.SetRecords(ctx, zone, recs)
	if err != nil {
		return nil, errors.New("DDNS provider update failed")
	}
	return out, nil
}
func (p *legacyDDNSSetter) DeleteRecords(ctx context.Context, zone string, recs []libdns.Record) ([]libdns.Record, error) {
	provider, err := p.provider(ctx, zone, recs)
	if err != nil {
		return nil, err
	}
	deleter, ok := provider.(libdns.RecordDeleter)
	if !ok {
		return nil, nil
	}
	out, err := deleter.DeleteRecords(ctx, zone, recs)
	if err != nil {
		return nil, errors.New("DDNS provider deletion failed")
	}
	return out, nil
}
