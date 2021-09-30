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
	Source  string
	Env     environment.Env
	Handler handler.Handler
}

func Make(request Request) (*Store, error) {
	store := &Store{
		Source: request.GetSource(),
		Env:    request.GetEnv(),
	}

	if err := store.build(request.GetDriver()); err != nil {
		return nil, err
	}

	return store, nil
}

func (current Store) ExchangeRates() (model.Currencies, error) {
	return current.Handler.ExchangeRates()
}

func (current *Store) build(driver string) error {
	switch driver {
	case CurrencyLayerDriverName:
		(*current).Handler = current.currencyLayerHandler()
	default:
		return errors.New(fmt.Sprintf("The given driver [%s] is invalid", driver))
	}

	return nil
}

func (current Store) currencyLayerHandler() handler.Handler {
	if current.Env.IsLive() {
		return currencyLayer.Handler{Source: current.Source, Env: current.Env}
	}

	return currencyLayer.Mock{Source: current.Source, Env: current.Env}
}
