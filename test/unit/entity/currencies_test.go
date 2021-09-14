package test_entity

import (
	"github.com/voyago/converter/pkg/entity"
	testSupport "github.com/voyago/converter/test/support"
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
	collection.Add(mockCurrency(t))

	if collection.Count() != 1 || len(collection.All()) != 1 {
		t.Errorf("The given currencies size is invalid")
	}
}

func TestItFindsCurrencies(t *testing.T) {
	collection := mockCurrencies()

	if _, err := collection.Find("SGD"); err == nil {
		t.Errorf("The given collection found a missing [SGD] item")
	}

	currency := mockCurrency(t)

	collection.Add(currency)
	byCode, _ := collection.Find("SGD")
	byEntity, _ := collection.Find(currency)

	if byCode.Code != "SGD" || byEntity.Code != "SGD" {
		t.Errorf("The given collection holds an invalid [SGD] item")
	}
}

func TestItRemovesCurrencies(t *testing.T) {
	collection := mockCurrencies()
	currency := mockCurrency(t)

	collection.Add(currency)

	if collection.Count() != 1 || len(collection.All()) != 1 {
		t.Errorf("The given currencies size is invalid")
	}

	//by entity
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

func mockCurrencies() entity.Currencies {
	items := make(map[string]entity.Currency)

	return entity.Currencies{Items: &items}
}

func mockCurrency(t *testing.T) entity.Currency {
	currency := entity.Currency{}

	if err := testSupport.ParseJson(stub, &currency); err != nil {
		t.Errorf("There was an error [%d] opening the given currency stub", err)
	}

	return currency
}
