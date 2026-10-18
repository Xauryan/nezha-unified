package unified

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/certimate-go/certimate/pkg/embedded"
	"github.com/nezhahq/nezha/pkg/agentproto"
	"github.com/nezhahq/nezha/pkg/certdeploy"
	"gorm.io/gorm"
)

func (s *Service) authorizeTargetCredential(ctx context.Context, use embedded.CredentialUse) error {
	p, ok := principalFromContext(ctx)
	if !ok || p.UserID != use.ActorID || use.Purpose != "deploy.target" || use.OperationID == "" {
		return ErrDenied
	}
	var d Deployment
	if err := s.db.WithContext(ctx).Where("operation_id = ? AND target_id = ?", use.OperationID, use.ResourceID).First(&d).Error; err != nil {
		return ErrDenied
	}
	if d.ActorID != use.ActorID || d.WorkflowID != use.WorkflowID || d.WorkflowRunID != use.RunID || d.State != "dispatching" || d.Code == "cancel_requested" {
		return ErrDenied
	}
	var target DeploymentTarget
	if err := s.db.WithContext(ctx).First(&target, d.TargetID).Error; err != nil {
		return ErrDenied
	}
	if target.Kind != "certimate" || target.Version != d.TargetVersion || target.CredentialID != use.CredentialID || target.CredentialVersion != use.Version || target.DesiredVersion != d.ID {
		return ErrDenied
	}
	if err := checkDeploymentFence(s.db.WithContext(ctx), target.ID, d.ID); err != nil {
		return err
	}
	return s.deploymentAuthorization(ctx, p, target, d.CertificateID, d.WorkflowID, d.WorkflowRunID)
}

func (s *Service) processCloudDeployment(ctx context.Context, d Deployment, p Principal, target, latest DeploymentTarget, material *embedded.CertificateMaterial, reconcile bool, auth func(context.Context) error, finish func(string, string, agentproto.Receipt)) (Deployment, error) {
	// A provider error or interrupted request can follow a remote side effect.
	// Keep it fenced for operator review instead of replaying a cloud mutation.
	if reconcile {
		finish("reconciliation_required", "provider_result_unknown", d.Receipt)
		d.State = "reconciliation_required"
		d.Code = "provider_result_unknown"
		return d, ErrTargetConflict
	}
	if material == nil {
		m, err := s.engine.CertificateMaterial(ctx, d.CertificateID)
		if err != nil {
			finish("failed", "certificate", d.Receipt)
			return d, err
		}
		material = &m
	}
	if material.Version != d.CertificateVersion || material.Fingerprint != d.Fingerprint || certdeploy.ValidateCertificate(material.CertificatePEM, material.PrivateKeyPEM, material.Fingerprint, target.AllowedNames) != nil {
		finish("failed", "certificate_version", d.Receipt)
		return d, ErrTargetConflict
	}
	config, err := s.engine.OpenTargetConfiguration(ctx, fmt.Sprint(latest.ID), latest.Version, latest.ConfigurationEnvelope)
	if err != nil {
		finish("failed", "configuration", d.Receipt)
		return d, err
	}
	if err = auth(ctx); err != nil {
		finish("cancelled", "authorization", d.Receipt)
		return d, err
	}
	if err = s.db.WithContext(ctx).Model(&Deployment{}).Where("id = ?", d.ID).Updates(map[string]any{"state": "dispatching", "lease_until": time.Now().Add(11 * time.Minute)}).Error; err != nil {
		return d, err
	}
	use := embedded.CredentialUse{ActorID: p.UserID, CredentialID: target.CredentialID, Version: target.CredentialVersion, Purpose: "deploy.target", ResourceKind: "deployment_target", ResourceID: fmt.Sprint(target.ID), OperationID: d.OperationID, WorkflowID: d.WorkflowID, RunID: d.WorkflowRunID}
	bounded, cancel := context.WithTimeout(WithPrincipal(ctx, p), 10*time.Minute)
	done := make(chan struct{})
	watched := make(chan struct{})
	go func() {
		defer close(watched)
		timer := time.NewTicker(250 * time.Millisecond)
		defer timer.Stop()
		for {
			select {
			case <-done:
				return
			case <-bounded.Done():
				return
			case <-timer.C:
				if auth(bounded) != nil {
					cancel()
					return
				}
			}
		}
	}()
	err = s.engine.NativeDeployProvider(bounded, use, target.Provider, config, *material)
	close(done)
	cancel()
	<-watched
	if err != nil {
		finish("reconciliation_required", "provider_result_unknown", d.Receipt)
		d.State = "reconciliation_required"
		d.Code = "provider_result_unknown"
		return d, errors.New("cloud deployment outcome requires reconciliation")
	}
	// The provider acknowledged the mutation. No served TLS fingerprint is implied.
	receipt := agentproto.Receipt{OperationID: d.OperationID, TargetID: fmt.Sprint(target.ID), TargetVersion: target.Version, Sequence: d.ID, Fingerprint: d.Fingerprint, Phase: "applied", Code: "provider_accepted"}
	err = s.db.WithContext(context.WithoutCancel(ctx)).Transaction(func(tx *gorm.DB) error {
		if e := tx.Model(&Deployment{}).Where("id = ?", d.ID).Updates(map[string]any{"state": "applied", "code": "provider_accepted", "receipt": mustDeploymentJSON(receipt), "lease_until": time.Time{}}).Error; e != nil {
			return e
		}
		if e := tx.Model(&DeploymentTarget{}).Where("id = ? AND version = ? AND applied_version < ?", target.ID, target.Version, d.ID).Updates(map[string]any{"applied_version": d.ID, "applied_fingerprint": d.Fingerprint}).Error; e != nil {
			return e
		}
		return tx.Create(&Audit{ActorID: p.UserID, Action: "deployment.provider_accepted", Resource: fmt.Sprint(d.ID)}).Error
	})
	if err != nil {
		return d, err
	}
	d.State = "applied"
	d.Code = "provider_accepted"
	d.Receipt = receipt
	d.LeaseUntil = time.Time{}
	return d, nil
}
