package embedded

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
)

var ErrInvalid = errors.New("invalid resource")
var ErrNotFound = errors.New("resource not found")

var collections = map[string]string{"certificate": "certificate", "credential": "access", "workflow": "workflow"}

// Resource is an explicit metadata projection. Never serialize raw PocketBase
// records: workflows, access configs and certificate records contain secrets.
type Resource struct {
	ID       string `json:"id"`
	Kind     string `json:"kind"`
	Name     string `json:"name"`
	Provider string `json:"provider,omitempty"`
	Names    string `json:"names,omitempty"`
	Expires  string `json:"expires,omitempty"`
	Enabled  bool   `json:"enabled"`
	Version  int    `json:"version,omitempty"`
}

type Page struct {
	Items   []Resource `json:"items"`
	Total   int        `json:"total"`
	Page    int        `json:"page"`
	PerPage int        `json:"per_page"`
}

func project(kind string, r *core.Record) Resource {
	v := Resource{ID: r.Id, Kind: kind, Name: r.GetString("name")}
	switch kind {
	case "credential":
		v.Provider = r.GetString("provider")
		v.Enabled = r.GetBool("unified_enabled")
		v.Version = r.GetInt("unified_version")
	case "certificate":
		v.Name = r.GetString("subjectName")
		v.Names = r.GetString("subjectAltNames")
		v.Expires = r.GetString("validityNotAfter")
		v.Enabled = !r.GetBool("isRevoked")
	case "workflow":
		v.Enabled = r.GetBool("enabled")
	}
	return v
}

// List applies resource authorization before COUNT/LIMIT/OFFSET. all must only
// come from the host's verified administrator session, never from query input.
func (e *Engine) List(ctx context.Context, kind string, ids []string, all bool, page, perPage int) (Page, error) {
	e.mu.RLock()
	defer e.mu.RUnlock()
	result := Page{Items: []Resource{}, Page: page, PerPage: perPage}
	if err := e.check(ctx); err != nil {
		return result, err
	}
	collection, ok := collections[kind]
	if !ok || page < 1 || perPage < 1 || perPage > 100 || page > 100000 {
		return result, ErrInvalid
	}
	if !all && len(ids) == 0 {
		return result, nil
	}
	query := e.app.RecordQuery(collection).WithContext(ctx).AndWhere(dbx.HashExp{"deleted": ""})
	if !all {
		query.AndWhere(dbx.In("id", slicesToAny(ids)...))
	}
	countQuery := *query
	if err := countQuery.Select("COUNT(*)").Row(&result.Total); err != nil {
		return result, err
	}
	var records []*core.Record
	if err := query.OrderBy("created DESC", "id ASC").Limit(int64(perPage)).Offset(int64((page - 1) * perPage)).All(&records); err != nil {
		return result, err
	}
	for _, r := range records {
		result.Items = append(result.Items, project(kind, r))
	}
	return result, nil
}

func slicesToAny(ids []string) []any {
	v := make([]any, len(ids))
	for i, id := range ids {
		v[i] = id
	}
	return v
}

func (e *Engine) Exists(ctx context.Context, kind, id string) bool {
	e.mu.RLock()
	defer e.mu.RUnlock()
	collection, ok := collections[kind]
	if !ok || e.check(ctx) != nil {
		return false
	}
	r, err := e.app.FindRecordById(collection, id)
	return err == nil && r.GetDateTime("deleted").IsZero()
}

type CredentialInput struct {
	Name     string          `json:"name"`
	Provider string          `json:"provider"`
	Config   json.RawMessage `json:"config"`
}

func (e *Engine) CreateCredential(ctx context.Context, actor string, in CredentialInput) (Resource, error) {
	e.mu.RLock()
	defer e.mu.RUnlock()
	if err := e.check(ctx); err != nil {
		return Resource{}, err
	}
	var config map[string]any
	if actor == "" || strings.TrimSpace(in.Name) == "" || len(in.Name) > 120 || in.Provider == "" || len(in.Provider) > 100 || len(in.Config) > 65536 || json.Unmarshal(in.Config, &config) != nil || len(config) == 0 {
		return Resource{}, ErrInvalid
	}
	var result Resource
	err := e.app.RunInTransaction(func(tx core.App) error {
		c, err := tx.FindCollectionByNameOrId("access")
		if err != nil {
			return err
		}
		r := core.NewRecord(c)
		// Saving once gives the record its stable ID; the local transaction keeps
		// the empty config invisible until its encrypted replacement is committed.
		r.Set("name", strings.TrimSpace(in.Name))
		r.Set("provider", in.Provider)
		r.Set("config", map[string]any{})
		r.Set("unified_enabled", true)
		r.Set("unified_version", 1)
		if err := tx.Save(r); err != nil {
			return err
		}
		r.Set("unified_envelope", e.seal("credential:"+r.Id+":1", in.Config))
		if err := tx.Save(r); err != nil {
			return err
		}
		result = project("credential", r)
		return audit(tx, actor, "credential.create", "credential:"+r.Id)
	})
	return result, err
}

type CertificateInput struct {
	Certificate string `json:"certificate"`
	PrivateKey  string `json:"private_key"`
}

func (e *Engine) ImportCertificate(ctx context.Context, actor string, in CertificateInput) (Resource, error) {
	e.mu.RLock()
	defer e.mu.RUnlock()
	if err := e.check(ctx); err != nil {
		return Resource{}, err
	}
	if actor == "" || len(in.Certificate) > 131072 || len(in.PrivateKey) > 32768 {
		return Resource{}, ErrInvalid
	}
	pair, err := tls.X509KeyPair([]byte(in.Certificate), []byte(in.PrivateKey))
	if err != nil {
		return Resource{}, fmt.Errorf("%w: certificate and private key must match", ErrInvalid)
	}
	leaf, err := x509.ParseCertificate(pair.Certificate[0])
	if err != nil {
		return Resource{}, ErrInvalid
	}
	names := slices.Clone(leaf.DNSNames)
	for _, ip := range leaf.IPAddresses {
		names = append(names, ip.String())
	}
	if len(names) == 0 || leaf.IsCA {
		return Resource{}, fmt.Errorf("%w: a leaf certificate with SANs is required", ErrInvalid)
	}
	// tls.X509KeyPair tolerates unrelated PEM blocks. Re-encode only parsed
	// certificates so an accidentally concatenated private key can never enter
	// the public-chain download or the unencrypted certificate field.
	var publicChain strings.Builder
	for _, der := range pair.Certificate {
		if _, err := x509.ParseCertificate(der); err != nil {
			return Resource{}, ErrInvalid
		}
		publicChain.Write(pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: der}))
	}
	var result Resource
	err = e.app.RunInTransaction(func(tx core.App) error {
		c, err := tx.FindCollectionByNameOrId("certificate")
		if err != nil {
			return err
		}
		r := core.NewRecord(c)
		r.Set("source", "upload")
		r.Set("certificate", publicChain.String())
		r.Set("subjectName", leaf.Subject.CommonName)
		if leaf.Subject.CommonName == "" {
			r.Set("subjectName", names[0])
		}
		r.Set("subjectAltNames", strings.Join(names, ";"))
		r.Set("serialNumber", strings.ToUpper(leaf.SerialNumber.Text(16)))
		r.Set("issuerName", leaf.Issuer.CommonName)
		r.Set("validityNotBefore", leaf.NotBefore)
		r.Set("validityNotAfter", leaf.NotAfter)
		r.Set("validityInterval", int64(leaf.NotAfter.Sub(leaf.NotBefore).Seconds()))
		if err := tx.Save(r); err != nil {
			return err
		}
		r.Set("unified_private_key", e.seal("certificate:"+r.Id, []byte(in.PrivateKey)))
		if err := tx.Save(r); err != nil {
			return err
		}
		result = project("certificate", r)
		return audit(tx, actor, "certificate.import", "certificate:"+r.Id)
	})
	return result, err
}

// CertificatePEM returns the public chain only. Private key export requires a
// separate future action and is intentionally absent from this API surface.
func (e *Engine) CertificatePEM(ctx context.Context, id string) (string, error) {
	e.mu.RLock()
	defer e.mu.RUnlock()
	if err := e.check(ctx); err != nil {
		return "", err
	}
	r, err := e.app.FindRecordById("certificate", id)
	if err != nil || !r.GetDateTime("deleted").IsZero() {
		return "", ErrNotFound
	}
	return r.GetString("certificate"), nil
}
