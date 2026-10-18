// Package embedcrypto keeps secrets encrypted at the persistence boundary while
// allowing the unchanged workflow engine to operate on in-memory plaintext.
package embedcrypto

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"errors"
	"math/big"
	"strings"

	"github.com/pocketbase/pocketbase/core"
)

const storeKey = "nezha_embedded_crypto"
const envelopeKey = "unified_envelope"

// Hooks belong to one database and must not be modified after Configure.
type Hooks struct {
	Seal          func(string, []byte) string
	Open          func(string, string) ([]byte, error)
	ResolveAccess func(context.Context, string) (map[string]any, error)
}

func Configure(app core.App, hooks *Hooks) { app.Store().Set(storeKey, hooks) }

func configured(app core.App) *Hooks {
	hooks, _ := app.Store().Get(storeKey).(*Hooks)
	return hooks
}

// Enabled is true for an embedded database even before hooks are configured, so
// missing initialization fails closed instead of persisting plaintext.
func Enabled(app core.App) bool {
	return app.Store().Get("nezha_embedded") == true || configured(app) != nil
}

func Seal(app core.App, aad string, value []byte) (string, error) {
	hooks := configured(app)
	if hooks == nil || hooks.Seal == nil {
		return "", errors.New("embedded encryption is unavailable")
	}
	return hooks.Seal(aad, value), nil
}

func Open(app core.App, aad, value string) ([]byte, error) {
	hooks := configured(app)
	if hooks == nil || hooks.Open == nil || !strings.HasPrefix(value, "v1:") {
		return nil, errors.New("embedded decryption is unavailable")
	}
	plain, err := hooks.Open(aad, value)
	if err != nil {
		return nil, errors.New("encrypted data could not be authenticated")
	}
	return plain, nil
}

func ResolveAccess(app core.App, ctx context.Context, id string) (map[string]any, error) {
	hooks := configured(app)
	if hooks == nil || hooks.ResolveAccess == nil {
		return nil, errors.New("embedded credential authorization is unavailable")
	}
	return hooks.ResolveAccess(ctx, id)
}

// EnsureID assigns the same lower-case alphanumeric form used by PocketBase
// before binding ciphertext to the record identity. No plaintext pre-save is needed.
func EnsureID(record *core.Record) error {
	if record.Id != "" {
		return nil
	}
	const alphabet = "abcdefghijklmnopqrstuvwxyz0123456789"
	var id [15]byte
	for i := range id {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(alphabet))))
		if err != nil {
			return errors.New("record identity generation failed")
		}
		id[i] = alphabet[n.Int64()]
	}
	record.Id = string(id[:])
	return nil
}

func SealJSON(app core.App, aad string, value any) (any, error) {
	if !Enabled(app) {
		return value, nil
	}
	plain, err := json.Marshal(value)
	if err != nil {
		return nil, errors.New("protected JSON is invalid")
	}
	envelope, err := Seal(app, aad, plain)
	if err != nil {
		return nil, err
	}
	return map[string]string{envelopeKey: envelope}, nil
}

func OpenJSON(app core.App, aad string, value any, target any) error {
	raw, err := json.Marshal(value)
	if err != nil {
		return errors.New("protected JSON is invalid")
	}
	if Enabled(app) {
		var envelope map[string]string
		if json.Unmarshal(raw, &envelope) != nil || len(envelope) != 1 || envelope[envelopeKey] == "" {
			return errors.New("protected JSON envelope is missing")
		}
		raw, err = Open(app, aad, envelope[envelopeKey])
		if err != nil {
			return err
		}
	}
	if json.Unmarshal(raw, target) != nil {
		return errors.New("protected JSON is invalid")
	}
	return nil
}

func SetJSON(app core.App, record *core.Record, field string, value any) error {
	if Enabled(app) {
		if err := EnsureID(record); err != nil {
			return err
		}
	}
	value, err := SealJSON(app, record.Collection().Name+":"+record.Id+":"+field, value)
	if err != nil {
		return err
	}
	record.Set(field, value)
	return nil
}

func ReadJSON(app core.App, record *core.Record, field string, target any) error {
	return OpenJSON(app, record.Collection().Name+":"+record.Id+":"+field, record.Get(field), target)
}

// Provider errors and log fields can contain credential values, commands, PEM
// keys or HTTP payloads. Persist only a neutral event and its structured identity.
func SafeError(app core.App, message string) string {
	if Enabled(app) && message != "" {
		return "Workflow execution failed; sensitive provider details were omitted."
	}
	return message
}

func SafeLog(app core.App, message string, data map[string]any) (string, map[string]any) {
	if Enabled(app) {
		return "Workflow node event (provider details omitted).", map[string]any{}
	}
	return message, data
}
