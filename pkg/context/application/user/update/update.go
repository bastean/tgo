package update

import (
	"github.com/bastean/tgo/pkg/context/domain/aggregate/user"
	"github.com/bastean/tgo/pkg/context/domain/errors"
	"github.com/bastean/tgo/pkg/context/domain/repository"
	"github.com/bastean/tgo/pkg/context/domain/usecase"
)

type Update struct {
	repository.User
}

func (update *Update) Run(primitive *user.Primitive) error {
	account, err := user.New(primitive)

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	_, err = update.User.Search(&repository.UserSearchCriteria{
		Username: account.Username,
	})

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	err = update.User.Update(account)

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	return nil
}

func New(repository repository.User) usecase.UserUpdate {
	return &Update{
		User: repository,
	}
}
