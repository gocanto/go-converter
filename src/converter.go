package src

import (
	"github.com/voyago/converter/src/model"
)

type Converter struct {
	Price model.Price
}

func MakeConverter(price model.Price) Converter {
	return Converter{Price: price}
}

func (current Converter) ConvertTo(currency model.Currency) (model.Price, error) {
	rate, err := currency.ToRatePrice()

	if err != nil {
		return model.Price{}, err
	}

	target := float64(current.Price.Amount) / rate.ToFloat()
	price, _ := model.MakePrice(currency, target)

	return price, nil
}
