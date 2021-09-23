package store

import (
	"errors"
	"fmt"
	"github.com/voyago/converter/environment"
	"github.com/voyago/converter/pkg/store/handler"
	"github.com/voyago/converter/pkg/store/handler/currencyLayer"
)

type Store struct {
	Currency string
	Env      environment.Env
	Handler  handler.Handler
}

func Make(driver string, currency string) (*Store, error) {
	env, err := environment.Make()

	if err != nil {
		return nil, errors.New("the store was unable to load the environment information")
	}

	return resolve(env, driver, currency)
}

func Mock(driver string, currency string) (*Store, error) {
	env, err := environment.MakeWith("testing")

	if err != nil {
		return nil, errors.New("the store was unable to load the environment information")
	}

	return resolve(env, driver, currency)
}

func (current *Store) build(driver string) error {
	switch driver {
	case "currency-layer":
		(*current).Handler = current.CurrencyLayerHandler()
	default:
		return errors.New(fmt.Sprintf("The given driver [%s] is invalid", driver))
	}

	return nil
}

func resolve(env *environment.Env, driver string, currency string) (*Store, error) {
	store := &Store{
		Currency: currency,
		Env: *env,
	}

	if err := store.build(driver); err != nil {
		return nil, err
	}

	return store, nil
}

func (current Store) CurrencyLayerHandler() handler.Handler {
	if current.Env.IsLive() {
		return currencyLayer.Handler{Source: current.Currency, Env: current.Env}
	}

	return currencyLayer.Mock{Source: current.Currency, Env: current.Env}
}
