package unified

import (
	"context"
	"errors"
	"testing"
	"time"

	"gorm.io/gorm"
)

func TestCloudRetirementRechecksStateAfterWaitingForOperationLock(t *testing.T) {
	s, p, target, cert, _ := cloudDeploymentFixture(t)
	ctx := context.Background()
	d, err := s.EnqueueDeployment(ctx, p, target.ID, DeploymentInput{OperationID: "synthetic-cloud-retirement-race", TargetVersion: target.Version, CertificateID: cert.ID})
	if err != nil {
		t.Fatal(err)
	}
	if err = s.db.Model(&Deployment{}).Where("id = ?", d.ID).Update("state", "unknown").Error; err != nil {
		t.Fatal(err)
	}
	read := make(chan struct{}, 1)
	if err = s.db.Callback().Query().After("gorm:query").Register("synthetic:retirement-read", func(query *gorm.DB) {
		if query.Statement.Schema != nil && query.Statement.Schema.Table == (Deployment{}).TableName() {
			select {
			case read <- struct{}{}:
			default:
			}
		}
	}); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { s.db.Callback().Query().Remove("synthetic:retirement-read") })
	lock := &s.deployments.locks[d.ID%uint64(len(s.deployments.locks))]
	lock.Lock()
	locked := true
	defer func() {
		if locked {
			lock.Unlock()
		}
	}()
	result := make(chan error, 1)
	go func() {
		_, err := s.RetireDeployment(ctx, p, d.ID, true, d.Fingerprint)
		result <- err
	}()
	select {
	case <-read:
	case <-time.After(5 * time.Second):
		t.Fatal("retirement did not load the initial unknown state")
	}
	// A provider can return its acknowledgement while the manual retirement
	// request is waiting for this operation's execution lock.
	if err = s.db.Model(&Deployment{}).Where("id = ?", d.ID).Updates(map[string]any{"state": "applied", "code": "provider_accepted"}).Error; err != nil {
		t.Fatal(err)
	}
	lock.Unlock()
	locked = false
	select {
	case err = <-result:
		if !errors.Is(err, ErrTargetConflict) {
			t.Fatal("retirement overwrote a completed cloud operation", err)
		}
	case <-time.After(5 * time.Second):
		t.Fatal("retirement did not release the operation lock")
	}
	var saved Deployment
	if err = s.db.First(&saved, d.ID).Error; err != nil || saved.State != "applied" || saved.Code != "provider_accepted" {
		t.Fatal("retirement changed the accepted provider outcome", err)
	}
	var audits int64
	if err = s.db.Model(&Audit{}).Where("action = ? AND resource = ?", "deployment.retire", d.ID).Count(&audits).Error; err != nil || audits != 0 {
		t.Fatal("rejected retirement recorded a successful retirement audit", err)
	}
}
