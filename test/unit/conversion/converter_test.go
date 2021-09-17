package test_conversion

import (
	"github.com/voyago/converter/pkg"
	"github.com/voyago/converter/pkg/model"
	"github.com/voyago/converter/test"
	"testing"
)

func TestItConvertsFromSgdToUsd(t *testing.T) {
	//1 SGD to USD = 0.74
	//Exchange rate = 1/0.74

	sgd, usd := createCurrencies(t)
	price, _ := model.MakePrice(sgd, 1)
	converter := pkg.MakeConverter(price)

	result, err := converter.ConvertTo(usd)

	if err != nil {
		t.Errorf("%v", err)
	}

	if result.ToFloat() != 0.74 {
		t.Errorf("The given [SGD to USD] conversion is invalid")
	}

	if result.ToString() != "$ 0.74" {
		t.Errorf("The given [SGD to USD] format is invalid")
	}
}

func TestItConvertsFromUsdToSgd(t *testing.T) {
	//1 USD to SGD = 1.34
	//Exchange rate = 1/1.34
	sgd, usd := createCurrencies(t)
	usd.Rate = 1
	sgd.Rate = 0.7462

	price, _ := model.MakePrice(usd, 1)
	converter := pkg.MakeConverter(price)

	result, err := converter.ConvertTo(sgd)

	if err != nil {
		t.Errorf("%v", err)
	}

	if result.ToFloat() != 1.35 {
		t.Errorf("The given [USD to SGD] conversion is invalid")
	}

	if result.ToString() != "$ 1.35" {
		t.Errorf("The given [USD to SGD] format is invalid")
	}
}

func createCurrencies(t *testing.T) (model.Currency, model.Currency) {
	sgd := test.Currency(t)
	sgd.Code = "SGD"
	sgd.Rate = 1
	sgd.IsoMinorUnit = 2

	usd := test.Currency(t)
	usd.Code = "USD"
	usd.Rate = 1.34
	usd.IsoMinorUnit = 2
	usd.IsoCode = 840

	return sgd, usd
}
