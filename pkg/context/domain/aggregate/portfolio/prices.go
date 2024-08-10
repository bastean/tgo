package portfolio

type Prices struct {
	Value map[string]float64
}

func NewPrices(value map[string]float64) *Prices {
	return &Prices{
		Value: value,
	}
}
