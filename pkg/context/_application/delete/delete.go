package delete

import (
	"github.com/bastean/tgo/pkg/context/domain/aggregate/user"
	"github.com/bastean/tgo/pkg/context/domain/errors"
	"github.com/bastean/tgo/pkg/context/domain/hashing"
	"github.com/bastean/tgo/pkg/context/domain/repository"
	"github.com/bastean/tgo/pkg/context/domain/service"
)

type Delete struct {
	repository.User
	hashing.Hashing
}

func (delete *Delete) Run(id *user.Id, password *user.Password) error {
	found, err := delete.User.Search(&repository.UserSearchCriteria{
		Id: id,
	})

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	err = service.IsPasswordInvalid(delete.Hashing, found.Password.Value, password.Value)

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	err = delete.User.Delete(found.Id)

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	return nil
}
