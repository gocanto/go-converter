package store

import (
	"github.com/voyago/converter/environment"
	"github.com/voyago/converter/pkg/store"
	"github.com/voyago/converter/pkg/store/handler/currencyLayer"
	"testing"
)

func TestItExposesExchangesRates(t *testing.T) {
	t.Parallel()

	env, _ := environment.MakeWith("converter", ".env.example")
	manager, err := store.Make(store.NewCurrencyLayerRequest(env, "SGD"))

	if err != nil {
		t.Errorf("%v", err)
		t.FailNow()
	}

	if rates, err := manager.ExchangeRates(); rates.Count() != 179 || err != nil {
		t.Errorf("The [currency layer - mock] store exaches rates are invalid")
	}
}

func TestItProperlyBuildsTheCurrencyLayerDriver(t *testing.T) {
	t.Parallel()

	env, _ := environment.MakeWith("converter", ".env.example")
	manager, err := store.Make(store.NewCurrencyLayerRequest(env, "SGD"))

	if err != nil {
		t.Errorf("%v", err)
	}

	if manager.Env.IsLive() {
		t.Errorf("The given store environment is invalid.")
	}

	if manager.Source != "SGD" {
		t.Errorf("The given store base currency is invalid")
	}

	switch value := manager.Handler.(type) {
	case currencyLayer.Mock:
		if value.ApiKey() != "" {
			t.Errorf("The given key is invalid, %s", value.ApiKey())
		}

		if value.Source != "SGD" {
			t.Errorf("The given currency is invalid")
		}
	default:
		t.Errorf("The given store handler is invalid")
		t.FailNow()
	}
}
