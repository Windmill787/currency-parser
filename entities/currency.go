package entities

type Currency struct {
	Code   string
	Number int
}

func NewCurrency(code string, number int) *Currency {
	return &Currency{code, number}
}

func USD() *Currency {
	return NewCurrency("USD", 840)
}

func EUR() *Currency {
	return NewCurrency("EUR", 978)
}

func UAH() *Currency {
	return NewCurrency("UAH", 980)
}
