package engine

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"github.com/certimate-go/certimate/internal/domain"
	"github.com/certimate-go/certimate/internal/embeddedhooks"
	"github.com/certimate-go/certimate/pkg/logging"
	"io"
	"log/slog"
	"math/big"
	"net"
	"strings"
	"testing"
	"time"
)

func syntheticCertificate(t *testing.T) *domain.Certificate {
	t.Helper()
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	tpl := &x509.Certificate{SerialNumber: big.NewInt(1), Subject: pkix.Name{CommonName: "untrusted-cn.example.test"}, DNSNames: []string{"example.test", "*.example.test"}, IPAddresses: []net.IP{net.ParseIP("192.0.2.7")}, NotBefore: time.Now().Add(-time.Hour), NotAfter: time.Now().Add(time.Hour)}
	raw, err := x509.CreateCertificate(rand.Reader, tpl, tpl, &key.PublicKey, key)
	if err != nil {
		t.Fatal(err)
	}
	private, err := x509.MarshalPKCS8PrivateKey(key)
	if err != nil {
		t.Fatal(err)
	}
	return &domain.Certificate{Meta: domain.Meta{Id: "cert"}, Certificate: string(pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: raw})), PrivateKey: string(pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: private})), SubjectAltNames: "forged-metadata.example.test"}
}
func TestCertificateAuthorizationUsesActualCompleteSANs(t *testing.T) {
	var seen embeddedhooks.Operation
	embeddedhooks.Configure(&embeddedhooks.Hooks{Authorize: func(_ context.Context, op embeddedhooks.Operation) error { seen = op; return nil }})
	t.Cleanup(func() { embeddedhooks.Configure(nil) })
	ctx := (&NodeExecutionContext{}).SetContext(embeddedhooks.Bind(context.Background())).SetExecutingNode(&Node{Id: "node"}).SetExecutingWorkflow("workflow", "run", nil)
	if err := authorizeCertificate(ctx, "certificate.deploy", syntheticCertificate(t)); err != nil {
		t.Fatal(err)
	}
	if strings.Join(seen.Domains, ",") != "example.test,*.example.test,192.0.2.7" {
		t.Fatalf("incomplete or metadata-derived SANs: %#v", seen.Domains)
	}
}

type probeExecutor struct {
	nodeExecutor
	calls int
	fail  bool
}

func (p *probeExecutor) Execute(ctx *NodeExecutionContext) (*NodeExecutionResult, error) {
	p.calls++
	p.logger.Info("synthetic-provider-secret", slog.String("key", "synthetic-key"))
	if p.fail {
		return nil, errors.New("synthetic-provider-secret")
	}
	return newNodeExecutionResult(ctx.Node), nil
}
func TestEngineRechecksNodesAndSanitizesDiagnostics(t *testing.T) {
	revoked := false
	embeddedhooks.Configure(&embeddedhooks.Hooks{Authorize: func(_ context.Context, op embeddedhooks.Operation) error {
		if revoked && op.Action == "node.execute" {
			return errors.New("revoked")
		}
		return nil
	}})
	t.Cleanup(func() { embeddedhooks.Configure(nil) })
	probe := &probeExecutor{fail: true}
	engine := &workflowEngine{executors: map[NodeType]NodeExecutor{NodeTypeStart: probe}, syslog: slog.New(slog.NewTextHandler(io.Discard, nil))}
	var logged logging.Record
	engine.OnNodeLogging(func(_ context.Context, _ *Node, r logging.Record) error { logged = r; return nil })
	graph := &Graph{Nodes: []*Node{{Id: "node", Type: NodeTypeStart}}}
	err := engine.Invoke(context.Background(), WorkflowExecution{WorkflowId: "workflow", RunId: "run", Graph: graph})
	if !errors.Is(err, embeddedhooks.ErrExecution) {
		t.Fatalf("unsafe error returned: %v", err)
	}
	if strings.Contains(logged.Message, "secret") || logged.NumAttrs() != 0 {
		t.Fatal("provider diagnostic leaked")
	}
	revoked = true
	err = engine.Invoke(context.Background(), WorkflowExecution{WorkflowId: "workflow", RunId: "run", Graph: graph})
	if !errors.Is(err, embeddedhooks.ErrDenied) || probe.calls != 1 {
		t.Fatal("revoked node executed")
	}
}

func TestAgentDeploymentBindsVersionFingerprintAndRetryIdentity(t *testing.T) {
	var requests []embeddedhooks.TargetDeployment
	embeddedhooks.Configure(&embeddedhooks.Hooks{
		Authorize: func(context.Context, embeddedhooks.Operation) error { return nil },
		DeployTarget: func(_ context.Context, request embeddedhooks.TargetDeployment) error {
			requests = append(requests, request)
			return nil
		},
	})
	t.Cleanup(func() { embeddedhooks.Configure(nil) })
	ctx := (&NodeExecutionContext{}).SetContext(embeddedhooks.Bind(context.Background())).SetExecutingNode(&Node{Id: "node"}).SetExecutingWorkflow("workflow", "run", nil)
	cfg := domain.WorkflowNodeConfigForBizDeploy{Provider: "nezha-agent", ProviderAccessId: "42", ProviderConfig: map[string]any{"targetVersion": 2}}
	cert := syntheticCertificate(t)
	for range 2 {
		if err := deployNezhaTarget(ctx, cfg, cert); err != nil {
			t.Fatal(err)
		}
	}
	if len(requests) != 2 || requests[0].OperationID != requests[1].OperationID || requests[0].TargetVersion != 2 || len(requests[0].Domains) != 3 || len(requests[0].Fingerprint) != 64 {
		t.Fatal("deployment retry binding incomplete")
	}
	cfg.ProviderConfig["targetVersion"] = 3
	if err := deployNezhaTarget(ctx, cfg, cert); err != nil {
		t.Fatal(err)
	}
	if requests[0].OperationID == requests[2].OperationID {
		t.Fatal("target version change reused idempotency key")
	}
}
