package main

import (
	converter "converter/entity"
	"fmt"
	"reflect"
)

func main()  {
	currency := converter.Currency{
		Code: "SGD",
		Name: "Singapore Dollar",
		Symbol: "$",
		IsoCode: "1",
		IsoMinorUnit: 1,
	}

	//if reflect.Type(converter.Currency{}) {
	//	fmt.Println("YES")
	//}

	name := "Gus"

	fmt.Println(reflect.TypeOf(currency))
	fmt.Println(reflect.TypeOf(name))
	fmt.Println(reflect.TypeOf(currency))
}
