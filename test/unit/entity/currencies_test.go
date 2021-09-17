package test_entity

import (
	"github.com/voyago/converter/src/model"
	"github.com/voyago/converter/test"
	"testing"
)

func TestItHoldsValidData(t *testing.T) {
	collection := mockCurrencies()

	if collection.Count() != 0 || len(collection.All()) != 0 {
		t.Errorf("The given currencies size is invalid")
	}
}

func TestItAddsCurrencies(t *testing.T) {
	collection := mockCurrencies()
	collection.Add(test.Currency(t))

	if collection.Count() != 1 || len(collection.All()) != 1 {
		t.Errorf("The given currencies size is invalid")
	}
}

func TestItFindsCurrencies(t *testing.T) {
	collection := mockCurrencies()

	if _, err := collection.Find("SGD"); err == nil {
		t.Errorf("The given collection found a missing [SGD] item")
	}

	currency := test.Currency(t)

	collection.Add(currency)
	byCode, _ := collection.Find("SGD")
	byEntity, _ := collection.Find(currency)

	if byCode.Code != "SGD" || byEntity.Code != "SGD" {
		t.Errorf("The given collection holds an invalid [SGD] item")
	}
}

func TestItRemovesCurrencies(t *testing.T) {
	collection := mockCurrencies()
	currency := test.Currency(t)

	collection.Add(currency)

	if collection.Count() != 1 || len(collection.All()) != 1 {
		t.Errorf("The given currencies size is invalid")
	}

	//by model
	if err := collection.Forget(currency); err != nil {
		t.Errorf("The given currency code [%s] could not be removed", currency.Code)
	}

	collection.Add(currency)

	//by code
	if err := collection.Forget(currency.Code); err != nil {
		t.Errorf("The given currency code [%s] could not be removed", currency.Code)
	}

	if collection.Count() != 0 || len(collection.All()) != 0 {
		t.Errorf("The given currencies size is invalid")
	}
}

func mockCurrencies() model.Currencies {
	items := make(map[string]model.Currency)

	return model.Currencies{Items: &items}
}
