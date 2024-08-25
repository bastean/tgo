package create

import (
	"github.com/bastean/tgo/pkg/context/domain/aggregate/user"
	"github.com/bastean/tgo/pkg/context/domain/errors"
	"github.com/bastean/tgo/pkg/context/domain/repository"
	"github.com/bastean/tgo/pkg/context/domain/usecase"
)

type Create struct {
	repository.User
}

func (create *Create) Run(primitive *user.Primitive) error {
	new, err := user.New(primitive)

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	err = create.User.Create(new)

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	return nil
}

func New(repository repository.User) usecase.UserCreate {
	return &Create{
		User: repository,
	}
}
