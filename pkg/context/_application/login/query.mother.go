package login

import (
	"github.com/bastean/tgo/pkg/context/domain/aggregate/user"
)

func RandomQuery() *Query {
	email := user.EmailWithValidValue()
	password := user.PasswordWithValidValue()

	return &Query{
		Email:    email.Value,
		Password: password.Value,
	}
}
