package embedded

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/pocketbase/pocketbase/core"
)

// CredentialUse carries the actual execution identity and target. The host
// authorizer re-reads current grants for every resolution, including background work.
type CredentialUse struct {
	ActorID      uint64 `json:"actor_id"`
	CredentialID string `json:"credential_id"`
	Version      int    `json:"version"`
	Purpose      string `json:"purpose"`
	ResourceKind string `json:"resource_kind"`
	ResourceID   string `json:"resource_id"`
	WorkflowID   string `json:"workflow_id,omitempty"`
	RunID        string `json:"run_id,omitempty"`
}

type CredentialAuthorizer func(context.Context, CredentialUse) error

func (e *Engine) SetCredentialAuthorizer(authorize CredentialAuthorizer) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.authorizeCredential = authorize
}

func (e *Engine) unseal(aad, value string) ([]byte, error) {
	blob, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(value, "v1:"))
	if err != nil || !strings.HasPrefix(value, "v1:") || len(blob) < e.aead.NonceSize()+e.aead.Overhead() {
		return nil, errors.New("invalid encrypted value")
	}
	return e.aead.Open(nil, blob[:e.aead.NonceSize()], blob[e.aead.NonceSize():], []byte(aad))
}

// ResolveCredential is server-only. It never returns a secret without a host
// authorization decision, active state, exact purpose/target and current version.
func (e *Engine) ResolveCredential(ctx context.Context, use CredentialUse) (map[string]any, error) {
	e.mu.RLock()
	defer e.mu.RUnlock()
	if err := e.check(ctx); err != nil {
		return nil, err
	}
	if e.authorizeCredential == nil || use.ActorID == 0 || use.Version < 1 || use.Purpose == "" || use.ResourceKind == "" || use.ResourceID == "" {
		return nil, errors.New("credential execution context is incomplete")
	}
	if err := e.authorizeCredential(ctx, use); err != nil {
		return nil, err
	}
	r, err := e.app.FindRecordById("access", use.CredentialID)
	if err != nil || !nativeActive(r) || !r.GetDateTime("deleted").IsZero() {
		return nil, ErrNotFound
	}
	if !r.GetBool("unified_enabled") || r.GetInt("unified_version") != use.Version {
		return nil, errors.New("credential is disabled or its version changed")
	}
	plaintext, err := e.unseal(fmt.Sprintf("credential:%s:%d", r.Id, use.Version), r.GetString("unified_envelope"))
	if err != nil {
		return nil, err
	}
	var config map[string]any
	if err := json.Unmarshal(plaintext, &config); err != nil {
		return nil, err
	}
	return config, nil
}

// UpdateCredential uses compare-and-swap versioning so concurrent rotation never
// silently replaces newer secrets. A nil config changes enabled state only.
func (e *Engine) UpdateCredential(ctx context.Context, actor, id string, expected int, enabled bool, config json.RawMessage) (Resource, error) {
	e.mu.RLock()
	defer e.mu.RUnlock()
	if err := e.check(ctx); err != nil {
		return Resource{}, err
	}
	if actor == "" || expected < 1 {
		return Resource{}, ErrInvalid
	}
	if config != nil {
		var m map[string]any
		if len(config) > 65536 || json.Unmarshal(config, &m) != nil || len(m) == 0 {
			return Resource{}, ErrInvalid
		}
	}
	var result Resource
	err := e.app.RunInTransaction(func(tx core.App) error {
		r, err := tx.FindRecordById("access", id)
		if err != nil {
			return ErrNotFound
		}
		if r.GetInt("unified_version") != expected {
			return fmt.Errorf("%w: credential version conflict", ErrInvalid)
		}
		if config != nil {
			r.Set("unified_version", expected+1)
			r.Set("unified_envelope", e.seal(fmt.Sprintf("credential:%s:%d", id, expected+1), config))
		}
		r.Set("unified_enabled", enabled)
		if err := tx.Save(r); err != nil {
			return err
		}
		result = project("credential", r)
		return audit(tx, actor, "credential.update", "credential:"+id)
	})
	return result, err
}

func (e *Engine) CredentialInfo(ctx context.Context, id string) (Resource, error) {
	e.mu.RLock()
	defer e.mu.RUnlock()
	if err := e.check(ctx); err != nil {
		return Resource{}, err
	}
	r, err := e.app.FindRecordById("access", id)
	if err != nil || !r.GetDateTime("deleted").IsZero() {
		return Resource{}, ErrNotFound
	}
	return project("credential", r), nil
}
