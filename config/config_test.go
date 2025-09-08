package config

import (
	"github.com/rotisserie/eris"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Error(eris.Wrap(r.(error), "Error"))
		}
	}()
	config := GetAppConfig()
	if config.DbName == "" {
		t.Error("config not loaded")
	}
}
