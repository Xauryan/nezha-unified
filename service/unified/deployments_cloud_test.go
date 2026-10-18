package unified

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"testing"

	"github.com/certimate-go/certimate/pkg/embedded"
)

func TestCloudTargetSealsConfigurationAndPinsCredentialAndTargetVersions(t *testing.T) {
	s, p, _, cert := deploymentFixture(t)
	ctx := context.Background()
	credential, err := s.CreateCredential(ctx, p, embedded.CredentialInput{Name: "synthetic cloud", Provider: "cloudflare", Config: json.RawMessage(`{"apiToken":"synthetic-cloud-token"}`)})
	if err != nil {
		t.Fatal(err)
	}
	target, err := s.SaveTarget(ctx, p, 0, DeploymentTarget{Name: "synthetic cloud", Kind: "certimate", Provider: "synthetic-provider-without-network", CredentialID: credential.ID, CredentialVersion: credential.Version, AllowedNames: []string{"synthetic.example.test"}, Enabled: true, Config: map[string]any{"endpoint": "synthetic.invalid", "material": "synthetic-target-setting"}})
	if err != nil {
		t.Fatal(err)
	}
	encoded, _ := json.Marshal(target)
	if strings.Contains(string(encoded), "synthetic-target-setting") || strings.Contains(string(encoded), "configuration_envelope") || target.Config != nil || !target.Configured {
		t.Fatal("cloud metadata exposed secret settings")
	}
	var saved DeploymentTarget
	if err = s.db.First(&saved, target.ID).Error; err != nil {
		t.Fatal(err)
	}
	if saved.ConfigurationEnvelope == "" || strings.Contains(saved.ConfigurationEnvelope, "synthetic-target-setting") {
		t.Fatal("target configuration was not encrypted")
	}
	oldVersion := target.Version
	target, err = s.SaveTarget(ctx, p, target.ID, target)
	if err != nil {
		t.Fatal(err)
	}
	config, err := s.engine.OpenTargetConfiguration(ctx, strconv.FormatUint(target.ID, 10), target.Version, target.ConfigurationEnvelope)
	if err != nil || config["material"] != "synthetic-target-setting" {
		t.Fatal("blank metadata edit lost saved settings")
	}
	if _, err = s.engine.OpenTargetConfiguration(ctx, strconv.FormatUint(target.ID, 10), oldVersion, target.ConfigurationEnvelope); err == nil {
		t.Fatal("configuration crossed a target version")
	}
	d, err := s.EnqueueDeployment(ctx, p, target.ID, DeploymentInput{OperationID: "synthetic-cloud-unknown", TargetVersion: target.Version, CertificateID: cert.ID})
	if err != nil {
		t.Fatal(err)
	}
	s.deployments = &deploymentRuntime{wake: make(chan struct{}, 1)}
	if err = s.db.Model(&Deployment{}).Where("id = ?", d.ID).Update("state", "dispatching").Error; err != nil {
		t.Fatal(err)
	}
	use := embedded.CredentialUse{ActorID: p.UserID, CredentialID: credential.ID, Version: credential.Version, Purpose: "deploy.target", ResourceKind: "deployment_target", ResourceID: strconv.FormatUint(target.ID, 10), OperationID: d.OperationID}
	if _, err = s.engine.ResolveCredential(WithPrincipal(ctx, p), use); err != nil {
		t.Fatal(err)
	}
	if err = s.db.Model(&Deployment{}).Where("id = ?", d.ID).Updates(map[string]any{"state": "unknown"}).Error; err != nil {
		t.Fatal(err)
	}
	result, err := s.processDeployment(ctx, d.ID, nil, true)
	if err == nil || result.State != "reconciliation_required" {
		t.Fatal("uncertain cloud mutation was replayed")
	}
	if _, err = s.engine.ResolveCredential(WithPrincipal(ctx, p), use); err == nil {
		t.Fatal("reconciliation fence allowed credential execution")
	}
}

func cloudDeploymentFixture(t *testing.T) (*Service, Principal, DeploymentTarget, embedded.Resource, embedded.Resource) {
	t.Helper()
	s, p, _, cert := deploymentFixture(t)
	ctx := context.Background()
	credential, err := s.CreateCredential(ctx, p, embedded.CredentialInput{Name: "synthetic-cloud-credential", Provider: "cloudflare", Config: json.RawMessage(`{"apiToken":"synthetic-cloud-material"}`)})
	if err != nil {
		t.Fatal(err)
	}
	target, err := s.SaveTarget(ctx, p, 0, DeploymentTarget{Name: "synthetic-cloud-target", Kind: "certimate", Provider: "synthetic-provider-without-network", CredentialID: credential.ID, CredentialVersion: credential.Version, AllowedNames: []string{"synthetic.example.test"}, Enabled: true})
	if err != nil {
		t.Fatal(err)
	}
	return s, p, target, cert, credential
}

func TestCloudDeploymentCurrentAuthorizationRejectsDisabledOrRotatedCredential(t *testing.T) {
	for _, rotate := range []bool{false, true} {
		name := "disabled"
		if rotate {
			name = "rotated"
		}
		t.Run(name, func(t *testing.T) {
			s, p, target, cert, credential := cloudDeploymentFixture(t)
			ctx := context.Background()
			d, err := s.EnqueueDeployment(ctx, p, target.ID, DeploymentInput{OperationID: "synthetic-cloud-running", TargetVersion: target.Version, CertificateID: cert.ID})
			if err != nil {
				t.Fatal(err)
			}
			if err = s.db.Model(&Deployment{}).Where("id = ?", d.ID).Update("state", "dispatching").Error; err != nil {
				t.Fatal(err)
			}
			use := embedded.CredentialUse{ActorID: p.UserID, CredentialID: credential.ID, Version: credential.Version, Purpose: "deploy.target", ResourceKind: "deployment_target", ResourceID: strconv.FormatUint(target.ID, 10), OperationID: d.OperationID}
			if err = s.deploymentAuthorization(ctx, p, target, cert.ID); err != nil {
				t.Fatal("current cloud authorization rejected baseline", err)
			}
			if err = s.authorizeTargetCredential(WithPrincipal(ctx, p), use); err != nil {
				t.Fatal("cloud credential authorization rejected baseline", err)
			}
			var replacement json.RawMessage
			if rotate {
				replacement = json.RawMessage(`{"apiToken":"synthetic-cloud-replacement"}`)
			}
			if _, err = s.UpdateCredential(ctx, p, credential.ID, credential.Version, rotate, replacement); err != nil {
				t.Fatal(err)
			}
			// The in-flight cloud watchdog calls this same live guard after a
			// provider has retained its already-decrypted configuration.
			if err = s.deploymentAuthorization(ctx, p, target, cert.ID); !errors.Is(err, ErrDenied) {
				t.Fatal("live cloud deployment guard accepted changed credential", err)
			}
			if err = s.authorizeTargetCredential(WithPrincipal(ctx, p), use); !errors.Is(err, ErrDenied) {
				t.Fatal("changed cloud credential passed execution authorization", err)
			}
			blocked, err := s.processDeployment(ctx, d.ID, nil, false)
			if !errors.Is(err, ErrDenied) || blocked.State != "reconciliation_required" {
				t.Fatal("interrupted cloud execution did not retain unresolved outcome", err)
			}
		})
	}
}
