package usecase

import (
	"github.com/bastean/tgo/pkg/context/domain/aggregate/user"
)

type (
	UserCreate interface {
		Run(*user.User) error
	}
	UserRead interface {
		Run(*user.Username) (*user.User, error)
	}
	UserUpdate interface {
		Run(*user.User) error
	}
	UserDelete interface {
		Run(*user.Username) error
	}
)

type (
	CoinsPrice interface {
		Run(*user.Username) (map[string]float32, error)
	}
)
