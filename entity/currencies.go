package converter

import (
	"errors"
	"fmt"
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

func (current *Currencies) Forget(abstract interface{}) error {
	target := ""
	items := *current.Items

	switch value := abstract.(type) {
		case string:
			target = value
		case Currency:
			target = value.Code
		default:
			return errors.New(fmt.Sprintf("The given abstract [%v] is invalid", abstract))
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
