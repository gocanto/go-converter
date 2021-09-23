package conversion

import (
	"github.com/voyago/converter/pkg/model"
)

func Convert(price model.Price, currency model.Currency) (model.Price, error)  {
	rate, err := currency.ToRatePrice()

	if err != nil {
		return model.Price{}, err
	}

	target := float64(price.Amount) / rate.ToFloat()
	result, _ := model.MakePrice(currency, target)

	return result, nil
}
