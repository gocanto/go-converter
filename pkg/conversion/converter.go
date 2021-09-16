package conversion

import (
	"fmt"
	"github.com/voyago/converter/pkg/contracts"
	"github.com/voyago/converter/pkg/entity"
)

type Converter struct {
	Amount entity.Price
	repository contracts.CurrenciesRepository
}

func (current *Converter) Using (repository contracts.CurrenciesRepository)  {
	(*current).repository = repository
}

func (current Converter) ConvertTo(currency entity.Currency) (entity.Price, error) {
	rate := float32(current.Amount.Amount) / currency.Rate

	fmt.Println("rate:", rate)

	//amount := current.Amount.ToFloat() * rate
	//
	//price, err := entity.MakePrice(currency, amount)
	//
	//if err != nil {
	//	return entity.Price{}, err
	//}

	return entity.Price{}, nil
}
