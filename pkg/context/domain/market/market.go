package market

import (
	"github.com/bastean/tgo/pkg/context/domain/aggregate/coins"
)

type Market interface {
	Price(*coins.List) (map[string]float32, error)
}
