//go:build certimate_full

package embedded

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/certimate-go/certimate/internal/certmgmt/deployers"
	"github.com/certimate-go/certimate/internal/domain"
	"github.com/certimate-go/certimate/pkg/core"
)

var nativeDeployTestSequence atomic.Int32

type nativeSyntheticDeployer struct {
	logger *slog.Logger
	called *int
	fail   *bool
	t      *testing.T
}

func (d *nativeSyntheticDeployer) SetLogger(logger *slog.Logger) { d.logger = logger }
func (d *nativeSyntheticDeployer) Deploy(ctx context.Context, cert, key string) (*core.DeployerDeployResult, error) {
	*d.called++
	deadline, ok := ctx.Deadline()
	if !ok || time.Until(deadline) > 10*time.Minute {
		d.t.Error("provider did not receive the bounded context")
	}
	if d.logger == nil {
		d.t.Error("provider logger was not set")
	} else {
		d.logger.Info("synthetic-provider-secret", slog.String("key", key), slog.String("certificate", cert))
	}
	if *d.fail {
		return nil, errors.New("synthetic-sdk-secret")
	}
	return &core.DeployerDeployResult{}, nil
}

func TestNativeDeployUsesRegistryWithCurrentCredentialAndNoRawDiagnostics(t *testing.T) {
	e := testEngine(t)
	ctx := context.Background()
	credential, err := e.CreateCredential(ctx, "synthetic", CredentialInput{Name: "synthetic", Provider: "synthetic", Config: json.RawMessage(`{"token":"synthetic-provider-token"}`)})
	if err != nil {
		t.Fatal(err)
	}
	certificate, err := e.ImportCertificate(ctx, "synthetic", testCertificate(t))
	if err != nil {
		t.Fatal(err)
	}
	material, err := e.CertificateMaterial(ctx, certificate.ID)
	if err != nil {
		t.Fatal(err)
	}
	use := CredentialUse{ActorID: 1, CredentialID: credential.ID, Version: 1, Purpose: "deploy.target", ResourceKind: "deployment_target", ResourceID: "target-1"}
	provider := domain.DeploymentProviderType(fmt.Sprintf("synthetic-native-%d", nativeDeployTestSequence.Add(1)))
	constructed, deployed := 0, 0
	fail := false
	if err := deployers.Registries.Register(provider, func(options *deployers.ProviderFactoryOptions) (core.Deployer, error) {
		constructed++
		if options.ProviderAccessConfig["token"] != "synthetic-provider-token" && options.ProviderAccessConfig["token"] != "synthetic-provider-rotated" {
			t.Error("resolved credential missing")
		}
		if options.ProviderExtendedConfig["endpoint"] != "synthetic.invalid" {
			t.Error("target config missing")
		}
		options.ProviderExtendedConfig["providerMutation"] = "synthetic"
		return &nativeSyntheticDeployer{called: &deployed, fail: &fail, t: t}, nil
	}); err != nil {
		t.Fatal(err)
	}
	config := map[string]any{"endpoint": "synthetic.invalid"}
	checks := 0
	e.SetCredentialAuthorizer(func(_ context.Context, requested CredentialUse) error {
		checks++
		if requested.ResourceID != use.ResourceID || requested.Purpose != "deploy.target" || checks == 2 {
			return errors.New("synthetic revoked grant")
		}
		return nil
	})
	if err := e.NativeDeployProvider(ctx, use, string(provider), config, material); err == nil || constructed != 0 || deployed != 0 {
		t.Fatalf("revoked target reached provider: %v constructors=%d effects=%d", err, constructed, deployed)
	}
	e.SetCredentialAuthorizer(func(context.Context, CredentialUse) error { return nil })
	var diagnostics bytes.Buffer
	previous := slog.Default()
	slog.SetDefault(slog.New(slog.NewTextHandler(&diagnostics, nil)))
	defer slog.SetDefault(previous)
	if err := e.NativeDeployProvider(ctx, use, string(provider), config, material); err != nil {
		t.Fatal(err)
	}
	if constructed != 1 || deployed != 1 {
		t.Fatalf("original registry path not invoked: %d %d", constructed, deployed)
	}
	if _, mutated := config["providerMutation"]; mutated {
		t.Fatal("provider mutated immutable target configuration")
	}
	if diagnostics.Len() != 0 {
		t.Fatal("provider diagnostics reached the default logger")
	}
	bad := material
	bad.Domains = []string{"metadata-forgery.invalid"}
	if err := e.NativeDeployProvider(ctx, use, string(provider), config, bad); err == nil || deployed != 1 {
		t.Fatal("forged certificate names reached provider")
	}
	if _, err := e.UpdateCredential(ctx, "synthetic", credential.ID, 1, false, nil); err != nil {
		t.Fatal(err)
	}
	if err := e.NativeDeployProvider(ctx, use, string(provider), config, material); err == nil || deployed != 1 {
		t.Fatal("disabled credential reached provider")
	}
	if _, err := e.UpdateCredential(ctx, "synthetic", credential.ID, 1, true, json.RawMessage(`{"token":"synthetic-provider-rotated"}`)); err != nil {
		t.Fatal(err)
	}
	if err := e.NativeDeployProvider(ctx, use, string(provider), config, material); err == nil || deployed != 1 {
		t.Fatal("old credential version reached provider")
	}
	use.Version = 2
	fail = true
	err = e.NativeDeployProvider(ctx, use, string(provider), config, material)
	if err == nil || strings.Contains(err.Error(), "synthetic-sdk-secret") || deployed != 2 {
		t.Fatalf("provider failure was not safely reported: %v", err)
	}
}
