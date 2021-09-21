package mock

import (
	"github.com/voyago/converter/src/model"
	"github.com/voyago/converter/src/support"
	"runtime"
	"strings"
	"testing"
)

func Currency(t *testing.T) model.Currency {
	currency := model.Currency{}
	_, fileName, _, _ := runtime.Caller(0)

	baseDir := strings.Split(fileName, "/mock/mock.go")[0]
	fullPath := baseDir + "/__fixtures__/currency.json"

	if err := support.ParseJson(fullPath, &currency); err != nil {
		t.Errorf("There was an error mocking a new currency: %v", err)
		t.FailNow()
	}

	return currency
}
