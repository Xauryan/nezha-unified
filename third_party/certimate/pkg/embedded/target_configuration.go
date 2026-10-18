package embedded

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
)

func targetConfigurationAAD(id string, version uint64) (string, error) {
	if strings.TrimSpace(id) == "" || len(id) > 128 || version == 0 {
		return "", ErrInvalid
	}
	return fmt.Sprintf("deployment_target:%s:%d", id, version), nil
}

// SealTargetConfiguration creates a server-only immutable configuration value.
// The host persists only the envelope and authorizes the target before use.
func (e *Engine) SealTargetConfiguration(ctx context.Context, id string, version uint64, config map[string]any) (string, error) {
	e.mu.RLock()
	defer e.mu.RUnlock()
	if err := e.check(ctx); err != nil {
		return "", err
	}
	aad, err := targetConfigurationAAD(id, version)
	if err != nil {
		return "", err
	}
	if config == nil {
		config = map[string]any{}
	}
	encoded, err := json.Marshal(config)
	if err != nil || len(encoded) > 2<<20 {
		return "", ErrInvalid
	}
	return e.seal(aad, encoded), nil
}

// OpenTargetConfiguration binds decrypted settings to both target identity and
// exact version. This value must never be serialized into target metadata.
func (e *Engine) OpenTargetConfiguration(ctx context.Context, id string, version uint64, envelope string) (map[string]any, error) {
	e.mu.RLock()
	defer e.mu.RUnlock()
	if err := e.check(ctx); err != nil {
		return nil, err
	}
	aad, err := targetConfigurationAAD(id, version)
	if err != nil {
		return nil, err
	}
	encoded, err := e.unseal(aad, envelope)
	if err != nil || len(encoded) > 2<<20 {
		return nil, ErrInvalid
	}
	var config map[string]any
	if json.Unmarshal(encoded, &config) != nil || config == nil {
		return nil, ErrInvalid
	}
	return config, nil
}
