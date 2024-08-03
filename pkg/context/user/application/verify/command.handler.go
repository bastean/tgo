package verify

import (
	"github.com/bastean/tgo/pkg/context/shared/domain/errors"
	"github.com/bastean/tgo/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/tgo/pkg/context/user/domain/usecase"
)

type Handler struct {
	usecase.Verify
}

func (handler *Handler) Handle(command *Command) error {
	id, err := user.NewId(command.Id)

	if err != nil {
		return errors.BubbleUp(err, "Handle")
	}

	err = handler.Verify.Run(id)

	if err != nil {
		return errors.BubbleUp(err, "Handle")
	}

	return nil
}
