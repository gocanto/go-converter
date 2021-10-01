package currencyLayer

import (
	"fmt"
	"github.com/voyago/converter/pkg/model"
)

func mapExchangeRates(response Response, value map[string]interface{}, currencies *model.Currencies) {
	code := fmt.Sprintf("%s", value["code"])

	rate, err := response.RateFor(code)

	if err != nil {
		return
	}

	currencies.Add(model.Currency{
		Code:         code,
		Name:         fmt.Sprintf("%s", value["name"]),
		IsoMinorUnit: int8(value["iso_minor_unit"].(float64)),
		IsoCode:      int16(value["iso_code"].(float64)),
		Rate:         float32(rate),
	})
}
