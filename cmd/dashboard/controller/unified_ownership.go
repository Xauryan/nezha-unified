package controller

import (
	"github.com/certimate-go/certimate/pkg/embedded"
	"github.com/gin-gonic/gin"
	"github.com/nezhahq/nezha/service/singleton"
	"github.com/nezhahq/nezha/service/unified"
)

func bindUnifiedOwnershipRoutes(r *gin.RouterGroup) {
	r.GET("/ownership", func(c *gin.Context) {
		rows, err := singleton.Unified.Ownership(c.Request.Context(), unifiedPrincipal(c))
		unifiedResult(c, rows, err)
	})
	r.POST("/ownership/transfer", func(c *gin.Context) {
		var in struct {
			Kind            string `json:"kind"`
			ResourceID      string `json:"resource_id"`
			OwnerID         uint64 `json:"owner_id"`
			ExpectedVersion uint64 `json:"expected_version"`
		}
		if c.ShouldBindJSON(&in) != nil || in.Kind == "" || in.ResourceID == "" || in.OwnerID == 0 || in.ExpectedVersion == 0 {
			unifiedResult(c, nil, embedded.ErrInvalid)
			return
		}
		p, err := singleton.Unified.CurrentPrincipal(c.Request.Context(), unifiedPrincipal(c))
		if err != nil {
			unifiedResult(c, nil, err)
			return
		}
		row, err := singleton.Unified.TransferOwnership(c.Request.Context(), p, in.Kind, in.ResourceID, in.OwnerID, in.ExpectedVersion)
		unifiedResult(c, row, err)
	})
	r.GET("/workflow-policies", func(c *gin.Context) {
		rows, err := singleton.Unified.WorkflowPolicies(c.Request.Context(), unifiedPrincipal(c))
		unifiedResult(c, rows, err)
	})
	r.PUT("/workflow-policies/:id", func(c *gin.Context) {
		var in struct {
			unified.WorkflowPolicy
			ExpectedVersion *uint64 `json:"expected_version"`
		}
		if c.ShouldBindJSON(&in) != nil || in.ExpectedVersion == nil || (in.WorkflowID != "" && in.WorkflowID != c.Param("id")) {
			unifiedResult(c, nil, embedded.ErrInvalid)
			return
		}
		p, err := singleton.Unified.CurrentPrincipal(c.Request.Context(), unifiedPrincipal(c))
		if err != nil {
			unifiedResult(c, nil, err)
			return
		}
		row, err := singleton.Unified.SaveWorkflowPolicy(c.Request.Context(), p, c.Param("id"), *in.ExpectedVersion, in.WorkflowPolicy)
		unifiedResult(c, row, err)
	})
}
