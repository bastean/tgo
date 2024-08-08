package coins

import (
	"github.com/bastean/tgo/pkg/context/domain/errors"
	"github.com/bastean/tgo/pkg/context/domain/service"
)

type List struct {
	Value []string `validate:"gt=0,unique,dive,alphanum"`
}

func NewList(value []string) (*List, error) {
	valueObj := &List{
		Value: value,
	}

	if service.IsValueObjectInvalid(valueObj) {
		return nil, errors.NewInvalidValue(&errors.Bubble{
			Where: "NewList",
			What:  "List must be greater than zero, unique and alphanumeric only",
			Why: errors.Meta{
				"List": value,
			},
		})
	}

	return valueObj, nil
}
