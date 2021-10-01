package currencyLayer

import (
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

func (current Mock) GetExchangeRates() (model.Currencies, error) {
	currencies := model.MakeCurrencies()
	currenciesBlueprint, err := blueprint.MakeCurrenciesBlueprint()

	if err != nil {
		return currencies, err
	}

	if err := current.FetchRates(); err != nil {
		return currencies, err
	}

	for _, value := range currenciesBlueprint.Items() {
		mapExchangeRates(current.response, value, &currencies)
	}

	return currencies, nil
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
