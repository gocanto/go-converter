package currencyLayer

type Response struct {
	Success bool `json:"success"`
	Timestamp int64 `json:"timestamp"`
	Source string `json:"source"`
	Quotes map[string]float64 `json:"quotes"`
}

func (current Response) RateFor(currency string) float64 {
	for index, value := range current.Quotes {
		if index[3:] == currency {
			return value
		}
	}

	return 1
}
