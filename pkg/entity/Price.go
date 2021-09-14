package entity

type Price struct {
	Amount int64 `json:"amount"`
	Currency Currency
}

func (price Price) ToFloat() float64  {
	currency := price.Currency

	return float64(price.Amount / int64(currency.IsoMinorUnit))
}
