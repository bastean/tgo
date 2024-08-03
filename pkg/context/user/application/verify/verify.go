package verify

import (
	"github.com/bastean/tgo/pkg/context/shared/domain/errors"
	"github.com/bastean/tgo/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/tgo/pkg/context/user/domain/repository"
)

type Verify struct {
	repository.User
}

func (verify *Verify) Run(id *user.Id) error {
	found, err := verify.User.Search(&repository.UserSearchCriteria{
		Id: id,
	})

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	if found.Verified.Value {
		return nil
	}

	err = verify.User.Verify(id)

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	return nil
}
