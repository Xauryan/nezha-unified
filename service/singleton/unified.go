package singleton

import (
	"github.com/certimate-go/certimate/pkg/embedded"
	"github.com/nezhahq/nezha/service/unified"
)

var Unified *unified.Service

func InitUnified(dir string) error {
	engine, err := embedded.Open(dir)
	if err != nil {
		return err
	}
	service, err := unified.New(DB, engine)
	if err != nil {
		_ = engine.Close()
		return err
	}
	Unified = service
	return nil
}
