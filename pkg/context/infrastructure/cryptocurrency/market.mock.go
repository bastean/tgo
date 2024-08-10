package cryptocurrency

import (
	"github.com/bastean/tgo/pkg/context/domain/aggregate/portfolio"
	"github.com/stretchr/testify/mock"
)

type MarketMock struct {
	mock.Mock
}

func (market *MarketMock) Tracker(wallet *portfolio.Portfolio) (*portfolio.Prices, error) {
	args := market.Called(wallet)
	return args.Get(0).(*portfolio.Prices), nil
}
