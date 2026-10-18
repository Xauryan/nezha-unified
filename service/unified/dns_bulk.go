package unified

import (
	"context"
	"fmt"
	"github.com/certimate-go/certimate/pkg/embedded"
	"github.com/nezhahq/nezha/service/unified/dnsprovider"
)

type DNSOptions struct {
	Capabilities dnsprovider.Capabilities `json:"capabilities"`
	Lines        []dnsprovider.Line       `json:"lines"`
	Warnings     []string                 `json:"warnings"`
}

func (s *Service) DNSZoneOptions(ctx context.Context, p Principal, id uint64) (DNSOptions, error) {
	if !s.RestrictPrincipal(ctx, p, "dns_zone", fmt.Sprint(id), "read") {
		return DNSOptions{}, ErrDenied
	}
	result := DNSOptions{Lines: []dnsprovider.Line{}, Warnings: []string{}}
	z, c, err := s.dnsContext(ctx, p, id)
	if err != nil {
		return result, err
	}
	result.Capabilities = c.Capabilities()
	if c.Capabilities().NSOnly && z.Mode != "full" {
		result.Warnings = append(result.Warnings, "edgeone_requires_ns_mode")
		return result, nil
	}
	if z.Status != "" && z.Status != "active" {
		result.Warnings = append(result.Warnings, "zone_not_active_check_nameservers")
	}
	if lines, ok := c.(dnsprovider.LineProvider); ok {
		result.Lines, err = lines.Lines(ctx, providerZone(z))
	}
	return result, err
}

type DNSBatchResult struct {
	Index     int           `json:"index"`
	Operation *DNSOperation `json:"operation,omitempty"`
	Error     string        `json:"error,omitempty"`
}

func (s *Service) BatchDNS(ctx context.Context, p Principal, id uint64, changes []DNSChange) ([]DNSBatchResult, error) {
	if !validPrincipal(p) {
		return nil, ErrDenied
	}
	if len(changes) == 0 || len(changes) > 500 {
		return nil, embedded.ErrInvalid
	}
	result := make([]DNSBatchResult, 0, len(changes))
	for i, in := range changes {
		if err := ctx.Err(); err != nil {
			return result, err
		}
		op, err := s.ChangeDNS(ctx, p, id, in)
		row := DNSBatchResult{Index: i}
		if err != nil {
			row.Error = "change_rejected"
		} else {
			row.Operation = &op
		}
		result = append(result, row)
	}
	return result, nil
}

type DNSExport struct {
	Version int                  `json:"version"`
	Zone    string               `json:"zone"`
	Records []dnsprovider.Record `json:"records"`
}

func (s *Service) ExportDNS(ctx context.Context, p Principal, id uint64) (DNSExport, error) {
	if !s.RestrictPrincipal(ctx, p, "dns_zone", fmt.Sprint(id), "read") {
		return DNSExport{}, ErrDenied
	}
	z, c, err := s.dnsContext(ctx, p, id)
	if err != nil {
		return DNSExport{}, err
	}
	records, err := c.Records(ctx, providerZone(z))
	if err != nil {
		return DNSExport{}, err
	}
	out := DNSExport{Version: 1, Zone: z.Name, Records: []dnsprovider.Record{}}
	for _, r := range records {
		if s.canDNS(ctx, p, z, r, "read") {
			out.Records = append(out.Records, r)
		}
	}
	return out, nil
}
func (s *Service) ImportDNS(ctx context.Context, p Principal, id uint64, key string, records []dnsprovider.Record) ([]DNSBatchResult, error) {
	if len(key) < 16 || len(key) > 100 || len(records) == 0 || len(records) > 500 {
		return nil, embedded.ErrInvalid
	}
	changes := make([]DNSChange, len(records))
	for i, r := range records {
		r.ID = ""
		changes[i] = DNSChange{OperationKey: fmt.Sprintf("%s:%d", key, i), Action: "create", Record: r}
	}
	return s.BatchDNS(ctx, p, id, changes)
}
