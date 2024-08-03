package login

import (
	"github.com/bastean/tgo/pkg/context/shared/domain/errors"
	"github.com/bastean/tgo/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/tgo/pkg/context/user/domain/hashing"
	"github.com/bastean/tgo/pkg/context/user/domain/repository"
	"github.com/bastean/tgo/pkg/context/user/domain/service"
)

type Login struct {
	repository.User
	hashing.Hashing
}

func (login *Login) Run(email *user.Email, password *user.Password) (*user.User, error) {
	found, err := login.User.Search(&repository.UserSearchCriteria{
		Email: email,
	})

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	err = service.IsPasswordInvalid(login.Hashing, found.Password.Value, password.Value)

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	return found, nil
}
