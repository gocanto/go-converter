package converter

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

type Currencies struct {
	Items *map[string]Currency
}

func Make() Currencies  {
	items := make(map[string]Currency)

	return Currencies{Items: &items}
}

func (current *Currencies) Add(currency Currency)  {
	(*current.Items)[currency.Code] = currency
}

func (current *Currencies) Remove(target interface{}) (error error) {
	items := *current.Items
	scope := reflect.TypeOf(target).String()

	if scope == "string" {
		delete(items, fmt.Sprintf("%v", target))
		current.Items = &items

		return nil
	}

	currency, err := resolveCurrencyFrom(target)

	if err != nil {
		return err
	}

	delete(items, currency.Code)
	current.Items = &items

	return nil
}

func (current *Currencies) All() map[string]Currency {
	return *current.Items
}

func (current *Currencies) Count() int {
	return len(*current.Items)
}

func resolveCurrencyFrom(input interface{}) (currency Currency, err error) {
	target := Currency{}
	data, _ := json.Marshal(input)
	err = json.Unmarshal(data, &target)

	if err != nil {
		return  Currency{}, err
	}

	if target.Code == "" {
		return Currency{}, errors.New("the given currency is invalid")
	}

	return target, nil
}
