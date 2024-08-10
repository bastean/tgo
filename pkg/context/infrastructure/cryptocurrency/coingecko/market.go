package coingecko

import (
	"github.com/bastean/tgo/pkg/context/domain/aggregate/portfolio"
	"github.com/bastean/tgo/pkg/context/domain/errors"
)

type Market struct {
	*CoinGecko
}

func (market *Market) Tracker(wallet *portfolio.Portfolio) (*portfolio.Prices, error) {
	err := market.IsCurrencySupported(wallet.Currency.Value)

	if err != nil {
		errors.BubbleUp(err, "Tracker")
	}

	prices, err := market.CoinPrices(wallet.Currency.Value, wallet.Coins.Value)

	if err != nil {
		errors.BubbleUp(err, "Tracker")
	}

	return portfolio.NewPrices(prices), nil
}

func NewMarket(coingecko *CoinGecko) *Market {
	return &Market{
		CoinGecko: coingecko,
	}
}
