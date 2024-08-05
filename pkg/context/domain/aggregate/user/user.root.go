package user

import (
	"github.com/bastean/tgo/pkg/context/domain/aggregate/coins"
	"github.com/bastean/tgo/pkg/context/domain/errors"
)

type User struct {
	*Username
	Coins *coins.Coins
}

type Primitive struct {
	Username string
	Coins    *coins.Primitive
}

func create(primitive *Primitive) (*User, error) {
	username, errUsername := NewUsername(primitive.Username)
	coins, errCoins := coins.New(primitive.Coins)

	if err := errors.Join(errUsername, errCoins); err != nil {
		return nil, errors.BubbleUp(err, "create")
	}

	return &User{
		Username: username,
		Coins:    coins,
	}, nil
}

func (user *User) ToPrimitive() *Primitive {
	return &Primitive{
		Username: user.Username.Value,
		Coins:    user.Coins.ToPrimitive(),
	}
}

func FromPrimitive(primitive *Primitive) (*User, error) {
	user, err := create(primitive)

	if err != nil {
		return nil, errors.BubbleUp(err, "FromPrimitive")
	}

	return user, nil
}

func New(primitive *Primitive) (*User, error) {
	user, err := create(primitive)

	if err != nil {
		return nil, errors.BubbleUp(err, "New")
	}

	return user, nil
}
