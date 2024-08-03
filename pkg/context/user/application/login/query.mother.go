package login

import (
	"github.com/bastean/tgo/pkg/context/user/domain/aggregate/user"
)

func RandomQuery() *Query {
	email := user.EmailWithValidValue()
	password := user.PasswordWithValidValue()

	return &Query{
		Email:    email.Value,
		Password: password.Value,
	}
}
