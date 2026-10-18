package unified

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/certimate-go/certimate/pkg/embedded"
	"github.com/nezhahq/nezha/service/unified/dnsprovider"
	"github.com/nezhahq/nezha/service/unified/policy"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (s *Service) StartRuntime(ctx context.Context) error {
	if err := s.RecoverIntents(ctx); err != nil {
		return err
	}
	return s.engine.StartRuntime(ctx, s.RuntimeHooks())
}
func (s *Service) StopRuntime(ctx context.Context) error { return s.engine.StopRuntime(ctx) }
func (s *Service) RuntimeHooks() embedded.RuntimeHooks {
	return embedded.RuntimeHooks{
		Authorize: s.authorizeRuntime, OnRunCreated: s.bindRun, ResolveCredential: s.resolveWorkflowCredential,
		DeployTarget: func(ctx context.Context, in embedded.TargetDeployment) error {
			p, wp, _, err := s.executionIdentity(ctx, in.WorkflowID, in.RunID)
			if err != nil {
				return err
			}
			if wp.Targets[in.TargetID] != in.TargetVersion {
				return ErrDenied
			}
			return s.DeployTarget(ctx, p, in)
		},
	}
}

// updateWorkflowApproval only approves a changed graph for an unrestricted
// administrator. A delegated edit invalidates the prior executable version.
func (s *Service) updateWorkflowApproval(ctx context.Context, p Principal, id string) error {
	var binding ResourceBinding
	err := s.db.WithContext(ctx).First(&binding, "kind = ? AND resource_id = ?", "workflow", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		if requireAdmin(p) != nil {
			return ErrDenied
		}
		binding = ResourceBinding{Kind: "workflow", ResourceID: id, OwnerID: p.UserID, Version: 1, Active: true}
		if err = s.db.WithContext(ctx).Create(&binding).Error; err != nil {
			return err
		}
	} else if err != nil {
		return err
	}
	var wp WorkflowPolicy
	err = s.db.WithContext(ctx).First(&wp, "workflow_id = ?", id).Error
	fresh := errors.Is(err, gorm.ErrRecordNotFound)
	if err != nil && !fresh {
		return err
	}
	digest, digestErr := s.engine.WorkflowGraphDigest(ctx, id)
	if digestErr != nil && !errors.Is(digestErr, embedded.ErrInvalid) {
		return digestErr
	}
	if !fresh && digest == wp.GraphDigest {
		return nil
	}
	if fresh {
		wp = WorkflowPolicy{WorkflowID: id, RunAsUserID: binding.OwnerID, AllowInline: requireAdmin(p) == nil, Credentials: map[string]int{}, Targets: map[string]uint64{}}
	}
	wp.Version++
	if requireAdmin(p) != nil {
		wp.Enabled = false
		return s.db.WithContext(ctx).Save(&wp).Error
	}
	wp.GraphDigest = digest
	wp.Enabled = digest != ""
	wp.Credentials = map[string]int{}
	wp.Targets = map[string]uint64{}
	record, err := s.engine.NativeGet(ctx, "workflow", id)
	if errors.Is(err, embedded.ErrNotFound) && binding.OperationID != "" {
		record, err = s.engine.FindNativeOperation(ctx, "workflow", binding.OperationID)
	}
	if err != nil {
		return err
	}
	var visit func(any) error
	visit = func(value any) error {
		switch node := value.(type) {
		case map[string]any:
			provider, _ := node["provider"].(string)
			for key, item := range node {
				if key == "providerAccessId" || key == "caProviderAccessId" {
					ref, _ := item.(string)
					if ref == "" {
						continue
					}
					if provider == "nezha-agent" && key == "providerAccessId" {
						targetID, err := strconv.ParseUint(ref, 10, 64)
						if err != nil {
							return embedded.ErrInvalid
						}
						var target DeploymentTarget
						if err = s.db.WithContext(ctx).First(&target, targetID).Error; err != nil {
							return embedded.ErrInvalid
						}
						wp.Targets[ref] = target.Version
					} else {
						info, err := s.engine.CredentialInfo(ctx, ref)
						if err != nil || !info.Enabled {
							return embedded.ErrInvalid
						}
						wp.Credentials[ref] = info.Version
					}
				}
				if err := visit(item); err != nil {
					return err
				}
			}
		case []any:
			for _, item := range node {
				if err := visit(item); err != nil {
					return err
				}
			}
		}
		return nil
	}
	if err = visit(record["graphContent"]); err != nil {
		return err
	}
	return s.db.WithContext(ctx).Save(&wp).Error
}
func (s *Service) executionIdentity(ctx context.Context, workflowID, runID string) (Principal, WorkflowPolicy, ResourceBinding, error) {
	var wp WorkflowPolicy
	var owner ResourceBinding
	var p Principal
	if workflowID == "" {
		return p, wp, owner, ErrDenied
	}
	if s.db.WithContext(ctx).First(&wp, "workflow_id = ? AND enabled = ?", workflowID, true).Error != nil {
		return p, wp, owner, ErrDenied
	}
	if s.db.WithContext(ctx).First(&owner, "kind = ? AND resource_id = ? AND active = ?", "workflow", workflowID, true).Error != nil {
		return p, wp, owner, ErrDenied
	}
	if runID != "" {
		var run RunBinding
		if s.db.WithContext(ctx).First(&run, "run_id = ? AND workflow_id = ?", runID, workflowID).Error != nil || run.Revoked || run.PolicyVersion != wp.Version || run.OwnerVersion != owner.Version || run.GraphDigest != wp.GraphDigest {
			return p, wp, owner, ErrDenied
		}
		p = Principal{UserID: run.ActorID, TokenID: run.TokenID, ServiceID: run.ServiceID, IdentityVersion: run.IdentityVersion}
	} else if request, ok := principalFromContext(ctx); ok {
		p = request
	} else {
		p = Principal{UserID: wp.RunAsUserID, ServiceID: wp.ServiceID, IdentityVersion: wp.IdentityVersion}
	}
	var err error
	p, err = s.CurrentPrincipal(ctx, p)
	if err != nil || !s.Can(ctx, p, "workflow", workflowID, "execute") {
		return p, wp, owner, ErrDenied
	}
	return p, wp, owner, nil
}
func (s *Service) bindRun(ctx context.Context, workflowID, runID string) error {
	p, wp, owner, err := s.executionIdentity(ctx, workflowID, "")
	if err != nil {
		return err
	}
	run := RunBinding{RunID: runID, WorkflowID: workflowID, PolicyVersion: wp.Version, OwnerVersion: owner.Version, ActorID: p.UserID, TokenID: p.TokenID, ServiceID: p.ServiceID, IdentityVersion: p.IdentityVersion, GraphDigest: wp.GraphDigest}
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&run).Error; err != nil {
			return err
		}
		return tx.Create(&Audit{ActorID: p.UserID, Action: "workflow.run.bind", Resource: workflowID + ":" + runID}).Error
	})
}
func (s *Service) AuthorizeRun(ctx context.Context, workflowID, runID string) error {
	_, _, _, err := s.executionIdentity(ctx, workflowID, runID)
	return err
}
func (s *Service) authorizeRuntime(ctx context.Context, op embedded.RuntimeOperation) error {
	p, wp, _, err := s.executionIdentity(ctx, op.WorkflowID, op.RunID)
	if err != nil {
		return err
	}
	if op.GraphDigest == "" || op.GraphDigest != wp.GraphDigest {
		return ErrDenied
	}
	// Provider clients retain their decrypted configuration during ACME waits.
	// Revalidate every bound version at the effect boundary, including CA and
	// HTTP-01 credentials that are not named by the current operation.
	for credential, version := range wp.Credentials {
		info, err := s.engine.CredentialInfo(ctx, credential)
		if err != nil || !info.Enabled || info.Version != version || !s.Can(ctx, p, "credential", credential, "use") {
			return ErrDenied
		}
	}
	if len(op.Domains) > 0 && op.Action != "dns.challenge" && requireAdmin(p) != nil && !policy.AllowsAllCertificateNames(wp.Names, op.Domains) {
		return ErrDenied
	}
	switch op.Action {
	case "workflow.run", "node.execute":
		return nil
	case "credential.use", "notification.send":
		if wp.Credentials[op.ResourceID] < 1 || !s.Can(ctx, p, "credential", op.ResourceID, "use") {
			return ErrDenied
		}
		return nil
	case "certificate.issue", "certificate.monitor":
		return nil
	case "certificate.import":
		if requireAdmin(p) != nil && !wp.AllowImport {
			return ErrDenied
		}
		return nil
	case "certificate.use", "certificate.deploy":
		if op.ResourceID == "" {
			return ErrDenied
		}
		if s.Can(ctx, p, "certificate", op.ResourceID, "use") {
			return nil
		}
		record, err := s.engine.NativeGet(ctx, "certificate", op.ResourceID)
		if err != nil {
			return err
		}
		if record["workflowRef"] == op.WorkflowID {
			return nil
		}
		return ErrDenied
	case "deployment.execute":
		if op.ResourceKind == "deployment_target" {
			if wp.Targets[op.ResourceID] < 1 || !s.Can(ctx, p, "deployment_target", op.ResourceID, "execute") {
				return ErrDenied
			}
			return nil
		}
		if !wp.AllowInline {
			return ErrDenied
		}
		if op.ResourceID == "" {
			if requireAdmin(p) != nil {
				return ErrDenied
			}
			return nil
		}
		if wp.Credentials[op.ResourceID] < 1 || !s.Can(ctx, p, "credential", op.ResourceID, "use") {
			return ErrDenied
		}
		return nil
	case "dns.challenge":
		return s.authorizeChallenge(ctx, p, wp, op.ResourceID, op.Domains)
	default:
		return ErrDenied
	}
}
func (s *Service) authorizeChallenge(ctx context.Context, p Principal, wp WorkflowPolicy, credential string, names []string) error {
	if len(names) == 0 || wp.Credentials[credential] < 1 || !s.Can(ctx, p, "credential", credential, "use") {
		return ErrDenied
	}
	info, err := s.engine.CredentialInfo(ctx, credential)
	if err != nil || !info.Enabled || info.Version != wp.Credentials[credential] {
		return ErrDenied
	}
	if p.Admin && p.TokenID == 0 {
		return nil
	}
	var accounts []DNSAccount
	if err := s.db.WithContext(ctx).Where("credential_id = ? AND credential_version = ? AND enabled = ?", credential, wp.Credentials[credential], true).Find(&accounts).Error; err != nil {
		return err
	}
	for _, name := range names {
		normalized, err := policy.NormalizeName(name)
		if err != nil {
			return ErrDenied
		}
		allowed := false
		for _, account := range accounts {
			var zones []DNSZone
			if err = s.db.WithContext(ctx).Where("account_id = ? AND active = ?", account.ID, true).Find(&zones).Error; err != nil {
				return err
			}
			for _, zone := range zones {
				if normalized != zone.Name && !strings.HasSuffix(normalized, "."+zone.Name) {
					continue
				}
				if !s.RestrictPrincipal(ctx, p, "dns_zone", fmt.Sprint(zone.ID), "create") || !s.RestrictPrincipal(ctx, p, "dns_zone", fmt.Sprint(zone.ID), "delete") {
					continue
				}
				record := dnsprovider.Record{Name: normalized, Type: "TXT"}
				if s.canDNS(ctx, p, zone, record, "create") && s.canDNS(ctx, p, zone, record, "delete") {
					allowed = true
				}
			}
		}
		if !allowed {
			return ErrDenied
		}
	}
	return nil
}
func (s *Service) resolveWorkflowCredential(ctx context.Context, id, purpose string) (map[string]any, error) {
	op, ok := embedded.RuntimeOperationFromContext(ctx)
	if !ok {
		return nil, ErrDenied
	}
	p, wp, _, err := s.executionIdentity(ctx, op.WorkflowID, op.RunID)
	if err != nil {
		return nil, err
	}
	if wp.Credentials[id] < 1 {
		return nil, ErrDenied
	}
	return s.engine.ResolveCredential(WithPrincipal(ctx, p), embedded.CredentialUse{ActorID: p.UserID, CredentialID: id, Version: wp.Credentials[id], Purpose: purpose, ResourceKind: "workflow", ResourceID: op.WorkflowID, WorkflowID: op.WorkflowID, RunID: op.RunID})
}
func (s *Service) authorizeWorkflowCredential(ctx context.Context, use embedded.CredentialUse) error {
	p, wp, _, err := s.executionIdentity(ctx, use.WorkflowID, use.RunID)
	if err != nil {
		return err
	}
	if use.ActorID != p.UserID || use.ResourceID != use.WorkflowID || wp.Credentials[use.CredentialID] != use.Version || !s.Can(ctx, p, "credential", use.CredentialID, "use") {
		return ErrDenied
	}
	return nil
}
func (s *Service) StartWorkflow(ctx context.Context, p Principal, id string) (string, error) {
	if !s.Can(ctx, p, "workflow", id, "execute") {
		return "", ErrDenied
	}
	return s.engine.StartWorkflow(WithPrincipal(ctx, p), id)
}
func (s *Service) CancelWorkflow(ctx context.Context, p Principal, id, runID string) error {
	if !s.Can(ctx, p, "workflow", id, "cancel") {
		return ErrDenied
	}
	// Persist cancellation first; even a currently blocked external call cannot
	// issue a new effect if its dispatcher cancellation races with this request.
	if err := s.db.WithContext(ctx).Model(&RunBinding{}).Where("run_id = ? AND workflow_id = ?", runID, id).Update("revoked", true).Error; err != nil {
		return err
	}
	return s.engine.CancelWorkflow(WithPrincipal(ctx, p), id, runID)
}

func (s *Service) WorkflowStats(ctx context.Context, p Principal) (map[string]any, error) {
	p, err := s.CurrentPrincipal(ctx, p)
	if err != nil {
		return nil, err
	}
	stats, err := s.engine.NativeWorkflowStats(ctx)
	if err != nil {
		return nil, err
	}
	visible := func(ids []string) []string {
		out := []string{}
		for _, id := range ids {
			if s.nativeCan(ctx, p, "workflow_run", id, "read") {
				out = append(out, id)
			}
		}
		return out
	}
	return map[string]any{"concurrency": stats.Concurrency, "pendingRunIds": visible(stats.PendingRunIDs), "processingRunIds": visible(stats.ProcessingRunIDs)}, nil
}
func (s *Service) Runs(ctx context.Context, p Principal) (map[string]any, error) {
	workflows, err := s.NativeList(ctx, p, "workflow_run", embedded.NativeOptions{Page: 1, PerPage: 100, Sort: "-created"})
	if err != nil {
		return nil, err
	}
	dns, err := s.DNSOperations(ctx, p, 100)
	if err != nil {
		return nil, err
	}
	deployments, err := s.Deployments(ctx, p)
	if err != nil {
		return nil, err
	}
	return map[string]any{"workflows": workflows, "dns": dns, "deployments": deployments}, nil
}

// Adopted legacy resources are linked only through explicit administrator
// actions; this helper does not infer background execution authority.
func (s *Service) bindImported(ctx context.Context, p Principal, kind, id string) error {
	binding := ResourceBinding{Kind: kind, ResourceID: id, OwnerID: p.UserID, Active: true, Version: 1}
	return s.db.WithContext(ctx).Clauses(clause.OnConflict{DoNothing: true}).Create(&binding).Error
}
