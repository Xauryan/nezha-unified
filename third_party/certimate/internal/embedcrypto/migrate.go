package embedcrypto

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/pocketbase/pocketbase/core"
)

// Migrate protects records created by earlier embedded versions and upstream
// migrations. It must run after the host configures the database's recovery key
// and before exposing native services or starting any workflow workers.
func Migrate(app core.App) error {
	if !Enabled(app) {
		return nil
	}
	return app.RunInTransaction(func(tx core.App) error {
		for collection, fields := range map[string][]string{
			"workflow":        {"graphDraft", "graphContent"},
			"workflow_run":    {"graph"},
			"workflow_output": {"nodeConfig", "outputs"},
			"settings":        {"content"},
			"access":          {"config"},
		} {
			records, err := tx.FindAllRecords(collection)
			if err != nil {
				return err
			}
			for _, record := range records {
				changed := false
				for _, field := range fields {
					if collection == "access" && record.GetString("unified_envelope") != "" {
						record.Set("config", map[string]any{})
						changed = true
						continue
					}
					raw, err := json.Marshal(record.Get(field))
					if err != nil {
						return errors.New("protected migration JSON is invalid")
					}
					var envelope map[string]json.RawMessage
					if json.Unmarshal(raw, &envelope) == nil && envelope[envelopeKey] != nil {
						var value any
						if err := ReadJSON(app, record, field, &value); err != nil {
							return err
						}
						continue
					}
					if err := SetJSON(app, record, field, record.Get(field)); err != nil {
						return err
					}
					changed = true
				}
				if collection == "workflow_run" && record.GetString("error") != "" {
					record.Set("error", SafeError(app, record.GetString("error")))
					changed = true
				}
				if changed {
					if err := tx.Save(record); err != nil {
						return err
					}
				}
			}
		}
		for _, collection := range []string{"certificate", "acme_accounts"} {
			records, err := tx.FindAllRecords(collection)
			if err != nil {
				return err
			}
			for _, record := range records {
				privateKey := record.GetString("privateKey")
				field := "privateKey"
				if collection == "certificate" {
					field = "unified_private_key"
				}
				aad := "certificate:" + record.Id
				if collection == "acme_accounts" {
					aad = "acme_account:" + record.Id
				}
				encrypted := record.GetString(field)
				if strings.HasPrefix(encrypted, "v1:") {
					if _, err := Open(app, aad, encrypted); err != nil {
						return err
					}
					if collection != "certificate" || privateKey == "" {
						continue
					}
				} else {
					encrypted, err = Seal(app, aad, []byte(privateKey))
					if err != nil {
						return err
					}
				}
				record.Set(field, encrypted)
				if collection == "certificate" {
					record.Set("privateKey", "")
				}
				if err := tx.Save(record); err != nil {
					return err
				}
			}
		}
		logs, err := tx.FindAllRecords("workflow_logs")
		if err != nil {
			return err
		}
		for _, record := range logs {
			message, data := SafeLog(app, record.GetString("message"), nil)
			record.Set("message", message)
			record.Set("data", data)
			if err := tx.Save(record); err != nil {
				return err
			}
		}
		return nil
	})
}
