package handler

import "github.com/voyago/converter/pkg/model"

type Handler interface {
	ExchangeRates() (model.Currencies, error)
}
