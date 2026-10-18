package settings

import (
	"github.com/certimate-go/certimate/internal/app"
	"github.com/certimate-go/certimate/internal/domain"
	"github.com/certimate-go/certimate/internal/embedcrypto"
	"github.com/pocketbase/pocketbase/core"
)

func Setup() {
	initPbSettings()

	registerSettingsStoreByName(domain.SettingsNameSSLProvider)
	registerSettingsStoreByName(domain.SettingsNamePersistence)
	registerSettingsRecordEvents()
}

// InitEmbedded loads only application settings, without public API/request hooks.
func InitEmbedded() error {
	for _, name := range []string{domain.SettingsNameSSLProvider, domain.SettingsNamePersistence} {
		if err := registerSettingsStoreByName(name); err != nil && !domain.IsRecordNotFoundError(err) {
			return err
		}
	}
	return nil
}
func RefreshEmbedded(record *core.Record, deleted bool) error {
	pb := app.GetApp()
	name := record.GetString("name")
	if deleted {
		pb.Store().Remove(buildPbStoreKey(name))
		return nil
	}
	var content domain.SettingsContent
	if err := embedcrypto.ReadJSON(pb, record, "content", &content); err != nil {
		return err
	}
	pb.Store().Set(buildPbStoreKey(name), content)
	return nil
}
