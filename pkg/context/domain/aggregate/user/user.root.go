package user

import (
	"github.com/bastean/tgo/pkg/context/domain/aggregate/portfolio"
	"github.com/bastean/tgo/pkg/context/domain/errors"
)

type User struct {
	*Username
	*portfolio.Portfolio
}

type Primitive struct {
	Username  string
	Portfolio *portfolio.Primitive
}

func create(primitive *Primitive) (*User, error) {
	username, errUsername := NewUsername(primitive.Username)
	portfolio, errPortfolio := portfolio.New(primitive.Portfolio)

	if err := errors.Join(errUsername, errPortfolio); err != nil {
		return nil, errors.BubbleUp(err, "create")
	}

	return &User{
		Username:  username,
		Portfolio: portfolio,
	}, nil
}

func (user *User) ToPrimitive() *Primitive {
	return &Primitive{
		Username:  user.Username.Value,
		Portfolio: user.Portfolio.ToPrimitive(),
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
