package handler

import (
	"github.com/bastean/tgo/internal/pkg/service/errors"
	"github.com/bastean/tgo/internal/pkg/service/logger/log"
	tele "gopkg.in/telebot.v3"
)

func Error(c tele.Context, err error) error {
	var (
		errInvalidValue *errors.ErrInvalidValue
		errAlreadyExist *errors.ErrAlreadyExist
		errNotExist     *errors.ErrNotExist
		errFailure      *errors.ErrFailure
		errInternal     *errors.ErrInternal
	)

	var reply string

	switch {
	case errors.As(err, &errInvalidValue):
		reply = errInvalidValue.What
	case errors.As(err, &errAlreadyExist):
		reply = errAlreadyExist.What
	case errors.As(err, &errNotExist):
		reply = errNotExist.What
	case errors.As(err, &errFailure):
		reply = errFailure.What
	case errors.As(err, &errInternal):
		reply = "Bot error. Try again later."
		fallthrough
	case err != nil:
		log.Error(err.Error())
	}

	if reply != "" {
		return c.Send(reply)
	}

	return nil
}
