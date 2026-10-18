//go:build certimate_full

package embedded

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"

	"github.com/certimate-go/certimate/internal/certificate"
	"github.com/certimate-go/certimate/internal/domain"
	"github.com/certimate-go/certimate/internal/domain/dtos"
	"github.com/certimate-go/certimate/internal/notify"
	"github.com/certimate-go/certimate/internal/repository"
)

// NativeCertificateDownload is an explicit private-key export, not a metadata
// read. The Dashboard separately authorizes export before entering this method.
func (e *Engine) NativeCertificateDownload(ctx context.Context, actor, id string, opts NativeCertificateDownloadOptions) ([]byte, error) {
	e.mu.RLock()
	defer e.mu.RUnlock()
	e.runtimeMu.Lock()
	defer e.runtimeMu.Unlock()
	if err := e.check(ctx); err != nil {
		return nil, err
	}
	if actor == "" {
		return nil, ErrInvalid
	}
	if e.app.Store().Get("nezha_runtime_ready") != true {
		return nil, errors.New("certificate runtime is unavailable")
	}
	if _, err := e.nativeRecord("certificate", id); err != nil {
		return nil, err
	}
	if err := audit(e.app, actor, "certificate.key_export", "certificate:"+id); err != nil {
		return nil, err
	}
	service := certificate.NewCertificateService(repository.NewACMEAccountRepository(), repository.NewCertificateRepository())
	res, err := service.DownloadCertificate(ctx, &dtos.CertificateDownloadReq{CertificateId: id, FileFormat: domain.CertificateFormatType(opts.FileFormat), PfxPassword: opts.PfxPassword, PfxEncoder: opts.PfxEncoder, JksAlias: opts.JksAlias, JksKeypass: opts.JksKeypass, JksStorepass: opts.JksStorepass})
	if err != nil {
		return nil, errors.New("certificate conversion failed")
	}
	return res.ZipBytes, nil
}

func (e *Engine) NativeRevokeCertificate(ctx context.Context, actor, id string) error {
	e.mu.RLock()
	defer e.mu.RUnlock()
	e.runtimeMu.Lock()
	defer e.runtimeMu.Unlock()
	if err := e.check(ctx); err != nil {
		return err
	}
	if actor == "" {
		return ErrInvalid
	}
	if e.app.Store().Get("nezha_runtime_ready") != true {
		return errors.New("certificate runtime is unavailable")
	}
	if _, err := e.nativeRecord("certificate", id); err != nil {
		return err
	}
	if err := audit(e.app, actor, "certificate.revoke.request", "certificate:"+id); err != nil {
		return err
	}
	service := certificate.NewCertificateService(repository.NewACMEAccountRepository(), repository.NewCertificateRepository())
	if _, err := service.RevokeCertificate(ctx, &dtos.CertificateRevokeReq{CertificateId: id}); err != nil {
		return errors.New("certificate revocation failed")
	}
	return audit(e.app, actor, "certificate.revoke.completed", "certificate:"+id)
}

func (e *Engine) NativeTestNotification(ctx context.Context, use CredentialUse, provider string) error {
	// ResolveCredential performs a fresh host decision. Use purpose and target
	// are fixed here so a read grant cannot be disguised as notification use.
	if use.Purpose != "notification.test" || use.ResourceKind != "credential" || use.ResourceID != use.CredentialID {
		return ErrInvalid
	}
	config, err := e.ResolveCredential(ctx, use)
	if err != nil {
		return err
	}
	e.mu.RLock()
	defer e.mu.RUnlock()
	if err := e.check(ctx); err != nil {
		return err
	}
	r, err := e.nativeRecord("access", use.CredentialID)
	if err != nil {
		return err
	}
	if provider != r.GetString("provider") || r.GetString("reserve") != "notif" || !r.GetBool("unified_enabled") || r.GetInt("unified_version") != use.Version {
		return ErrInvalid
	}
	actor := fmt.Sprintf("user:%d", use.ActorID)
	if err := audit(e.app, actor, "notification.test.request", "credential:"+use.CredentialID); err != nil {
		return err
	}
	_, err = notify.NewClient(notify.WithLogger(slog.New(slog.NewTextHandler(io.Discard, nil)))).SendNotification(ctx, &notify.SendNotificationRequest{Provider: domain.NotificationProviderType(provider), ProviderAccessConfig: config, ProviderExtendedConfig: map[string]any{}, Subject: "[Nezha Unified] Notification Testing", Message: "Notification delivery test from Nezha Unified."})
	if err != nil {
		return errors.New("notification delivery failed")
	}
	return audit(e.app, actor, "notification.test.completed", "credential:"+use.CredentialID)
}
