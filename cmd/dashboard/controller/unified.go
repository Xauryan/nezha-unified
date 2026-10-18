package controller

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/certimate-go/certimate/pkg/embedded"
	"github.com/gin-gonic/gin"
	"github.com/nezhahq/nezha/model"
	"github.com/nezhahq/nezha/service/singleton"
	"github.com/nezhahq/nezha/service/unified"
)

func unifiedPrincipal(c *gin.Context) unified.Principal {
	if value, ok := c.Get(unifiedScopedPrincipalCtxKey); ok {
		if p, ok := value.(unified.Principal); ok {
			return p
		}
	}
	u, ok := c.Get(model.CtxKeyAuthorizedUser)
	if !ok {
		return unified.Principal{}
	}
	user, ok := u.(*model.User)
	if !ok || user == nil {
		return unified.Principal{}
	}
	_, pat := c.Get(model.CtxKeyAPIToken)
	return unified.Principal{UserID: user.ID, Admin: user.Role.IsAdmin(), PAT: pat}
}

func unifiedResult(c *gin.Context, value any, err error) {
	if err == nil {
		c.JSON(http.StatusOK, model.CommonResponse[any]{Success: true, Data: value})
		return
	}
	status, msg := http.StatusInternalServerError, "unified storage operation failed"
	switch {
	case errors.Is(err, unified.ErrDenied):
		status, msg = http.StatusForbidden, "permission denied"
	case errors.Is(err, embedded.ErrInvalid):
		status, msg = http.StatusBadRequest, "invalid resource or request"
	case errors.Is(err, unified.ErrTargetConflict):
		status, msg = http.StatusConflict, "resource version or operation conflict"
	case errors.Is(err, embedded.ErrNotFound):
		status, msg = http.StatusNotFound, "resource not found"
	default:
		log.Printf("UNIFIED>> storage operation failed (%T)", err)
	}
	c.JSON(status, model.CommonResponse[any]{Success: false, Error: msg})
}

func bindUnifiedRoutes(auth *gin.RouterGroup) {
	// Existing Nezha PAT scopes never imply new-module permission. Explicit
	// resource tokens use a separate bearer-only entry with the same handlers.
	r := auth.Group("/unified", restPATForbiddenMiddleware(), unifiedReadyMiddleware())
	bindUnifiedResources(r)
}
func bindUnifiedScopedRoutes(api *gin.RouterGroup) {
	r := api.Group("/unified-api", unifiedScopedTokenAuthMiddleware(), unifiedReadyMiddleware())
	bindUnifiedResources(r)
}
func unifiedReadyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Cache-Control", "no-store")
		if singleton.Unified == nil {
			c.AbortWithStatusJSON(http.StatusServiceUnavailable, newErrorResponse(errors.New("unified module unavailable")))
			return
		}
		p, err := singleton.Unified.CurrentPrincipal(c.Request.Context(), unifiedPrincipal(c))
		if err != nil {
			unifiedResult(c, nil, err)
			c.Abort()
			return
		}
		c.Set(unifiedScopedPrincipalCtxKey, p)
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, 8<<20)
		c.Next()
	}
}
func bindUnifiedResources(r *gin.RouterGroup) {
	bindUnifiedDNSRoutes(r)
	bindUnifiedCertimateRoutes(r)
	bindUnifiedOwnershipRoutes(r)
	bindUnifiedIdentityRoutes(r)
	bindUnifiedTargetRoutes(r)
	r.GET("/runs", func(c *gin.Context) {
		value, err := singleton.Unified.Runs(c.Request.Context(), unifiedPrincipal(c))
		unifiedResult(c, value, err)
	})
	r.GET("/status", func(c *gin.Context) {
		unifiedResult(c, gin.H{"version": "0.2.0-dev", "certimate": "0.4.35", "storage": "ready", "workflow_execution": embedded.RuntimeAvailable(), "dns_management": true, "agent_deployment": true, "scoped_tokens": true, "service_identities": true}, nil)
	})
	for path, kind := range map[string]string{"/credentials": "credential", "/certificates": "certificate", "/workflows": "workflow"} {
		r.GET(path, func(c *gin.Context) {
			page, e1 := strconv.Atoi(c.DefaultQuery("page", "1"))
			size, e2 := strconv.Atoi(c.DefaultQuery("per_page", "20"))
			if e1 != nil || e2 != nil {
				unifiedResult(c, nil, embedded.ErrInvalid)
				return
			}
			v, err := singleton.Unified.List(c.Request.Context(), unifiedPrincipal(c), kind, page, size)
			unifiedResult(c, v, err)
		})
	}
	r.POST("/credentials", func(c *gin.Context) {
		var in embedded.CredentialInput
		if c.ShouldBindJSON(&in) != nil {
			unifiedResult(c, nil, embedded.ErrInvalid)
			return
		}
		v, err := singleton.Unified.CreateCredential(c.Request.Context(), unifiedPrincipal(c), in)
		unifiedResult(c, v, err)
	})
	r.POST("/certificates/import", func(c *gin.Context) {
		var in embedded.CertificateInput
		if c.ShouldBindJSON(&in) != nil {
			unifiedResult(c, nil, embedded.ErrInvalid)
			return
		}
		v, err := singleton.Unified.ImportCertificate(c.Request.Context(), unifiedPrincipal(c), in)
		unifiedResult(c, v, err)
	})
	r.GET("/certificates/:id/pem", func(c *gin.Context) {
		pem, err := singleton.Unified.CertificatePEM(c.Request.Context(), unifiedPrincipal(c), c.Param("id"))
		if err != nil {
			unifiedResult(c, nil, err)
			return
		}
		c.Header("Content-Disposition", `attachment; filename="certificate.pem"`)
		c.Data(http.StatusOK, "application/x-pem-file", []byte(pem))
	})
	r.GET("/grants", func(c *gin.Context) {
		v, err := singleton.Unified.ListGrants(c.Request.Context(), unifiedPrincipal(c))
		unifiedResult(c, v, err)
	})
	r.POST("/grants", func(c *gin.Context) {
		p := unifiedPrincipal(c)
		if !p.Admin || p.PAT {
			unifiedResult(c, nil, unified.ErrDenied)
			return
		}
		var g unified.Grant
		if c.ShouldBindJSON(&g) != nil {
			unifiedResult(c, nil, embedded.ErrInvalid)
			return
		}
		var count int64
		if err := singleton.DB.Model(&model.User{}).Where("id = ?", g.UserID).Count(&count).Error; err != nil {
			unifiedResult(c, nil, err)
			return
		}
		if count != 1 {
			unifiedResult(c, nil, embedded.ErrInvalid)
			return
		}
		v, err := singleton.Unified.CreateGrant(c.Request.Context(), p, g)
		unifiedResult(c, v, err)
	})
	r.DELETE("/grants/:id", func(c *gin.Context) {
		id, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil || id == 0 {
			unifiedResult(c, nil, embedded.ErrInvalid)
			return
		}
		err = singleton.Unified.RevokeGrant(c.Request.Context(), unifiedPrincipal(c), id)
		unifiedResult(c, gin.H{"revoked": err == nil}, err)
	})
}
