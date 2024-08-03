package delete

import (
	"github.com/bastean/tgo/pkg/context/shared/domain/errors"
	"github.com/bastean/tgo/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/tgo/pkg/context/user/domain/usecase"
)

type Handler struct {
	usecase.Delete
}

func (handler *Handler) Handle(command *Command) error {
	id, errId := user.NewId(command.Id)
	password, errPassword := user.NewPassword(command.Password)

	err := errors.Join(errId, errPassword)

	if err != nil {
		return errors.BubbleUp(err, "Handle")
	}

	err = handler.Delete.Run(id, password)

	if err != nil {
		return errors.BubbleUp(err, "Handle")
	}

	return nil
}
