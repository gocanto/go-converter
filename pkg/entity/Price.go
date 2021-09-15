package entity

import "fmt"

type Price struct {
	Amount int64 `json:"amount"`
	Currency Currency
}

func (price Price) ToFloat() float64  {
	//currency := price.Currency

	fmt.Println("---->", price.Amount)

	return float64(price.Amount * 2 / 100)
}
