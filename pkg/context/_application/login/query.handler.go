package login

import (
	"github.com/bastean/tgo/pkg/context/domain/aggregate/user"
	"github.com/bastean/tgo/pkg/context/domain/errors"
	"github.com/bastean/tgo/pkg/context/domain/usecase"
)

type Handler struct {
	usecase.Login
}

func (handler *Handler) Handle(query *Query) (*Response, error) {
	email, errEmail := user.NewEmail(query.Email)
	password, errPassword := user.NewPassword(query.Password)

	err := errors.Join(errEmail, errPassword)

	if err != nil {
		return nil, errors.BubbleUp(err, "Handle")
	}

	found, err := handler.Login.Run(email, password)

	if err != nil {
		return nil, errors.BubbleUp(err, "Handle")
	}

	response := Response(*found.ToPrimitive())

	return &response, nil
}
