package dnsprovider

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

type tencent struct{ *client }
type scalarID string

func (id *scalarID) UnmarshalJSON(data []byte) error {
	var s string
	if len(data) > 0 && data[0] == '"' {
		if err := json.Unmarshal(data, &s); err != nil {
			return err
		}
	} else {
		var n json.Number
		if err := json.Unmarshal(data, &n); err != nil {
			return err
		}
		s = n.String()
	}
	*id = scalarID(s)
	return nil
}
func (c *tencent) Capabilities() Capabilities {
	if c.kind == "edgeone" {
		return Capabilities{Types: []string{"A", "AAAA", "CNAME", "MX", "TXT", "NS", "CAA", "SRV"}, Fields: []string{"line", "priority", "weight", "enabled"}, MinTTL: 60, NSOnly: true}
	}
	return Capabilities{Types: []string{"A", "AAAA", "CNAME", "MX", "TXT", "NS", "CAA", "SRV", "HTTPS", "SVCB"}, Fields: []string{"line", "priority", "weight", "comment", "enabled"}, MinTTL: 60}
}
func (c *tencent) Zones(ctx context.Context) ([]Zone, error) {
	result := []Zone{}
	for offset := 0; offset < 1000000; offset += 100 {
		if c.kind == "edgeone" {
			var out struct {
				Zones []struct {
					ZoneID      string `json:"ZoneId"`
					ZoneName    string
					Type        string
					Status      string
					NameServers []string
				}
				TotalCount int
			}
			err := c.tc(ctx, "DescribeZones", map[string]any{"Offset": offset, "Limit": 100}, &out)
			if err != nil {
				return nil, err
			}
			for _, z := range out.Zones {
				if z.ZoneName == "" || (z.Type != "full" && z.Type != "partial" && z.Type != "dnsPodAccess") {
					continue
				}
				result = append(result, Zone{ID: z.ZoneID, Name: z.ZoneName, Mode: z.Type, Status: z.Status, NameServers: z.NameServers})
			}
			if offset+len(out.Zones) >= out.TotalCount {
				return result, nil
			}
		} else {
			var out struct {
				DomainList []struct {
					DomainID scalarID `json:"DomainId"`
					Name     string
				}
				DomainCountInfo struct{ DomainTotal int }
			}
			if err := c.tc(ctx, "DescribeDomainList", map[string]any{"Offset": offset, "Limit": 100}, &out); err != nil {
				return nil, err
			}
			for _, z := range out.DomainList {
				result = append(result, Zone{ID: string(z.DomainID), Name: z.Name, Mode: "full"})
			}
			if offset+len(out.DomainList) >= out.DomainCountInfo.DomainTotal {
				return result, nil
			}
		}
	}
	return nil, errors.New("Tencent zone pagination limit exceeded")
}

type tcRecord struct {
	RecordID scalarID `json:"RecordId"`
	Name     string
	Type     string
	Value    string
	Content  string
	TTL      int
	MX       int
	Priority int
	LineID   string `json:"LineId"`
	Location string
	Status   string
	Weight   *int
	Remark   string
}

func (c *tencent) project(z Zone, r tcRecord) Record {
	on := r.Status == "ENABLE"
	value, line, priority := r.Value, r.LineID, r.MX
	if c.kind == "edgeone" {
		on = r.Status == "enable"
		value, line, priority = r.Content, r.Location, r.Priority
	}
	return Record{ID: string(r.RecordID), Name: fqdn(r.Name, z.Name), Type: r.Type, Value: value, TTL: r.TTL, Priority: priority, Line: line, Weight: r.Weight, Comment: r.Remark, Enabled: &on}
}
func (c *tencent) Records(ctx context.Context, z Zone) ([]Record, error) {
	result := []Record{}
	if c.kind == "edgeone" && z.Mode != "full" {
		return nil, errors.New("EdgeOne DNS requires NS/full mode")
	}
	for offset := 0; offset < 1000000; offset += 100 {
		payload := map[string]any{"Offset": offset, "Limit": 100}
		if c.kind == "edgeone" {
			payload["ZoneId"] = z.ID
			var out struct {
				DnsRecords []tcRecord
				TotalCount int
			}
			if err := c.tc(ctx, "DescribeDnsRecords", payload, &out); err != nil {
				return nil, err
			}
			for _, r := range out.DnsRecords {
				result = append(result, c.project(z, r))
			}
			if offset+len(out.DnsRecords) >= out.TotalCount {
				return result, nil
			}
		} else {
			payload["Domain"] = z.Name
			var out struct {
				RecordList      []tcRecord
				RecordCountInfo struct{ TotalCount int }
			}
			if err := c.tc(ctx, "DescribeRecordList", payload, &out); err != nil {
				var apiErr *APIError
				if errors.As(err, &apiErr) && apiErr.Code == "ResourceNotFound.NoDataOfRecord" {
					return result, nil
				}
				return nil, err
			}
			for _, r := range out.RecordList {
				result = append(result, c.project(z, r))
			}
			if offset+len(out.RecordList) >= out.RecordCountInfo.TotalCount {
				return result, nil
			}
		}
	}
	return nil, errors.New("Tencent record pagination limit exceeded")
}
func recordNumber(id string) (json.Number, error) {
	n, err := strconv.ParseUint(id, 10, 64)
	if err != nil || n == 0 {
		return "", errors.New("invalid DNSPod record ID")
	}
	return json.Number(id), nil
}
func (c *tencent) Get(ctx context.Context, z Zone, id string) (Record, error) {
	if c.kind == "edgeone" {
		if z.Mode != "full" {
			return Record{}, errors.New("EdgeOne DNS requires NS/full mode")
		}
		var out struct{ DnsRecords []tcRecord }
		if err := c.tc(ctx, "DescribeDnsRecords", map[string]any{"ZoneId": z.ID, "Filters": []any{map[string]any{"Name": "id", "Values": []string{id}}}}, &out); err != nil {
			return Record{}, err
		}
		if len(out.DnsRecords) != 1 || string(out.DnsRecords[0].RecordID) != id {
			return Record{}, &APIError{Code: "not_found"}
		}
		return c.project(z, out.DnsRecords[0]), nil
	}
	n, err := recordNumber(id)
	if err != nil {
		return Record{}, err
	}
	var out struct {
		RecordInfo struct {
			ID           scalarID `json:"Id"`
			SubDomain    string
			RecordType   string
			Value        string
			RecordLineID string `json:"RecordLineId"`
			TTL          int
			MX           int
			Enabled      int
			Weight       *int
			Remark       string
		}
	}
	if err := c.tc(ctx, "DescribeRecord", map[string]any{"Domain": z.Name, "RecordId": n}, &out); err != nil {
		return Record{}, err
	}
	r := out.RecordInfo
	enabled := r.Enabled == 1
	if string(r.ID) != id {
		return Record{}, &APIError{Code: "identity_mismatch"}
	}
	return Record{ID: id, Name: fqdn(r.SubDomain, z.Name), Type: r.RecordType, Value: r.Value, Line: r.RecordLineID, TTL: r.TTL, Priority: r.MX, Enabled: &enabled, Weight: r.Weight, Comment: r.Remark}, nil
}
func (c *tencent) write(ctx context.Context, z Zone, r Record, create bool) (Record, error) {
	if r.Proxied != nil || r.Data != nil || r.Settings != nil || len(r.Tags) > 0 {
		return Record{}, errors.New("Tencent DNS does not support Cloudflare proxy/data/tags fields")
	}
	payload := map[string]any{"TTL": r.TTL}
	action := ""
	if c.kind == "edgeone" {
		if z.Mode != "full" {
			return Record{}, errors.New("EdgeOne DNS requires NS/full mode")
		}
		if r.Comment != "" {
			return Record{}, errors.New("EdgeOne DNS does not support record comments")
		}
		payload["ZoneId"], payload["Name"], payload["Type"], payload["Content"] = z.ID, r.Name, r.Type, r.Value
		line := r.Line
		if line == "" {
			line = "Default"
		}
		payload["Location"] = line
		payload["Weight"] = -1
		if r.Weight != nil {
			payload["Weight"] = *r.Weight
		}
		if r.Type == "MX" {
			payload["Priority"] = r.Priority
		}
		action = "ModifyDnsRecord"
		payload["DnsRecordId"] = r.ID
		if create {
			action = "CreateDnsRecord"
			delete(payload, "DnsRecordId")
		}
	} else {
		payload["Domain"], payload["SubDomain"], payload["RecordType"], payload["Value"] = z.Name, relative(r.Name, z.Name), r.Type, r.Value
		line := r.Line
		if line == "" {
			line = "0"
		}
		payload["RecordLine"] = "默认"
		payload["RecordLineId"] = line
		if r.Type == "MX" {
			payload["MX"] = r.Priority
		}
		if r.Weight != nil {
			payload["Weight"] = *r.Weight
		}
		action = "ModifyRecord"
		if create {
			action = "CreateRecord"
		} else {
			n, err := recordNumber(r.ID)
			if err != nil {
				return Record{}, err
			}
			payload["RecordId"] = n
		}
	}
	var out struct {
		RecordID scalarID `json:"RecordId"`
	}
	if err := c.tc(ctx, action, payload, &out); err != nil {
		return Record{}, err
	}
	if create {
		r.ID = string(out.RecordID)
		if r.ID == "" {
			return Record{}, &APIError{Code: "missing_created_id", Unknown: true}
		}
	}
	// These are separate upstream operations. Any failure after the first write
	// is ambiguous/partial and must be reconciled, never blindly reissued.
	if r.Enabled != nil {
		var body map[string]any
		var action string
		if c.kind == "edgeone" {
			action = "ModifyDnsRecordsStatus"
			body = map[string]any{"ZoneId": z.ID}
			key := "RecordsToDisable"
			if *r.Enabled {
				key = "RecordsToEnable"
			}
			body[key] = []string{r.ID}
		} else {
			action = "ModifyRecordStatus"
			status := "DISABLE"
			if *r.Enabled {
				status = "ENABLE"
			}
			n, _ := recordNumber(r.ID)
			body = map[string]any{"Domain": z.Name, "RecordId": n, "Status": status}
		}
		if err := c.tc(ctx, action, body, &map[string]any{}); err != nil {
			return r, &APIError{Code: "partial_status_update", Unknown: true}
		}
	}
	if c.kind == "dnspod" {
		n, _ := recordNumber(r.ID)
		if err := c.tc(ctx, "ModifyRecordRemark", map[string]any{"Domain": z.Name, "RecordId": n, "Remark": r.Comment}, &map[string]any{}); err != nil {
			return r, &APIError{Code: "partial_comment_update", Unknown: true}
		}
	}
	actual, err := c.Get(ctx, z, r.ID)
	if err != nil {
		return r, &APIError{Code: "write_readback_failed", Unknown: true}
	}
	return actual, nil
}
func (c *tencent) Create(ctx context.Context, z Zone, r Record) (Record, error) {
	return c.write(ctx, z, r, true)
}
func (c *tencent) Update(ctx context.Context, z Zone, r Record) (Record, error) {
	return c.write(ctx, z, r, false)
}
func (c *tencent) Delete(ctx context.Context, z Zone, id string) error {
	if c.kind == "edgeone" {
		if z.Mode != "full" {
			return errors.New("EdgeOne DNS requires NS/full mode")
		}
		return c.tc(ctx, "DeleteDnsRecords", map[string]any{"ZoneId": z.ID, "RecordIds": []string{id}}, &map[string]any{})
	}
	n, err := recordNumber(id)
	if err != nil {
		return fmt.Errorf("DNSPod record: %w", err)
	}
	return c.tc(ctx, "DeleteRecord", map[string]any{"Domain": z.Name, "RecordId": n}, &map[string]any{})
}

// Lines queries DNSPod's entitlement-aware line tree; unsupported lines are not offered.
func (c *tencent) Lines(ctx context.Context, z Zone) ([]Line, error) {
	if c.kind == "edgeone" {
		return []Line{{ID: "Default", Name: "默认"}}, nil
	}
	type item struct {
		LineID   string `json:"LineId"`
		LineName string
		Useful   bool
		Lines    []item
		SubGroup []item
	}
	var out struct{ LineList []item }
	if err := c.tc(ctx, "DescribeRecordLineCategoryList", map[string]any{"Domain": z.Name}, &out); err != nil {
		return nil, err
	}
	result := []Line{}
	seen := map[string]bool{}
	var appendLines func([]item)
	appendLines = func(items []item) {
		for _, x := range items {
			if x.Useful && x.LineID != "" && !seen[x.LineID] {
				seen[x.LineID] = true
				result = append(result, Line{ID: x.LineID, Name: x.LineName})
			}
			appendLines(x.Lines)
			appendLines(x.SubGroup)
		}
	}
	appendLines(out.LineList)
	return result, nil
}
