package entity

type Currency struct {
	Code         string  `json:"code"`
	Name         string  `json:"name"`
	Symbol       string  `json:"symbol"`
	Rate         float32 `json:"rate"`
	IsoCode      int     `json:"iso_code"`
	IsoMinorUnit int     `json:"iso_minor_unit"`
}
