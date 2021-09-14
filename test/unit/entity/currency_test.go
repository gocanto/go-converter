package test_entity

import (
	"github.com/voyago/converter/test"
	"testing"
)

func TestCurrencyJsonFormat(t *testing.T) {
	currency := test.Currency(t)

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
