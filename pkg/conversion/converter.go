package conversion

import (
	"github.com/voyago/converter/pkg/contracts"
	"github.com/voyago/converter/pkg/entity"
	store "github.com/voyago/converter/pkg/store/repository"
)

type Converter struct {
	Price      entity.Price
	repository contracts.CurrenciesRepository
}

func MakeConverter(price entity.Price) Converter {
	return Converter{Price: price, repository: store.DBRepository{}}
}

// WithRepository @todo check whether we need this at all
func (current *Converter) WithRepository(repository contracts.CurrenciesRepository) {
	(*current).repository = repository
}

func (current *Converter) WithPrice(price entity.Price) {
	(*current).Price = price
}

func (current Converter) ConvertTo(currency entity.Currency) (entity.Price, error) {
	rate, err := currency.ToRatePrice()

	if err != nil {
		return entity.Price{}, err
	}

	target := float64(current.Price.Amount) / rate.ToFloat()
	price, _ := entity.MakePrice(currency, target)

	return price, nil
}
