package portfolio

import (
	"github.com/bastean/tgo/pkg/context/application/portfolio/price"
	"github.com/bastean/tgo/pkg/context/domain/aggregate/portfolio"
	"github.com/bastean/tgo/pkg/context/domain/market"
	"github.com/bastean/tgo/pkg/context/domain/repository"
)

type (
	Primitive = portfolio.Primitive
)

func Start(repository repository.User, market market.Market) {
	Price = &price.Price{
		User:   repository,
		Market: market,
	}
}
