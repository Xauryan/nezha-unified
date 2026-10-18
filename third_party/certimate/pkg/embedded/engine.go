// Package embedded is the public Nezha integration boundary. It owns Certimate's
// internal database without starting PocketBase HTTP, authentication or cron.
package embedded

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/certimate-go/certimate/internal/embedcrypto"
	"github.com/certimate-go/certimate/internal/embeddedhooks"
	_ "github.com/certimate-go/certimate/migrations"
	"github.com/pocketbase/pocketbase/core"
)

type Engine struct {
	app                 core.App
	aead                cipher.AEAD
	mu                  sync.RWMutex
	closed              bool
	closing             atomic.Bool
	runtimeMu           sync.Mutex
	runtime             *runtimeState
	authorizeCredential CredentialAuthorizer
	authorizerVersion   uint64
	effects             sync.WaitGroup
}

// Open uses a dedicated directory. A missing key on an existing database is a
// recovery error; generating a replacement would silently lose decryption access.
func Open(dir string) (*Engine, error) {
	if dir == "" {
		return nil, errors.New("Certimate data directory is required")
	}
	if err := os.MkdirAll(dir, 0700); err != nil {
		return nil, err
	}
	if err := os.Chmod(dir, 0700); err != nil {
		return nil, err
	}
	keyPath := filepath.Join(dir, "unified.key")
	key, err := os.ReadFile(keyPath)
	if errors.Is(err, os.ErrNotExist) {
		if _, statErr := os.Stat(filepath.Join(dir, "data.db")); !errors.Is(statErr, os.ErrNotExist) {
			return nil, errors.New("Certimate database exists but unified.key is missing; restore the original key")
		}
		key = make([]byte, 32)
		if _, err = rand.Read(key); err != nil {
			return nil, err
		}
		f, openErr := os.OpenFile(keyPath, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0600)
		if openErr != nil {
			return nil, openErr
		}
		_, writeErr := f.Write(key)
		syncErr := f.Sync()
		err = errors.Join(writeErr, syncErr, f.Close())
	}
	if err != nil {
		return nil, err
	}
	if len(key) != 32 {
		return nil, errors.New("unified.key must contain exactly 32 bytes")
	}
	if err := os.Chmod(keyPath, 0600); err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	aead, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	pb := core.NewBaseApp(core.BaseAppConfig{DataDir: dir, QueryTimeout: 15 * time.Second})
	pb.Store().Set("nezha_embedded", true)
	cleanup := func(err error) (*Engine, error) { return nil, errors.Join(err, pb.ResetBootstrapState()) }
	if err := pb.Bootstrap(); err != nil {
		return cleanup(err)
	}
	if err := pb.RunAppMigrations(); err != nil {
		return cleanup(err)
	}
	if err := migrateUnified(pb); err != nil {
		return cleanup(err)
	}
	e := &Engine{app: pb, aead: aead}
	embedcrypto.Configure(pb, &embedcrypto.Hooks{Seal: e.seal, Open: e.unseal, ResolveAccess: func(ctx context.Context, id string) (map[string]any, error) {
		return embeddedhooks.Credential(ctx, id, "workflow")
	}})
	if err := e.verifyRecoveryKey(); err != nil {
		return cleanup(err)
	}
	if err := embedcrypto.Migrate(pb); err != nil {
		return cleanup(err)
	}
	return e, nil
}

// Verify the recovery key against stored material before accepting requests.
// This detects restoring the DB and key from unrelated installations.
func (e *Engine) verifyRecoveryKey() error {
	for collection, field := range map[string]string{"access": "unified_envelope", "certificate": "unified_private_key"} {
		records, err := e.app.FindRecordsByFilter(collection, field+" != ''", "id", 0, 0)
		if err != nil {
			return err
		}
		for _, r := range records {
			value := r.GetString(field)
			blob, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(value, "v1:"))
			if err != nil || !strings.HasPrefix(value, "v1:") || len(blob) < e.aead.NonceSize()+e.aead.Overhead() {
				return errors.New("invalid encrypted resource; restore a consistent backup")
			}
			aad := "certificate:" + r.Id
			if collection == "access" {
				aad = fmt.Sprintf("credential:%s:%d", r.Id, r.GetInt("unified_version"))
			}
			if _, err := e.aead.Open(nil, blob[:e.aead.NonceSize()], blob[e.aead.NonceSize():], []byte(aad)); err != nil {
				return errors.New("unified.key cannot decrypt existing resources; restore the matching recovery key")
			}
		}
	}
	return nil
}

// Close waits for current service calls; the host shuts down HTTP before Close.
func (e *Engine) Close() error {
	e.closing.Store(true)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := e.StopRuntime(ctx); err != nil {
		return err
	}
	e.mu.Lock()
	defer e.mu.Unlock()
	if e.closed {
		return nil
	}
	e.effects.Wait()
	e.closed = true
	return e.app.ResetBootstrapState()
}

func (e *Engine) check(ctx context.Context) error {
	if e.closed || e.closing.Load() {
		return errors.New("Certimate storage is closed")
	}
	return ctx.Err()
}

func (e *Engine) seal(id string, data []byte) string {
	nonce := make([]byte, e.aead.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		panic(err)
	}
	return "v1:" + base64.StdEncoding.EncodeToString(e.aead.Seal(nonce, nonce, data, []byte(id)))
}

func migrateUnified(app core.App) error {
	return app.RunInTransaction(func(tx core.App) error {
		for name, fields := range map[string][]core.Field{
			"access":        {&core.TextField{Name: "unified_envelope", Hidden: true}, &core.BoolField{Name: "unified_enabled"}, &core.NumberField{Name: "unified_version", OnlyInt: true}},
			"certificate":   {&core.TextField{Name: "unified_private_key", Hidden: true}},
			"workflow":      {},
			"acme_accounts": {},
		} {
			c, err := tx.FindCollectionByNameOrId(name)
			if err != nil {
				return err
			}
			changed := false
			// New imports keep key material exclusively in the authenticated
			// envelope. Only the embedded application service writes this store.
			if name == "certificate" {
				keyField, ok := c.Fields.GetByName("privateKey").(*core.TextField)
				if !ok {
					return errors.New("unexpected certificate privateKey schema")
				}
				if keyField.Required || !keyField.Hidden {
					keyField.Required = false
					keyField.Hidden = true
					changed = true
				}
			}
			if name == "acme_accounts" {
				if f, ok := c.Fields.GetByName("privateKey").(*core.TextField); ok && !f.Hidden {
					f.Hidden = true
					changed = true
				}
			}
			if name == "access" || name == "workflow" || name == "certificate" {
				fields = append(fields, &core.TextField{Name: "unified_operation_id", Hidden: true}, &core.BoolField{Name: "unified_active", Hidden: true})
			}
			for _, f := range fields {
				if c.Fields.GetByName(f.GetName()) == nil {
					c.Fields.Add(f)
					changed = true
				}
			}
			if name == "access" || name == "workflow" || name == "certificate" {
				index := "CREATE UNIQUE INDEX idx_" + name + "_unified_operation ON " + name + " (unified_operation_id) WHERE unified_operation_id != ''"
				found := false
				for _, existing := range c.Indexes {
					if existing == index {
						found = true
					}
				}
				if !found {
					c.Indexes = append(c.Indexes, index)
					changed = true
				}
			}
			if changed {
				if err := tx.Save(c); err != nil {
					return err
				}
			}
		}
		if _, err := tx.FindCollectionByNameOrId("unified_audit"); err != nil {
			c := core.NewBaseCollection("unified_audit")
			c.Fields.Add(&core.TextField{Name: "actor", Required: true}, &core.TextField{Name: "action", Required: true}, &core.TextField{Name: "resource", Required: true}, &core.AutodateField{Name: "created", OnCreate: true})
			if err := tx.Save(c); err != nil {
				return fmt.Errorf("audit migration: %w", err)
			}
		}
		return nil
	})
}

func audit(tx core.App, actor, action, resource string) error {
	c, err := tx.FindCollectionByNameOrId("unified_audit")
	if err != nil {
		return err
	}
	r := core.NewRecord(c)
	r.Set("actor", actor)
	r.Set("action", action)
	r.Set("resource", resource)
	return tx.Save(r)
}
