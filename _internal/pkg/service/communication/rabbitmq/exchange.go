package rabbitmq

import (
	"github.com/bastean/tgo/pkg/context/shared/domain/messages"
)

func Exchange(name string) *messages.Router {
	return &messages.Router{
		Name: name,
	}
}
