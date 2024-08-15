package user

import (
	"github.com/bastean/tgo/pkg/context/application/user/create"
	"github.com/bastean/tgo/pkg/context/application/user/delete"
	"github.com/bastean/tgo/pkg/context/application/user/read"
	"github.com/bastean/tgo/pkg/context/application/user/update"
	"github.com/bastean/tgo/pkg/context/domain/aggregate/user"
	"github.com/bastean/tgo/pkg/context/domain/repository"
)

type (
	Primitive = user.Primitive
)

func Start(repository repository.User) {
	Create = &create.Create{
		User: repository,
	}

	Read = &read.Read{
		User: repository,
	}

	Update = &update.Update{
		User: repository,
	}

	Delete = &delete.Delete{
		User: repository,
	}
}
