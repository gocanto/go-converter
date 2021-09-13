package test_entity

import (
	"converter/entity"
	"testing"
)

func TestCurrency(t *testing.T) {
	currency := entity.Currency{
		Code:         "SGD",
		Name:         "Singapore Dollar",
		Symbol:       "$",
		Rate:         2,
		IsoCode:      "1",
		IsoMinorUnit: 1,
	}

	if currency.Code != "SGD" {
		t.Errorf("The given currency code [%s] is invalid", currency.Code)
	}
}
