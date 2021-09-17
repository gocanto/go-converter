package store

import (
	"errors"
	"fmt"
	"github.com/voyago/converter/src/store/handler"
)

type Store struct {
	Handler      handler.Handler
	baseCurrency string
}

func Make(driver string, baseCurrency string) (Store, error) {
	store := Store{
		baseCurrency: baseCurrency,
	}

	if store.doesntHaveDriver(driver) {
		return Store{}, errors.New(fmt.Sprintf("The given driver [%string] is invalid", driver))
	}

	if err := store.buildDriver(driver); err != nil {
		return Store{}, err
	}

	return store, nil
}

func (current Store) doesntHaveDriver(driver string) bool  {
	return !current.hasDriver(driver)
}

func (current Store) hasDriver(driver string) bool  {
	allowed := []string{
		"currencylayer",
	}

	for _, item := range allowed {
		if item == driver {
			return true
		}
	}

	return false
}

func (current *Store) buildDriver(driver string) error  {
	switch driver {
	case "currencylayer":
		(*current).Handler = handler.CurrencyLayer{ApiKey: ""}

		return nil
	}

	return errors.New(fmt.Sprintf("The given driver [%string] is invalid", driver))
}
