// Package dnsprovider implements authoritative DNS management APIs. DNS-01
// challenge providers are kept separately in Certimate's existing engine.
package dnsprovider

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type Zone struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Mode        string   `json:"mode"`
	Status      string   `json:"status,omitempty"`
	NameServers []string `json:"name_servers,omitempty"`
}
type Record struct {
	ID       string         `json:"id,omitempty"`
	Name     string         `json:"name"`
	Type     string         `json:"type"`
	Value    string         `json:"value"`
	TTL      int            `json:"ttl"`
	Priority int            `json:"priority,omitempty"`
	Proxied  *bool          `json:"proxied,omitempty"`
	Line     string         `json:"line,omitempty"`
	Weight   *int           `json:"weight,omitempty"`
	Enabled  *bool          `json:"enabled,omitempty"`
	Comment  string         `json:"comment,omitempty"`
	Data     map[string]any `json:"data,omitempty"`
	Settings map[string]any `json:"settings,omitempty"`
	Tags     []string       `json:"tags,omitempty"`
}
type Capabilities struct {
	Types   []string `json:"types"`
	Fields  []string `json:"fields"`
	MinTTL  int      `json:"min_ttl"`
	NSOnly  bool     `json:"ns_only"`
	AutoTTL bool     `json:"auto_ttl"`
}
type Line struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// LineProvider exposes the account-specific available routing lines.
type LineProvider interface {
	Lines(context.Context, Zone) ([]Line, error)
}

type APIError struct {
	Code    string
	Unknown bool
}

func (e *APIError) Error() string { return "DNS provider request failed: " + e.Code }
func Ambiguous(err error) bool    { var e *APIError; return errors.As(err, &e) && e.Unknown }

type Provider interface {
	Capabilities() Capabilities
	Zones(context.Context) ([]Zone, error)
	Records(context.Context, Zone) ([]Record, error)
	Get(context.Context, Zone, string) (Record, error)
	Create(context.Context, Zone, Record) (Record, error)
	Update(context.Context, Zone, Record) (Record, error)
	Delete(context.Context, Zone, string) error
}

type client struct {
	kind   string
	config map[string]any
	http   *http.Client
}

func New(kind string, config map[string]any) (Provider, error) {
	return NewWithClient(kind, config, &http.Client{Timeout: 30 * time.Second, CheckRedirect: func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse }})
}
func NewWithClient(kind string, config map[string]any, httpClient *http.Client) (Provider, error) {
	c := &client{kind: kind, config: config, http: httpClient}
	switch kind {
	case "cloudflare":
		if c.secret("apiToken") == "" {
			return nil, errors.New("Cloudflare API Token is required")
		}
		return &cloudflare{c}, nil
	case "dnspod", "edgeone":
		if c.secret("secretId") == "" || c.secret("secretKey") == "" {
			return nil, errors.New("Tencent SecretId/SecretKey credentials are required; DNSPod legacy tokens are a separate format")
		}
		return &tencent{c}, nil
	default:
		return nil, errors.New("unsupported DNS management provider")
	}
}
func (c *client) secret(key string) string { v, _ := c.config[key].(string); return v }
func hash(b []byte) string                 { s := sha256.Sum256(b); return hex.EncodeToString(s[:]) }
func mac(key []byte, s string) []byte {
	h := hmac.New(sha256.New, key)
	h.Write([]byte(s))
	return h.Sum(nil)
}
func (c *client) request(ctx context.Context, method, endpoint string, payload any, headers map[string]string, out any) error {
	var data []byte
	var err error
	if payload != nil {
		data, err = json.Marshal(payload)
		if err != nil {
			return err
		}
	}
	req, err := http.NewRequestWithContext(ctx, method, endpoint, bytes.NewReader(data))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := c.http.Do(req)
	if err != nil {
		return &APIError{Code: "transport", Unknown: method != "GET"}
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return &APIError{Code: fmt.Sprintf("http_%d", resp.StatusCode), Unknown: resp.StatusCode >= 500 && method != "GET"}
	}
	decoder := json.NewDecoder(io.LimitReader(resp.Body, 16<<20))
	decoder.UseNumber()
	if err := decoder.Decode(out); err != nil {
		return &APIError{Code: "invalid_response", Unknown: method != "GET"}
	}
	return nil
}
func (c *client) tc(ctx context.Context, action string, payload map[string]any, out any) error {
	service, version := "teo", "2022-09-01"
	if c.kind == "dnspod" {
		service, version = "dnspod", "2021-03-23"
	}
	host := service + ".tencentcloudapi.com"
	stamp := time.Now().UTC()
	date := stamp.Format("2006-01-02")
	timestamp := strconv.FormatInt(stamp.Unix(), 10)
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	canonical := "POST\n/\n\ncontent-type:application/json; charset=utf-8\nhost:" + host + "\n\ncontent-type;host\n" + hash(data)
	scope := date + "/" + service + "/tc3_request"
	signing := mac(mac(mac([]byte("TC3"+c.secret("secretKey")), date), service), "tc3_request")
	signature := hex.EncodeToString(mac(signing, "TC3-HMAC-SHA256\n"+timestamp+"\n"+scope+"\n"+hash([]byte(canonical))))
	headers := map[string]string{"X-TC-Action": action, "X-TC-Timestamp": timestamp, "X-TC-Version": version, "Authorization": "TC3-HMAC-SHA256 Credential=" + c.secret("secretId") + "/" + scope + ", SignedHeaders=content-type;host, Signature=" + signature}
	var envelope struct{ Response json.RawMessage }
	if err := c.request(ctx, "POST", "https://"+host+"/", payload, headers, &envelope); err != nil {
		return err
	}
	var failure struct{ Error *struct{ Code string } }
	if err := json.Unmarshal(envelope.Response, &failure); err != nil {
		return &APIError{Code: "invalid_response", Unknown: !strings.HasPrefix(action, "Describe")}
	}
	if failure.Error != nil {
		return &APIError{Code: failure.Error.Code, Unknown: strings.HasPrefix(failure.Error.Code, "InternalError") && !strings.HasPrefix(action, "Describe")}
	}
	return json.Unmarshal(envelope.Response, out)
}
func pathID(id string) string { return url.PathEscape(id) }
func fqdn(name, zone string) string {
	if name == "@" || name == "" {
		return zone
	}
	if strings.EqualFold(name, zone) || strings.HasSuffix(strings.ToLower(name), "."+strings.ToLower(zone)) {
		return name
	}
	return name + "." + zone
}
func relative(name, zone string) string {
	if strings.EqualFold(name, zone) {
		return "@"
	}
	return strings.TrimSuffix(name, "."+zone)
}
