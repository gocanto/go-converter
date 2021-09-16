package store_repository

import "github.com/voyago/converter/pkg/entity"

type DBRepository struct {

}

func (repository DBRepository) FetchCurrency(currency interface{}) entity.Currency {
	return entity.Currency{}
}
