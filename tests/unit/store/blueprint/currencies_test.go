package blueprint

import (
	"github.com/voyago/converter/pkg/store/blueprint"
	"strings"
	"testing"
)

func TestItReturnsValidCurrenciesBlueprint(t *testing.T) {
	currencies, err := blueprint.MakeCurrenciesBlueprint()

	if err != nil {
		t.Errorf("%v", err)
		t.FailNow()
	}

	if len(currencies.Items()) < 1 {
		t.Errorf("The currencies blueprint has invalid items")
	}

	if !strings.Contains(currencies.GetSourcePath(), "resources/currencies.json") {
		t.Errorf("The currencies blueprint has an invalid source path")
	}

	if currencies.Items()[0]["code"] != "AFN" {
		t.Errorf("The currencies blueprint for [AFN] is invalid")
	}
}
