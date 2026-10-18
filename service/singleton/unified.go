package singleton

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/certimate-go/certimate/pkg/embedded"
	"github.com/nezhahq/nezha/model"
	"github.com/nezhahq/nezha/pkg/agentproto"
	"github.com/nezhahq/nezha/service/unified"
	"gorm.io/gorm"
)

var Unified *unified.Service

func InitUnified(dir string) error {
	engine, err := embedded.Open(dir)
	if err != nil {
		return err
	}
	service, err := unified.New(DB, engine)
	if err != nil {
		_ = engine.Close()
		return err
	}
	Unified = service
	return nil
}

type unifiedHostIP struct {
	IP         model.IP
	ObservedAt time.Time
}

var unifiedHostReports = struct {
	sync.RWMutex
	hosts map[uint64]unifiedHostIP
}{hosts: make(map[uint64]unifiedHostIP)}

// ReportUnifiedHostIP receives a copied authenticated RPC observation. Keeping
// this snapshot separate avoids reading the legacy mutable GeoIP pointer.
func ReportUnifiedHostIP(id uint64, ip model.IP, observedAt time.Time) {
	unifiedHostReports.Lock()
	previous, exists := unifiedHostReports.hosts[id]
	if !exists || !observedAt.Before(previous.ObservedAt) {
		unifiedHostReports.hosts[id] = unifiedHostIP{IP: ip, ObservedAt: observedAt}
	}
	unifiedHostReports.Unlock()
	if Unified != nil {
		Unified.NotifyDDNSHost(id)
	}
}

func ReportUnifiedHostAddresses(id uint64, addresses []agentproto.HostAddress, observedAt time.Time) {
	if Unified == nil {
		return
	}
	values := make([]unified.DDNSAddress, len(addresses))
	for i, address := range addresses {
		values[i] = unified.DDNSAddress{Interface: address.Interface, Address: address.Address, Temporary: address.Temporary, Deprecated: address.Deprecated}
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_ = Unified.ReportDDNSHostAddresses(ctx, id, values, observedAt)
}

func UnifiedHostSource(ctx context.Context, id uint64) (unified.DDNSHostSnapshot, error) {
	if err := ctx.Err(); err != nil {
		return unified.DDNSHostSnapshot{}, err
	}
	if ServerShared == nil {
		return unified.DDNSHostSnapshot{}, unified.ErrDDNSUnavailable
	}
	server, ok := ServerShared.Get(id)
	if !ok || server == nil {
		return unified.DDNSHostSnapshot{}, embedded.ErrNotFound
	}
	runtime := server.RuntimeSnapshot()
	unifiedHostReports.RLock()
	observed, known := unifiedHostReports.hosts[id]
	unifiedHostReports.RUnlock()
	snapshot := unified.DDNSHostSnapshot{OwnerID: server.GetUserID(), Online: !runtime.LastActive.IsZero() && time.Since(runtime.LastActive) < 3*time.Minute}
	if known {
		snapshot.IPv4, snapshot.IPv6, snapshot.ObservedAt = observed.IP.IPv4Addr, observed.IP.IPv6Addr, observed.ObservedAt
	}
	if Unified != nil {
		report, err := Unified.DDNSHostAddressReport(ctx, id)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return unified.DDNSHostSnapshot{}, err
		}
		if err == nil {
			snapshot.Addresses, snapshot.AddressesObservedAt = report.Addresses, report.ObservedAt
		}
	}
	return snapshot, nil
}
