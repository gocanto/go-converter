package currencyLayer

import (
	"encoding/json"
	"fmt"
	"github.com/voyago/converter/environment"
	"github.com/voyago/converter/pkg/model"
	"github.com/voyago/converter/pkg/store/blueprint"
	"io"
	"net/http"
)

const (
	ApiEndpoint = "http://api.currencylayer.com/live"
	ApiKeyName  = "CONVERTER_CURRENCY_LAYER_KEY"
)

type Handler struct {
	source   string
	env      environment.Env
	response Response
}

func (current Handler) ExchangeRates() (model.Currencies, error) {
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

func (current *Handler) FetchRates() error {
	response, err := http.Get(current.BaseUrl())

	if err != nil {
		return err
	}

	if err := current.parseResponse(response); err != nil {
		return err
	}

	return nil
}

func (current *Handler) parseResponse(response *http.Response) error {
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return err
	}

	if err = json.Unmarshal(body, &current.response); err != nil {
		return err
	}

	return nil
}

func (current Handler) BaseUrl() string {
	return ApiEndpoint + "?access_key=" + current.ApiKey() + "&source=" + current.source
}

func (current Handler) ApiKey() string {
	return current.env.Get(ApiKeyName)
}

func (current *Handler) SetSource(source string) {
	current.source = source
}

func (current Handler) GetSource() string {
	return current.source
}

func (current *Handler) SetEnv(env environment.Env) {
	current.env = env
}

func (current Handler) GetEnv() environment.Env {
	return current.env
}

func (current Handler) GetResponse() Response {
	return current.response
}
