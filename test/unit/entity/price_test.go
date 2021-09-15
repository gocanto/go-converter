package test_entity

import (
	"github.com/voyago/converter/pkg/entity"
	"github.com/voyago/converter/test"
	"testing"
)

func TestItAllowsForBasicFormatting(t *testing.T) {
	currency := test.Currency(t) //SGD

	price, err := entity.MakePrice(currency, "0.99")

	if err != nil {
		t.Errorf("%v", err)
		t.FailNow()
	}

	if price.ToString() != "$ 0.99" {
		t.Error("The given price formatting is invalid")
	}
}

func TestItAllowsCreationFromIntegers(t *testing.T) {
	currency := test.Currency(t)

	price, err := entity.MakePrice(currency, 199)
	if err != nil || price.Amount != 199 || price.ToFloat() != 1.99 {
		t.Errorf("The given price [1.99] was not parsed properly")
	}

	price, err = entity.MakePrice(currency, 1099)
	if err != nil || price.Amount != 1099 || price.ToFloat() != 10.99 {
		t.Errorf("The given price [10.99] was not parsed properly")
	}
}

func TestItAllowsCreationFromStrings(t *testing.T) {
	currency := test.Currency(t)

	price, err := entity.MakePrice(currency, "0.99")
	if err != nil || price.Amount != 99 || price.ToFloat() != 0.99 {
		t.Errorf("The given price [0.99] was not parsed properly")
	}

	price, err = entity.MakePrice(currency, "1.99")
	if err != nil || price.Amount != 199 || price.ToFloat() != 1.99 {
		t.Errorf("The given price [1.99] was not parsed properly")
	}

	price, err = entity.MakePrice(currency, "10.99")
	if err != nil || price.Amount != 1099 || price.ToFloat() != 10.99 {
		t.Errorf("The given price [10.99] was not parsed properly")
	}
}

func TestItAllowsCreationFromFloaters(t *testing.T) {
	currency := test.Currency(t)

	price, err := entity.MakePrice(currency, 1.99)
	if err != nil || price.Amount != 199 || price.ToFloat() != 1.99 {
		t.Errorf("The given price [1.99] was not parsed properly")
	}

	price, err = entity.MakePrice(currency, 10.99)
	if err != nil || price.Amount != 1099 || price.ToFloat() != 10.99 {
		t.Errorf("The given price [10.99] was not parsed properly")
	}
}
