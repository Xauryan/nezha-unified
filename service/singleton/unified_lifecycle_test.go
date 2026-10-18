package singleton

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/nezhahq/nezha/model"
)

func TestUnifiedHostSourceCopiesAuthenticatedObservation(t *testing.T) {
	oldServers, oldUnified := ServerShared, Unified
	Unified = nil
	server := &model.Server{Common: model.Common{ID: 71}, LastActive: time.Now()}
	server.SetUserID(5)
	ServerShared = &ServerClass{class: class[uint64, *model.Server]{list: map[uint64]*model.Server{71: server}}}
	t.Cleanup(func() {
		ServerShared, Unified = oldServers, oldUnified
		unifiedHostReports.Lock()
		delete(unifiedHostReports.hosts, 71)
		unifiedHostReports.Unlock()
	})
	now := time.Now()
	ReportUnifiedHostIP(71, model.IP{IPv4Addr: "192.0.2.1"}, now)
	ReportUnifiedHostIP(71, model.IP{IPv4Addr: "192.0.2.99"}, now.Add(-time.Minute))
	snapshot, err := UnifiedHostSource(context.Background(), 71)
	if err != nil {
		t.Fatal(err)
	}
	if snapshot.OwnerID != 5 || !snapshot.Online || snapshot.IPv4 != "192.0.2.1" || !snapshot.ObservedAt.Equal(now) {
		t.Fatalf("snapshot did not retain newest authenticated report: %#v", snapshot)
	}
	server.SetUserID(6)
	snapshot, err = UnifiedHostSource(context.Background(), 71)
	if err != nil || snapshot.OwnerID != 6 {
		t.Fatal("ownership rotation was not reflected")
	}
	if _, err := UnifiedHostSource(context.Background(), 72); err == nil {
		t.Fatal("missing host accepted")
	}
	cancelled, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := UnifiedHostSource(cancelled, 71); err == nil {
		t.Fatal("cancelled source accepted")
	}
	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				ReportUnifiedHostIP(71, model.IP{IPv4Addr: "192.0.2.2"}, now.Add(time.Second))
				if _, err := UnifiedHostSource(context.Background(), 71); err != nil {
					t.Error(err)
				}
			}
		}()
	}
	wg.Wait()
}
