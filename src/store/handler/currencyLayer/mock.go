package currencyLayer

import (
	"fmt"
	"github.com/voyago/converter/environment"
	"github.com/voyago/converter/src/model"
)

type Mock struct {
	Currency string
	Env environment.Env
}

func (current Mock) ExchangeRates() (model.Currencies, error) {
	collection := model.MakeCurrencies()
	response, err := model.MakeCurrenciesBody()

	if err != nil {
		return collection, err
	}

	for _, value := range response.Items() {
		collection.Add(model.Currency{
			Code: fmt.Sprintf("%s", value["code"]),
			Name: fmt.Sprintf("%s", value["name"]),
			IsoMinorUnit: int(value["iso_minor_unit"].(float64)),
			IsoCode: int(value["iso_code"].(float64)),
			Symbol: "",
			Rate: 1.0,
		})
	}

	return collection, nil
}

func (current Mock) ApiKey() string  {
	return current.Env.Get("CONVERTER_CURRENCY_LAYER_KEY")
}
