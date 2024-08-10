package market

import (
	"github.com/bastean/tgo/pkg/context/domain/aggregate/portfolio"
)

type Market interface {
	Tracker(*portfolio.Portfolio) (*portfolio.Prices, error)
}
