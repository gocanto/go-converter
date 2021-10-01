package handler

import (
	"github.com/voyago/converter/environment"
	"github.com/voyago/converter/pkg/store/handler/currencyLayer"
	"testing"
)

func TestItProperlyFindsRates(t *testing.T) {
	t.Parallel()

	env, err := environment.MakeWith("converter", ".env.example")

	if err != nil {
		t.Errorf("%v", err)
		t.FailNow()
	}

	driver := currencyLayer.Mock{}
	driver.SetSource("USD")
	driver.SetEnv(env)

	rates, err := driver.GetExchangeRates()

	if err != nil {
		t.Errorf("%v", err)
		t.FailNow()
	}

	currency, _ := rates.Find("SGD")

	if currency.Code != "SGD" {
		t.Errorf("The given currency code [%s] is invalid", currency.Code)
	}

	if currency.Rate != 1.34258 {
		t.Errorf("The given currency [%s] rate [%v] is invalid", currency.Code, currency.Rate)
	}
}
