package store

import (
	"github.com/voyago/converter/src/store/handler"
)

type Store struct {
	Handler      handler.Handler
	baseCurrency string
}

func Make(handler string, baseCurrency string) (Store, error) {
	manager := Driver{Id: handler}
	driver, err := manager.Build()

	if err != nil {
		return Store{}, err
	}

	return Store{Handler: driver, baseCurrency: baseCurrency}, nil
}
