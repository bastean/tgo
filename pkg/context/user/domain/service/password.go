package service

import (
	"github.com/bastean/tgo/pkg/context/shared/domain/errors"
	"github.com/bastean/tgo/pkg/context/user/domain/hashing"
)

func IsPasswordInvalid(hashing hashing.Hashing, hashed, plain string) error {
	if hashing.IsNotEqual(hashed, plain) {
		return errors.NewFailure(&errors.Bubble{
			Where: "IsPasswordInvalid",
			What:  "Invalid password",
		})
	}

	return nil
}
