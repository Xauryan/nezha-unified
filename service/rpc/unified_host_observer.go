package rpc

import (
	"github.com/nezhahq/nezha/model"
	"github.com/nezhahq/nezha/pkg/agentproto"
	"sync/atomic"
	"time"
)

var unifiedHostObserver atomic.Pointer[func(uint64)]
var unifiedHostIPObserver atomic.Pointer[func(uint64, model.IP, time.Time)]
var unifiedHostAddressesObserver atomic.Pointer[func(uint64, []agentproto.HostAddress, time.Time)]

func SetUnifiedHostAddressesObserver(fn func(uint64, []agentproto.HostAddress, time.Time)) {
	unifiedHostAddressesObserver.Store(&fn)
}
func notifyUnifiedHostAddresses(id uint64, addresses []agentproto.HostAddress) {
	if p := unifiedHostAddressesObserver.Load(); p != nil && *p != nil {
		(*p)(id, addresses, time.Now())
	}
}

func SetUnifiedHostObserver(fn func(uint64))                        { unifiedHostObserver.Store(&fn) }
func SetUnifiedHostIPObserver(fn func(uint64, model.IP, time.Time)) { unifiedHostIPObserver.Store(&fn) }
func notifyUnifiedHost(id uint64) {
	if p := unifiedHostObserver.Load(); p != nil && *p != nil {
		(*p)(id)
	}
}
func notifyUnifiedHostIP(id uint64, ip model.IP) {
	if p := unifiedHostIPObserver.Load(); p != nil && *p != nil {
		(*p)(id, ip, time.Now())
	}
}
