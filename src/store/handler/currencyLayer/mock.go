package currencyLayer

import (
	"fmt"
	"github.com/voyago/converter/environment"
	"github.com/voyago/converter/src/model"
	"github.com/voyago/converter/src/store/blueprint"
	"github.com/voyago/converter/src/support"
	"runtime"
	"strings"
)

type Mock struct {
	Source string
	Env    environment.Env
	Response Response
}

func (current Mock) ExchangeRates() (model.Currencies, error) {
	collection := model.MakeCurrencies()
	currenciesBlueprint, err := blueprint.MakeCurrenciesBlueprint()

	if err != nil {
		return collection, err
	}

	if err := current.FetchRates(); err != nil {
		return collection, err
	}

	for _, value := range currenciesBlueprint.Items() {
		code := fmt.Sprintf("%s", value["code"])

		collection.Add(model.Currency{
			Code: code,
			Name: fmt.Sprintf("%s", value["name"]),
			IsoMinorUnit: int8(value["iso_minor_unit"].(float64)),
			IsoCode: int16(value["iso_code"].(float64)),
			Rate: float32(current.Response.RateFor(code)),
		})
	}

	return collection, nil
}

func (current Mock) ApiKey() string  {
	return ""
}

func (current *Mock) FetchRates() error {
	_, fileName, _, _ := runtime.Caller(0)

	dir := strings.Split(fileName, "/src/store/handler/currencyLayer/mock.go")[0]
	payload := dir + "/resources/currency_layer_live_response.json"

	if err := support.ParseJson(payload, &current.Response); err != nil {
		return err
	}

	return nil
}
