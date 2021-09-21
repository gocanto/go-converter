package handler

import (
	"fmt"
	"github.com/voyago/converter/environment"
	"github.com/voyago/converter/src/store/handler/currencyLayer"
	"testing"
)

func TestName(t *testing.T) {
	env, err := environment.MakeWith("testing")

	if err != nil {
		t.Errorf("%v", err)
		t.FailNow()
	}

	driver := currencyLayer.Mock{Currency: "SGD", Env: *env}

	rates, _ := driver.ExchangeRates()

	fmt.Println(rates.Find("SGD"))
}
