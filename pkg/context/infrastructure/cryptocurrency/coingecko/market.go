package coingecko

import (
	"github.com/bastean/tgo/pkg/context/domain/aggregate/portfolio"
	"github.com/bastean/tgo/pkg/context/domain/errors"
)

type Market struct {
	*CoinGecko
}

func (market *Market) Tracker(portfolio *portfolio.Portfolio) (map[string]float64, error) {
	err := market.IsCurrencySupported(portfolio.Currency.Value)

	if err != nil {
		errors.BubbleUp(err, "Tracker")
	}

	prices, err := market.CoinPrices(portfolio.Currency.Value, portfolio.Coins.Value)

	if err != nil {
		errors.BubbleUp(err, "Tracker")
	}

	return prices, nil
}

func NewMarket(coingecko *CoinGecko) *Market {
	return &Market{
		CoinGecko: coingecko,
	}
}
