package main

import (
	"fmt"
	"github.com/voyago/converter/pkg/entity"
)

func main() {
	usd := entity.Currency{
		Code:         "USD",
		Name:         "USD",
		Symbol:       "$",
		Rate:         1,
		IsoCode:      "1",
		IsoMinorUnit: 1,
	}

	sgd := entity.Currency{
		Code:         "SGD",
		Name:         "Singapore Dollar",
		Symbol:       "$",
		Rate:         2,
		IsoCode:      "1",
		IsoMinorUnit: 1,
	}

	collection := entity.Make()

	collection.Add(sgd)
	collection.Add(usd)

	if err := collection.Forget(usd); err != nil {
		fmt.Println("error", err)
		return
	}

	currency, err := collection.Find(sgd)

	if err != nil {
		fmt.Println("error", err)
		return
	}

	fmt.Println(currency)
	fmt.Println(collection.Count())
	fmt.Println(collection.All())
}
