package app

import (
	"errors"
	"log/slog"
	"sync"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

var (
	instance      core.App
	instanceMu    sync.Mutex
	embeddedOwner core.App
)

func GetApp() core.App {
	instanceMu.Lock()
	defer instanceMu.Unlock()
	if instance == nil {
		pb := pocketbase.NewWithConfig(pocketbase.Config{
			HideStartBanner: true,
		})

		pb.RootCmd.Flags().MarkHidden("encryptionEnv")
		pb.RootCmd.Flags().MarkHidden("queryTimeout")

		instance = pb
	}

	return instance
}

func GetDB() dbx.Builder {
	return GetApp().DB()
}

func GetLogger() *slog.Logger {
	app := GetApp()
	if !app.IsBootstrapped() {
		panic("MUST NOT USE THIS BEFORE APP BOOTSTRAPPED!")
	}

	return app.Logger()
}

// BindEmbedded installs the host-owned database without bootstrapping a server.
// Exactly one engine may own the original process-global Certimate services.
func BindEmbedded(pb core.App) error {
	instanceMu.Lock()
	defer instanceMu.Unlock()
	if embeddedOwner != nil && embeddedOwner != pb {
		return errors.New("another Certimate runtime is already active")
	}
	if instance != nil && instance != pb && instance.IsBootstrapped() {
		return errors.New("another Certimate application is already active")
	}
	instance, embeddedOwner = pb, pb
	return nil
}
func ReleaseEmbedded(pb core.App) {
	instanceMu.Lock()
	defer instanceMu.Unlock()
	if embeddedOwner == pb {
		instance, embeddedOwner = nil, nil
	}
}
