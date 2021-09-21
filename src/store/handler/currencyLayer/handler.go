package currencyLayer

import (
	"github.com/voyago/converter/environment"
	"github.com/voyago/converter/src/model"
)

type Handler struct {
	Currency string
	Env environment.Env
}

func (current Handler) ExchangeRates() (model.Currencies, error) {
	items := make(map[string]model.Currency)

	return model.Currencies{Items: &items}, nil
}

func (current Handler) ApiKey() string  {
	return current.Env.Get("CONVERTER_CURRENCY_LAYER_KEY")
}
