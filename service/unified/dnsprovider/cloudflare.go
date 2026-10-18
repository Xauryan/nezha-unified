package dnsprovider

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type cloudflare struct{ *client }
type cfRecord struct {
	ID       string         `json:"id"`
	Name     string         `json:"name"`
	Type     string         `json:"type"`
	Content  string         `json:"content"`
	TTL      int            `json:"ttl"`
	Priority int            `json:"priority"`
	Proxied  *bool          `json:"proxied"`
	Comment  string         `json:"comment"`
	Data     map[string]any `json:"data"`
	Settings map[string]any `json:"settings"`
	Tags     []string       `json:"tags"`
}

func cfRecordValue(r cfRecord) Record {
	return Record{ID: r.ID, Name: r.Name, Type: r.Type, Value: r.Content, TTL: r.TTL, Priority: r.Priority, Proxied: r.Proxied, Comment: r.Comment, Data: r.Data, Settings: r.Settings, Tags: r.Tags}
}

type cfEnvelope[T any] struct {
	Success bool `json:"success"`
	Result  T    `json:"result"`
	Errors  []struct {
		Code int `json:"code"`
	} `json:"errors"`
	Info struct {
		TotalPages int `json:"total_pages"`
	} `json:"result_info"`
}

func cfCall[T any](c *cloudflare, ctx context.Context, method, path string, payload any) (cfEnvelope[T], error) {
	var out cfEnvelope[T]
	token := c.secret("apiToken")
	if strings.HasPrefix(path, "/zones?") && c.secret("apiTokenForZone") != "" {
		token = c.secret("apiTokenForZone")
	}
	if err := c.request(ctx, method, "https://api.cloudflare.com/client/v4"+path, payload, map[string]string{"Authorization": "Bearer " + token}, &out); err != nil {
		return out, err
	}
	if !out.Success {
		code := "rejected"
		if len(out.Errors) > 0 {
			code = strconv.Itoa(out.Errors[0].Code)
		}
		return out, &APIError{Code: code}
	}
	return out, nil
}
func (*cloudflare) Capabilities() Capabilities {
	return Capabilities{Types: []string{"A", "AAAA", "CAA", "CERT", "CNAME", "DNSKEY", "DS", "HTTPS", "LOC", "MX", "NAPTR", "NS", "OPENPGPKEY", "PTR", "SMIMEA", "SRV", "SSHFP", "SVCB", "TLSA", "TXT", "URI"}, Fields: []string{"proxied", "comment", "priority", "data", "settings", "tags"}, MinTTL: 60, AutoTTL: true}
}
func (c *cloudflare) Zones(ctx context.Context) ([]Zone, error) {
	result := []Zone{}
	for page := 1; page <= 10000; page++ {
		out, err := cfCall[[]Zone](c, ctx, "GET", fmt.Sprintf("/zones?per_page=50&page=%d", page), nil)
		if err != nil {
			return nil, err
		}
		result = append(result, out.Result...)
		if page >= out.Info.TotalPages {
			return result, nil
		}
	}
	return nil, errors.New("Cloudflare zone pagination limit exceeded")
}
func (c *cloudflare) Records(ctx context.Context, z Zone) ([]Record, error) {
	result := []Record{}
	for page := 1; page <= 10000; page++ {
		out, err := cfCall[[]cfRecord](c, ctx, "GET", fmt.Sprintf("/zones/%s/dns_records?per_page=100&page=%d", pathID(z.ID), page), nil)
		if err != nil {
			return nil, err
		}
		for _, r := range out.Result {
			result = append(result, cfRecordValue(r))
		}
		if page >= out.Info.TotalPages {
			return result, nil
		}
	}
	return nil, errors.New("Cloudflare record pagination limit exceeded")
}
func (c *cloudflare) Get(ctx context.Context, z Zone, id string) (Record, error) {
	out, err := cfCall[cfRecord](c, ctx, "GET", "/zones/"+pathID(z.ID)+"/dns_records/"+pathID(id), nil)
	return cfRecordValue(out.Result), err
}
func cfPayload(r Record) (map[string]any, error) {
	if r.Line != "" || r.Weight != nil || r.Enabled != nil {
		return nil, errors.New("Cloudflare does not support DNSPod line/weight/enabled fields")
	}
	if r.Type == "SRV" && r.Data != nil {
		name, _ := r.Data["name"].(string)
		service, _ := r.Data["service"].(string)
		proto, _ := r.Data["proto"].(string)
		if name != "" && !strings.EqualFold(strings.TrimSuffix(service+"."+proto+"."+name, "."), strings.TrimSuffix(r.Name, ".")) {
			return nil, errors.New("SRV data name does not match authorized record name")
		}
	}
	value := map[string]any{"name": r.Name, "type": r.Type, "ttl": r.TTL, "comment": r.Comment}
	if r.Tags != nil {
		value["tags"] = r.Tags
	}
	if r.Settings != nil {
		value["settings"] = r.Settings
	}
	if r.Data != nil {
		value["data"] = r.Data
	} else {
		value["content"] = r.Value
	}
	if r.Proxied != nil {
		value["proxied"] = *r.Proxied
	}
	if r.Type == "MX" || r.Type == "URI" {
		value["priority"] = r.Priority
	}
	return value, nil
}
func (c *cloudflare) Create(ctx context.Context, z Zone, r Record) (Record, error) {
	body, err := cfPayload(r)
	if err != nil {
		return Record{}, err
	}
	out, err := cfCall[cfRecord](c, ctx, "POST", "/zones/"+url.PathEscape(z.ID)+"/dns_records", body)
	return cfRecordValue(out.Result), err
}
func (c *cloudflare) Update(ctx context.Context, z Zone, r Record) (Record, error) {
	body, err := cfPayload(r)
	if err != nil {
		return Record{}, err
	}
	out, err := cfCall[cfRecord](c, ctx, "PATCH", "/zones/"+pathID(z.ID)+"/dns_records/"+pathID(r.ID), body)
	return cfRecordValue(out.Result), err
}
func (c *cloudflare) Delete(ctx context.Context, z Zone, id string) error {
	_, err := cfCall[map[string]any](c, ctx, "DELETE", "/zones/"+pathID(z.ID)+"/dns_records/"+pathID(id), nil)
	return err
}
