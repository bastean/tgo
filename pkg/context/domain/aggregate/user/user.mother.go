package user

import (
	"github.com/bastean/tgo/pkg/context/domain/aggregate/portfolio"
	"github.com/bastean/tgo/pkg/context/domain/errors"
)

func Random() *User {
	username := UsernameWithValidValue()
	portfolio := portfolio.Random()

	user, err := New(&Primitive{
		Username:  username.Value,
		Portfolio: portfolio.ToPrimitive(),
	})

	if err != nil {
		errors.Panic(err.Error(), "Random")
	}

	return user
}

func RandomPrimitive() *Primitive {
	return Random().ToPrimitive()
}
