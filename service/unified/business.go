package unified

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/certimate-go/certimate/pkg/embedded"
	"gorm.io/gorm/clause"
)

func (s *Service) DownloadCertificate(ctx context.Context, p Principal, id string, opts embedded.NativeCertificateDownloadOptions) ([]byte, error) {
	if !s.Can(ctx, p, "certificate", id, "download_key") {
		return nil, ErrDenied
	}
	return s.engine.NativeCertificateDownload(WithPrincipal(ctx, p), fmt.Sprint(p.UserID), id, opts)
}
func (s *Service) RevokeCertificate(ctx context.Context, p Principal, id string) error {
	if !s.Can(ctx, p, "certificate", id, "revoke") {
		return ErrDenied
	}
	return s.engine.NativeRevokeCertificate(WithPrincipal(ctx, p), fmt.Sprint(p.UserID), id)
}
func (s *Service) TestNotification(ctx context.Context, p Principal, id, provider string) error {
	if !s.Can(ctx, p, "credential", id, "use") {
		return ErrDenied
	}
	info, err := s.engine.CredentialInfo(ctx, id)
	if err != nil {
		return err
	}
	return s.engine.NativeTestNotification(WithPrincipal(ctx, p), embedded.CredentialUse{ActorID: p.UserID, CredentialID: id, Version: info.Version, Purpose: "notification.test", ResourceKind: "credential", ResourceID: id}, provider)
}
func (s *Service) Statistics(ctx context.Context, p Principal) (map[string]int, error) {
	certs, err := s.resourceScope(ctx, p, "certificate")
	if err != nil {
		return nil, err
	}
	workflows, err := s.resourceScope(ctx, p, "workflow")
	if err != nil {
		return nil, err
	}
	return s.engine.NativeStatistics(ctx, certs, workflows)
}
func (s *Service) CronJobs(ctx context.Context, p Principal) ([]map[string]string, error) {
	p, err := s.CurrentPrincipal(ctx, p)
	if err != nil {
		return nil, err
	}
	if requireAdmin(p) != nil {
		return nil, ErrDenied
	}
	return s.engine.NativeCronJobs(ctx)
}
func (s *Service) SystemLogs(ctx context.Context, p Principal, page, perPage int) (embedded.NativePage, error) {
	p, err := s.CurrentPrincipal(ctx, p)
	if err != nil {
		return embedded.NativePage{}, err
	}
	if requireAdmin(p) != nil {
		return embedded.NativePage{}, ErrDenied
	}
	return s.engine.NativeSystemLogs(ctx, page, perPage)
}
func (s *Service) authorizeStandaloneCredential(ctx context.Context, use embedded.CredentialUse) error {
	p, ok := principalFromContext(ctx)
	if !ok || p.UserID != use.ActorID || use.ResourceID != use.CredentialID {
		return ErrDenied
	}
	action := map[string]string{"notification.test": "use", "credential.clone": "clone", "credential.export": "export"}[use.Purpose]
	if action == "" || !s.Can(ctx, p, "credential", use.CredentialID, action) {
		return ErrDenied
	}
	return nil
}
func (s *Service) ExportCredential(ctx context.Context, p Principal, id string) (map[string]any, error) {
	if !s.Can(ctx, p, "credential", id, "export") {
		return nil, ErrDenied
	}
	info, err := s.engine.CredentialInfo(ctx, id)
	if err != nil {
		return nil, err
	}
	if err = s.db.WithContext(ctx).Create(&Audit{ActorID: p.UserID, Action: "credential.export", Resource: id}).Error; err != nil {
		return nil, err
	}
	return s.engine.ResolveCredential(WithPrincipal(ctx, p), embedded.CredentialUse{ActorID: p.UserID, CredentialID: id, Version: info.Version, Purpose: "credential.export", ResourceKind: "credential", ResourceID: id})
}
func (s *Service) CloneCredential(ctx context.Context, p Principal, id, name, op string) (map[string]any, error) {
	p, err := s.CurrentPrincipal(ctx, p)
	if err != nil {
		return nil, err
	}
	if !s.Can(ctx, p, "credential", id, "read") || !s.Can(ctx, p, "credential", id, "clone") || !s.Can(ctx, p, "credential", "new", "create") {
		return nil, ErrDenied
	}
	info, err := s.engine.CredentialInfo(ctx, id)
	if err != nil {
		return nil, err
	}
	if op == "" {
		op, err = operationID()
		if err != nil {
			return nil, err
		}
	}
	if len(op) > 64 {
		return nil, embedded.ErrInvalid
	}
	raw, _ := json.Marshal([]any{id, name, info.Version})
	sum := sha256.Sum256(raw)
	digest := hex.EncodeToString(sum[:])
	s.gate.Lock()
	defer s.gate.Unlock()
	intent := ResourceIntent{OperationID: op, ActorID: p.UserID, TokenID: p.TokenID, ServiceID: p.ServiceID, IdentityVersion: p.IdentityVersion, Kind: "credential", Collection: "access", Digest: digest, State: "prepared"}
	created := s.db.WithContext(ctx).Clauses(clause.OnConflict{DoNothing: true}).Create(&intent)
	if created.Error != nil {
		return nil, created.Error
	}
	if created.RowsAffected == 0 {
		if err = s.db.WithContext(ctx).First(&intent, "operation_id = ?", op).Error; err != nil {
			return nil, err
		}
		if intent.ActorID != p.UserID || intent.TokenID != p.TokenID || intent.ServiceID != p.ServiceID || intent.IdentityVersion != p.IdentityVersion || intent.Collection != "access" || intent.Digest != digest {
			return nil, embedded.ErrInvalid
		}
	}
	if intent.State == "active" {
		return s.NativeGet(ctx, p, "access", intent.ResourceID)
	}
	record, err := s.engine.FindNativeOperation(ctx, "access", op)
	if errors.Is(err, embedded.ErrNotFound) {
		record, err = s.engine.NativeCloneCredential(WithPrincipal(ctx, p), embedded.CredentialUse{ActorID: p.UserID, CredentialID: id, Version: info.Version, Purpose: "credential.clone", ResourceKind: "credential", ResourceID: id}, fmt.Sprint(p.UserID), name, op)
	}
	if err != nil {
		return nil, err
	}
	intent.ResourceID, _ = record["id"].(string)
	if err = s.commitNativeOwnership(ctx, p, intent); err != nil {
		return nil, err
	}
	return s.engine.NativeGet(ctx, "access", intent.ResourceID)
}
