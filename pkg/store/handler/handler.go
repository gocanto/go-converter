package handler

import "github.com/voyago/converter/pkg/model"

type Handler interface {
	GetExchangeRates() (model.Currencies, error)
}
