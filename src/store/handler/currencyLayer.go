package handler

import (
	"github.com/voyago/converter/environment"
	"github.com/voyago/converter/src/model"
)

type CurrencyLayer struct {
	Currency string
	Env environment.Env
}

func (current CurrencyLayer) ExchangeRates() model.Currencies {
	items := make(map[string]model.Currency)

	return model.Currencies{Items: &items}
}

func (current CurrencyLayer) ApiKey() string  {
	return current.Env.Get("CONVERTER_CURRENCY_LAYER_KEY")
}
