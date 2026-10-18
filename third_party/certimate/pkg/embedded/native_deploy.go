//go:build certimate_full

package embedded

import (
	"context"
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"errors"
	"io"
	"log/slog"
	"slices"
	"strings"
	"time"

	"github.com/certimate-go/certimate/internal/certmgmt"
	"github.com/certimate-go/certimate/internal/domain"
)

// NativeDeployProvider executes any registered original Certimate deployer.
// Reusable host targets require an explicitly bound versioned credential; inline
// credential-free deployment remains available through approved workflow nodes.
func (e *Engine) NativeDeployProvider(ctx context.Context, use CredentialUse, provider string, config map[string]any, material CertificateMaterial) error {
	if use.Purpose != "deploy.target" || use.ResourceKind != "deployment_target" || use.ResourceID == "" || use.CredentialID == "" || provider == "" || len(provider) > 100 {
		return ErrInvalid
	}
	ctx, cancel := context.WithTimeout(ctx, 10*time.Minute)
	defer cancel()
	if err := ctx.Err(); err != nil {
		return err
	}
	if len(material.CertificatePEM) > 131072 || len(material.PrivateKeyPEM) > 32768 {
		return ErrInvalid
	}
	pair, err := tls.X509KeyPair([]byte(material.CertificatePEM), []byte(material.PrivateKeyPEM))
	if err != nil || len(pair.Certificate) == 0 {
		return ErrInvalid
	}
	leaf, err := x509.ParseCertificate(pair.Certificate[0])
	if err != nil || leaf.IsCA {
		return ErrInvalid
	}
	digest := sha256.Sum256(pair.Certificate[0])
	fingerprint := hex.EncodeToString(digest[:])
	if material.Fingerprint != fingerprint || material.Version != fingerprint {
		return ErrInvalid
	}
	names := append([]string(nil), leaf.DNSNames...)
	for _, ip := range leaf.IPAddresses {
		names = append(names, ip.String())
	}
	names = append(names, leaf.EmailAddresses...)
	for _, uri := range leaf.URIs {
		names = append(names, uri.String())
	}
	if len(names) == 0 && leaf.Subject.CommonName != "" {
		names = append(names, leaf.Subject.CommonName)
	}
	if !slices.Equal(names, material.Domains) {
		return ErrInvalid
	}
	var publicChain strings.Builder
	for _, der := range pair.Certificate {
		if _, err := x509.ParseCertificate(der); err != nil {
			return ErrInvalid
		}
		publicChain.Write(pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: der}))
	}
	// Provider factories may normalize settings; do not let them mutate the
	// configuration snapshot retained by the host operation.
	encoded, err := json.Marshal(config)
	if err != nil || len(encoded) > 2<<20 {
		return ErrInvalid
	}
	var providerConfig map[string]any
	if json.Unmarshal(encoded, &providerConfig) != nil {
		return ErrInvalid
	}
	if providerConfig == nil {
		providerConfig = map[string]any{}
	}
	access, err := e.ResolveCredential(ctx, use)
	if err != nil {
		return errors.New("deployment credential authorization denied")
	}
	e.mu.RLock()
	if err := e.check(ctx); err != nil {
		e.mu.RUnlock()
		return err
	}
	authorize, authorizerVersion := e.authorizeCredential, e.authorizerVersion
	e.mu.RUnlock()
	if authorize == nil || authorize(ctx, use) != nil {
		return errors.New("deployment target authorization denied")
	}
	e.mu.RLock()
	if err := e.check(ctx); err != nil {
		e.mu.RUnlock()
		return err
	}
	if e.authorizerVersion != authorizerVersion {
		e.mu.RUnlock()
		return errors.New("deployment target authorization changed")
	}
	// Rotation, disabling and host authorization are rechecked immediately
	// before construction of the provider that can perform external work.
	r, err := e.nativeRecord("access", use.CredentialID)
	if err != nil || !r.GetBool("unified_enabled") || r.GetInt("unified_version") != use.Version {
		e.mu.RUnlock()
		return errors.New("deployment credential authorization denied")
	}
	e.effects.Add(1)
	e.mu.RUnlock()
	defer e.effects.Done()
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	_, err = certmgmt.NewClient(certmgmt.WithLogger(logger)).DeployCertificate(ctx, &certmgmt.DeployCertificateRequest{
		Provider: domain.DeploymentProviderType(provider), ProviderAccessConfig: access, ProviderExtendedConfig: providerConfig,
		CertificatePEM: publicChain.String(), PrivateKeyPEM: material.PrivateKeyPEM,
	})
	if err != nil {
		if ctx.Err() != nil {
			return ctx.Err()
		}
		return errors.New("certificate deployment failed")
	}
	return nil
}
