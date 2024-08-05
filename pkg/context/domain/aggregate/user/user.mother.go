package user

import (
	"github.com/bastean/tgo/pkg/context/domain/aggregate/coins"
	"github.com/bastean/tgo/pkg/context/domain/errors"
)

func Random() *User {
	username := UsernameWithValidValue()
	coins := coins.Random()

	user, err := New(&Primitive{
		Username: username.Value,
		Coins:    coins.ToPrimitive(),
	})

	if err != nil {
		errors.Panic(err.Error(), "Random")
	}

	return user
}

func RandomPrimitive() *Primitive {
	return Random().ToPrimitive()
}
