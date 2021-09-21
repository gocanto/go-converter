package store

import (
	"github.com/voyago/converter/src/store"
	"github.com/voyago/converter/src/store/handler"
	"testing"
)

func TestItReturnsNilForInvalidDrivers(t *testing.T) {
	t.Parallel()

	if manager, err := store.Make("foo", "bar"); manager != nil {
		t.Errorf("%v", err)
	}
}

func TestItProperlyBuildsTheCurrencyLayerDriver(t *testing.T) {
	t.Parallel()

	manager, err := store.Mock("currency-layer", "SGD")

	if err != nil {
		t.Errorf("%v", err)
	}

	if manager.Env.IsLive() {
		t.Errorf("The given store environment is invalid.")
	}

	if manager.Currency != "SGD" {
		t.Errorf("The given store base currency is invalid")
	}

	assert := func(current handler.CurrencyLayer) {
		if current.ApiKey() != "" {
			t.Errorf("The given key is invalid, %s", current.ApiKey())
		}

		if current.Currency != "SGD" {
			t.Errorf("The given currency is invalid")
		}
	}

	switch value := manager.Handler.(type) {
	case handler.CurrencyLayer:
		assert(value)
	default:
		t.Errorf("The given store handler is invalid")
		t.FailNow()
	}
}
