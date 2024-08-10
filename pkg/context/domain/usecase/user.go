package usecase

import (
	"github.com/bastean/tgo/pkg/context/domain/aggregate/user"
)

type (
	UserCreate interface {
		Run(*user.Primitive) error
	}
	UserRead interface {
		Run(username string) (*user.Primitive, error)
	}
	UserUpdate interface {
		Run(*user.Primitive) error
	}
	UserDelete interface {
		Run(username string) error
	}
)

type (
	PortfolioPrice interface {
		Run(username string) (map[string]float64, error)
	}
)
