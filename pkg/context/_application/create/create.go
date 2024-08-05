package create

import (
	"github.com/bastean/tgo/pkg/context/domain/aggregate/user"
	"github.com/bastean/tgo/pkg/context/domain/errors"
	"github.com/bastean/tgo/pkg/context/domain/repository"
)

type Create struct {
	repository.User
}

func (create *Create) Run(user *user.User) error {
	err := create.User.Save(user)

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	return nil
}
