package controller

import (
	"encoding/json"
	"github.com/nezhahq/nezha/service/unified/dnsprovider"
	"strconv"
	"strings"

	"github.com/certimate-go/certimate/pkg/embedded"
	"github.com/gin-gonic/gin"
	"github.com/nezhahq/nezha/service/singleton"
	"github.com/nezhahq/nezha/service/unified"
)

func unifiedID(c *gin.Context) (uint64, bool) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil || id == 0 {
		unifiedResult(c, nil, embedded.ErrInvalid)
		return 0, false
	}
	return id, true
}
func bindUnifiedDNSRoutes(r *gin.RouterGroup) {
	r.POST("/ddns/bindings/:id/service", func(c *gin.Context) {
		id, ok := unifiedID(c)
		if !ok {
			return
		}
		var in struct {
			ServiceID uint64 `json:"service_id"`
			Version   uint64 `json:"version"`
		}
		if c.ShouldBindJSON(&in) != nil {
			unifiedResult(c, nil, embedded.ErrInvalid)
			return
		}
		v, e := singleton.Unified.AssignDDNSService(c.Request.Context(), unifiedPrincipal(c), id, in.ServiceID, in.Version)
		unifiedResult(c, v, e)
	})
	r.POST("/dns/operations/:id/reconcile", func(c *gin.Context) {
		id, ok := unifiedID(c)
		if !ok {
			return
		}
		var in struct {
			RecordID string `json:"record_id"`
		}
		if c.ShouldBindJSON(&in) != nil {
			unifiedResult(c, nil, embedded.ErrInvalid)
			return
		}
		v, e := singleton.Unified.ReconcileDNSOperation(c.Request.Context(), unifiedPrincipal(c), id, in.RecordID)
		unifiedResult(c, v, e)
	})
	r.PATCH("/credentials/:credential", func(c *gin.Context) {
		var in struct {
			Version int             `json:"version"`
			Enabled bool            `json:"enabled"`
			Config  json.RawMessage `json:"config"`
		}
		if c.ShouldBindJSON(&in) != nil {
			unifiedResult(c, nil, embedded.ErrInvalid)
			return
		}
		v, err := singleton.Unified.UpdateCredential(c.Request.Context(), unifiedPrincipal(c), c.Param("credential"), in.Version, in.Enabled, in.Config)
		unifiedResult(c, v, err)
	})
	r.GET("/dns/accounts", func(c *gin.Context) {
		v, err := singleton.Unified.DNSAccounts(c.Request.Context(), unifiedPrincipal(c))
		unifiedResult(c, v, err)
	})
	r.POST("/dns/accounts", func(c *gin.Context) {
		var in unified.DNSAccount
		if c.ShouldBindJSON(&in) != nil {
			unifiedResult(c, nil, embedded.ErrInvalid)
			return
		}
		v, err := singleton.Unified.CreateDNSAccount(c.Request.Context(), unifiedPrincipal(c), in)
		unifiedResult(c, v, err)
	})
	r.PATCH("/dns/accounts/:id", func(c *gin.Context) {
		id, ok := unifiedID(c)
		if !ok {
			return
		}
		var in unified.DNSAccount
		if c.ShouldBindJSON(&in) != nil {
			unifiedResult(c, nil, embedded.ErrInvalid)
			return
		}
		v, err := singleton.Unified.UpdateDNSAccount(c.Request.Context(), unifiedPrincipal(c), id, in)
		unifiedResult(c, v, err)
	})
	r.POST("/dns/accounts/:id/refresh", func(c *gin.Context) {
		id, ok := unifiedID(c)
		if !ok {
			return
		}
		v, err := singleton.Unified.RefreshDNSZones(c.Request.Context(), unifiedPrincipal(c), id)
		unifiedResult(c, v, err)
	})
	r.GET("/dns/zones", func(c *gin.Context) {
		v, err := singleton.Unified.DNSZones(c.Request.Context(), unifiedPrincipal(c))
		unifiedResult(c, v, err)
	})
	r.GET("/dns/zones/:id/records", func(c *gin.Context) {
		id, ok := unifiedID(c)
		if !ok {
			return
		}
		page, e1 := strconv.Atoi(c.DefaultQuery("page", "1"))
		size, e2 := strconv.Atoi(c.DefaultQuery("per_page", "50"))
		if e1 != nil || e2 != nil {
			unifiedResult(c, nil, embedded.ErrInvalid)
			return
		}
		v, err := singleton.Unified.DNSRecords(c.Request.Context(), unifiedPrincipal(c), id, page, size)
		unifiedResult(c, v, err)
	})
	r.POST("/dns/zones/:id/changes", func(c *gin.Context) {
		id, ok := unifiedID(c)
		if !ok {
			return
		}
		var in unified.DNSChange
		if c.ShouldBindJSON(&in) != nil {
			unifiedResult(c, nil, embedded.ErrInvalid)
			return
		}
		v, err := singleton.Unified.ChangeDNS(c.Request.Context(), unifiedPrincipal(c), id, in)
		unifiedResult(c, v, err)
	})
	r.GET("/dns/policies", func(c *gin.Context) {
		v, err := singleton.Unified.DNSPolicies(c.Request.Context(), unifiedPrincipal(c))
		unifiedResult(c, v, err)
	})
	r.POST("/dns/policies", func(c *gin.Context) {
		var in unified.DNSPolicy
		if c.ShouldBindJSON(&in) != nil {
			unifiedResult(c, nil, embedded.ErrInvalid)
			return
		}
		v, err := singleton.Unified.CreateDNSPolicy(c.Request.Context(), unifiedPrincipal(c), in)
		unifiedResult(c, v, err)
	})
	r.DELETE("/dns/policies/:id", func(c *gin.Context) {
		id, ok := unifiedID(c)
		if !ok {
			return
		}
		err := singleton.Unified.RevokeDNSPolicy(c.Request.Context(), unifiedPrincipal(c), id)
		unifiedResult(c, gin.H{"revoked": err == nil}, err)
	})
	r.GET("/dns/operations", func(c *gin.Context) {
		v, err := singleton.Unified.DNSOperations(c.Request.Context(), unifiedPrincipal(c), 100)
		unifiedResult(c, v, err)
	})
	r.GET("/dns/zones/:id/options", func(c *gin.Context) {
		id, ok := unifiedID(c)
		if !ok {
			return
		}
		v, e := singleton.Unified.DNSZoneOptions(c.Request.Context(), unifiedPrincipal(c), id)
		unifiedResult(c, v, e)
	})
	r.GET("/dns/zones/:id/export", func(c *gin.Context) {
		id, ok := unifiedID(c)
		if !ok {
			return
		}
		v, e := singleton.Unified.ExportDNS(c.Request.Context(), unifiedPrincipal(c), id)
		unifiedResult(c, v, e)
	})
	r.POST("/dns/zones/:id/batch", func(c *gin.Context) {
		id, ok := unifiedID(c)
		if !ok {
			return
		}
		var in struct {
			Changes []unified.DNSChange `json:"changes"`
		}
		if c.ShouldBindJSON(&in) != nil {
			unifiedResult(c, nil, embedded.ErrInvalid)
			return
		}
		v, e := singleton.Unified.BatchDNS(c.Request.Context(), unifiedPrincipal(c), id, in.Changes)
		unifiedResult(c, v, e)
	})
	r.POST("/dns/zones/:id/import", func(c *gin.Context) {
		id, ok := unifiedID(c)
		if !ok {
			return
		}
		var in struct {
			OperationKey string               `json:"operation_key"`
			Records      []dnsprovider.Record `json:"records"`
		}
		if c.ShouldBindJSON(&in) != nil {
			unifiedResult(c, nil, embedded.ErrInvalid)
			return
		}
		v, e := singleton.Unified.ImportDNS(c.Request.Context(), unifiedPrincipal(c), id, in.OperationKey, in.Records)
		unifiedResult(c, v, e)
	})
	r.GET("/ddns/bindings", func(c *gin.Context) {
		v, e := singleton.Unified.DDNSBindings(c.Request.Context(), unifiedPrincipal(c))
		unifiedResult(c, v, e)
	})
	r.POST("/ddns/bindings", func(c *gin.Context) {
		var in unified.DDNSBinding
		if c.ShouldBindJSON(&in) != nil {
			unifiedResult(c, nil, embedded.ErrInvalid)
			return
		}
		v, e := singleton.Unified.SaveDDNSBinding(c.Request.Context(), unifiedPrincipal(c), 0, in)
		unifiedResult(c, v, e)
	})
	r.PATCH("/ddns/bindings/:id", func(c *gin.Context) {
		id, ok := unifiedID(c)
		if !ok {
			return
		}
		var in unified.DDNSBinding
		if c.ShouldBindJSON(&in) != nil {
			unifiedResult(c, nil, embedded.ErrInvalid)
			return
		}
		v, e := singleton.Unified.SaveDDNSBinding(c.Request.Context(), unifiedPrincipal(c), id, in)
		unifiedResult(c, v, e)
	})
	r.DELETE("/ddns/bindings/:id", func(c *gin.Context) {
		id, ok := unifiedID(c)
		if !ok {
			return
		}
		e := singleton.Unified.DeleteDDNSBinding(c.Request.Context(), unifiedPrincipal(c), id)
		unifiedResult(c, gin.H{"deleted": e == nil}, e)
	})
	r.POST("/ddns/bindings/:id/run", func(c *gin.Context) {
		id, ok := unifiedID(c)
		if !ok {
			return
		}
		v, e := singleton.Unified.RunDDNS(c.Request.Context(), unifiedPrincipal(c), id)
		unifiedResult(c, v, e)
	})
	r.GET("/ddns/hosts/:id/addresses", func(c *gin.Context) {
		id, ok := unifiedID(c)
		if !ok {
			return
		}
		v, e := singleton.Unified.DDNSHostAddresses(c.Request.Context(), unifiedPrincipal(c), id)
		unifiedResult(c, v, e)
	})
	r.GET("/ddns/devices", func(c *gin.Context) {
		v, e := singleton.Unified.DDNSDevices(c.Request.Context(), unifiedPrincipal(c))
		unifiedResult(c, v, e)
	})
	r.POST("/ddns/devices", func(c *gin.Context) {
		var in struct {
			BindingID uint64 `json:"binding_id"`
		}
		if c.ShouldBindJSON(&in) != nil {
			unifiedResult(c, nil, embedded.ErrInvalid)
			return
		}
		d, token, e := singleton.Unified.CreateDDNSDevice(c.Request.Context(), unifiedPrincipal(c), in.BindingID)
		unifiedResult(c, gin.H{"device": d, "token": token}, e)
	})
	r.DELETE("/ddns/devices/:id", func(c *gin.Context) {
		id, ok := unifiedID(c)
		if !ok {
			return
		}
		e := singleton.Unified.RevokeDDNSDevice(c.Request.Context(), unifiedPrincipal(c), id)
		unifiedResult(c, gin.H{"revoked": e == nil}, e)
	})

}

// This exact route authenticates a scoped device token instead of a user session.
func bindUnifiedDDNSDeviceRoute(r *gin.RouterGroup) {
	r.POST("/unified/ddns/device/update", func(c *gin.Context) {
		if singleton.Unified == nil {
			unifiedResult(c, nil, unified.ErrDDNSUnavailable)
			return
		}
		var in struct {
			IP string `json:"ip"`
		}
		auth := c.GetHeader("Authorization")
		if !strings.HasPrefix(auth, "Bearer ") || c.ShouldBindJSON(&in) != nil {
			unifiedResult(c, nil, unified.ErrDenied)
			return
		}
		b, e := singleton.Unified.ReportDDNSDevice(c.Request.Context(), strings.TrimPrefix(auth, "Bearer "), in.IP)
		// Device callers never receive account, host or credential metadata.
		unifiedResult(c, gin.H{"state": b.State, "last_success": b.LastSuccess, "error": b.Error}, e)
	})
}
