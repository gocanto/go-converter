package conversion

import (
	"github.com/voyago/converter/environment"
	"github.com/voyago/converter/pkg/conversion"
	"github.com/voyago/converter/pkg/model"
	"github.com/voyago/converter/pkg/store"
	"github.com/voyago/converter/tests/mock"
	"testing"
)

func TestItConvertsFromSgdToUsd(t *testing.T) {
	t.Parallel()

	//1 SGD to USD = 0.74
	//Exchange rate = 1/0.74
	sgd, usd := createCurrencies(t)

	price, _ := model.MakePrice(sgd, 1)
	result, err := conversion.ConvertTo(price, usd)

	if err != nil {
		t.Errorf("%v", err)
	}

	if result.ToFloat() != 0.74 {
		t.Errorf("The given [SGD to USD] conversion is invalid")
	}

	if result.ToString() != "USD 0.74" {
		t.Errorf("The given [SGD to USD] format is invalid")
	}
}

func TestItConvertsFromUsdToSgd(t *testing.T) {
	t.Parallel()

	//1 USD to SGD = 1.34
	//Exchange rate = 1/1.34
	sgd, usd := createCurrencies(t)
	usd.Rate = 1
	sgd.Rate = 0.7462

	price, _ := model.MakePrice(usd, 1)
	result, err := conversion.ConvertTo(price, sgd)

	if err != nil {
		t.Errorf("%v", err)
	}

	if result.ToFloat() != 1.35 {
		t.Errorf("The given [USD to SGD] conversion is invalid")
	}

	if result.ToString() != "SGD 1.35" {
		t.Errorf("The given [USD to SGD] format is invalid")
	}
}

func TestItCanConvertValuesByUsingTheStore(t *testing.T) {
	env, _ := environment.MakeWith("converter", ".env.example")
	manager, _ := store.Make(store.NewCurrencyLayerRequest(env, "SGD"))

	converter := conversion.Make(*manager)

	_, usd := createCurrencies(t)
	usd.Rate = 1

	price, _ := model.MakePrice(usd, 1)
	result, err := converter.Convert(price)

	if err != nil {
		t.Errorf("%v", err)
	}

	if result.ToFloat() != 0.74 {
		t.Errorf("The given [SGD to USD] conversion is invalid")
	}

	if result.ToString() != "SGD 0.74" {
		t.Errorf("The given [USD to SGD] format is invalid")
	}
}

func createCurrencies(t *testing.T) (model.Currency, model.Currency) {
	sgd := mock.Currency(t)
	sgd.Code = "SGD"
	sgd.Rate = 1
	sgd.IsoMinorUnit = 2

	usd := mock.Currency(t)
	usd.Code = "USD"
	usd.Rate = 1.34
	usd.IsoMinorUnit = 2
	usd.IsoCode = 840

	return sgd, usd
}
