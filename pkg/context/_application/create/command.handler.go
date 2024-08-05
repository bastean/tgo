package create

import (
	"github.com/bastean/tgo/pkg/context/domain/aggregate/user"
	"github.com/bastean/tgo/pkg/context/domain/errors"
	"github.com/bastean/tgo/pkg/context/domain/messages"
	"github.com/bastean/tgo/pkg/context/domain/usecase"
)

type Handler struct {
	usecase.Create
	messages.Broker
}

func (handler *Handler) Handle(command *Command) error {
	new, err := user.New(&user.Primitive{
		Id:       command.Id,
		Email:    command.Email,
		Username: command.Username,
		Password: command.Password,
	})

	if err != nil {
		return errors.BubbleUp(err, "Handle")
	}

	err = handler.Create.Run(new)

	if err != nil {
		return errors.BubbleUp(err, "Handle")
	}

	handler.Broker.PublishMessages(new.PullMessages())

	return nil
}
