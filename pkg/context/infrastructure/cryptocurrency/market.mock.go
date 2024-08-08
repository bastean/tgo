package cryptocurrency

import (
	"github.com/bastean/tgo/pkg/context/domain/aggregate/portfolio"
	"github.com/stretchr/testify/mock"
)

type MarketMock struct {
	mock.Mock
}

func (market *MarketMock) Tracker(portfolio *portfolio.Portfolio) (map[string]float64, error) {
	args := market.Called(portfolio)
	return args.Get(0).(map[string]float64), nil
}
