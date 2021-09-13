package test_entity

import (
	"github.com/voyago/converter/pkg/entity"
	"github.com/voyago/converter/test/support"
	"testing"
)

const stub = "./../../__fixtures__/currency.json"

func TestCurrencyJsonFormat(t *testing.T) {
	currency := entity.Currency{}

	if err := test_support.ParseJson(stub, &currency); err != nil {
		t.Errorf("There was an error [%d] opening the given currency stub", err)
	}

	if currency.Code != "SGD" {
		t.Errorf("Currency code [%s] is invalid", currency.Code)
	}

	if currency.Name != "Singapore Dollar" {
		t.Errorf("Currency name [%s] is invalid", currency.Name)
	}

	if currency.Symbol != "$" {
		t.Errorf("Currency smbold [%s] is invalid", currency.Symbol)
	}

	if currency.Rate != 0.74 {
		t.Errorf("Currency smbold [%f] is invalid", currency.Rate)
	}

	if currency.IsoCode != 702 {
		t.Errorf("Currency iso code [%d] is invalid", currency.IsoCode)
	}

	if currency.IsoMinorUnit != 2 {
		t.Errorf("Currency iso minor unit [%d] is invalid", currency.IsoMinorUnit)
	}
}
