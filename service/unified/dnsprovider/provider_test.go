package dnsprovider

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"
)

type roundTrip func(*http.Request) (*http.Response, error)

func (f roundTrip) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }
func response(body string) *http.Response {
	return &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(body)), Header: make(http.Header)}
}
func TestCloudflarePaginationAndFieldPreservation(t *testing.T) {
	calls := 0
	p, err := NewWithClient("cloudflare", map[string]any{"apiToken": "synthetic"}, &http.Client{Transport: roundTrip(func(r *http.Request) (*http.Response, error) {
		calls++
		if r.Host != "api.cloudflare.com" || r.Header.Get("Authorization") != "Bearer synthetic" {
			t.Fatal("wrong endpoint or auth")
		}
		switch r.Method {
		case "GET":
			if r.URL.Query().Get("page") == "1" {
				return response(`{"success":true,"result":[{"id":"r1","name":"www.example.com","type":"A","content":"192.0.2.1","ttl":1,"proxied":true,"comment":"keep","tags":["owner:test"],"settings":{"ipv4_only":true}}],"result_info":{"total_pages":2}}`), nil
			}
			return response(`{"success":true,"result":[],"result_info":{"total_pages":2}}`), nil
		case "PATCH":
			var body map[string]any
			json.NewDecoder(r.Body).Decode(&body)
			if body["proxied"] != true || body["comment"] != "keep" || body["settings"].(map[string]any)["ipv4_only"] != true {
				t.Fatal("lost provider fields")
			}
			return response(`{"success":true,"result":{"id":"r1","name":"www.example.com","type":"A","content":"192.0.2.2","ttl":1,"proxied":true}}`), nil
		}
		t.Fatalf("unexpected method %s", r.Method)
		return nil, nil
	})})
	if err != nil {
		t.Fatal(err)
	}
	records, err := p.Records(context.Background(), Zone{ID: "z", Name: "example.com"})
	if err != nil || len(records) != 1 || calls != 2 {
		t.Fatalf("pagination: %v %d", err, calls)
	}
	records[0].Value = "192.0.2.2"
	if _, err := p.Update(context.Background(), Zone{ID: "z"}, records[0]); err != nil {
		t.Fatal(err)
	}
}
func TestTencentAuthenticationFormatsAndLargeRecordIDs(t *testing.T) {
	if _, err := New("dnspod", map[string]any{"token": "legacy"}); err == nil {
		t.Fatal("mixed legacy and Tencent formats")
	}
	p, err := NewWithClient("dnspod", map[string]any{"secretId": "synthetic-id", "secretKey": "synthetic-key"}, &http.Client{Transport: roundTrip(func(r *http.Request) (*http.Response, error) {
		if !strings.HasPrefix(r.Header.Get("Authorization"), "TC3-HMAC-SHA256 Credential=synthetic-id/") || r.Header.Get("X-TC-Version") != "2021-03-23" {
			t.Fatal("invalid signing headers")
		}
		body, _ := io.ReadAll(r.Body)
		if !strings.Contains(string(body), `"RecordId":9007199254740993`) {
			t.Fatalf("ID rounded: %s", body)
		}
		return response(`{"Response":{"RecordInfo":{"Id":9007199254740993,"SubDomain":"www","RecordType":"A","Value":"192.0.2.1","TTL":600,"Enabled":1,"RecordLineId":"0","Remark":"retained"}}}`), nil
	})})
	if err != nil {
		t.Fatal(err)
	}
	r, err := p.Get(context.Background(), Zone{Name: "example.com"}, "9007199254740993")
	if err != nil || r.ID != "9007199254740993" || r.Name != "www.example.com" || r.Comment != "retained" {
		t.Fatalf("record projection: %#v %v", r, err)
	}
}
func TestEdgeOneNSModeAndSRVTargetConfusion(t *testing.T) {
	p, _ := New("edgeone", map[string]any{"secretId": "synthetic", "secretKey": "synthetic"})
	if _, err := p.Records(context.Background(), Zone{ID: "z", Mode: "partial"}); err == nil {
		t.Fatal("non-NS EdgeOne zone allowed")
	}
	if _, err := cfPayload(Record{Name: "_sip._tcp.allowed.example", Type: "SRV", Data: map[string]any{"service": "_sip", "proto": "_tcp", "name": "other.example"}}); err == nil {
		t.Fatal("SRV data changed authorized name")
	}
}

func TestTencentEntitledLinesAndEdgeOneNameserverWarnings(t *testing.T) {
	c, err := NewWithClient("dnspod", map[string]any{"secretId": "synthetic", "secretKey": "synthetic"}, &http.Client{Transport: roundTrip(func(r *http.Request) (*http.Response, error) {
		if r.Header.Get("X-TC-Action") != "DescribeRecordLineCategoryList" {
			t.Fatal("wrong line API")
		}
		return response(`{"Response":{"LineList":[{"LineId":"0","LineName":"默认","Useful":true},{"LineName":"地域","SubGroup":[{"Lines":[{"LineId":"10","LineName":"允许线路","Useful":true},{"LineId":"11","LineName":"未购买","Useful":false}]}]}]}}`), nil
	})})
	if err != nil {
		t.Fatal(err)
	}
	lines, err := c.(LineProvider).Lines(context.Background(), Zone{Name: "example.com"})
	if err != nil || len(lines) != 2 || lines[0].ID != "0" || lines[1].ID != "10" {
		t.Fatalf("entitled line tree %#v %v", lines, err)
	}
	e, err := NewWithClient("edgeone", map[string]any{"secretId": "synthetic", "secretKey": "synthetic"}, &http.Client{Transport: roundTrip(func(r *http.Request) (*http.Response, error) {
		if r.Header.Get("X-TC-Action") != "DescribeZones" {
			t.Fatal("wrong zone API")
		}
		return response(`{"Response":{"Zones":[{"ZoneId":"full","ZoneName":"example.com","Type":"full","Status":"pending","NameServers":["ns1.example.net"]},{"ZoneId":"partial","ZoneName":"other.example.com","Type":"partial","Status":"active"},{"ZoneId":"no-domain","ZoneName":"","Type":"partialComposite"}],"TotalCount":3}}`), nil
	})})
	if err != nil {
		t.Fatal(err)
	}
	zones, err := e.Zones(context.Background())
	if err != nil || len(zones) != 2 || zones[0].Status != "pending" || len(zones[0].NameServers) != 1 || zones[1].Mode != "partial" {
		t.Fatalf("mode and nameserver metadata %#v %v", zones, err)
	}
}
