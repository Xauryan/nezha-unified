package unified

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/certimate-go/certimate/pkg/embedded"
	"github.com/nezhahq/nezha/pkg/agentproto"
	"github.com/nezhahq/nezha/pkg/certdeploy"
	"github.com/nezhahq/nezha/service/agentcall"
	"gorm.io/gorm"
	"regexp"
	"strconv"
	"sync"
	"time"
)

var ErrTargetConflict = errors.New("deployment target version or operation conflict")
var deploymentKey = regexp.MustCompile(`^[A-Za-z0-9_-]{1,128}$`)

type Deployment struct {
	ID                 uint64             `json:"id" gorm:"primaryKey"`
	OperationID        string             `json:"operation_id" gorm:"uniqueIndex;size:128"`
	TargetID           uint64             `json:"target_id" gorm:"index"`
	TargetVersion      uint64             `json:"target_version"`
	CertificateID      string             `json:"certificate_id"`
	CertificateVersion string             `json:"certificate_version"`
	Fingerprint        string             `json:"fingerprint"`
	WorkflowRunID      string             `json:"workflow_run_id,omitempty"`
	WorkflowID         string             `json:"workflow_id,omitempty"`
	NodeID             string             `json:"node_id,omitempty"`
	ActorID            uint64             `json:"actor_id"`
	PrincipalJSON      string             `json:"-"`
	TargetJSON         string             `json:"-"`
	Digest             string             `json:"-"`
	State              string             `json:"state" gorm:"index"`
	Attempts           int                `json:"attempts"`
	LeaseUntil         time.Time          `json:"lease_until"`
	NextAttempt        time.Time          `json:"next_attempt"`
	Receipt            agentproto.Receipt `json:"receipt" gorm:"serializer:json"`
	Code               string             `json:"code,omitempty"`
	CreatedAt          time.Time          `json:"created_at"`
	UpdatedAt          time.Time          `json:"updated_at"`
}

func (Deployment) TableName() string { return "unified_deployments" }

type DeploymentInput struct {
	OperationID        string `json:"operation_id"`
	TargetVersion      uint64 `json:"target_version"`
	CertificateID      string `json:"certificate_id"`
	CertificateVersion string `json:"certificate_version"`
	WorkflowRunID      string `json:"-"`
	WorkflowID         string `json:"-"`
	NodeID             string `json:"-"`
}
type deploymentRuntime struct {
	calls  *agentcall.Broker
	cancel context.CancelFunc
	wg     sync.WaitGroup
	mu     sync.Mutex
	locks  [64]sync.Mutex
	wake   chan struct{}
}

func (s *Service) SetAgentCalls(calls *agentcall.Broker) {
	if s.deployments == nil {
		s.deployments = &deploymentRuntime{wake: make(chan struct{}, 1)}
	}
	r := s.deployments
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.cancel != nil {
		return
	}
	r.calls = calls
	if calls == nil {
		return
	}
	ctx, cancel := context.WithCancel(context.Background())
	r.cancel = cancel
	r.wg.Add(1)
	go func() {
		defer r.wg.Done()
		timer := time.NewTicker(5 * time.Second)
		defer timer.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-timer.C:
			case <-r.wake:
			}
			s.resumeDeployments(ctx)
		}
	}()
}
func (s *Service) CloseDeployments() {
	if s.deployments == nil {
		return
	}
	r := s.deployments
	r.mu.Lock()
	if r.cancel != nil {
		r.cancel()
	}
	r.mu.Unlock()
	r.wg.Wait()
}
func (s *Service) CloseDeploymentsContext(ctx context.Context) error {
	if s.deployments == nil {
		return nil
	}
	r := s.deployments
	r.mu.Lock()
	if r.cancel != nil {
		r.cancel()
	}
	r.mu.Unlock()
	done := make(chan struct{})
	go func() { r.wg.Wait(); close(done) }()
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
func (s *Service) deploymentAuthorization(ctx context.Context, p Principal, t DeploymentTarget, cert string, run ...string) error {
	current, err := s.CurrentPrincipal(ctx, p)
	if err != nil {
		return err
	}
	if !s.Can(ctx, current, "deployment_target", fmt.Sprint(t.ID), "execute") {
		return ErrDenied
	}
	if len(run) == 2 && run[0] != "" && run[1] != "" {
		_, wp, _, err := s.executionIdentity(ctx, run[0], run[1])
		if err != nil {
			return err
		}
		if wp.Targets[fmt.Sprint(t.ID)] != t.Version {
			return ErrDenied
		}
		if err = s.authorizeRuntime(ctx, embedded.RuntimeOperation{WorkflowID: run[0], RunID: run[1], GraphDigest: wp.GraphDigest, Action: "certificate.deploy", ResourceID: cert}); err != nil {
			return err
		}
	} else if !s.Can(ctx, current, "certificate", cert, "use") {
		return ErrDenied
	}
	var fresh DeploymentTarget
	if err = s.db.WithContext(ctx).First(&fresh, t.ID).Error; err != nil {
		return err
	}
	if !fresh.Enabled || fresh.Version != t.Version || fresh.ServerID != t.ServerID {
		return ErrTargetConflict
	}
	return nil
}
func (s *Service) createDeployment(ctx context.Context, p Principal, target uint64, in DeploymentInput, m embedded.CertificateMaterial) (Deployment, error) {
	var t DeploymentTarget
	if err := s.db.WithContext(ctx).First(&t, target).Error; err != nil {
		return Deployment{}, err
	}
	if err := s.deploymentAuthorization(ctx, p, t, in.CertificateID, in.WorkflowID, in.WorkflowRunID); err != nil {
		return Deployment{}, err
	}
	if !deploymentKey.MatchString(in.OperationID) || in.TargetVersion != t.Version || in.CertificateVersion != m.Version || in.CertificateID == "" {
		return Deployment{}, embedded.ErrInvalid
	}
	check := agentproto.Deployment{OperationID: in.OperationID, Sequence: 1, Target: t.agentTarget(), CertificatePEM: m.CertificatePEM, PrivateKeyPEM: m.PrivateKeyPEM, Fingerprint: m.Fingerprint}
	if certdeploy.Validate(check) != nil {
		return Deployment{}, embedded.ErrInvalid
	}
	actor, _ := json.Marshal(p)
	snapshot, _ := json.Marshal(t)
	body, _ := json.Marshal(struct {
		Input    DeploymentInput
		Target   uint64
		Actor    string
		Workflow string
		Run      string
		Node     string
	}{in, target, string(actor), in.WorkflowID, in.WorkflowRunID, in.NodeID})
	hash := sha256.Sum256(body)
	d := Deployment{OperationID: in.OperationID, TargetID: target, TargetVersion: t.Version, CertificateID: in.CertificateID, CertificateVersion: m.Version, Fingerprint: m.Fingerprint, WorkflowRunID: in.WorkflowRunID, NodeID: in.NodeID, ActorID: p.UserID, PrincipalJSON: string(actor), TargetJSON: string(snapshot), Digest: hex.EncodeToString(hash[:]), State: "queued"}
	d.WorkflowID = in.WorkflowID
	err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var prior Deployment
		if err := tx.Where("operation_id = ?", in.OperationID).First(&prior).Error; err == nil {
			if prior.Digest != d.Digest {
				return ErrTargetConflict
			}
			d = prior
			return nil
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		if err := tx.Create(&d).Error; err != nil {
			return err
		}
		res := tx.Model(&DeploymentTarget{}).Where("id = ? AND version = ? AND enabled = ?", target, t.Version, true).Update("desired_version", d.ID)
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected != 1 {
			return ErrTargetConflict
		}
		return tx.Create(&Audit{ActorID: p.UserID, Action: "deployment.enqueue", Resource: fmt.Sprint(d.ID)}).Error
	})
	return d, err
}
func (s *Service) EnqueueDeployment(ctx context.Context, p Principal, target uint64, in DeploymentInput) (Deployment, error) {
	var t DeploymentTarget
	if err := s.db.WithContext(ctx).First(&t, target).Error; err != nil {
		return Deployment{}, err
	}
	if err := s.deploymentAuthorization(ctx, p, t, in.CertificateID); err != nil {
		return Deployment{}, err
	}
	m, err := s.engine.CertificateMaterial(ctx, in.CertificateID)
	if err != nil {
		return Deployment{}, err
	}
	if in.CertificateVersion == "" {
		in.CertificateVersion = m.Version
	}
	d, err := s.createDeployment(ctx, p, target, in, m)
	if err == nil && s.deployments != nil {
		select {
		case s.deployments.wake <- struct{}{}:
		default:
		}
	}
	return d, err
}
func (s *Service) DeployTarget(ctx context.Context, p Principal, in embedded.TargetDeployment) error {
	target, err := strconv.ParseUint(in.TargetID, 10, 64)
	if err != nil {
		return embedded.ErrInvalid
	}
	m := embedded.CertificateMaterial{CertificatePEM: in.CertificatePEM, PrivateKeyPEM: in.PrivateKeyPEM, Domains: in.Domains, Fingerprint: in.Fingerprint, Version: in.CertificateVersion}
	d, err := s.createDeployment(ctx, p, target, DeploymentInput{OperationID: in.OperationID, TargetVersion: in.TargetVersion, CertificateID: in.CertificateID, CertificateVersion: in.CertificateVersion, WorkflowID: in.WorkflowID, WorkflowRunID: in.RunID, NodeID: in.NodeID}, m)
	if err != nil {
		return err
	}
	_, err = s.processDeployment(ctx, d.ID, &m, false)
	return err
}
func (s *Service) processDeployment(ctx context.Context, id uint64, m *embedded.CertificateMaterial, reconcile bool) (Deployment, error) {
	if s.deployments == nil || s.deployments.calls == nil {
		return Deployment{}, agentcall.ErrOffline
	}
	r := s.deployments
	r.locks[id%uint64(len(r.locks))].Lock()
	defer r.locks[id%uint64(len(r.locks))].Unlock()
	var d Deployment
	if err := s.db.WithContext(ctx).First(&d, id).Error; err != nil {
		return d, err
	}
	if d.State == "verified" {
		return d, nil
	}
	if d.State == "cancelled" || d.State == "superseded" || d.State == "rolled_back" || d.State == "failed" || d.State == "reconciliation_required" {
		return d, ErrTargetConflict
	}
	var p Principal
	var t DeploymentTarget
	if json.Unmarshal([]byte(d.PrincipalJSON), &p) != nil || json.Unmarshal([]byte(d.TargetJSON), &t) != nil {
		return d, embedded.ErrInvalid
	}
	targetLock := &s.recordLocks[d.TargetID%uint64(len(s.recordLocks))]
	targetLock.Lock()
	defer targetLock.Unlock()
	auth := func(ctx context.Context) error {
		var current Deployment
		if err := s.db.WithContext(ctx).First(&current, d.ID).Error; err != nil {
			return err
		}
		if current.Code == "cancel_requested" || current.State == "cancelled" || current.State == "reconciliation_required" {
			return ErrDenied
		}
		return s.deploymentAuthorization(ctx, p, t, d.CertificateID, d.WorkflowID, d.WorkflowRunID)
	}
	finish := func(state, code string, receipt agentproto.Receipt) {
		d.State = state
		d.Code = code
		d.Receipt = receipt
		d.LeaseUntil = time.Time{}
		_ = s.db.WithContext(context.WithoutCancel(ctx)).Model(&Deployment{}).Where("id = ?", d.ID).Updates(map[string]any{"state": state, "code": code, "receipt": mustDeploymentJSON(receipt), "lease_until": time.Time{}, "next_attempt": time.Now().Add(30 * time.Second)}).Error
	}
	if err := auth(ctx); err != nil {
		state := "cancelled"
		switch d.State {
		case "dispatching", "staging", "validating", "applying", "unknown", "reconciling":
			state = "reconciliation_required"
		}
		finish(state, "authorization", d.Receipt)
		return d, err
	}
	var latest DeploymentTarget
	if err := s.db.WithContext(ctx).First(&latest, t.ID).Error; err != nil {
		return d, err
	}
	reconcile = reconcile || d.State == "unknown" || d.State == "reconciling" || (!d.LeaseUntil.IsZero() && time.Now().After(d.LeaseUntil))
	if !reconcile && latest.DesiredVersion != d.ID {
		finish("superseded", "newer_version", d.Receipt)
		return d, ErrTargetConflict
	}
	if !d.LeaseUntil.IsZero() && time.Now().Before(d.LeaseUntil) {
		return d, ErrTargetConflict
	}
	now := time.Now()
	claim := s.db.WithContext(ctx).Model(&Deployment{}).Where("id = ? AND (lease_until IS NULL OR lease_until = ? OR lease_until < ?)", id, time.Time{}, now).Updates(map[string]any{"lease_until": now.Add(3 * time.Minute), "attempts": gorm.Expr("attempts + 1")})
	if claim.Error != nil || claim.RowsAffected != 1 {
		return d, ErrTargetConflict
	}
	cap, err := r.calls.Capabilities(ctx, t.ServerID, auth)
	if err != nil {
		state := "waiting_online"
		if reconcile {
			state = "unknown"
		}
		finish(state, "offline", d.Receipt)
		return d, err
	}
	if !cap.CertificateDeploy {
		finish("failed", "capability", d.Receipt)
		return d, ErrTargetConflict
	}
	request := agentproto.Deployment{OperationID: d.OperationID, Sequence: d.ID, Target: t.agentTarget(), Fingerprint: d.Fingerprint}
	method := "reconcile"
	state := "reconciling"
	if !reconcile {
		if m == nil {
			material, e := s.engine.CertificateMaterial(ctx, d.CertificateID)
			if e != nil {
				finish("failed", "certificate", d.Receipt)
				return d, e
			}
			m = &material
		}
		if m.Version != d.CertificateVersion || m.Fingerprint != d.Fingerprint {
			finish("failed", "certificate_version", d.Receipt)
			return d, ErrTargetConflict
		}
		request.CertificatePEM = m.CertificatePEM
		request.PrivateKeyPEM = m.PrivateKeyPEM
		method = "deploy"
		state = "dispatching"
	}
	if err = s.db.WithContext(ctx).Model(&Deployment{}).Where("id = ?", id).Update("state", state).Error; err != nil {
		return d, err
	}
	payload, _ := json.Marshal(request)
	raw, callErr := r.calls.Call(ctx, t.ServerID, agentproto.TaskCall, "certificate", agentproto.Request{Protocol: agentproto.Protocol, OperationID: d.OperationID, Method: method, Payload: payload}, auth, func(out agentproto.Response) {
		if out.Final {
			return
		}
		if out.Phase == "ready" {
			permitAuth := func(ctx context.Context) error {
				if err := auth(ctx); err != nil {
					return err
				}
				var target DeploymentTarget
				if err := s.db.WithContext(ctx).First(&target, t.ID).Error; err != nil {
					return err
				}
				if target.DesiredVersion != d.ID {
					return ErrTargetConflict
				}
				return nil
			}
			if _, err := r.calls.Call(ctx, t.ServerID, agentproto.TaskCall, "certificate_permit", agentproto.Request{Protocol: agentproto.Protocol, OperationID: d.OperationID, Method: "permit"}, permitAuth, nil); err != nil {
				s.cancelIssuedDeployment(t.ServerID, d.OperationID)
			}
			return
		}
		switch out.Phase {
		case "staging", "validating", "applying":
			_ = s.db.Model(&Deployment{}).Where("id = ?", id).Update("state", out.Phase).Error
		}
	})
	if errors.Is(callErr, agentcall.ErrDenied) || errors.Is(callErr, context.Canceled) {
		s.cancelIssuedDeployment(t.ServerID, d.OperationID)
	}
	var receipt agentproto.Receipt
	if len(raw) > 0 && json.Unmarshal(raw, &receipt) == nil && receipt.OperationID == d.OperationID && receipt.TargetID == fmt.Sprint(t.ID) && receipt.TargetVersion == t.Version && receipt.Sequence == d.ID && receipt.Fingerprint == d.Fingerprint {
		if callErr == nil && receipt.Verified && receipt.Phase == "verified" {
			err := s.db.WithContext(context.WithoutCancel(ctx)).Transaction(func(tx *gorm.DB) error {
				if err := tx.Model(&Deployment{}).Where("id = ?", d.ID).Updates(map[string]any{"state": "verified", "code": "", "receipt": mustDeploymentJSON(receipt), "lease_until": time.Time{}}).Error; err != nil {
					return err
				}
				if err := tx.Model(&DeploymentTarget{}).Where("id = ? AND version = ? AND applied_version < ?", t.ID, t.Version, d.ID).Updates(map[string]any{"applied_version": d.ID, "applied_fingerprint": d.Fingerprint}).Error; err != nil {
					return err
				}
				return tx.Create(&Audit{ActorID: d.ActorID, Action: "deployment.verified", Resource: fmt.Sprint(d.ID)}).Error
			})
			if err != nil {
				return d, err
			}
			d.State = "verified"
			d.Code = ""
			d.Receipt = receipt
			d.LeaseUntil = time.Time{}
			return d, nil
		}
		switch receipt.Phase {
		case "failed", "rolled_back":
			finish(receipt.Phase, receipt.Code, receipt)
			return d, callErr
		}
	}
	finish("unknown", "result_unknown", receipt)
	if callErr == nil {
		callErr = agentcall.ErrProtocol
	}
	return d, callErr
}

// Cancellation is a bounded compensation for this exact already-issued
// operation. It cannot change paths, commands, certificates or target scope.
func (s *Service) cancelIssuedDeployment(server uint64, operation string) {
	if s.deployments == nil || s.deployments.calls == nil {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, _ = s.deployments.calls.Call(ctx, server, agentproto.TaskCall, "certificate_cancel", agentproto.Request{Protocol: agentproto.Protocol, OperationID: operation, Method: "cancel"}, func(context.Context) error { return nil }, nil)
}
func mustDeploymentJSON(v any) string { b, _ := json.Marshal(v); return string(b) }
func (s *Service) resumeDeployments(ctx context.Context) {
	var jobs []Deployment
	now := time.Now()
	_ = s.db.WithContext(ctx).Model(&Deployment{}).Where("state IN ? AND lease_until < ?", []string{"dispatching", "staging", "validating", "applying", "reconciling"}, now).Update("state", "unknown").Error
	if s.db.WithContext(ctx).Where("state IN ? AND next_attempt <= ?", []string{"queued", "waiting_online", "unknown"}, now).Order("id ASC").Limit(4).Find(&jobs).Error != nil {
		return
	}
	for _, job := range jobs {
		if ctx.Err() != nil {
			return
		}
		_, _ = s.processDeployment(ctx, job.ID, nil, job.State == "unknown")
	}
}
func (s *Service) Deployment(ctx context.Context, p Principal, id uint64) (Deployment, error) {
	var d Deployment
	if err := s.db.WithContext(ctx).First(&d, id).Error; err != nil {
		return d, err
	}
	if !s.Can(ctx, p, "deployment_target", fmt.Sprint(d.TargetID), "read") {
		return Deployment{}, ErrDenied
	}
	return d, nil
}
func (s *Service) Deployments(ctx context.Context, p Principal) ([]Deployment, error) {
	scope, err := s.resourceScope(ctx, p, "deployment_target")
	if err != nil {
		return nil, err
	}
	q := s.db.WithContext(ctx).Order("id DESC")
	if !scope.All {
		q = q.Where("target_id IN ?", scope.IDs)
	}
	var rows []Deployment
	if err = q.Find(&rows).Error; err != nil {
		return nil, err
	}
	out := []Deployment{}
	for _, d := range rows {
		if s.Can(ctx, p, "deployment_target", fmt.Sprint(d.TargetID), "read") {
			out = append(out, d)
			if len(out) == 100 {
				break
			}
		}
	}
	return out, nil
}
func (s *Service) CancelDeployment(ctx context.Context, p Principal, id uint64) (Deployment, error) {
	d, err := s.Deployment(ctx, p, id)
	if err != nil {
		return d, err
	}
	if !s.Can(ctx, p, "deployment_target", fmt.Sprint(d.TargetID), "execute") {
		return d, ErrDenied
	}
	state := "cancelled"
	switch d.State {
	case "verified", "failed", "rolled_back", "superseded", "cancelled":
		return d, nil
	case "dispatching", "staging", "validating", "applying", "unknown", "reconciling":
		state = "unknown"
	}
	if err = s.db.WithContext(ctx).Model(&Deployment{}).Where("id = ?", id).Updates(map[string]any{"state": state, "code": "cancel_requested"}).Error; err != nil {
		return d, err
	}
	d.State = state
	d.Code = "cancel_requested"
	if state == "unknown" && s.deployments != nil && s.deployments.calls != nil {
		var t DeploymentTarget
		_ = json.Unmarshal([]byte(d.TargetJSON), &t)
		_, _ = s.deployments.calls.Call(ctx, t.ServerID, agentproto.TaskCall, "certificate_cancel", agentproto.Request{Protocol: agentproto.Protocol, OperationID: d.OperationID, Method: "cancel"}, func(ctx context.Context) error {
			if !s.Can(ctx, p, "deployment_target", fmt.Sprint(t.ID), "execute") {
				return ErrDenied
			}
			return nil
		}, nil)
	}
	return d, nil
}
func (s *Service) ReconcileDeployment(ctx context.Context, p Principal, id uint64) (Deployment, error) {
	d, err := s.Deployment(ctx, p, id)
	if err != nil {
		return d, err
	}
	if !s.Can(ctx, p, "deployment_target", fmt.Sprint(d.TargetID), "execute") {
		return d, ErrDenied
	}
	return s.processDeployment(ctx, id, nil, true)
}
