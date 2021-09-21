package handler

import "github.com/voyago/converter/src/model"

type Handler interface {
	ExchangeRates() (model.Currencies, error)
}
