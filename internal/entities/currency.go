package entities

type CurrencyRate struct {
	BaseCode string
	Code     string
	Rate     float64
}

func NewCurrency(base_code, code string, rate float64) *CurrencyRate {
	return &CurrencyRate{base_code, code, rate}
}
