package valueobjs

import (
	"github.com/bastean/tgo/pkg/context/shared/domain/errors"
	"github.com/bastean/tgo/pkg/context/shared/domain/services"
)

func TypeWithValidValue() *Type {
	value, err := NewType(services.Create.RandomString([]string{"event", "command"}))

	if err != nil {
		errors.Panic(err.Error(), "TypeWithValidValue")
	}

	return value
}

func TypeWithInvalidValue() (string, error) {
	value := "x"

	_, err := NewType(value)

	return value, err
}
