package repository

import (
	"github.com/bastean/tgo/pkg/context/domain/aggregate/user"
)

type UserSearchCriteria struct {
	*user.Username
}

type User interface {
	Create(*user.User) error
	Update(*user.User) error
	Delete(*user.Username) error
	Search(*UserSearchCriteria) (*user.User, error)
}
