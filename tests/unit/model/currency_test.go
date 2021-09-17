package model

import (
	"github.com/voyago/converter/src/model"
	"github.com/voyago/converter/tests/mock"
	"testing"
)

func TestItComputesValidMultipliers(t *testing.T) {
	a := model.Currency{IsoMinorUnit: 1}
	b := model.Currency{IsoMinorUnit: 2}
	c := model.Currency{IsoMinorUnit: 3}

	if value := a.Multiplier(); value != 10 {
		t.Errorf("The given multiplier [%v] is invalid", value)
	}

	if value := b.Multiplier(); value != 100 {
		t.Errorf("The given multiplier [%v] is invalid", value)
	}

	if value := c.Multiplier(); value != 1000 {
		t.Errorf("The given multiplier [%v] is invalid", value)
	}
}

func TestItCanConvertItSelfToRatePrice(t *testing.T) {
	currency := mock.Currency(t)
	currency.Rate = 100

	rate, err := currency.ToRatePrice()

	if err != nil {
		t.Errorf("%v", err)
	}

	if rate.ToString() != "$ 100.00" {
		t.Errorf("The given rate amount [100] is invalid")
	}
}

func TestItHoldsBasicInfo(t *testing.T) {
	currency := mock.Currency(t)

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
