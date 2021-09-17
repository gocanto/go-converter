package test

import (
	"github.com/voyago/converter/pkg/model"
	"path/filepath"
	"runtime"
	"testing"
)

func Currency(t *testing.T) model.Currency {
	currency := model.Currency{}

	_, fileName, _, _ := runtime.Caller(0)
	fullPath := filepath.Dir(fileName) + "/__fixtures__/currency.json"

	if err := ParseJson(fullPath, &currency); err != nil {
		t.Errorf("There was an error mocking a new currency: %v", err)
		t.FailNow()
	}

	return currency
}
