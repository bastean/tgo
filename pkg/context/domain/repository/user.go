package repository

import (
	"github.com/bastean/tgo/pkg/context/domain/aggregate/user"
)

type UserSearchCriteria struct {
	*user.Username
}

type User interface {
	Save(*user.User) error
	Update(*user.Username) error
	Delete(*user.Username) error
	Search(*UserSearchCriteria) (*user.User, error)
}
