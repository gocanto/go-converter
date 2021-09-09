package main

import (
	converter "converter/entity"
	"fmt"
)

func main()  {
	usd := converter.Currency{
		Code: "USD",
		Name: "USD",
		Symbol: "$",
		IsoCode: "1",
		IsoMinorUnit: 1,
	}

	sgd := converter.Currency{
		Code: "SGD",
		Name: "Singapore Dollar",
		Symbol: "$",
		IsoCode: "1",
		IsoMinorUnit: 1,
	}

	collection := converter.Make()

	collection.Add(sgd)
	collection.Add(usd)

	//collection.RemoveByCode("USD")
	err := collection.Remove(usd)

	if err != nil {
		fmt.Println("error", err)
		return 
	}

	fmt.Println(collection.Count())
	fmt.Println(collection.All())
}
