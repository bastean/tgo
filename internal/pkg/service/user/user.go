package user

import (
	"github.com/bastean/tgo/pkg/context/application/portfolio/price"
	"github.com/bastean/tgo/pkg/context/application/user/create"
	"github.com/bastean/tgo/pkg/context/application/user/delete"
	"github.com/bastean/tgo/pkg/context/application/user/read"
	"github.com/bastean/tgo/pkg/context/application/user/update"
	"github.com/bastean/tgo/pkg/context/domain/market"
	"github.com/bastean/tgo/pkg/context/domain/repository"
)

func Start(repository repository.User, market market.Market) {
	Create = &create.Create{
		User: repository,
	}

	Read = &read.Read{
		User: repository,
	}

	Update = &update.Update{
		User: repository,
	}

	Delete = &delete.Delete{
		User: repository,
	}

	Price = &price.Price{
		User:   repository,
		Market: market,
	}
}
