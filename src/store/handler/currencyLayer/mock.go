package currencyLayer

import (
	"github.com/voyago/converter/environment"
	"github.com/voyago/converter/src/model"
)

type Mock struct {
	Currency string
	Env environment.Env
}

func (current Mock) ExchangeRates() model.Currencies {
	items := make(map[string]model.Currency)

	return model.Currencies{Items: &items}
}

func (current Mock) ApiKey() string  {
	return current.Env.Get("CONVERTER_CURRENCY_LAYER_KEY")
}
