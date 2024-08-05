package verify

import (
	"github.com/bastean/tgo/pkg/context/domain/aggregate/user"
)

func RandomCommand() *Command {
	id := user.IdWithValidValue()

	return &Command{
		Id: id.Value,
	}
}
