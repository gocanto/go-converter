package model

import (
	"github.com/voyago/converter/src/support"
	"runtime"
	"strings"
)

type CurrenciesBody struct {
	items []map[string]interface{}
}

func MakeCurrenciesBody() (CurrenciesBody, error) {
	body := CurrenciesBody{}
	var response []map[string]interface{}

	if err := support.ParseJson(body.GetSourcePath(), &response); err != nil {
		return CurrenciesBody{}, err
	}

	return CurrenciesBody{items: response}, nil
}

func (body CurrenciesBody) GetSourcePath() string {
	_, fileName, _, _ := runtime.Caller(0)
	baseDir := strings.Split(fileName, "/src/store/handler/currencyLayer/mock.go")[0]

	return baseDir + "/resources/currencies.json"
}

func (body CurrenciesBody) Items() []map[string]interface{} {
	return body.items
}
