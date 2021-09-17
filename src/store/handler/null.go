package handler

import "github.com/voyago/converter/src/model"

type Null struct {

}

func (receiver Null) ExchangeRates() model.Currencies  {
	return model.MakeCurrencies()
}
