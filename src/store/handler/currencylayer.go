package handler

import "github.com/voyago/converter/src/model"

type CurrencyLayer struct {
	ApiKey string
	Currency string
}

func (handler CurrencyLayer) ExchangeRates() model.Currencies {
	items := make(map[string]model.Currency)

	return model.Currencies{Items: &items}
}
