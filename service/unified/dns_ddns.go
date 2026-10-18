package unified

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"net/netip"
	"slices"
	"strings"
	"sync"
	"time"

	"github.com/certimate-go/certimate/pkg/embedded"
	"github.com/nezhahq/nezha/model"
	"github.com/nezhahq/nezha/service/unified/dnsprovider"
	"github.com/nezhahq/nezha/service/unified/policy"
	"gorm.io/gorm"
)

var ErrDDNSUnavailable = errors.New("DDNS host source is unavailable")
var ErrDDNSManaged = errors.New("record has an active DDNS binding; pause it explicitly before manual changes")

type DDNSBinding struct {
	TokenID             uint64 `json:"-" `
	ServiceID           uint64 `json:"service_id,omitempty"`
	IdentityVersion     uint64 `json:"identity_version,omitempty"`
	ApprovedHostOwnerID uint64 `json:"-"`

	ID                  uint64    `json:"id" gorm:"primaryKey"`
	ZoneID              uint64    `json:"zone_id" gorm:"uniqueIndex:ddns_record"`
	RecordID            string    `json:"record_id" gorm:"uniqueIndex:ddns_record"`
	Name                string    `json:"name"`
	RecordType          string    `json:"record_type"`
	OwnerID             uint64    `json:"owner_id" gorm:"index"`
	Version             uint64    `json:"version"`
	SourceKind          string    `json:"source_kind"`
	HostID              uint64    `json:"host_id"`
	Interface           string    `json:"interface"`
	Address             string    `json:"address"`
	StabilitySeconds    int       `json:"stability_seconds"`
	MaxAgeSeconds       int       `json:"max_age_seconds"`
	Enabled             bool      `json:"enabled"`
	NetworkEnabled      bool      `json:"network_enabled"`
	ObservationDeviceID uint64    `json:"-"`
	LastObserved        string    `json:"last_observed"`
	LastSuccess         string    `json:"last_success"`
	CandidateSince      time.Time `json:"candidate_since"`
	LastObservedAt      time.Time `json:"last_observed_at"`
	LastUpdatedAt       time.Time `json:"last_updated_at"`
	LastAttemptAt       time.Time `json:"last_attempt_at"`
	State               string    `json:"state"`
	Error               string    `json:"error,omitempty"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}

func (DDNSBinding) TableName() string { return "unified_ddns_bindings" }

type DDNSDevice struct {
	ID         uint64    `json:"id" gorm:"primaryKey"`
	BindingID  uint64    `json:"binding_id" gorm:"index"`
	TokenHash  string    `json:"-" gorm:"uniqueIndex"`
	Enabled    bool      `json:"enabled"`
	CreatedAt  time.Time `json:"created_at"`
	LastUsedAt time.Time `json:"last_used_at"`
}

func (DDNSDevice) TableName() string { return "unified_ddns_devices" }

type DDNSAddress struct {
	Interface  string `json:"interface"`
	Address    string `json:"address"`
	Temporary  bool   `json:"temporary"`
	Deprecated bool   `json:"deprecated"`
}
type DDNSHostSnapshot struct {
	OwnerID             uint64        `json:"owner_id"`
	Online              bool          `json:"online"`
	ObservedAt          time.Time     `json:"observed_at"`
	IPv4                string        `json:"ipv4"`
	IPv6                string        `json:"ipv6"`
	Addresses           []DDNSAddress `json:"addresses"`
	AddressesObservedAt time.Time     `json:"addresses_observed_at"`
}

// DDNSHostReport is written only by the authenticated Agent RPC adapter. It is
// independent of the default public-IP report so fresh heartbeat traffic cannot
// keep stale interface addresses eligible indefinitely.
type DDNSHostReport struct {
	HostID     uint64        `json:"host_id" gorm:"primaryKey"`
	ObservedAt time.Time     `json:"observed_at"`
	Addresses  []DDNSAddress `json:"addresses" gorm:"serializer:json"`
}

func (DDNSHostReport) TableName() string { return "unified_ddns_host_reports" }

func (s *Service) ReportDDNSHostAddresses(ctx context.Context, id uint64, addresses []DDNSAddress, observedAt time.Time) error {
	if id == 0 || len(addresses) > 4096 || observedAt.IsZero() || observedAt.After(time.Now().Add(time.Minute)) {
		return embedded.ErrInvalid
	}
	var host model.Server
	if s.db.WithContext(ctx).First(&host, id).Error != nil {
		return ErrDenied
	}
	clean := make([]DDNSAddress, 0, len(addresses))
	for _, a := range addresses {
		ip, err := netip.ParseAddr(a.Address)
		if err != nil || !ip.IsGlobalUnicast() || ip.Is4In6() || len(a.Interface) > 128 || a.Interface == "" {
			continue
		}
		a.Address = ip.String()
		clean = append(clean, a)
	}
	report := DDNSHostReport{HostID: id, ObservedAt: observedAt, Addresses: clean}
	err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var old DDNSHostReport
		if err := tx.First(&old, "host_id = ?", id).Error; err == nil {
			if !observedAt.After(old.ObservedAt) {
				return nil
			}
			return tx.Save(&report).Error
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		return tx.Create(&report).Error
	})
	if err == nil {
		s.NotifyDDNSHost(id)
	}
	return err
}
func (s *Service) DDNSHostAddressReport(ctx context.Context, id uint64) (DDNSHostReport, error) {
	var report DDNSHostReport
	err := s.db.WithContext(ctx).First(&report, "host_id = ?", id).Error
	return report, err
}

type DDNSHostSource func(context.Context, uint64) (DDNSHostSnapshot, error)

type DDNSWorker struct {
	service       *Service
	mu            sync.Mutex
	source        DDNSHostSource
	cancel        context.CancelFunc
	done          chan struct{}
	wake          chan struct{}
	bindingLocks  [64]sync.Mutex
	interval      time.Duration
	stopping      bool
	stopDone      chan struct{}
	effectContext context.Context
	effectCancel  context.CancelFunc
	active        sync.WaitGroup
}

func newDDNSWorker(s *Service) *DDNSWorker {
	ctx, cancel := context.WithCancel(context.Background())
	return &DDNSWorker{service: s, wake: make(chan struct{}, 1), interval: 30 * time.Second, effectContext: ctx, effectCancel: cancel}
}

// StartDDNS and StopDDNS are idempotent and owned by Dashboard's lifecycle.
func (s *Service) StartDDNS(ctx context.Context, source DDNSHostSource) error {
	w := s.ddns
	if w == nil || source == nil {
		return embedded.ErrInvalid
	}
	w.mu.Lock()
	defer w.mu.Unlock()
	if w.stopping {
		if w.stopDone == nil {
			return ErrDDNSUnavailable
		}
		select {
		case <-w.stopDone:
			w.cancel = nil
			w.stopDone = nil
		default:
			return ErrDDNSUnavailable
		}
	}
	if w.cancel != nil {
		return nil
	}
	if w.stopping {
		w.effectContext, w.effectCancel = context.WithCancel(context.Background())
		w.stopping = false
	}
	w.source = source
	run, cancel := context.WithCancel(ctx)
	w.cancel = cancel
	w.done = make(chan struct{})
	go func(done chan struct{}) {
		defer close(done)
		ticker := time.NewTicker(w.interval)
		defer ticker.Stop()
		for {
			select {
			case <-run.Done():
				return
			case <-ticker.C:
			case <-w.wake:
			}
			var rows []DDNSBinding
			if s.db.WithContext(run).Where("enabled = ? AND network_enabled = ?", true, true).Find(&rows).Error != nil {
				continue
			}
			for _, b := range rows {
				if run.Err() != nil {
					return
				}
				_, _ = s.executeDDNS(run, b.ID)
			}
		}
	}(w.done)
	select {
	case w.wake <- struct{}{}:
	default:
	}
	return nil
}
func (s *Service) StopDDNS() { _ = s.StopDDNSContext(context.Background()) }
func (s *Service) StopDDNSContext(ctx context.Context) error {
	if s.ddns == nil {
		return nil
	}
	w := s.ddns
	w.mu.Lock()
	if !w.stopping {
		w.stopping = true
		w.effectCancel()
		if w.cancel != nil {
			w.cancel()
		}
		w.stopDone = make(chan struct{})
		schedulerDone, stopDone := w.done, w.stopDone
		go func() {
			if schedulerDone != nil {
				<-schedulerDone
			}
			w.active.Wait()
			close(stopDone)
		}()
	}
	stopDone := w.stopDone
	w.mu.Unlock()
	select {
	case <-stopDone:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
func (s *Service) NotifyDDNSHost(_ uint64) {
	if s.ddns == nil {
		return
	}
	select {
	case s.ddns.wake <- struct{}{}:
	default:
	}
}
func (s *Service) ddnsHost(ctx context.Context, id uint64) (DDNSHostSnapshot, error) {
	if s.ddns == nil {
		return DDNSHostSnapshot{}, ErrDDNSUnavailable
	}
	s.ddns.mu.Lock()
	source := s.ddns.source
	s.ddns.mu.Unlock()
	if source == nil {
		return DDNSHostSnapshot{}, ErrDDNSUnavailable
	}
	return source(ctx, id)
}
func (s *Service) currentDNSPrincipal(ctx context.Context, p Principal) (Principal, error) {
	return s.CurrentPrincipal(ctx, p)
}
func (s *Service) canDDNS(ctx context.Context, p Principal, b DDNSBinding, actions ...string) bool {
	action := "execute"
	if len(actions) > 0 {
		action = actions[0]
	}
	if !s.RestrictPrincipal(ctx, p, "ddns_binding", fmt.Sprint(b.ID), action) {
		return false
	}
	if !validPrincipal(p) || (!p.Admin && p.UserID != b.OwnerID) {
		return false
	}
	var z DNSZone
	return s.db.WithContext(ctx).First(&z, b.ZoneID).Error == nil && s.canDNS(ctx, p, z, dnsprovider.Record{Name: b.Name, Type: b.RecordType}, "ddns")
}
func (s *Service) DDNSBindings(ctx context.Context, p Principal) ([]DDNSBinding, error) {
	p, err := s.currentDNSPrincipal(ctx, p)
	if err != nil {
		return nil, err
	}
	var all []DDNSBinding
	q := s.db.WithContext(ctx)
	if !p.Admin {
		q = q.Where("owner_id = ?", p.UserID)
	}
	if err := q.Order("id").Find(&all).Error; err != nil {
		return nil, err
	}
	out := []DDNSBinding{}
	for _, b := range all {
		if s.canDDNS(ctx, p, b, "read") {
			out = append(out, b)
		}
	}
	return out, nil
}
func (s *Service) DDNSHostAddresses(ctx context.Context, p Principal, id uint64) (DDNSHostSnapshot, error) {
	p, err := s.currentDNSPrincipal(ctx, p)
	if err != nil {
		return DDNSHostSnapshot{}, err
	}
	var host model.Server
	if s.db.WithContext(ctx).First(&host, id).Error != nil || (!p.Admin && host.UserID != p.UserID) {
		return DDNSHostSnapshot{}, ErrDenied
	}
	return s.ddnsHost(ctx, id)
}
func (s *Service) validateDDNS(ctx context.Context, p Principal, b *DDNSBinding) error {
	if b.SourceKind != "host" && b.SourceKind != "device" || b.RecordID == "" || b.ZoneID == 0 || b.StabilitySeconds < 0 || b.StabilitySeconds > 86400 || b.MaxAgeSeconds < 30 || b.MaxAgeSeconds > 86400 {
		return embedded.ErrInvalid
	}
	z, c, err := s.dnsContext(ctx, p, b.ZoneID)
	if err != nil {
		return err
	}
	r, err := c.Get(ctx, providerZone(z), b.RecordID)
	if err != nil {
		return err
	}
	if r.ID != b.RecordID || (r.Type != "A" && r.Type != "AAAA") || !s.canDNS(ctx, p, z, r, "ddns") {
		return ErrDenied
	}
	normalized, err := policy.NormalizeName(r.Name)
	if err != nil {
		return embedded.ErrInvalid
	}
	b.Name = normalized
	b.RecordType = r.Type
	if b.SourceKind == "host" {
		var host model.Server
		if s.db.WithContext(ctx).First(&host, b.HostID).Error != nil || (!p.Admin && host.UserID != p.UserID) {
			return ErrDenied
		}
		if b.Address != "" {
			ip, e := netip.ParseAddr(b.Address)
			if e != nil || ip.Is4() != (r.Type == "A") {
				return embedded.ErrInvalid
			}
			b.Address = ip.String()
		}
	} else if b.HostID != 0 || b.Interface != "" || b.Address != "" {
		return embedded.ErrInvalid
	}
	return nil
}
func (s *Service) SaveDDNSBinding(ctx context.Context, p Principal, id uint64, in DDNSBinding) (DDNSBinding, error) {
	p, err := s.currentDNSPrincipal(ctx, p)
	if err != nil {
		return DDNSBinding{}, err
	}
	s.gate.Lock()
	defer s.gate.Unlock()
	var old DDNSBinding
	if id == 0 && !s.RestrictPrincipal(ctx, p, "ddns_binding", "new", "create") {
		return DDNSBinding{}, ErrDenied
	}
	if id > 0 {
		if err := s.db.WithContext(ctx).First(&old, id).Error; err != nil {
			return old, err
		}
		if !s.canDDNS(ctx, p, old, "write") {
			return DDNSBinding{}, ErrDenied
		}
		if in.Version != old.Version {
			return old, embedded.ErrInvalid
		}
		if in.ZoneID != old.ZoneID || in.RecordID != old.RecordID {
			return old, embedded.ErrInvalid
		}
	}
	if in.MaxAgeSeconds == 0 {
		in.MaxAgeSeconds = 300
	}
	if err := s.validateDDNS(ctx, p, &in); err != nil {
		return DDNSBinding{}, err
	}
	if in.Enabled && s.legacyDDNSOwnsName(ctx, in.Name, in.RecordType) {
		return DDNSBinding{}, ErrDDNSManaged
	}
	b := DDNSBinding{ID: id, ZoneID: in.ZoneID, RecordID: in.RecordID, Name: in.Name, RecordType: in.RecordType, OwnerID: p.UserID, TokenID: p.TokenID, ServiceID: p.ServiceID, IdentityVersion: p.IdentityVersion, Version: old.Version + 1, SourceKind: in.SourceKind, HostID: in.HostID, Interface: in.Interface, Address: in.Address, StabilitySeconds: in.StabilitySeconds, MaxAgeSeconds: in.MaxAgeSeconds, Enabled: in.Enabled, NetworkEnabled: in.NetworkEnabled, State: "pending", LastSuccess: old.LastSuccess, LastUpdatedAt: old.LastUpdatedAt, CreatedAt: old.CreatedAt}
	if id > 0 {
		b.OwnerID = old.OwnerID
		b.TokenID = old.TokenID
		b.ServiceID = old.ServiceID
		b.IdentityVersion = old.IdentityVersion
		b.ApprovedHostOwnerID = old.ApprovedHostOwnerID
		b.LastAttemptAt = old.LastAttemptAt
		if old.SourceKind == b.SourceKind && old.HostID == b.HostID && old.Interface == b.Interface && old.Address == b.Address {
			b.LastObserved = old.LastObserved
			b.LastObservedAt = old.LastObservedAt
			b.CandidateSince = old.CandidateSince
			b.ObservationDeviceID = old.ObservationDeviceID
		}
	}
	if !b.Enabled || !b.NetworkEnabled {
		b.State = "paused"
	}
	err = s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var e error
		if id == 0 {
			e = tx.Create(&b).Error
		} else {
			e = tx.Save(&b).Error
		}
		if e != nil {
			return e
		}
		return tx.Create(&Audit{ActorID: p.UserID, Action: "ddns.binding.save", Resource: fmt.Sprint(b.ID)}).Error
	})
	s.NotifyDDNSHost(b.HostID)
	return b, err
}

// AssignDDNSService is the explicit audited handover from a user's derived
// automation to a platform service identity. Exact DNS grants and identity
// scopes must already exist; an administrator cannot silently manufacture them.
func (s *Service) AssignDDNSService(ctx context.Context, p Principal, id, serviceID, version uint64) (DDNSBinding, error) {
	p, err := s.CurrentPrincipal(ctx, p)
	if err != nil {
		return DDNSBinding{}, err
	}
	if err := requireAdmin(p); err != nil {
		return DDNSBinding{}, err
	}
	s.gate.Lock()
	defer s.gate.Unlock()
	var b DDNSBinding
	var identity ServiceIdentity
	if s.db.WithContext(ctx).First(&b, id).Error != nil || b.Version != version || serviceID == 0 || s.db.WithContext(ctx).First(&identity, serviceID).Error != nil || !identity.Active {
		return b, embedded.ErrInvalid
	}
	service, err := s.CurrentPrincipal(ctx, Principal{UserID: identity.OwnerID, ServiceID: identity.ID, IdentityVersion: identity.Version})
	if err != nil {
		return b, err
	}
	candidate := b
	candidate.OwnerID = identity.OwnerID
	candidate.ServiceID = identity.ID
	candidate.IdentityVersion = identity.Version
	candidate.TokenID = 0
	if !s.canDDNS(ctx, service, candidate) {
		return b, ErrDenied
	}
	if b.SourceKind == "host" {
		var host model.Server
		if s.db.WithContext(ctx).First(&host, b.HostID).Error != nil {
			return b, ErrDenied
		}
		candidate.ApprovedHostOwnerID = host.UserID
	}
	candidate.Version++
	candidate.State = "pending"
	candidate.Error = ""
	candidate.LastObserved = ""
	candidate.LastObservedAt = time.Time{}
	candidate.CandidateSince = time.Time{}
	err = s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(&candidate).Error; err != nil {
			return err
		}
		if err := tx.Model(&DDNSDevice{}).Where("binding_id = ?", id).Update("enabled", false).Error; err != nil {
			return err
		}
		return tx.Create(&Audit{ActorID: p.UserID, Action: "ddns.service.assign", Resource: fmt.Sprintf("%d:%d", id, serviceID)}).Error
	})
	s.NotifyDDNSHost(candidate.HostID)
	return candidate, err
}

func ddnsHostOwnerAllowed(p Principal, b DDNSBinding, owner uint64) bool {
	if p.ServiceID != 0 && b.ApprovedHostOwnerID != 0 {
		return b.ApprovedHostOwnerID == owner
	}
	return p.Admin || owner == p.UserID
}
func (s *Service) DeleteDDNSBinding(ctx context.Context, p Principal, id uint64) error {
	p, err := s.currentDNSPrincipal(ctx, p)
	if err != nil {
		return err
	}
	s.gate.Lock()
	defer s.gate.Unlock()
	var b DDNSBinding
	if s.db.WithContext(ctx).First(&b, id).Error != nil || !s.canDDNS(ctx, p, b, "delete") {
		return ErrDenied
	}
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("binding_id = ?", id).Delete(&DDNSDevice{}).Error; err != nil {
			return err
		}
		if err := tx.Delete(&b).Error; err != nil {
			return err
		}
		return tx.Create(&Audit{ActorID: p.UserID, Action: "ddns.binding.delete", Resource: fmt.Sprint(id)}).Error
	})
}
func (s *Service) RunDDNS(ctx context.Context, p Principal, id uint64) (DDNSBinding, error) {
	p, err := s.currentDNSPrincipal(ctx, p)
	if err != nil {
		return DDNSBinding{}, err
	}
	var b DDNSBinding
	if s.db.WithContext(ctx).First(&b, id).Error != nil || !s.canDDNS(ctx, p, b) {
		return b, ErrDenied
	}
	return s.executeDDNS(ctx, id)
}
func selectDDNSAddress(b DDNSBinding, h DDNSHostSnapshot) string {
	if b.Interface == "" && b.Address == "" {
		v := h.IPv6
		if b.RecordType == "A" {
			v = h.IPv4
		}
		if validDDNSIP(v, b.RecordType) {
			ip, _ := netip.ParseAddr(v)
			return ip.String()
		}
		return ""
	}
	candidates := []string{}
	for _, a := range h.Addresses {
		if a.Deprecated || a.Temporary || (b.Interface != "" && a.Interface != b.Interface) || !validDDNSIP(a.Address, b.RecordType) {
			continue
		}
		ip, _ := netip.ParseAddr(a.Address)
		value := ip.String()
		if b.Address != "" && value != b.Address {
			continue
		}
		candidates = append(candidates, value)
	}
	if slices.Contains(candidates, b.LastSuccess) {
		return b.LastSuccess
	}
	slices.Sort(candidates)
	if len(candidates) > 0 {
		return candidates[0]
	}
	return ""
}
func validDDNSIP(v, typ string) bool {
	ip, err := netip.ParseAddr(v)
	return err == nil && ip.Zone() == "" && ip.IsGlobalUnicast() && ip.Is4() == (typ == "A") && !ip.Is4In6()
}
func (s *Service) saveDDNSState(ctx context.Context, b DDNSBinding, state, code string) (DDNSBinding, error) {
	b.State = state
	b.Error = code
	err := s.db.WithContext(context.WithoutCancel(ctx)).Model(&DDNSBinding{}).Where("id = ? AND version = ?", b.ID, b.Version).Where("source_kind <> ? OR last_observed_at = ?", "device", b.LastObservedAt).Updates(map[string]any{"last_observed": b.LastObserved, "last_success": b.LastSuccess, "candidate_since": b.CandidateSince, "last_observed_at": b.LastObservedAt, "last_updated_at": b.LastUpdatedAt, "last_attempt_at": b.LastAttemptAt, "state": state, "error": code}).Error
	return b, err
}
func (s *Service) executeDDNS(ctx context.Context, id uint64) (DDNSBinding, error) {
	w := s.ddns
	if w == nil {
		return DDNSBinding{}, ErrDDNSUnavailable
	}
	w.mu.Lock()
	if w.stopping {
		w.mu.Unlock()
		return DDNSBinding{}, ErrDDNSUnavailable
	}
	w.active.Add(1)
	lifetime := w.effectContext
	w.mu.Unlock()
	defer w.active.Done()
	run, cancel := context.WithCancel(ctx)
	detach := context.AfterFunc(lifetime, cancel)
	defer func() { detach(); cancel() }()
	ctx = run
	lock := &w.bindingLocks[id%64]
	lock.Lock()
	defer lock.Unlock()
	var b DDNSBinding
	if err := s.db.WithContext(ctx).First(&b, id).Error; err != nil {
		return b, err
	}
	if !b.Enabled || !b.NetworkEnabled {
		return s.saveDDNSState(ctx, b, "paused", "")
	}
	p, err := s.currentDNSPrincipal(ctx, Principal{UserID: b.OwnerID, TokenID: b.TokenID, ServiceID: b.ServiceID, IdentityVersion: b.IdentityVersion})
	if err != nil || !s.canDDNS(ctx, p, b) {
		return s.saveDDNSState(ctx, b, "blocked", "authorization_revoked")
	}
	if b.SourceKind == "device" {
		var d DDNSDevice
		if b.ObservationDeviceID == 0 || s.db.WithContext(ctx).First(&d, b.ObservationDeviceID).Error != nil || !d.Enabled || d.BindingID != b.ID {
			return s.saveDDNSState(ctx, b, "waiting", "device_report_required")
		}
	}
	now := time.Now().UTC()
	value := b.LastObserved
	observed := b.LastObservedAt
	if b.SourceKind == "host" {
		h, e := s.ddnsHost(ctx, b.HostID)
		var host model.Server
		if e != nil || s.db.WithContext(ctx).First(&host, b.HostID).Error != nil || host.UserID != h.OwnerID || !ddnsHostOwnerAllowed(p, b, host.UserID) || !ddnsHostOwnerAllowed(p, b, h.OwnerID) {
			return s.saveDDNSState(ctx, b, "blocked", "host_unavailable")
		}
		if !h.Online {
			return s.saveDDNSState(ctx, b, "waiting", "host_offline")
		}
		value = selectDDNSAddress(b, h)
		observed = h.ObservedAt
		if b.Interface != "" || b.Address != "" {
			observed = h.AddressesObservedAt
		}
	}
	if observed.IsZero() || observed.After(now.Add(time.Minute)) || now.Sub(observed) > time.Duration(b.MaxAgeSeconds)*time.Second {
		return s.saveDDNSState(ctx, b, "waiting", "source_stale")
	}
	if !validDDNSIP(value, b.RecordType) {
		return s.saveDDNSState(ctx, b, "waiting", "no_eligible_address")
	}
	if b.LastObserved != value || b.CandidateSince.IsZero() {
		b.CandidateSince = now
	}
	b.LastObserved = value
	b.LastObservedAt = observed
	if now.Sub(b.CandidateSince) < time.Duration(b.StabilitySeconds)*time.Second {
		return s.saveDDNSState(ctx, b, "stabilizing", "")
	}
	z, c, err := s.dnsContext(ctx, p, b.ZoneID)
	if err != nil {
		return s.saveDDNSState(ctx, b, "blocked", "credential_unavailable")
	}
	r, err := c.Get(ctx, providerZone(z), b.RecordID)
	if err != nil {
		return s.saveDDNSState(ctx, b, "failed", "record_read_failed")
	}
	normalized, nameErr := policy.NormalizeName(r.Name)
	if nameErr != nil || r.ID != b.RecordID || normalized != b.Name || r.Type != b.RecordType || !s.canDNS(ctx, p, z, r, "ddns") {
		return s.saveDDNSState(ctx, b, "blocked", "record_identity_changed")
	}
	if r.Value == value {
		b.LastSuccess = value
		// A provider readback may reconcile an intent left by a process crash.
		// This records observed state, without claiming when the write happened.
		if err := s.db.WithContext(ctx).Model(&DNSOperation{}).Where("actor_id = ? AND zone_id = ? AND record_id = ? AND action = ? AND state IN ?", b.OwnerID, b.ZoneID, b.RecordID, "update", []string{"applying", "unknown"}).Updates(map[string]any{"state": "reconciled", "error": ""}).Error; err != nil {
			return b, err
		}
		return s.saveDDNSState(ctx, b, "current", "")
	}
	// An uncertain write must be reconciled by readback. A differing remote value
	// does not prove that repeating an interrupted provider operation is safe.
	if b.State == "unknown" || b.Error == "reconciliation_required" {
		return s.saveDDNSState(ctx, b, "unknown", "reconciliation_required")
	}
	var pending int64
	if err := s.db.WithContext(ctx).Model(&DNSOperation{}).Where("zone_id = ? AND record_id = ? AND action = ? AND state IN ?", b.ZoneID, b.RecordID, "update", []string{"applying", "unknown"}).Count(&pending).Error; err != nil {
		return b, err
	}
	if pending > 0 {
		return s.saveDDNSState(ctx, b, "unknown", "reconciliation_required")
	}
	r.Value = value
	b.LastAttemptAt = now
	op, err := s.changeDNS(ctx, p, b.ZoneID, DNSChange{OperationKey: fmt.Sprintf("ddns-%d-%d-%d", b.ID, b.Version, now.UnixNano()), Action: "update", Record: r}, &b)
	if err != nil {
		return s.saveDDNSState(ctx, b, "blocked", "execution_denied")
	}
	if op.State != "succeeded" {
		return s.saveDDNSState(ctx, b, op.State, op.Error)
	}
	b.LastSuccess = value
	b.LastUpdatedAt = now
	return s.saveDDNSState(ctx, b, "current", "")
}
func (s *Service) CreateDDNSDevice(ctx context.Context, p Principal, bindingID uint64) (DDNSDevice, string, error) {
	p, err := s.currentDNSPrincipal(ctx, p)
	if err != nil {
		return DDNSDevice{}, "", err
	}
	s.gate.Lock()
	defer s.gate.Unlock()
	var b DDNSBinding
	if s.db.WithContext(ctx).First(&b, bindingID).Error != nil || b.SourceKind != "device" || !s.canDDNS(ctx, p, b, "write") {
		return DDNSDevice{}, "", ErrDenied
	}
	raw := make([]byte, 32)
	if _, err := rand.Read(raw); err != nil {
		return DDNSDevice{}, "", err
	}
	token := hex.EncodeToString(raw)
	sum := sha256.Sum256([]byte(token))
	d := DDNSDevice{BindingID: b.ID, TokenHash: hex.EncodeToString(sum[:]), Enabled: true}
	err = s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&d).Error; err != nil {
			return err
		}
		return tx.Create(&Audit{ActorID: p.UserID, Action: "ddns.device.create", Resource: fmt.Sprint(d.ID)}).Error
	})
	return d, token, err
}
func (s *Service) DDNSDevices(ctx context.Context, p Principal) ([]DDNSDevice, error) {
	bindings, err := s.DDNSBindings(ctx, p)
	if err != nil {
		return nil, err
	}
	ids := []uint64{}
	for _, b := range bindings {
		ids = append(ids, b.ID)
	}
	rows := []DDNSDevice{}
	err = s.db.WithContext(ctx).Where("binding_id IN ?", ids).Find(&rows).Error
	return rows, err
}
func (s *Service) RevokeDDNSDevice(ctx context.Context, p Principal, id uint64) error {
	p, err := s.currentDNSPrincipal(ctx, p)
	if err != nil {
		return err
	}
	s.gate.Lock()
	defer s.gate.Unlock()
	var d DDNSDevice
	var b DDNSBinding
	if s.db.WithContext(ctx).First(&d, id).Error != nil || s.db.WithContext(ctx).First(&b, d.BindingID).Error != nil || !s.canDDNS(ctx, p, b, "write") {
		return ErrDenied
	}
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&d).Update("enabled", false).Error; err != nil {
			return err
		}
		return tx.Create(&Audit{ActorID: p.UserID, Action: "ddns.device.revoke", Resource: fmt.Sprint(id)}).Error
	})
}

// ReportDDNSDevice accepts only an IP. Account, name, type and record identity are
// immutable server-side binding facts and cannot be supplied by a device token.
func (s *Service) ReportDDNSDevice(ctx context.Context, token, value string) (DDNSBinding, error) {
	if len(token) != 64 || strings.TrimSpace(token) != token {
		return DDNSBinding{}, ErrDenied
	}
	sum := sha256.Sum256([]byte(token))
	s.gate.RLock()
	var d DDNSDevice
	var b DDNSBinding
	if s.db.WithContext(ctx).Where("token_hash = ? AND enabled = ?", hex.EncodeToString(sum[:]), true).First(&d).Error != nil || s.db.WithContext(ctx).First(&b, d.BindingID).Error != nil || b.SourceKind != "device" || !b.Enabled || !b.NetworkEnabled {
		s.gate.RUnlock()
		return b, ErrDenied
	}
	p, err := s.currentDNSPrincipal(ctx, Principal{UserID: b.OwnerID, TokenID: b.TokenID, ServiceID: b.ServiceID, IdentityVersion: b.IdentityVersion})
	if err != nil || !s.canDDNS(ctx, p, b) {
		s.gate.RUnlock()
		return b, ErrDenied
	}
	if !validDDNSIP(value, b.RecordType) {
		s.gate.RUnlock()
		return b, embedded.ErrInvalid
	}
	ip, _ := netip.ParseAddr(value)
	value = ip.String()
	now := time.Now().UTC()
	changes := map[string]any{"observation_device_id": d.ID, "last_observed": value, "last_observed_at": now}
	if b.LastObserved != value || b.CandidateSince.IsZero() {
		changes["candidate_since"] = now
	}
	err = s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&b).Updates(changes).Error; err != nil {
			return err
		}
		return tx.Model(&d).Update("last_used_at", now).Error
	})
	s.gate.RUnlock()
	if err != nil {
		return b, err
	}
	return s.executeDDNS(ctx, b.ID)
}

// HasDDNSName lets the legacy DDNS dispatcher avoid a second automatic writer.
func (s *Service) HasDDNSName(ctx context.Context, name string) bool {
	name = strings.ToLower(strings.TrimSuffix(name, "."))
	var count int64
	return s.db.WithContext(ctx).Model(&DDNSBinding{}).Where("name = ? AND enabled = ?", name, true).Count(&count).Error == nil && count > 0
}
func (s *Service) legacyDDNSOwnsName(ctx context.Context, name, typ string) bool {
	if !s.db.Migrator().HasTable(&model.Server{}) || !s.db.Migrator().HasTable(&model.DDNSProfile{}) {
		return false
	}
	var hosts []model.Server
	if s.db.WithContext(ctx).Where(&model.Server{EnableDDNS: true}).Find(&hosts).Error != nil {
		return true
	}
	for i := range hosts {
		host := &hosts[i]
		for _, id := range host.DDNSProfiles {
			var profile model.DDNSProfile
			if s.db.WithContext(ctx).First(&profile, id).Error != nil {
				continue
			}
			if typ == "A" && profile.EnableIPv4 != nil && !*profile.EnableIPv4 || typ == "AAAA" && profile.EnableIPv6 != nil && !*profile.EnableIPv6 {
				continue
			}
			domains := profile.Domains
			if override, ok := host.OverrideDDNSDomains[id]; ok && len(override) > 0 {
				domains = override
			}
			for _, domain := range domains {
				if strings.EqualFold(strings.TrimSuffix(domain, "."), name) {
					return true
				}
			}
		}
	}
	return false
}
