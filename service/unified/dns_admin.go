package unified

import (
	"context"
	"encoding/json"
	"fmt"
	"hash/fnv"
	"slices"

	"github.com/certimate-go/certimate/pkg/embedded"
	"github.com/nezhahq/nezha/service/unified/dnsprovider"
	"gorm.io/gorm"
)

type DNSReconciliation struct {
	Operation DNSOperation       `json:"operation"`
	Record    dnsprovider.Record `json:"record"`
}

// ReconcileDNSOperation acknowledges the current provider-fetched record. This
// explicit operator action clears uncertainty without replaying the original
// mutation or representing readback as proof of the original outcome.
func (s *Service) ReconcileDNSOperation(ctx context.Context, p Principal, id uint64, recordID string) (DNSReconciliation, error) {
	p, err := s.CurrentPrincipal(ctx, p)
	if err != nil {
		return DNSReconciliation{}, err
	}
	var hint DNSOperation
	if err := s.db.WithContext(ctx).First(&hint, id).Error; err != nil {
		return DNSReconciliation{}, err
	}
	hashID := fnv.New32a()
	fmt.Fprintf(hashID, "%d:%s", hint.ZoneID, hint.RecordID)
	lock := &s.recordLocks[hashID.Sum32()%uint32(len(s.recordLocks))]
	lock.Lock()
	defer lock.Unlock()
	s.gate.RLock()
	defer s.gate.RUnlock()
	var op DNSOperation
	if s.db.WithContext(ctx).First(&op, id).Error != nil || (!p.Admin && op.ActorID != p.UserID) {
		return DNSReconciliation{}, ErrDenied
	}
	if !slices.Contains([]string{"applying", "unknown"}, op.State) {
		return DNSReconciliation{}, embedded.ErrInvalid
	}
	z, c, err := s.dnsContext(ctx, p, op.ZoneID)
	if err != nil {
		return DNSReconciliation{}, err
	}
	action := op.Action
	if op.BindingID != 0 {
		var b DDNSBinding
		if s.db.WithContext(ctx).First(&b, op.BindingID).Error != nil || b.ZoneID != op.ZoneID || b.RecordID != op.RecordID || !s.canDDNS(ctx, p, b) {
			return DNSReconciliation{}, ErrDenied
		}
		action = "ddns"
	}
	if !s.canDNS(ctx, p, z, dnsprovider.Record{Name: op.Name, Type: op.RecordType}, action) {
		return DNSReconciliation{}, ErrDenied
	}
	if op.RecordID != "" {
		if recordID != "" && recordID != op.RecordID {
			return DNSReconciliation{}, embedded.ErrInvalid
		}
		recordID = op.RecordID
	}
	if recordID == "" {
		return DNSReconciliation{}, embedded.ErrInvalid
	}
	r, err := c.Get(ctx, providerZone(z), recordID)
	if err != nil {
		return DNSReconciliation{}, err
	}
	if r.ID != recordID || r.Name != op.Name || r.Type != op.RecordType || !s.canDNS(ctx, p, z, r, action) {
		return DNSReconciliation{}, ErrDenied
	}
	op.State = "reconciled"
	op.Error = ""
	op.RecordID = recordID
	err = s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&DNSOperation{}).Where("id = ? AND state IN ?", op.ID, []string{"applying", "unknown"}).Updates(map[string]any{"state": op.State, "error": "", "record_id": recordID}).Error; err != nil {
			return err
		}
		if err := tx.Model(&DDNSBinding{}).Where("zone_id = ? AND record_id = ? AND state = ?", op.ZoneID, recordID, "unknown").Updates(map[string]any{"state": "pending", "error": ""}).Error; err != nil {
			return err
		}
		return tx.Create(&Audit{ActorID: p.UserID, Action: "dns.operation.reconcile", Resource: fmt.Sprint(op.ID)}).Error
	})
	return DNSReconciliation{Operation: op, Record: r}, err
}

func (s *Service) UpdateCredential(ctx context.Context, p Principal, id string, version int, enabled bool, config json.RawMessage) (embedded.Resource, error) {
	current, e := s.CurrentPrincipal(ctx, p)
	if e != nil {
		return embedded.Resource{}, e
	}
	p = current
	if err := requireAdmin(p); err != nil {
		return embedded.Resource{}, err
	}
	s.gate.Lock()
	defer s.gate.Unlock()
	return s.engine.UpdateCredential(ctx, fmt.Sprint(p.UserID), id, version, enabled, config)
}

func (s *Service) UpdateDNSAccount(ctx context.Context, p Principal, id uint64, in DNSAccount) (DNSAccount, error) {
	current, e := s.CurrentPrincipal(ctx, p)
	if e != nil {
		return DNSAccount{}, e
	}
	p = current
	if err := requireAdmin(p); err != nil {
		return DNSAccount{}, err
	}
	s.gate.Lock()
	defer s.gate.Unlock()
	var account DNSAccount
	if err := s.db.WithContext(ctx).First(&account, id).Error; err != nil {
		return account, err
	}
	credential, err := s.engine.CredentialInfo(ctx, account.CredentialID)
	if err != nil {
		return account, err
	}
	if in.CredentialVersion != credential.Version {
		return account, embedded.ErrInvalid
	}
	account.CredentialVersion = in.CredentialVersion
	account.Enabled = in.Enabled
	if in.Name != "" {
		account.Name = in.Name
	}
	err = s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(&account).Error; err != nil {
			return err
		}
		return tx.Create(&Audit{ActorID: p.UserID, Action: "dns.account.update", Resource: fmt.Sprint(id)}).Error
	})
	return account, err
}

func (s *Service) DNSOperations(ctx context.Context, p Principal, limit int) ([]DNSOperation, error) {
	current, e := s.CurrentPrincipal(ctx, p)
	if e != nil {
		return nil, e
	}
	p = current
	if !validPrincipal(p) {
		return nil, ErrDenied
	}
	if limit < 1 || limit > 200 {
		return nil, embedded.ErrInvalid
	}
	rows := []DNSOperation{}
	// Filter against current name grants before applying the visible limit.
	var all []DNSOperation
	q := s.db.WithContext(ctx).Order("id DESC")
	if !p.Admin {
		q = q.Where("actor_id = ?", p.UserID)
	}
	if p.ServiceID != 0 {
		q = q.Where("service_id = ?", p.ServiceID)
	}
	if err := q.Find(&all).Error; err != nil {
		return nil, err
	}
	for _, op := range all {
		var z DNSZone
		if s.db.WithContext(ctx).First(&z, op.ZoneID).Error == nil && s.canDNS(ctx, p, z, dnsprovider.Record{Name: op.Name, Type: op.RecordType}, "read") {
			rows = append(rows, op)
			if len(rows) == limit {
				break
			}
		}
	}
	return rows, nil
}
