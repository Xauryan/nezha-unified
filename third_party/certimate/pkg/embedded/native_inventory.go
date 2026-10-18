package embedded

import (
	"context"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/certimate-go/certimate/internal/embedcrypto"
	"github.com/pocketbase/pocketbase/core"
)

func (e *Engine) NativeStatistics(ctx context.Context, certificateScope, workflowScope NativeScope) (map[string]int, error) {
	warningDays := 21
	e.mu.RLock()
	if err := e.check(ctx); err != nil {
		e.mu.RUnlock()
		return nil, err
	}
	if record, err := e.app.FindFirstRecordByFilter("settings", "name='persistence'"); err == nil {
		var content struct {
			WarningDays int `json:"certificatesWarningDaysBeforeExpire"`
		}
		if err := embedcrypto.ReadJSON(e.app, record, "content", &content); err != nil {
			e.mu.RUnlock()
			return nil, err
		}
		if content.WarningDays > 0 && content.WarningDays < 36500 {
			warningDays = content.WarningDays
		}
	}
	e.mu.RUnlock()
	now := time.Now().UTC()
	queries := []struct {
		key, collection, filter string
		scope                   NativeScope
	}{
		{"certificateTotal", "certificate", "", certificateScope},
		{"certificateExpiringSoon", "certificate", fmt.Sprintf("validityNotAfter>'%s' && validityNotAfter<='%s' && isRevoked=false", now.Format("2006-01-02 15:04:05"), now.Add(time.Duration(warningDays)*24*time.Hour).Format("2006-01-02 15:04:05")), certificateScope},
		{"certificateExpired", "certificate", fmt.Sprintf("validityNotAfter<='%s'", now.Format("2006-01-02 15:04:05")), certificateScope},
		{"workflowTotal", "workflow", "", workflowScope},
		{"workflowEnabled", "workflow", "enabled=true", workflowScope},
	}
	result := map[string]int{}
	for _, q := range queries {
		page, err := e.NativeList(ctx, q.collection, q.scope, NativeOptions{Page: 1, PerPage: 1, Filter: q.filter})
		if err != nil {
			return nil, err
		}
		result[q.key] = page.TotalItems
	}
	result["workflowDisabled"] = result["workflowTotal"] - result["workflowEnabled"]
	return result, nil
}

func (e *Engine) NativeCronJobs(ctx context.Context) ([]map[string]string, error) {
	e.mu.RLock()
	defer e.mu.RUnlock()
	if err := e.check(ctx); err != nil {
		return nil, err
	}
	result := []map[string]string{}
	for _, job := range e.app.Cron().Jobs() {
		if !strings.HasPrefix(job.Id(), "__pb") {
			result = append(result, map[string]string{"id": job.Id(), "expression": job.Expression()})
		}
	}
	return result, nil
}

// System logs expose the structured native audit ledger. Provider log payloads
// may contain credentials; they are never forwarded by this admin endpoint.
func (e *Engine) NativeSystemLogs(ctx context.Context, page, perPage int) (NativePage, error) {
	e.mu.RLock()
	defer e.mu.RUnlock()
	result := NativePage{Page: page, PerPage: perPage, Items: []map[string]any{}}
	if err := e.check(ctx); err != nil {
		return result, err
	}
	if page < 1 || page > 100000 || perPage < 1 || perPage > 200 {
		return result, ErrInvalid
	}
	query := e.app.RecordQuery("unified_audit").WithContext(ctx)
	count := *query
	if err := count.Select("COUNT(*)").Row(&result.TotalItems); err != nil {
		return result, err
	}
	var records []*core.Record
	if err := query.OrderBy("created DESC", "id DESC").Limit(int64(perPage)).Offset(int64((page - 1) * perPage)).All(&records); err != nil {
		return result, err
	}
	for _, r := range records {
		result.Items = append(result.Items, map[string]any{"id": r.Id, "created": r.Get("created"), "level": 0, "message": r.GetString("action"), "data": map[string]any{"actor": r.GetString("actor"), "resource": r.GetString("resource")}})
	}
	result.TotalPages = (result.TotalItems + perPage - 1) / perPage
	return result, nil
}

// CertificateMaterial is a server-only transfer value. Serialization is
// intentionally disabled; callers must build and authorize a deploy envelope.
type CertificateMaterial struct {
	CertificatePEM string   `json:"-"`
	PrivateKeyPEM  string   `json:"-"`
	Domains        []string `json:"-"`
	Fingerprint    string   `json:"-"`
	Version        string   `json:"-"`
}

func (e *Engine) CertificateMaterial(ctx context.Context, id string) (CertificateMaterial, error) {
	e.mu.RLock()
	defer e.mu.RUnlock()
	var result CertificateMaterial
	if err := e.check(ctx); err != nil {
		return result, err
	}
	r, err := e.nativeRecord("certificate", id)
	if err != nil {
		return result, err
	}
	if r.GetBool("isRevoked") {
		return result, errors.New("certificate is revoked")
	}
	result.CertificatePEM = r.GetString("certificate")
	block, _ := pem.Decode([]byte(result.CertificatePEM))
	if block == nil || block.Type != "CERTIFICATE" {
		return CertificateMaterial{}, ErrInvalid
	}
	leaf, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return CertificateMaterial{}, ErrInvalid
	}
	key, err := e.unseal("certificate:"+id, r.GetString("unified_private_key"))
	if err != nil {
		return CertificateMaterial{}, err
	}
	result.PrivateKeyPEM = string(key)
	result.Domains = append([]string{}, leaf.DNSNames...)
	for _, ip := range leaf.IPAddresses {
		result.Domains = append(result.Domains, ip.String())
	}
	result.Domains = append(result.Domains, leaf.EmailAddresses...)
	for _, uri := range leaf.URIs {
		result.Domains = append(result.Domains, uri.String())
	}
	if len(result.Domains) == 0 && leaf.Subject.CommonName != "" {
		result.Domains = append(result.Domains, leaf.Subject.CommonName)
	}
	sum := sha256.Sum256(block.Bytes)
	result.Fingerprint = hex.EncodeToString(sum[:])
	result.Version = result.Fingerprint
	return result, nil
}
