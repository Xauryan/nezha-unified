package controller

import (
	"github.com/certimate-go/certimate/pkg/embedded"
	"github.com/gin-gonic/gin"
	"github.com/nezhahq/nezha/service/singleton"
	"github.com/nezhahq/nezha/service/unified"
)

func bindUnifiedTargetRoutes(r *gin.RouterGroup) {
	r.POST("/deployments/:id/retire", func(c *gin.Context) {
		id, ok := unifiedID(c)
		if !ok {
			return
		}
		var in struct {
			Acknowledge bool   `json:"acknowledge_unknown"`
			Fingerprint string `json:"expected_fingerprint"`
		}
		if c.ShouldBindJSON(&in) != nil {
			unifiedResult(c, nil, embedded.ErrInvalid)
			return
		}
		v, e := singleton.Unified.RetireDeployment(c.Request.Context(), unifiedPrincipal(c), id, in.Acknowledge, in.Fingerprint)
		unifiedResult(c, v, e)
	})
	r.GET("/targets", func(c *gin.Context) {
		v, e := singleton.Unified.Targets(c.Request.Context(), unifiedPrincipal(c))
		unifiedResult(c, v, e)
	})
	r.POST("/targets", func(c *gin.Context) {
		var in unified.DeploymentTarget
		if c.ShouldBindJSON(&in) != nil {
			unifiedResult(c, nil, embedded.ErrInvalid)
			return
		}
		v, e := singleton.Unified.SaveTarget(c.Request.Context(), unifiedPrincipal(c), 0, in)
		unifiedResult(c, v, e)
	})
	r.GET("/targets/:id", func(c *gin.Context) {
		id, ok := unifiedID(c)
		if !ok {
			return
		}
		v, e := singleton.Unified.Target(c.Request.Context(), unifiedPrincipal(c), id)
		unifiedResult(c, v, e)
	})
	r.PATCH("/targets/:id", func(c *gin.Context) {
		id, ok := unifiedID(c)
		if !ok {
			return
		}
		var in unified.DeploymentTarget
		if c.ShouldBindJSON(&in) != nil {
			unifiedResult(c, nil, embedded.ErrInvalid)
			return
		}
		v, e := singleton.Unified.SaveTarget(c.Request.Context(), unifiedPrincipal(c), id, in)
		unifiedResult(c, v, e)
	})
	r.DELETE("/targets/:id", func(c *gin.Context) {
		id, ok := unifiedID(c)
		if !ok {
			return
		}
		e := singleton.Unified.DeleteTarget(c.Request.Context(), unifiedPrincipal(c), id)
		unifiedResult(c, gin.H{"deleted": e == nil}, e)
	})
	r.GET("/targets/:id/capabilities", func(c *gin.Context) {
		id, ok := unifiedID(c)
		if !ok {
			return
		}
		v, e := singleton.Unified.TargetCapabilities(c.Request.Context(), unifiedPrincipal(c), id)
		unifiedResult(c, v, e)
	})
	r.POST("/targets/:id/deployments", func(c *gin.Context) {
		id, ok := unifiedID(c)
		if !ok {
			return
		}
		var in unified.DeploymentInput
		if c.ShouldBindJSON(&in) != nil {
			unifiedResult(c, nil, embedded.ErrInvalid)
			return
		}
		v, e := singleton.Unified.EnqueueDeployment(c.Request.Context(), unifiedPrincipal(c), id, in)
		unifiedResult(c, v, e)
	})
	r.GET("/deployments", func(c *gin.Context) {
		v, e := singleton.Unified.Deployments(c.Request.Context(), unifiedPrincipal(c))
		unifiedResult(c, v, e)
	})
	r.GET("/deployments/:id", func(c *gin.Context) {
		id, ok := unifiedID(c)
		if !ok {
			return
		}
		v, e := singleton.Unified.Deployment(c.Request.Context(), unifiedPrincipal(c), id)
		unifiedResult(c, v, e)
	})
	r.POST("/deployments/:id/cancel", func(c *gin.Context) {
		id, ok := unifiedID(c)
		if !ok {
			return
		}
		v, e := singleton.Unified.CancelDeployment(c.Request.Context(), unifiedPrincipal(c), id)
		unifiedResult(c, v, e)
	})
	r.POST("/deployments/:id/reconcile", func(c *gin.Context) {
		id, ok := unifiedID(c)
		if !ok {
			return
		}
		v, e := singleton.Unified.ReconcileDeployment(c.Request.Context(), unifiedPrincipal(c), id)
		unifiedResult(c, v, e)
	})
}
