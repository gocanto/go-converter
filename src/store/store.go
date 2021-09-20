package store

import (
	"errors"
	"fmt"
	lib "github.com/voyago/converter/src"
	"github.com/voyago/converter/src/store/handler"
)

type Store struct {
	Handler  handler.Handler
	Currency string
}

func MakeStore(driver string, currency string) (*Store, error) {
	store := &Store{Currency: currency}

	if err := store.build(driver); err != nil {
		return nil, err
	}

	return store, nil
}

func (current *Store) build(driver string) error {
	switch driver {
	case "currencylayer":
		(*current).Handler = handler.CurrencyLayer{
			ApiKey:   fmt.Sprintf("%s", lib.Env("CURRENCY_LAYER_KEY")),
			Currency: current.Currency,
		}
	default:
		return errors.New(fmt.Sprintf("The given driver [%s] is invalid", driver))
	}

	return nil
}
