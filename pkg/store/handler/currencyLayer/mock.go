package currencyLayer

import (
	"fmt"
	"github.com/voyago/converter/environment"
	"github.com/voyago/converter/pkg/model"
	"github.com/voyago/converter/pkg/store/blueprint"
	"github.com/voyago/converter/pkg/support"
	"runtime"
	"strings"
)

type Mock struct {
	source   string
	env      environment.Env
	response Response
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
			Code:         code,
			Name:         fmt.Sprintf("%s", value["name"]),
			IsoMinorUnit: int8(value["iso_minor_unit"].(float64)),
			IsoCode:      int16(value["iso_code"].(float64)),
			Rate:         float32(current.response.RateFor(code)),
		})
	}

	return collection, nil
}

func (current *Mock) FetchRates() error {
	_, fileName, _, _ := runtime.Caller(0)

	dir := strings.Split(fileName, "/pkg/store/handler/currencyLayer/mock.go")[0]
	payload := dir + "/resources/currency_layer_live_response.json"

	if err := support.ParseJson(payload, &current.response); err != nil {
		return err
	}

	return nil
}

// ApiKey @tod remove if not needed
func (current Mock) ApiKey() string {
	return ""
}

func (current *Mock) SetSource(source string) {
	current.source = source
}

func (current Mock) GetSource() string {
	return current.source
}

func (current *Mock) SetEnv(env environment.Env) {
	current.env = env
}

func (current Mock) GetEnv() environment.Env {
	return current.env
}

func (current Mock) GetResponse() Response {
	return current.response
}
