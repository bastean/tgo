package price

import (
	"github.com/bastean/tgo/pkg/context/domain/aggregate/user"
	"github.com/bastean/tgo/pkg/context/domain/errors"
	"github.com/bastean/tgo/pkg/context/domain/market"
	"github.com/bastean/tgo/pkg/context/domain/repository"
	"github.com/bastean/tgo/pkg/context/domain/usecase"
)

type Price struct {
	repository.User
	market.Market
}

func (price *Price) Run(username string) (map[string]float64, error) {
	usernameVO, err := user.NewUsername(username)

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	found, err := price.User.Search(&repository.UserSearchCriteria{
		Username: usernameVO,
	})

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	prices, err := price.Market.Tracker(found.Portfolio)

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	return prices.Value, nil
}

func New(repository repository.User, market market.Market) usecase.PortfolioPrice {
	return &Price{
		User:   repository,
		Market: market,
	}
}
