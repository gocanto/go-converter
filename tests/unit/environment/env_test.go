package environment

import (
	"github.com/voyago/converter/environment"
	"github.com/voyago/converter/pkg/store/handler/currencyLayer"
	"os"
	"testing"
)

func TestItCanLoadDefaultEnvFiles(t *testing.T) {
	t.Parallel()

	env, err := environment.Make("converter")

	if err != nil {
		t.Errorf("%v", err)
		t.FailNow()
	}

	if env.GetRootDir() != "converter" {
		t.Errorf("The given env root directoy invalid")
	}

	if !env.IsLive() {
		t.Errorf("The given env is invalid")
	}

	if env.Get(environment.EnvKey) != environment.EnvLiveValue {
		t.Errorf("The given key [%s] is invalid", environment.EnvKey)
	}
}

func TestItCanLoadGivenEnvFiles(t *testing.T) {
	t.Parallel()

	env, err := environment.MakeWith("converter", ".env.example")

	if err != nil {
		t.Errorf("%v", err)
		t.FailNow()
	}

	if env.GetRootDir() != "converter" {
		t.Errorf("The given env root directoy invalid")
	}

	if !env.IsTest() {
		t.Errorf("The given env is invalid")
	}

	if env.Get(environment.EnvKey) != "test" {
		t.Errorf("The given key [%s] is invalid", environment.EnvKey)
	}

	if env.Get(currencyLayer.ApiKeyName) != "foo-bar-biz" {
		t.Errorf("The given key [%s] is invalid", currencyLayer.ApiKeyName)
	}
}

func TestItFallbackToOsEnvIfKeysAreNotFound(t *testing.T) {
	t.Parallel()

	key := "GITHUB_HANDLER"
	value := "gocanto"

	err := os.Setenv(key, value)

	if err != nil {
		t.Errorf("%v", err)
		t.FailNow()
	}

	env, _ := environment.Make("converter")

	if env.GetRootDir() != "converter" {
		t.Errorf("The given env root directoy invalid")
	}

	if env.Get(key) != value {
		t.Errorf("The given key [%s] is invalid", key)
	}
}
