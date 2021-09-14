package test

import (
	"github.com/voyago/converter/pkg/entity"
	"path/filepath"
	"runtime"
	"testing"
)

func Currency(t *testing.T) entity.Currency {
	currency := entity.Currency{}

	_, fileName, _, _ := runtime.Caller(0)
	fullPath := filepath.Dir(fileName) + "/__fixtures__/currency.json"

	if err := ParseJson(fullPath, &currency); err != nil {
		t.Errorf("There was an error mocking a new currency: %v", err)
		t.FailNow()
	}

	return currency
}
