package created

import (
	"github.com/bastean/tgo/pkg/context/domain/aggregate/user"
	"github.com/bastean/tgo/pkg/context/domain/errors"
	"github.com/bastean/tgo/pkg/context/domain/transfers"
)

type Created struct {
	transfers.Transfer
}

func (created *Created) Run(event *user.CreatedSucceeded) error {
	err := created.Transfer.Submit(event.Attributes)

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	return nil
}
