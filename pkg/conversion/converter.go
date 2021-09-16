package conversion

import (
	"github.com/voyago/converter/pkg/entity"
)

type Converter struct {
	Price      entity.Price
}

func MakeConverter(price entity.Price) Converter {
	return Converter{Price: price}
}

func (current Converter) ConvertTo(currency entity.Currency) (entity.Price, error) {
	rate, err := currency.ToRatePrice()

	if err != nil {
		return entity.Price{}, err
	}

	target := float64(current.Price.Amount) / rate.ToFloat()
	price, _ := entity.MakePrice(currency, target)

	return price, nil
}
