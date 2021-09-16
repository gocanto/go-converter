package test_conversion

import (
	"fmt"
	"github.com/voyago/converter/pkg/entity"
	"github.com/voyago/converter/test"
	"testing"
)

func TestConverter(t *testing.T) {

	//1 SGD to USD = 0.74
	//Exchange rate = 1/0.74
	sgd, usd := createCurrencies(t)
	price, _ := entity.MakePrice(sgd, 1)
	ratePrice, _ := usd.ToRatePrice()

	targetBaseAmount, _ := entity.MakePrice(usd, float64(price.Amount) / ratePrice.ToFloat())

	fmt.Println(price.Amount, "|", ratePrice.Amount, "|", targetBaseAmount.ToFloat())
	fmt.Println("====================================================================================================")

	//1 USD to SGD = 1.34
	//Exchange rate = 1/1.34
	usd.Rate = 1
	sgd.Rate = 0.7462

	price, _ = entity.MakePrice(usd, 1)
	ratePrice, _ = usd.ToRatePrice()

	targetBaseAmount, _ = entity.MakePrice(sgd, float64(price.Amount) / ratePrice.ToFloat())

	fmt.Println(price.Amount, "|", ratePrice.Amount, "|", targetBaseAmount.ToFloat())
}

func createCurrencies(t *testing.T) (entity.Currency, entity.Currency) {
	sgd := test.Currency(t)
	sgd.Code = "SGD"
	sgd.Rate = 1
	sgd.IsoMinorUnit = 2

	usd := test.Currency(t)
	usd.Code = "USD"
	usd.Rate = 1.34
	usd.IsoMinorUnit = 2
	usd.IsoCode = 840

	return sgd, usd
}
