package currencyLayer

import (
	"errors"
	"fmt"
)

type Response struct {
	Success   bool               `json:"success"`
	Timestamp int64              `json:"timestamp"`
	Source    string             `json:"source"`
	Quotes    map[string]float64 `json:"quotes"`
}

func (current Response) RateFor(currency string) (float64, error) {
	for index, value := range current.Quotes {
		if index[3:] == currency {
			return value, nil
		}
	}

	return 0, errors.New(fmt.Sprintf("No rates found for the given currency [%s]", currency))
}
