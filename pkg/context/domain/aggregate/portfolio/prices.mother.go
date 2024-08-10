package portfolio

import (
	"github.com/bastean/tgo/pkg/context/domain/service"
)

func PricesWithRandomValue() *Prices {
	return NewPrices(map[string]float64{
		"monero":   service.Create.Float64(),
		"bitcoin":  service.Create.Float64(),
		"ethereum": service.Create.Float64(),
	})
}
