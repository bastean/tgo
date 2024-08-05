package user

import (
	"github.com/bastean/tgo/pkg/context/shared/domain/messages"
	"github.com/bastean/tgo/pkg/context/shared/domain/transfers"
	"github.com/bastean/tgo/pkg/context/user/application/created"
)

var (
	Created *created.Consumer
)

func InitCreated(transfer transfers.Transfer, queue *messages.Queue) {
	Created = &created.Consumer{
		Created: &created.Created{
			Transfer: transfer,
		},
		Queues: []*messages.Queue{queue},
	}
}
