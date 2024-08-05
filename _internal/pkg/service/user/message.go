package user

import (
	"github.com/bastean/tgo/internal/pkg/service/communication/rabbitmq"
	"github.com/bastean/tgo/pkg/context/shared/domain/messages"
)

var QueueSendConfirmation = &messages.Queue{
	Name: messages.NewRecipientName(&messages.RecipientNameComponents{
		Service: "user",
		Entity:  "user",
		Action:  "send confirmation",
		Event:   "created",
		Status:  "succeeded",
	}),
	Bindings: []string{rabbitmq.BindingEventCreatedSucceeded},
}
