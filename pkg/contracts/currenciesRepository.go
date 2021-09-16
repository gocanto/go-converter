package contracts

import "github.com/voyago/converter/pkg/entity"

type CurrenciesRepository interface {
	FetchCurrency(interface{}) entity.Currency
}
