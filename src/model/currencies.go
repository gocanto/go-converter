package model

import (
	"errors"
	"fmt"
)

type Currencies struct {
	Items *map[string]Currency
}

func Make() Currencies {
	items := make(map[string]Currency)

	return Currencies{Items: &items}
}

func (current *Currencies) Add(currency Currency) {
	(*current.Items)[currency.Code] = currency
}

func (current *Currencies) Find(abstract interface{}) (Currency, error) {
	items := *current.Items
	target, err := resolveCurrencyCode(abstract)

	if err != nil {
		return Currency{}, err
	}

	if currency, exists := items[target]; exists {
		return currency, nil
	}

	return Currency{}, errors.New(fmt.Sprintf("The given target [%v] does not exist", target))
}

func (current *Currencies) Forget(abstract interface{}) error {
	items := *current.Items
	target, err := resolveCurrencyCode(abstract)

	if err != nil {
		return err
	}

	delete(items, target)
	current.Items = &items

	return nil
}

func (current *Currencies) All() map[string]Currency {
	return *current.Items
}

func (current *Currencies) Count() int {
	return len(*current.Items)
}

func resolveCurrencyCode(abstract interface{}) (string, error) {
	switch value := abstract.(type) {
	case string:
		return value, nil
	case Currency:
		return value.Code, nil
	default:
		return "", errors.New(fmt.Sprintf("The given abstract [%v] is invalid", abstract))
	}
}
