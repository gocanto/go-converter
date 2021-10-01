package store

import (
	"errors"
	"fmt"
	"github.com/voyago/converter/environment"
	"github.com/voyago/converter/pkg/model"
	"github.com/voyago/converter/pkg/store/handler"
	"github.com/voyago/converter/pkg/store/handler/currencyLayer"
)

type Store struct {
	source  string
	env     environment.Env
	handler handler.Handler
}

func Make(request Request) (*Store, error) {
	store := &Store{
		source: request.GetSource(),
		env:    request.GetEnv(),
	}

	if err := store.build(request.GetDriver()); err != nil {
		return nil, err
	}

	return store, nil
}

func (current Store) ExchangeRates() (model.Currencies, error) {
	return current.handler.ExchangeRates()
}

func (current Store) GetSource() string {
	return current.source
}

func (current Store) GetEnv() environment.Env {
	return current.env
}

func (current Store) GetHandler() handler.Handler {
	return current.handler
}

func (current *Store) build(driver string) error {
	switch driver {
	case CurrencyLayerDriverName:
		(*current).handler = current.currencyLayerHandler()
	default:
		return errors.New(fmt.Sprintf("The given driver [%s] is invalid", driver))
	}

	return nil
}

func (current Store) currencyLayerHandler() handler.Handler {
	if current.env.IsLive() {
		driver := currencyLayer.Handler{}
		driver.SetSource(current.source)
		driver.SetEnv(current.env)

		return driver
	}

	driver := currencyLayer.Mock{}
	driver.SetSource(current.source)
	driver.SetEnv(current.env)

	return driver
}
