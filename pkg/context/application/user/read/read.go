package read

import (
	"github.com/bastean/tgo/pkg/context/domain/aggregate/user"
	"github.com/bastean/tgo/pkg/context/domain/errors"
	"github.com/bastean/tgo/pkg/context/domain/repository"
	"github.com/bastean/tgo/pkg/context/domain/usecase"
)

type Read struct {
	repository.User
}

func (read *Read) Run(username string) (*user.Primitive, error) {
	usernameVO, err := user.NewUsername(username)

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	found, err := read.User.Search(&repository.UserSearchCriteria{
		Username: usernameVO,
	})

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	return found.ToPrimitive(), nil
}

func New(repository repository.User) usecase.UserRead {
	return &Read{
		User: repository,
	}
}
