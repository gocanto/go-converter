package environment

import (
	"github.com/voyago/converter/environment"
	"testing"
)

func TestItProperlyLoadsTheDefaultEnv(t *testing.T) {
	env, err := environment.Make()

	if err != nil {
		t.Errorf("%v", err)
		t.FailNow()
	}

	if env.Get("CONVERTER_ENV") != "dev" {
		t.Errorf("The given key [CONVERTER_ENV] is invalid")
	}
}
