package environment

import (
	"github.com/voyago/converter/environment"
	"testing"
)

func TestItLoadsTheLiveEnv(t *testing.T) {
	env, err := environment.Make()

	if err != nil {
		t.Errorf("%v", err)
		t.FailNow()
	}

	if env.IsTest() {
		t.Errorf("The given env is invalid")
	}

	if env.Get("CONVERTER_ENV") != "local" {
		t.Errorf("The given key [CONVERTER_ENV] is invalid")
	}
}

func TestItCanLoadDifferentFiles(t *testing.T) {
	env, err := environment.MakeWith("testing")

	if err != nil {
		t.Errorf("%v", err)
		t.FailNow()
	}

	if env.IsLive() {
		t.Errorf("The given env is invalid")
	}

	if env.Get("CONVERTER_ENV") != "local" {
		t.Errorf("The given key [CONVERTER_ENV] is invalid")
	}

	if env.Get("CONVERTER_CURRENCY_LAYER_KEY") != "" {
		t.Errorf("The given key [CONVERTER_CURRENCY_LAYER_KEY] is invalid")
	}
}
