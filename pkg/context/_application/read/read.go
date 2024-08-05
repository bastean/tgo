package read

import (
	"github.com/bastean/tgo/pkg/context/domain/aggregate/user"
	"github.com/bastean/tgo/pkg/context/domain/errors"
	"github.com/bastean/tgo/pkg/context/domain/repository"
)

type Read struct {
	repository.User
}

func (read *Read) Run(id *user.Id) (*user.User, error) {
	found, err := read.User.Search(&repository.UserSearchCriteria{
		Id: id,
	})

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	return found, nil
}
