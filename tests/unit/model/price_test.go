package model

import (
	"github.com/voyago/converter/src/model"
	"github.com/voyago/converter/tests/mock"
	"testing"
)

func TestItAllowsForBasicFormatting(t *testing.T) {
	t.Parallel()
	currency := mock.Currency(t) //SGD

	price, err := model.MakePrice(currency, "0.99")

	if err != nil {
		t.Errorf("%v", err)
		t.FailNow()
	}

	if price.ToString() != "$ 0.99" {
		t.Error("The given price formatting is invalid")
	}
}

func TestItAllowsCreationFromIntegers(t *testing.T) {
	t.Parallel()
	currency := mock.Currency(t)

	price, err := model.MakePrice(currency, 199)
	if err != nil || price.Amount != 199 || price.ToFloat() != 1.99 {
		t.Errorf("The given price [1.99] was not parsed properly")
	}

	price, err = model.MakePrice(currency, 1099)
	if err != nil || price.Amount != 1099 || price.ToFloat() != 10.99 {
		t.Errorf("The given price [10.99] was not parsed properly")
	}
}

func TestItAllowsCreationFromStrings(t *testing.T) {
	t.Parallel()
	currency := mock.Currency(t)

	price, err := model.MakePrice(currency, "0.99")
	if err != nil || price.Amount != 99 || price.ToFloat() != 0.99 {
		t.Errorf("The given price [0.99] was not parsed properly")
	}

	price, err = model.MakePrice(currency, "1.99")
	if err != nil || price.Amount != 199 || price.ToFloat() != 1.99 {
		t.Errorf("The given price [1.99] was not parsed properly")
	}

	price, err = model.MakePrice(currency, "10.99")
	if err != nil || price.Amount != 1099 || price.ToFloat() != 10.99 {
		t.Errorf("The given price [10.99] was not parsed properly")
	}
}

func TestItAllowsCreationFromFloaters(t *testing.T) {
	t.Parallel()
	currency := mock.Currency(t)

	price, err := model.MakePrice(currency, 1.99)
	if err != nil || price.Amount != 199 || price.ToFloat() != 1.99 {
		t.Errorf("The given price [1.99] was not parsed properly")
	}

	price, err = model.MakePrice(currency, 10.99)
	if err != nil || price.Amount != 1099 || price.ToFloat() != 10.99 {
		t.Errorf("The given price [10.99] was not parsed properly")
	}
}
