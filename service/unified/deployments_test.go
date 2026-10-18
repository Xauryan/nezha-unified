package unified

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"github.com/certimate-go/certimate/pkg/embedded"
	"github.com/nezhahq/nezha/model"
	"github.com/nezhahq/nezha/pkg/agentproto"
	pb "github.com/nezhahq/nezha/proto"
	"github.com/nezhahq/nezha/service/agentcall"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"math/big"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func deploymentFixture(t *testing.T) (*Service, Principal, DeploymentTarget, embedded.Resource) {
	t.Helper()
	dir := t.TempDir()
	engine, err := embedded.Open(filepath.Join(dir, "certimate"))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { engine.Close() })
	db, err := gorm.Open(sqlite.Open(filepath.Join(dir, "host.db")), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	sqlDB, _ := db.DB()
	t.Cleanup(func() { sqlDB.Close() })
	if err = db.AutoMigrate(&model.User{}, &model.Server{}); err != nil {
		t.Fatal(err)
	}
	if err = db.Create(&model.User{Common: model.Common{ID: 1}, Username: "synthetic-admin", Role: model.RoleAdmin, AgentSecret: "synthetic-agent-secret"}).Error; err != nil {
		t.Fatal(err)
	}
	if err = db.Create(&model.Server{Common: model.Common{ID: 10}, Name: "synthetic"}).Error; err != nil {
		t.Fatal(err)
	}
	s, err := New(db, engine)
	if err != nil {
		t.Fatal(err)
	}
	admin := Principal{UserID: 1, Admin: true}
	target, err := s.SaveTarget(context.Background(), admin, 0, DeploymentTarget{Name: "synthetic", Kind: "nezha", ServerID: 10, AllowedNames: []string{"synthetic.example.test"}, BundlePath: filepath.Join(dir, "bundle"), Enabled: true, CheckCommand: []string{"/synthetic/check"}, ReloadCommand: []string{"/synthetic/reload"}, HealthCommand: []string{"/synthetic/health"}})
	if err != nil {
		t.Fatal(err)
	}
	key, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	tmpl := x509.Certificate{SerialNumber: big.NewInt(1), DNSNames: []string{"synthetic.example.test"}, NotBefore: time.Now().Add(-time.Hour), NotAfter: time.Now().Add(time.Hour), KeyUsage: x509.KeyUsageDigitalSignature}
	der, _ := x509.CreateCertificate(rand.Reader, &tmpl, &tmpl, &key.PublicKey, key)
	keyDER, _ := x509.MarshalPKCS8PrivateKey(key)
	cert, err := s.ImportCertificate(context.Background(), admin, embedded.CertificateInput{Certificate: string(pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: der})), PrivateKey: string(pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: keyDER}))})
	if err != nil {
		t.Fatal(err)
	}
	return s, admin, target, cert
}
func TestDeploymentDurabilitySupersessionAndCurrentAuthority(t *testing.T) {
	s, p, target, cert := deploymentFixture(t)
	ctx := context.Background()
	in := DeploymentInput{OperationID: "synthetic-first", TargetVersion: target.Version, CertificateID: cert.ID}
	first, err := s.EnqueueDeployment(ctx, p, target.ID, in)
	if err != nil {
		t.Fatal(err)
	}
	again, err := s.EnqueueDeployment(ctx, p, target.ID, in)
	if err != nil || again.ID != first.ID {
		t.Fatal("idempotent enqueue", err)
	}
	encoded, _ := json.Marshal(first)
	if strings.Contains(string(encoded), "PRIVATE KEY") || strings.Contains(string(encoded), "PrincipalJSON") || strings.Contains(string(encoded), "certificate_pem") {
		t.Fatal("task exposed secret payload")
	}
	var calls int
	s.deployments = &deploymentRuntime{calls: agentcall.New(func(context.Context, uint64, *pb.Task) error { calls++; return nil }, 4), wake: make(chan struct{}, 1)}
	in.OperationID = "synthetic-second"
	second, err := s.EnqueueDeployment(ctx, p, target.ID, in)
	if err != nil {
		t.Fatal(err)
	}
	old, err := s.processDeployment(ctx, first.ID, nil, false)
	if err == nil || old.State != "superseded" || calls != 0 {
		t.Fatal("old task reached transport")
	}
	if err = s.db.Model(&model.User{}).Where("id = ?", p.UserID).Update("role", model.RoleMember).Error; err != nil {
		t.Fatal(err)
	}
	denied, err := s.processDeployment(ctx, second.ID, nil, false)
	if err == nil || denied.State != "cancelled" || calls != 0 {
		t.Fatal("persisted administrator bypassed current role")
	}
}
func TestDeploymentUnknownReconcilesWithoutReplay(t *testing.T) {
	s, p, target, cert := deploymentFixture(t)
	d, err := s.EnqueueDeployment(context.Background(), p, target.ID, DeploymentInput{OperationID: "synthetic-unknown", TargetVersion: target.Version, CertificateID: cert.ID})
	if err != nil {
		t.Fatal(err)
	}
	var broker *agentcall.Broker
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	deployCount, reconcileCount := 0, 0
	broker = agentcall.New(func(ctx context.Context, server uint64, task *pb.Task) error {
		if task.Type == agentproto.TaskCapabilities {
			b, _ := json.Marshal(agentproto.Capabilities{Protocol: 1, CertificateDeploy: true})
			broker.Deliver(server, &pb.TaskResult{Id: task.Id, Type: task.Type, Successful: true, Data: string(b)})
			return nil
		}
		var req agentproto.Request
		_ = json.Unmarshal([]byte(task.Data), &req)
		if req.Method == "deploy" {
			deployCount++
			cancel()
			return nil
		}
		if req.Method == "cancel" {
			out, _ := json.Marshal(agentproto.Response{Protocol: 1, OperationID: d.OperationID, Final: true, Phase: "cancel_requested"})
			broker.Deliver(server, &pb.TaskResult{Id: task.Id, Type: task.Type, Successful: true, Data: string(out)})
			return nil
		}
		if req.Method != "reconcile" {
			t.Errorf("unexpected method %s", req.Method)
			return nil
		}
		reconcileCount++
		var payload agentproto.Deployment
		_ = json.Unmarshal(req.Payload, &payload)
		if payload.CertificatePEM != "" || payload.PrivateKeyPEM != "" {
			t.Error("reconcile retransmitted key")
		}
		r := agentproto.Receipt{OperationID: d.OperationID, TargetID: fmt.Sprint(target.ID), TargetVersion: target.Version, Sequence: d.ID, Fingerprint: d.Fingerprint, Verified: true, Phase: "verified"}
		body, _ := json.Marshal(r)
		out, _ := json.Marshal(agentproto.Response{Protocol: 1, OperationID: d.OperationID, Final: true, Payload: body})
		broker.Deliver(server, &pb.TaskResult{Id: task.Id, Type: task.Type, Successful: true, Data: string(out)})
		return nil
	}, 4)
	s.deployments = &deploymentRuntime{calls: broker, wake: make(chan struct{}, 1)}
	unknown, err := s.processDeployment(ctx, d.ID, nil, false)
	if err == nil || unknown.State != "unknown" {
		t.Fatal("timeout was not unknown", err, unknown.State)
	}
	verified, err := s.processDeployment(context.Background(), d.ID, nil, true)
	if err != nil || verified.State != "verified" || deployCount != 1 || reconcileCount != 1 {
		t.Fatal("unsafe reconciliation", err, deployCount, reconcileCount)
	}
	var latest DeploymentTarget
	if err = s.db.First(&latest, target.ID).Error; err != nil || latest.AppliedVersion != d.ID {
		t.Fatal("verified target version not durable", err)
	}
}
func TestTargetRequiresExactVersionAndRestrictedToken(t *testing.T) {
	s, p, target, _ := deploymentFixture(t)
	changed := target
	changed.Version = target.Version + 1
	if _, err := s.SaveTarget(context.Background(), p, target.ID, changed); err == nil {
		t.Fatal("stale target editor overwritten")
	}
	if _, err := s.SaveTarget(context.Background(), Principal{UserID: p.UserID, Admin: true, PAT: true}, target.ID, target); err == nil {
		t.Fatal("admin PAT changed target")
	}
	if _, err := s.Targets(context.Background(), Principal{UserID: p.UserID, Admin: true, PAT: true}); err == nil {
		t.Fatal("admin PAT listed target")
	}
}

func TestDeploymentRevocationBeforeCommitStopsPermit(t *testing.T) {
	s, p, target, cert := deploymentFixture(t)
	d, err := s.EnqueueDeployment(context.Background(), p, target.ID, DeploymentInput{OperationID: "synthetic-before-commit", TargetVersion: target.Version, CertificateID: cert.ID})
	if err != nil {
		t.Fatal(err)
	}
	var broker *agentcall.Broker
	permits, cancels := 0, 0
	broker = agentcall.New(func(_ context.Context, server uint64, task *pb.Task) error {
		if task.Type == agentproto.TaskCapabilities {
			data, _ := json.Marshal(agentproto.Capabilities{Protocol: 1, CertificateDeploy: true})
			broker.Deliver(server, &pb.TaskResult{Id: task.Id, Type: task.Type, Successful: true, Data: string(data)})
			return nil
		}
		var req agentproto.Request
		_ = json.Unmarshal([]byte(task.Data), &req)
		switch req.Method {
		case "deploy":
			if err := s.db.Model(&model.User{}).Where("id = ?", p.UserID).Update("role", model.RoleMember).Error; err != nil {
				t.Fatal(err)
			}
			out, _ := json.Marshal(agentproto.Response{Protocol: 1, OperationID: d.OperationID, Phase: "ready"})
			broker.Deliver(server, &pb.TaskResult{Id: task.Id, Type: task.Type, Successful: true, Data: string(out)})
		case "permit":
			permits++
		case "cancel":
			cancels++
			out, _ := json.Marshal(agentproto.Response{Protocol: 1, OperationID: d.OperationID, Phase: "cancel_requested", Final: true})
			broker.Deliver(server, &pb.TaskResult{Id: task.Id, Type: task.Type, Successful: true, Data: string(out)})
		}
		return nil
	}, 4)
	s.deployments = &deploymentRuntime{calls: broker, wake: make(chan struct{}, 1)}
	out, err := s.processDeployment(context.Background(), d.ID, nil, false)
	if err == nil || out.State != "unknown" || permits != 0 || cancels != 1 {
		t.Fatalf("commit gate bypass: state=%s permits=%d cancels=%d error=%v", out.State, permits, cancels, err)
	}
}
