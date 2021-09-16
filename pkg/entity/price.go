package entity

import (
	"errors"
	"fmt"
	"strconv"
)

type Price struct {
	Amount int64 `json:"amount"`
	Currency Currency `json:"currency"`
}

func (current Price) ToFloat() float64  {
	amount := float64(current.Amount)
	multiplier := float64(current.Currency.GetMultiplier())

	return amount / multiplier
}

func (current Price) ToString() string {
	return fmt.Sprintf("%s %.2f", current.Currency.Symbol, current.ToFloat())
}

func MakePrice(currency Currency, abstract interface{}) (Price, error)  {
	switch value := abstract.(type) {
		case int64:
			return Price{Amount: value, Currency: currency}, nil
		case int32:
			return Price{Amount: int64(value), Currency: currency}, nil
		case int8:
			return Price{Amount: int64(value), Currency: currency}, nil
		case int:
			return Price{Amount: int64(value), Currency: currency}, nil
		case float32:
			return fromFloat32(currency, value), nil
		case float64:
			return fromFloat64(currency, value), nil
		case string:
			if price, err := fromString(currency, value); err != nil {
				return Price{}, err
			} else {
				return price, nil
			}
		default:
			return Price{},  errors.New(fmt.Sprintf("The given abstract [%v] is invalid", abstract))
	}
}

func fromString(currency Currency, value string) (Price, error) {
	target, err := strconv.ParseFloat(value, 64)

	if err != nil {
		return Price{}, err
	}

	return fromFloat64(currency, target), nil
}

func fromFloat32(currency Currency, target float32) Price {
	amount := int64(
		target * float32(currency.GetMultiplier()),
	)

	return Price{Amount: amount, Currency: currency}
}

func fromFloat64(currency Currency, target float64) Price {
	amount := int64(
		target * float64(currency.GetMultiplier()),
	)

	return Price{Amount: amount, Currency: currency}
}
