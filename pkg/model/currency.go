package model

import (
	"github.com/voyago/converter/pkg/support"
)

type Currency struct {
	Code         string  `json:"code"`
	Name         string  `json:"name"`
	Rate         float32 `json:"rate"`
	IsoCode      int16   `json:"iso_code"`
	IsoMinorUnit int8    `json:"iso_minor_unit"`
}

func (current Currency) Multiplier() int64 {
	target := support.Strings{Value: "1"}
	target.RightPad("0", current.IsoMinorUnit)

	if result, err := target.ToInt64(); err == nil {
		return result
	}

	return 1
}

func (current Currency) ToRatePrice() (Price, error) {
	price, err := MakePrice(current, current.Rate)

	if err != nil {
		return Price{}, err
	}

	return price, nil
}
