package engine

import (
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/certimate-go/certimate/internal/domain"
	"github.com/certimate-go/certimate/internal/embeddedhooks"
	xcert "github.com/certimate-go/certimate/pkg/utils/cert"
	"strconv"
)

func authorizeEffect(ctx *NodeExecutionContext, action, kind, id string, domains []string) error {
	return embeddedhooks.Check(ctx.Context(), embeddedhooks.Operation{
		WorkflowID: ctx.WorkflowId, RunID: ctx.RunId, NodeID: ctx.Node.Id,
		Action: action, ResourceKind: kind, ResourceID: id, Domains: domains,
	})
}

func certificateNames(cert *x509.Certificate) []string {
	names := append([]string(nil), cert.DNSNames...)
	for _, ip := range cert.IPAddresses {
		names = append(names, ip.String())
	}
	// Unusual SAN types also reach policy evaluation rather than disappearing.
	names = append(names, cert.EmailAddresses...)
	for _, uri := range cert.URIs {
		names = append(names, uri.String())
	}
	if len(names) == 0 && cert.Subject.CommonName != "" {
		names = append(names, cert.Subject.CommonName)
	}
	return names
}

func authorizeCertificate(ctx *NodeExecutionContext, action string, certificate *domain.Certificate) error {
	if !embeddedhooks.EnabledFor(ctx.Context()) {
		return nil
	}
	cert, err := xcert.ParseCertificateFromPEM(certificate.Certificate)
	if err != nil {
		return embeddedhooks.ErrDenied
	}
	return authorizeEffect(ctx, action, "certificate", certificate.Id, certificateNames(cert))
}

func resolveAccess(ctx *NodeExecutionContext, repo accessRepository, id, purpose string) (*domain.Access, error) {
	if !embeddedhooks.EnabledFor(ctx.Context()) {
		return repo.GetById(ctx.Context(), id)
	}
	config, err := embeddedhooks.Credential(ctx.Context(), id, purpose)
	if err != nil {
		return nil, err
	}
	return &domain.Access{Meta: domain.Meta{Id: id}, Config: config}, nil
}

func deployNezhaTarget(ctx *NodeExecutionContext, cfg domain.WorkflowNodeConfigForBizDeploy, certificate *domain.Certificate) error {
	cert, err := xcert.ParseCertificateFromPEM(certificate.Certificate)
	if err != nil {
		return embeddedhooks.ErrDenied
	}
	version, err := strconv.ParseUint(fmt.Sprint(cfg.ProviderConfig["targetVersion"]), 10, 64)
	if err != nil || version == 0 {
		return embeddedhooks.ErrDenied
	}
	sum := sha256.Sum256(cert.Raw)
	fingerprint := hex.EncodeToString(sum[:])
	identity, _ := json.Marshal([]string{ctx.WorkflowId, ctx.RunId, ctx.Node.Id, cfg.ProviderAccessId, strconv.FormatUint(version, 10), fingerprint})
	opSum := sha256.Sum256(identity)
	return embeddedhooks.DeployTarget(ctx.Context(), embeddedhooks.TargetDeployment{
		WorkflowID: ctx.WorkflowId, RunID: ctx.RunId, NodeID: ctx.Node.Id,
		TargetID: cfg.ProviderAccessId, TargetVersion: version, CertificateID: certificate.Id,
		CertificateVersion: fingerprint, Fingerprint: fingerprint, OperationID: hex.EncodeToString(opSum[:]),
		CertificatePEM: certificate.Certificate, PrivateKeyPEM: certificate.PrivateKey, Domains: certificateNames(cert),
	})
}
