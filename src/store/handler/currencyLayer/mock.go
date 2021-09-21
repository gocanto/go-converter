package currencyLayer

import (
	"fmt"
	"github.com/voyago/converter/environment"
	"github.com/voyago/converter/src/model"
	"github.com/voyago/converter/src/support"
	"runtime"
	"strings"
)

type Mock struct {
	Currency string
	Env environment.Env
	Response Response
}

func (current Mock) ExchangeRates() (model.Currencies, error) {
	collection := model.MakeCurrencies()
	currenciesBody, err := model.MakeCurrenciesBody()

	if err != nil {
		return collection, err
	}

	response, err := fetchRates()

	if err != nil {
		return collection, err
	}

	for _, value := range currenciesBody.Items() {
		code := fmt.Sprintf("%s", value["code"])

		collection.Add(model.Currency{
			Code: code,
			Name: fmt.Sprintf("%s", value["name"]),
			IsoMinorUnit: int(value["iso_minor_unit"].(float64)),
			IsoCode: int(value["iso_code"].(float64)),
			Symbol: "",
			Rate: float32(response.FindRateFor(code)),
		})
	}

	return collection, nil
}

func (current Mock) ApiKey() string  {
	return current.Env.Get("CONVERTER_CURRENCY_LAYER_KEY")
}

func fetchRates() (Response, error) {
	_, fileName, _, _ := runtime.Caller(0)

	baseDir := strings.Split(fileName, "/src/store/handler/currencyLayer/mock.go")[0]
	stubPath := baseDir + "/resources/currency_layer_live_response.json"

	response := Response{}
	if err := support.ParseJson(stubPath, &response); err != nil {
		return Response{}, err
	}

	return response, nil
}
