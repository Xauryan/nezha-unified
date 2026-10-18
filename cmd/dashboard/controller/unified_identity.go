package controller

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/certimate-go/certimate/pkg/embedded"
	"github.com/gin-gonic/gin"
	"github.com/nezhahq/nezha/model"
	"github.com/nezhahq/nezha/service/singleton"
	"github.com/nezhahq/nezha/service/unified"
)

const unifiedScopedPrincipalCtxKey = "unified.scoped-principal"

// This is mounted only on the separate unified API group. These credentials
// never enter Nezha's JWT/PAT authorization chain and carry no legacy API power.
func unifiedScopedTokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Cache-Control", "no-store")
		if singleton.Unified == nil {
			c.AbortWithStatusJSON(http.StatusServiceUnavailable, model.CommonResponse[any]{Success: false, Error: "unified module unavailable"})
			return
		}
		header := c.Request.Header.Values("Authorization")
		if len(header) != 1 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, model.CommonResponse[any]{Success: false, Error: "invalid scoped token"})
			return
		}
		parts := strings.Fields(header[0])
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, model.CommonResponse[any]{Success: false, Error: "invalid scoped token"})
			return
		}
		principal, err := singleton.Unified.AuthenticateScopedToken(c.Request.Context(), parts[1])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, model.CommonResponse[any]{Success: false, Error: "invalid scoped token"})
			return
		}
		c.Set(unifiedScopedPrincipalCtxKey, principal)
	}
}

func unifiedIdentityID(c *gin.Context) (uint64, bool) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil || id == 0 {
		unifiedResult(c, nil, embedded.ErrInvalid)
		return 0, false
	}
	return id, true
}

func bindUnifiedIdentityRoutes(r *gin.RouterGroup) {
	r.GET("/tokens", func(c *gin.Context) {
		rows, err := singleton.Unified.ListScopedTokens(c.Request.Context(), unifiedPrincipal(c))
		unifiedResult(c, rows, err)
	})
	r.POST("/tokens", func(c *gin.Context) {
		var in unified.ScopedTokenInput
		if c.ShouldBindJSON(&in) != nil {
			unifiedResult(c, nil, embedded.ErrInvalid)
			return
		}
		row, err := singleton.Unified.CreateScopedToken(c.Request.Context(), unifiedPrincipal(c), in)
		unifiedResult(c, row, err)
	})
	r.DELETE("/tokens/:id", func(c *gin.Context) {
		id, ok := unifiedIdentityID(c)
		if !ok {
			return
		}
		err := singleton.Unified.RevokeScopedToken(c.Request.Context(), unifiedPrincipal(c), id)
		unifiedResult(c, gin.H{"revoked": err == nil}, err)
	})
	r.GET("/identities", func(c *gin.Context) {
		rows, err := singleton.Unified.ListServiceIdentities(c.Request.Context(), unifiedPrincipal(c))
		unifiedResult(c, rows, err)
	})
	r.POST("/identities", func(c *gin.Context) {
		var in unified.ServiceIdentityInput
		if c.ShouldBindJSON(&in) != nil {
			unifiedResult(c, nil, embedded.ErrInvalid)
			return
		}
		row, err := singleton.Unified.CreateServiceIdentity(c.Request.Context(), unifiedPrincipal(c), in)
		unifiedResult(c, row, err)
	})
	r.PATCH("/identities/:id", func(c *gin.Context) {
		id, ok := unifiedIdentityID(c)
		if !ok {
			return
		}
		var in unified.ServiceIdentityUpdate
		if c.ShouldBindJSON(&in) != nil {
			unifiedResult(c, nil, embedded.ErrInvalid)
			return
		}
		row, err := singleton.Unified.UpdateServiceIdentity(c.Request.Context(), unifiedPrincipal(c), id, in)
		unifiedResult(c, row, err)
	})
}
