package handler

import (
	"github.com/voyago/converter/environment"
	"github.com/voyago/converter/pkg/store/handler/currencyLayer"
	"testing"
)

func TestItProperlyFetchesTheRatesFromTheAPI(t *testing.T) {
	t.Parallel()

	t.Skipf("The [currency layer] feature test was skipped")
	t.SkipNow()

	env, err := environment.Make()

	if err != nil {
		t.Errorf("%v", err)
		t.FailNow()
	}

	driver := currencyLayer.Handler{Source: "SGD", Env: *env}

	rates, err := driver.ExchangeRates()

	if err != nil {
		t.Errorf("%v", err)
		t.FailNow()
	}

	currency, _ := rates.Find("SGD")

	if currency.Code != "SGD" {
		t.Errorf("The given currency code [%s] is invalid", currency.Code)
	}

	if currency.Rate > 1 {
		t.Errorf("The given currency [%s] rate [%v] is invalid", currency.Code, currency.Rate)
	}
}
