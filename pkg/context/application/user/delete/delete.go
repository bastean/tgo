package delete

import (
	"github.com/bastean/tgo/pkg/context/domain/aggregate/user"
	"github.com/bastean/tgo/pkg/context/domain/errors"
	"github.com/bastean/tgo/pkg/context/domain/repository"
	"github.com/bastean/tgo/pkg/context/domain/usecase"
)

type Delete struct {
	repository.User
}

func (delete *Delete) Run(username string) error {
	usernameVO, err := user.NewUsername(username)

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	found, err := delete.User.Search(&repository.UserSearchCriteria{
		Username: usernameVO,
	})

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	err = delete.User.Delete(found.Username)

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	return nil
}

func New(repository repository.User) usecase.UserDelete {
	return &Delete{
		User: repository,
	}
}
