package controller

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"net/http"
	"strconv"

	"github.com/certimate-go/certimate/pkg/embedded"
	"github.com/gin-gonic/gin"
	"github.com/nezhahq/nezha/service/singleton"
	"github.com/nezhahq/nezha/service/unified"
)

// The native API shares Nezha authentication and response envelopes. It does
// not mount PocketBase HTTP handlers or expose a second login session.
func bindUnifiedCertimateRoutes(r *gin.RouterGroup) {
	r = r.Group("/certimate")
	r.GET("/resources/:collection", func(c *gin.Context) {
		page, perPage, ok := unifiedNativePagination(c)
		if !ok {
			return
		}
		options := embedded.NativeOptions{Page: page, PerPage: perPage, Filter: c.Query("filter"), Sort: c.Query("sort"), Fields: c.Query("fields"), Expand: c.Query("expand")}
		result, err := singleton.Unified.NativeList(c.Request.Context(), unifiedPrincipal(c), c.Param("collection"), options)
		unifiedResult(c, result, err)
	})
	r.GET("/resources/:collection/:id", func(c *gin.Context) {
		result, err := singleton.Unified.NativeGet(c.Request.Context(), unifiedPrincipal(c), c.Param("collection"), c.Param("id"))
		unifiedResult(c, result, err)
	})
	for _, method := range []string{http.MethodPost, http.MethodPatch} {
		path := "/resources/:collection"
		if method == http.MethodPatch {
			path += "/:id"
		}
		r.Handle(method, path, func(c *gin.Context) {
			var body map[string]any
			if c.ShouldBindJSON(&body) != nil || body == nil {
				unifiedResult(c, nil, embedded.ErrInvalid)
				return
			}
			operationID, ok := unifiedNativeOperationID(c)
			if !ok {
				return
			}
			result, err := singleton.Unified.NativeSave(c.Request.Context(), unifiedPrincipal(c), c.Param("collection"), c.Param("id"), operationID, body)
			unifiedResult(c, result, err)
		})
	}
	r.DELETE("/resources/:collection/:id", func(c *gin.Context) {
		err := singleton.Unified.NativeDelete(c.Request.Context(), unifiedPrincipal(c), c.Param("collection"), c.Param("id"))
		unifiedResult(c, err == nil, err)
	})
	r.POST("/batch", unifiedNativeBatch)
	r.POST("/credentials/:id/duplicate", func(c *gin.Context) {
		var in struct {
			Name string `json:"name"`
		}
		if c.ShouldBindJSON(&in) != nil || in.Name == "" {
			unifiedResult(c, nil, embedded.ErrInvalid)
			return
		}
		operationID, ok := unifiedNativeOperationID(c)
		if !ok {
			return
		}
		result, err := singleton.Unified.CloneCredential(c.Request.Context(), unifiedPrincipal(c), c.Param("id"), in.Name, operationID)
		unifiedResult(c, result, err)
	})
	r.POST("/credentials/:id/export", func(c *gin.Context) {
		config, err := singleton.Unified.ExportCredential(c.Request.Context(), unifiedPrincipal(c), c.Param("id"))
		unifiedResult(c, gin.H{"config": config}, err)
	})
	r.GET("/statistics", func(c *gin.Context) {
		result, err := singleton.Unified.Statistics(c.Request.Context(), unifiedPrincipal(c))
		unifiedResult(c, result, err)
	})
	r.GET("/workflows/stats", func(c *gin.Context) {
		result, err := singleton.Unified.WorkflowStats(c.Request.Context(), unifiedPrincipal(c))
		unifiedResult(c, result, err)
	})
	r.POST("/workflows/:id/runs", func(c *gin.Context) {
		var in struct {
			Trigger string `json:"trigger"`
		}
		if c.ShouldBindJSON(&in) != nil || in.Trigger != "manual" {
			unifiedResult(c, nil, embedded.ErrInvalid)
			return
		}
		runID, err := singleton.Unified.StartWorkflow(c.Request.Context(), unifiedPrincipal(c), c.Param("id"))
		unifiedResult(c, gin.H{"runId": runID}, err)
	})
	r.POST("/workflows/:id/runs/:run/cancel", func(c *gin.Context) {
		err := singleton.Unified.CancelWorkflow(c.Request.Context(), unifiedPrincipal(c), c.Param("id"), c.Param("run"))
		unifiedResult(c, gin.H{"cancelled": err == nil}, err)
	})
	r.POST("/certificates/:id/download", func(c *gin.Context) {
		var options embedded.NativeCertificateDownloadOptions
		if c.ShouldBindJSON(&options) != nil {
			unifiedResult(c, nil, embedded.ErrInvalid)
			return
		}
		content, err := singleton.Unified.DownloadCertificate(c.Request.Context(), unifiedPrincipal(c), c.Param("id"), options)
		if err != nil {
			unifiedResult(c, nil, err)
			return
		}
		unifiedResult(c, gin.H{"zipBytes": base64.StdEncoding.EncodeToString(content)}, nil)
	})
	r.POST("/certificates/:id/revoke", func(c *gin.Context) {
		err := singleton.Unified.RevokeCertificate(c.Request.Context(), unifiedPrincipal(c), c.Param("id"))
		unifiedResult(c, gin.H{"revoked": err == nil}, err)
	})
	r.POST("/notifications/test", func(c *gin.Context) {
		var in struct {
			Provider string `json:"provider"`
			AccessID string `json:"accessId"`
		}
		if c.ShouldBindJSON(&in) != nil || in.Provider == "" || in.AccessID == "" {
			unifiedResult(c, nil, embedded.ErrInvalid)
			return
		}
		err := singleton.Unified.TestNotification(c.Request.Context(), unifiedPrincipal(c), in.AccessID, in.Provider)
		unifiedResult(c, gin.H{"sent": err == nil}, err)
	})
	r.GET("/system/crons", func(c *gin.Context) {
		result, err := singleton.Unified.CronJobs(c.Request.Context(), unifiedPrincipal(c))
		unifiedResult(c, result, err)
	})
	r.GET("/system/logs", func(c *gin.Context) {
		page, perPage, ok := unifiedNativePagination(c)
		if !ok {
			return
		}
		result, err := singleton.Unified.SystemLogs(c.Request.Context(), unifiedPrincipal(c), page, perPage)
		unifiedResult(c, result, err)
	})
}

func unifiedNativePagination(c *gin.Context) (int, int, bool) {
	page, errPage := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage, errSize := strconv.Atoi(c.DefaultQuery("perPage", "20"))
	if errPage != nil || errSize != nil || page < 1 || page > 100000 || perPage < 1 || perPage > 200 {
		unifiedResult(c, nil, embedded.ErrInvalid)
		return 0, 0, false
	}
	return page, perPage, true
}

func unifiedNativeOperationID(c *gin.Context) (string, bool) {
	values := c.Request.Header.Values("Idempotency-Key")
	if len(values) > 1 || (len(values) == 1 && (len(values[0]) == 0 || len(values[0]) > 64)) {
		unifiedResult(c, nil, embedded.ErrInvalid)
		return "", false
	}
	if len(values) == 1 {
		for _, ch := range values[0] {
			if (ch < 'a' || ch > 'z') && (ch < 'A' || ch > 'Z') && (ch < '0' || ch > '9') && ch != '-' && ch != '_' {
				unifiedResult(c, nil, embedded.ErrInvalid)
				return "", false
			}
		}
		return values[0], true
	}
	return rand.Text(), true
}

type unifiedNativeBatchOperation struct {
	Collection string         `json:"collection"`
	Method     string         `json:"method"`
	ID         string         `json:"id"`
	Body       map[string]any `json:"body"`
}

// Each item is authorized and committed independently. A failed item never
// implies rollback of an earlier success; the client consumes every status.
func unifiedNativeBatch(c *gin.Context) {
	var in struct {
		Operations []unifiedNativeBatchOperation `json:"operations"`
	}
	if c.ShouldBindJSON(&in) != nil || len(in.Operations) == 0 || len(in.Operations) > 200 {
		unifiedResult(c, nil, embedded.ErrInvalid)
		return
	}
	for _, item := range in.Operations {
		if item.Collection == "" || (item.Method != http.MethodPost && item.Method != http.MethodPatch && item.Method != http.MethodDelete) ||
			(item.Method != http.MethodPost && item.ID == "") || (item.Method == http.MethodPost && item.ID != "") ||
			(item.Method != http.MethodDelete && item.Body == nil) {
			unifiedResult(c, nil, embedded.ErrInvalid)
			return
		}
	}
	operationID, ok := unifiedNativeOperationID(c)
	if !ok {
		return
	}
	result := make([]map[string]any, 0, len(in.Operations))
	for index, item := range in.Operations {
		var err error
		var data any
		if item.Method == http.MethodDelete {
			err = singleton.Unified.NativeDelete(c.Request.Context(), unifiedPrincipal(c), item.Collection, item.ID)
			data = err == nil
		} else {
			// Stable sub-operation keys allow retried batches to recover creates.
			digest := sha256.Sum256([]byte(operationID + ":" + strconv.Itoa(index)))
			data, err = singleton.Unified.NativeSave(c.Request.Context(), unifiedPrincipal(c), item.Collection, item.ID, hex.EncodeToString(digest[:]), item.Body)
		}
		status, message := unifiedNativeBatchError(err)
		if err == nil {
			result = append(result, map[string]any{"status": status, "body": data})
		} else {
			result = append(result, map[string]any{"status": status, "body": map[string]any{"message": message}})
		}
	}
	unifiedResult(c, result, nil)
}

func unifiedNativeBatchError(err error) (int, string) {
	switch {
	case err == nil:
		return http.StatusOK, ""
	case errors.Is(err, unified.ErrDenied):
		return http.StatusForbidden, "permission denied"
	case errors.Is(err, embedded.ErrInvalid):
		return http.StatusBadRequest, "invalid resource or request"
	case errors.Is(err, embedded.ErrNotFound):
		return http.StatusNotFound, "resource not found"
	default:
		return http.StatusInternalServerError, "unified storage operation failed"
	}
}
