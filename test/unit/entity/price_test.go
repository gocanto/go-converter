package test_entity

import (
	"fmt"
	"github.com/voyago/converter/test"
	"testing"
)

func TestTesting(t *testing.T) {
	currency := test.Currency(t)
	//price := entity.Price{Amount: 1099, Currency: currency} //SGD 10.99
	//
	//fmt.Println("Result: ", price.ToFloat())

	cost := 1.99 * 100
	stringValue := "100"

	var amount float32 = 57

	result, _ := fmt.Printf("%.2f\n", amount)
	fmt.Printf("%T\n", result)

	fmt.Println("cost", cost, "---", cost / 100)
	fmt.Println("string value", stringValue, "---", fmt.Sprint(stringValue))

	fmt.Println("--------------------------------------------------------------------")

	//value := 1
	fmt.Println(currency.GetMultiplier())
}
