package test_entity

import (
	"github.com/voyago/converter/pkg/entity"
	"github.com/voyago/converter/test"
	"testing"
)

func TestItComputesValidMultipliers(t *testing.T) {
	a := entity.Currency{IsoMinorUnit: 1}
	b := entity.Currency{IsoMinorUnit: 2}
	c := entity.Currency{IsoMinorUnit: 3}

	if value := a.GetMultiplier(); value != 10 {
		t.Errorf("The given multiplier [%v] is invalid", value)
	}

	if value := b.GetMultiplier(); value != 100 {
		t.Errorf("The given multiplier [%v] is invalid", value)
	}

	if value := c.GetMultiplier(); value != 1000 {
		t.Errorf("The given multiplier [%v] is invalid", value)
	}
}

func TestItHoldsBasicInfo(t *testing.T) {
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
