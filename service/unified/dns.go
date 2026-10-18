package unified

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"hash/fnv"
	"net/netip"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/certimate-go/certimate/pkg/embedded"
	"github.com/nezhahq/nezha/model"
	"github.com/nezhahq/nezha/service/unified/dnsprovider"
	"github.com/nezhahq/nezha/service/unified/policy"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type DNSAccount struct {
	ID                uint64    `json:"id" gorm:"primaryKey"`
	Name              string    `json:"name"`
	OwnerID           uint64    `json:"owner_id"`
	Provider          string    `json:"provider"`
	CredentialID      string    `json:"credential_id"`
	CredentialVersion int       `json:"credential_version"`
	Enabled           bool      `json:"enabled"`
	CreatedAt         time.Time `json:"created_at"`
}

func (DNSAccount) TableName() string { return "unified_dns_accounts" }

type DNSZone struct {
	ID          uint64    `json:"id" gorm:"primaryKey"`
	AccountID   uint64    `json:"account_id" gorm:"uniqueIndex:provider_zone"`
	ProviderID  string    `json:"provider_id" gorm:"uniqueIndex:provider_zone"`
	Name        string    `json:"name"`
	Mode        string    `json:"mode"`
	Status      string    `json:"status,omitempty"`
	NameServers []string  `json:"name_servers,omitempty" gorm:"serializer:json"`
	Active      bool      `json:"active"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (DNSZone) TableName() string { return "unified_dns_zones" }

type DNSPolicy struct {
	ID               uint64 `json:"id" gorm:"primaryKey"`
	UserID           uint64 `json:"user_id" gorm:"index"`
	ServiceID        uint64 `json:"service_id,omitempty" gorm:"index"`
	ZoneID           uint64 `json:"zone_id" gorm:"index"`
	policy.NameScope `gorm:"embedded"`
	Types            []string  `json:"types" gorm:"serializer:json"`
	Actions          []string  `json:"actions" gorm:"serializer:json"`
	CreatedAt        time.Time `json:"created_at"`
}

func (DNSPolicy) TableName() string { return "unified_dns_policies" }

type DNSOperation struct {
	ID              uint64    `json:"id" gorm:"primaryKey"`
	ActorID         uint64    `json:"actor_id" gorm:"uniqueIndex:dns_op_key"`
	TokenID         uint64    `json:"token_id,omitempty"`
	ServiceID       uint64    `json:"service_id,omitempty"`
	IdentityVersion uint64    `json:"identity_version,omitempty"`
	BindingID       uint64    `json:"binding_id,omitempty"`
	Key             string    `json:"key" gorm:"uniqueIndex:dns_op_key"`
	Digest          string    `json:"-"`
	ZoneID          uint64    `json:"zone_id"`
	Action          string    `json:"action"`
	RecordID        string    `json:"record_id"`
	RecordType      string    `json:"record_type"`
	Name            string    `json:"name"`
	State           string    `json:"state"`
	Error           string    `json:"error,omitempty"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func (DNSOperation) TableName() string { return "unified_dns_operations" }

type DNSChange struct {
	OperationKey string             `json:"operation_key"`
	Action       string             `json:"action"`
	Record       dnsprovider.Record `json:"record"`
	PauseDDNS    bool               `json:"pause_ddns,omitempty"`
}
type DNSPage struct {
	Items        []dnsprovider.Record     `json:"items"`
	Total        int                      `json:"total"`
	Page         int                      `json:"page"`
	PerPage      int                      `json:"per_page"`
	Capabilities dnsprovider.Capabilities `json:"capabilities"`
}

func (s *Service) authorizeCredential(ctx context.Context, use embedded.CredentialUse) error {
	if use.ResourceKind == "credential" {
		return s.authorizeStandaloneCredential(ctx, use)
	}
	if use.ResourceKind == "workflow" {
		return s.authorizeWorkflowCredential(ctx, use)
	}
	executionPrincipal := Principal{UserID: use.ActorID}
	if p, ok := principalFromContext(ctx); ok {
		executionPrincipal = p
		if p.UserID != use.ActorID {
			return ErrDenied
		}
		if _, err := s.CurrentPrincipal(ctx, p); err != nil {
			return ErrDenied
		}
	}
	var user model.User
	if err := s.db.WithContext(ctx).First(&user, use.ActorID).Error; err != nil {
		return ErrDenied
	}
	if use.ResourceKind == "dns_account" && use.Purpose == "dns.manage" {
		id, err := strconv.ParseUint(use.ResourceID, 10, 64)
		if err != nil {
			return ErrDenied
		}
		var a DNSAccount
		if err := s.db.WithContext(ctx).First(&a, id).Error; err != nil || !a.Enabled || a.CredentialID != use.CredentialID || a.CredentialVersion != use.Version {
			return ErrDenied
		}
		if user.Role.IsAdmin() && executionPrincipal.ServiceID == 0 {
			return nil
		}
		var count int64
		q := s.db.WithContext(ctx).Model(&DNSPolicy{}).Joins("JOIN unified_dns_zones z ON z.id = unified_dns_policies.zone_id").Where("z.account_id = ? AND z.active = ?", a.ID, true)
		if executionPrincipal.ServiceID != 0 {
			q = q.Where("unified_dns_policies.service_id = ?", executionPrincipal.ServiceID)
		} else {
			q = q.Where("unified_dns_policies.user_id = ? AND unified_dns_policies.service_id = 0", user.ID)
		}
		err = q.Count(&count).Error
		if err == nil && count > 0 {
			return nil
		}
	}
	return ErrDenied
}

func (s *Service) CreateDNSAccount(ctx context.Context, p Principal, a DNSAccount) (DNSAccount, error) {
	current, e := s.CurrentPrincipal(ctx, p)
	if e != nil {
		return a, e
	}
	p = current
	if err := requireAdmin(p); err != nil {
		return a, err
	}
	info, err := s.engine.CredentialInfo(ctx, a.CredentialID)
	if err != nil {
		return a, err
	}
	expected := "tencentcloud"
	if a.Provider == "cloudflare" {
		expected = "cloudflare"
	} else if a.Provider != "dnspod" && a.Provider != "edgeone" {
		return a, embedded.ErrInvalid
	}
	if info.Provider != expected || !info.Enabled || a.CredentialVersion != info.Version || strings.TrimSpace(a.Name) == "" {
		return a, embedded.ErrInvalid
	}
	a.OwnerID = p.UserID
	a.ID = 0
	a.Enabled = true
	a.CreatedAt = time.Time{}
	err = s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&a).Error; err != nil {
			return err
		}
		return tx.Create(&Audit{ActorID: p.UserID, Action: "dns.account.create", Resource: fmt.Sprint(a.ID)}).Error
	})
	return a, err
}
func (s *Service) DNSAccounts(ctx context.Context, p Principal) ([]DNSAccount, error) {
	current, e := s.CurrentPrincipal(ctx, p)
	if e != nil {
		return nil, e
	}
	p = current
	if !validPrincipal(p) {
		return nil, ErrDenied
	}
	rows := []DNSAccount{}
	q := s.db.WithContext(ctx)
	if !p.Admin {
		q = q.Where("id IN (?)", s.db.Model(&DNSZone{}).Select("account_id").Where("id IN (?)", s.dnsPolicySubject(ctx, p).Select("zone_id")))
	}
	err := q.Order("id").Find(&rows).Error
	allowed := []DNSAccount{}
	for _, a := range rows {
		if s.RestrictPrincipal(ctx, p, "dns_account", fmt.Sprint(a.ID), "read") {
			allowed = append(allowed, a)
			continue
		}
		var zones []DNSZone
		s.db.WithContext(ctx).Where("account_id = ? AND active = ?", a.ID, true).Find(&zones)
		for _, z := range zones {
			if s.RestrictPrincipal(ctx, p, "dns_zone", fmt.Sprint(z.ID), "read") {
				allowed = append(allowed, a)
				break
			}
		}
	}
	return allowed, err
}
func (s *Service) DNSZones(ctx context.Context, p Principal) ([]DNSZone, error) {
	current, e := s.CurrentPrincipal(ctx, p)
	if e != nil {
		return nil, e
	}
	p = current
	if !validPrincipal(p) {
		return nil, ErrDenied
	}
	rows := []DNSZone{}
	q := s.db.WithContext(ctx).Where("active = ?", true)
	if !p.Admin {
		q = q.Where("id IN (?)", s.dnsPolicySubject(ctx, p).Select("zone_id"))
	}
	err := q.Order("name").Find(&rows).Error
	allowed := []DNSZone{}
	for _, z := range rows {
		if s.RestrictPrincipal(ctx, p, "dns_zone", fmt.Sprint(z.ID), "read") {
			allowed = append(allowed, z)
		}
	}
	return allowed, err
}
func (s *Service) dnsClient(ctx context.Context, p Principal, a DNSAccount) (dnsprovider.Provider, error) {
	if !validPrincipal(p) || !a.Enabled {
		return nil, ErrDenied
	}
	cfg, err := s.engine.ResolveCredential(WithPrincipal(ctx, p), embedded.CredentialUse{ActorID: p.UserID, CredentialID: a.CredentialID, Version: a.CredentialVersion, Purpose: "dns.manage", ResourceKind: "dns_account", ResourceID: fmt.Sprint(a.ID)})
	if err != nil {
		return nil, err
	}
	return s.dnsFactory(a.Provider, cfg)
}
func (s *Service) RefreshDNSZones(ctx context.Context, p Principal, id uint64) ([]DNSZone, error) {
	current, e := s.CurrentPrincipal(ctx, p)
	if e != nil {
		return nil, e
	}
	p = current
	if err := requireAdmin(p); err != nil {
		return nil, err
	}
	var account DNSAccount
	if err := s.db.WithContext(ctx).First(&account, id).Error; err != nil {
		return nil, err
	}
	client, err := s.dnsClient(ctx, p, account)
	if err != nil {
		return nil, err
	}
	zones, err := client.Zones(ctx)
	if err != nil {
		return nil, err
	}
	out := make([]DNSZone, 0, len(zones))
	for _, z := range zones {
		n, err := policy.NormalizeName(z.Name)
		if err != nil || z.ID == "" || strings.Contains(n, "*") {
			return nil, errors.New("provider returned invalid zone identity")
		}
		out = append(out, DNSZone{AccountID: id, ProviderID: z.ID, Name: n, Mode: z.Mode, Status: z.Status, NameServers: z.NameServers, Active: true})
	}
	s.gate.Lock()
	defer s.gate.Unlock()
	err = s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&DNSZone{}).Where("account_id = ?", id).Update("active", false).Error; err != nil {
			return err
		}
		for i := range out {
			if err := tx.Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "account_id"}, {Name: "provider_id"}}, DoUpdates: clause.AssignmentColumns([]string{"name", "mode", "status", "name_servers", "active", "updated_at"})}).Create(&out[i]).Error; err != nil {
				return err
			}
			if err := tx.Where("account_id = ? AND provider_id = ?", id, out[i].ProviderID).First(&out[i]).Error; err != nil {
				return err
			}
		}
		return tx.Create(&Audit{ActorID: p.UserID, Action: "dns.zones.refresh", Resource: fmt.Sprint(id)}).Error
	})
	return out, err
}
func (s *Service) dnsContext(ctx context.Context, p Principal, zoneID uint64) (DNSZone, dnsprovider.Provider, error) {
	var z DNSZone
	if err := s.db.WithContext(ctx).First(&z, zoneID).Error; err != nil || !z.Active {
		return z, nil, ErrDenied
	}
	current, err := s.CurrentPrincipal(ctx, p)
	if err != nil {
		return z, nil, ErrDenied
	}
	p = current
	allowedIdentity := false
	for _, action := range []string{"read", "create", "write", "delete", "execute"} {
		if s.RestrictPrincipal(ctx, p, "dns_zone", fmt.Sprint(zoneID), action) {
			allowedIdentity = true
			break
		}
	}
	if !allowedIdentity {
		return z, nil, ErrDenied
	}
	if !p.Admin {
		var count int64
		if s.dnsPolicySubject(ctx, p).Where("zone_id = ?", zoneID).Count(&count).Error != nil || count == 0 {
			return z, nil, ErrDenied
		}
	}
	var a DNSAccount
	if err := s.db.WithContext(ctx).First(&a, z.AccountID).Error; err != nil {
		return z, nil, err
	}
	client, err := s.dnsClient(ctx, p, a)
	return z, client, err
}
func providerZone(z DNSZone) dnsprovider.Zone {
	return dnsprovider.Zone{ID: z.ProviderID, Name: z.Name, Mode: z.Mode}
}
func (s *Service) canDNS(ctx context.Context, p Principal, z DNSZone, r dnsprovider.Record, action string) bool {
	if !validPrincipal(p) || !z.Active {
		return false
	}
	tokenAction := action
	if action == "update" {
		tokenAction = "write"
	} else if action == "ddns" {
		tokenAction = "execute"
	}
	if !s.RestrictPrincipal(ctx, p, "dns_zone", fmt.Sprint(z.ID), tokenAction) {
		return false
	}
	if !(policy.NameScope{Root: z.Name, IncludeRoot: true, Descendants: true, AllowWildcard: true}).Allows(r.Name) {
		return false
	}
	var current model.User
	if s.db.WithContext(ctx).First(&current, p.UserID).Error != nil {
		return false
	}
	if current.Role.IsAdmin() && p.ServiceID == 0 {
		return true
	}
	var grants []DNSPolicy
	if s.dnsPolicySubject(ctx, p).Where("zone_id = ?", z.ID).Find(&grants).Error != nil {
		return false
	}
	for _, g := range grants {
		if slices.Contains(g.Actions, action) && slices.Contains(g.Types, r.Type) && g.NameScope.Allows(r.Name) {
			return true
		}
	}
	return false
}
func (s *Service) DNSRecords(ctx context.Context, p Principal, zoneID uint64, page, size int) (DNSPage, error) {
	if !s.RestrictPrincipal(ctx, p, "dns_zone", fmt.Sprint(zoneID), "read") {
		return DNSPage{}, ErrDenied
	}
	result := DNSPage{Items: []dnsprovider.Record{}, Page: page, PerPage: size}
	if page < 1 || size < 1 || size > 1000 || page > 100000 {
		return result, embedded.ErrInvalid
	}
	z, client, err := s.dnsContext(ctx, p, zoneID)
	if err != nil {
		return result, err
	}
	records, err := client.Records(ctx, providerZone(z))
	if err != nil {
		return result, err
	}
	allowed := make([]dnsprovider.Record, 0, len(records))
	for _, r := range records {
		if s.canDNS(ctx, p, z, r, "read") {
			allowed = append(allowed, r)
		}
	}
	result.Total = len(allowed)
	result.Capabilities = client.Capabilities()
	start := (page - 1) * size
	if start < len(allowed) {
		result.Items = allowed[start:min(start+size, len(allowed))]
	}
	return result, nil
}
func (s *Service) DNSPolicies(ctx context.Context, p Principal) ([]DNSPolicy, error) {
	current, e := s.CurrentPrincipal(ctx, p)
	if e != nil {
		return nil, e
	}
	p = current
	if !validPrincipal(p) {
		return nil, ErrDenied
	}
	rows := []DNSPolicy{}
	q := s.db.WithContext(ctx)
	if !p.Admin {
		q = s.dnsPolicySubject(ctx, p)
	}
	err := q.Order("id").Find(&rows).Error
	allowed := []DNSPolicy{}
	for _, g := range rows {
		if s.RestrictPrincipal(ctx, p, "dns_zone", fmt.Sprint(g.ZoneID), "read") {
			allowed = append(allowed, g)
		}
	}
	return allowed, err
}
func (s *Service) CreateDNSPolicy(ctx context.Context, p Principal, g DNSPolicy) (DNSPolicy, error) {
	current, e := s.CurrentPrincipal(ctx, p)
	if e != nil {
		return g, e
	}
	p = current
	if err := requireAdmin(p); err != nil {
		return g, err
	}
	var z DNSZone
	if s.db.WithContext(ctx).First(&z, g.ZoneID).Error != nil || !z.Active {
		return g, embedded.ErrInvalid
	}
	if g.ServiceID != 0 {
		var identity ServiceIdentity
		if g.UserID != 0 || s.db.WithContext(ctx).First(&identity, g.ServiceID).Error != nil || !identity.Active {
			return g, embedded.ErrInvalid
		}
		var owner model.User
		if s.db.WithContext(ctx).First(&owner, identity.OwnerID).Error != nil {
			return g, embedded.ErrInvalid
		}
	} else {
		var u model.User
		if g.UserID == 0 || s.db.WithContext(ctx).First(&u, g.UserID).Error != nil {
			return g, embedded.ErrInvalid
		}
	}
	n, err := policy.NormalizeName(g.Root)
	if err != nil || strings.Contains(n, "*") || !(policy.NameScope{Root: z.Name, IncludeRoot: true, Descendants: true}).Allows(n) {
		return g, embedded.ErrInvalid
	}
	g.Root = n
	if len(g.Types) == 0 || len(g.Actions) == 0 {
		return g, embedded.ErrInvalid
	}
	for _, action := range g.Actions {
		if !slices.Contains([]string{"read", "create", "update", "delete", "ddns"}, action) {
			return g, embedded.ErrInvalid
		}
	}
	for _, typ := range g.Types {
		if typ == "" || typ != strings.ToUpper(typ) || len(typ) > 16 {
			return g, embedded.ErrInvalid
		}
	}
	g.ID = 0
	g.CreatedAt = time.Time{}
	err = s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&g).Error; err != nil {
			return err
		}
		return tx.Create(&Audit{ActorID: p.UserID, Action: "dns.policy.create", Resource: fmt.Sprint(g.ID)}).Error
	})
	return g, err
}
func (s *Service) RevokeDNSPolicy(ctx context.Context, p Principal, id uint64) error {
	current, e := s.CurrentPrincipal(ctx, p)
	if e != nil {
		return e
	}
	p = current
	if err := requireAdmin(p); err != nil {
		return err
	}
	s.gate.Lock()
	defer s.gate.Unlock()
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var g DNSPolicy
		if err := tx.First(&g, id).Error; err != nil {
			return err
		}
		if err := tx.Delete(&g).Error; err != nil {
			return err
		}
		return tx.Create(&Audit{ActorID: p.UserID, Action: "dns.policy.revoke", Resource: fmt.Sprint(id)}).Error
	})
}
func (s *Service) ChangeDNS(ctx context.Context, p Principal, zoneID uint64, in DNSChange) (DNSOperation, error) {
	return s.changeDNS(ctx, p, zoneID, in, nil)
}

func (s *Service) changeDNS(ctx context.Context, p Principal, zoneID uint64, in DNSChange, binding *DDNSBinding) (DNSOperation, error) {
	if !validPrincipal(p) {
		return DNSOperation{}, ErrDenied
	}
	if len(in.OperationKey) < 16 || len(in.OperationKey) > 128 || !slices.Contains([]string{"create", "update", "delete"}, in.Action) {
		return DNSOperation{}, embedded.ErrInvalid
	}
	digestInput, _ := json.Marshal(struct {
		Zone   uint64
		Change DNSChange
	}{zoneID, in})
	digest := sha256.Sum256(digestInput)
	hashID := fnv.New32a()
	fmt.Fprintf(hashID, "%d:%s", zoneID, in.Record.ID)
	lock := &s.recordLocks[hashID.Sum32()%uint32(len(s.recordLocks))]
	lock.Lock()
	defer lock.Unlock()
	s.gate.RLock()
	defer s.gate.RUnlock()
	authAction := in.Action
	if binding != nil {
		authAction = "ddns"
		var current DDNSBinding
		if err := s.db.WithContext(ctx).First(&current, binding.ID).Error; err != nil || !current.Enabled || !current.NetworkEnabled || current.Version != binding.Version || current.OwnerID != p.UserID || current.ZoneID != zoneID || current.RecordID != in.Record.ID {
			return DNSOperation{}, ErrDenied
		}
		if current.SourceKind == "device" {
			var d DDNSDevice
			if current.ObservationDeviceID != binding.ObservationDeviceID || current.LastObserved != binding.LastObserved || !current.LastObservedAt.Equal(binding.LastObservedAt) || current.ObservationDeviceID == 0 || s.db.WithContext(ctx).First(&d, current.ObservationDeviceID).Error != nil || !d.Enabled || d.BindingID != current.ID {
				return DNSOperation{}, ErrDenied
			}
		}
	}
	var previous DNSOperation
	if err := s.db.WithContext(ctx).Where("actor_id = ? AND key = ?", p.UserID, in.OperationKey).First(&previous).Error; err == nil {
		if previous.ServiceID != p.ServiceID || previous.TokenID != p.TokenID || previous.IdentityVersion != p.IdentityVersion {
			return DNSOperation{}, ErrDenied
		}
		if previous.Digest != hex.EncodeToString(digest[:]) {
			return previous, embedded.ErrInvalid
		}
		var zone DNSZone
		if s.db.WithContext(ctx).First(&zone, previous.ZoneID).Error != nil || !s.canDNS(ctx, p, zone, dnsprovider.Record{Name: previous.Name, Type: previous.RecordType}, authAction) {
			return DNSOperation{}, ErrDenied
		}
		return previous, nil
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return previous, err
	}
	z, client, err := s.dnsContext(ctx, p, zoneID)
	if err != nil {
		return DNSOperation{}, err
	}
	record := in.Record
	if in.Action != "create" {
		if record.ID == "" {
			return DNSOperation{}, embedded.ErrInvalid
		}
		original, err := client.Get(ctx, providerZone(z), record.ID)
		if err != nil {
			return DNSOperation{}, err
		}
		if original.ID != record.ID || !s.canDNS(ctx, p, z, original, authAction) {
			return DNSOperation{}, ErrDenied
		}
		if in.Action == "delete" {
			record = original
		}
	}
	n, err := policy.NormalizeName(record.Name)
	if err != nil {
		return DNSOperation{}, embedded.ErrInvalid
	}
	record.Name = n
	if !s.canDNS(ctx, p, z, record, authAction) {
		return DNSOperation{}, ErrDenied
	}
	if in.Action != "delete" {
		cap := client.Capabilities()
		if !slices.Contains(cap.Types, record.Type) || record.TTL < cap.MinTTL && !(record.TTL == 1 && cap.AutoTTL) {
			return DNSOperation{}, embedded.ErrInvalid
		}
		if record.Type == "A" || record.Type == "AAAA" {
			ip, err := netip.ParseAddr(record.Value)
			if err != nil || ip.Is4() != (record.Type == "A") {
				return DNSOperation{}, embedded.ErrInvalid
			}
		}
	}
	if binding == nil && in.Action != "create" {
		var active []DDNSBinding
		if err := s.db.WithContext(ctx).Where("zone_id = ? AND record_id = ? AND enabled = ?", zoneID, record.ID, true).Find(&active).Error; err != nil {
			return DNSOperation{}, err
		}
		if len(active) > 0 && !in.PauseDDNS {
			return DNSOperation{}, ErrDDNSManaged
		}
		if len(active) > 0 {
			if err := s.db.WithContext(ctx).Model(&DDNSBinding{}).Where("zone_id = ? AND record_id = ?", zoneID, record.ID).Updates(map[string]any{"enabled": false, "state": "paused_manual", "version": gorm.Expr("version + 1")}).Error; err != nil {
				return DNSOperation{}, err
			}
		}
	}
	op := DNSOperation{ActorID: p.UserID, TokenID: p.TokenID, ServiceID: p.ServiceID, IdentityVersion: p.IdentityVersion, Key: in.OperationKey, Digest: hex.EncodeToString(digest[:]), ZoneID: z.ID, Action: in.Action, RecordID: record.ID, Name: record.Name, RecordType: record.Type, State: "applying"}
	if binding != nil {
		op.BindingID = binding.ID
	}
	if err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&op).Error; err != nil {
			return err
		}
		return tx.Create(&Audit{ActorID: p.UserID, Action: "dns." + in.Action, Resource: fmt.Sprintf("%d:%s", zoneID, record.ID)}).Error
	}); err != nil {
		return op, err
	}
	// The intent is durable before the external call. A process interruption leaves
	// applying, which is displayed as requiring reconciliation and is never replayed.
	var actual dnsprovider.Record
	switch in.Action {
	case "create":
		actual, err = client.Create(ctx, providerZone(z), record)
	case "update":
		actual, err = client.Update(ctx, providerZone(z), record)
	case "delete":
		err = client.Delete(ctx, providerZone(z), record.ID)
	}
	op.State = "succeeded"
	if err != nil {
		op.State = "failed"
		op.Error = "provider_rejected"
		if dnsprovider.Ambiguous(err) {
			op.State = "unknown"
			op.Error = "reconciliation_required"
		}
	} else if in.Action != "delete" {
		op.RecordID = actual.ID
		if actual.ID == "" || !s.canDNS(ctx, p, z, actual, authAction) {
			op.State = "unknown"
			op.Error = "unexpected_provider_result"
		}
	}
	// Persist outcomes independently of a disconnected browser; no second external
	// call occurs if recording the result fails.
	persistErr := s.db.WithContext(context.WithoutCancel(ctx)).Model(&DNSOperation{}).Where("id = ?", op.ID).Updates(map[string]any{"state": op.State, "error": op.Error, "record_id": op.RecordID}).Error
	if persistErr != nil {
		return op, persistErr
	}
	return op, nil
}

// dnsPolicySubject keeps explicitly approved platform service grants independent
// from the owner's delegated user grants while current identity scopes still intersect.
func (s *Service) dnsPolicySubject(ctx context.Context, p Principal) *gorm.DB {
	q := s.db.WithContext(ctx).Model(&DNSPolicy{})
	if p.ServiceID != 0 {
		return q.Where("service_id = ?", p.ServiceID)
	}
	return q.Where("user_id = ? AND service_id = 0", p.UserID)
}
