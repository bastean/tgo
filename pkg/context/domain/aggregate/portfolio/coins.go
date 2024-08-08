package portfolio

import (
	"github.com/bastean/tgo/pkg/context/domain/errors"
	"github.com/bastean/tgo/pkg/context/domain/service"
)

type Coins struct {
	Value []string `validate:"gt=0,unique,dive,alphanum"`
}

func NewCoins(value []string) (*Coins, error) {
	valueObj := &Coins{
		Value: value,
	}

	if service.IsValueObjectInvalid(valueObj) {
		return nil, errors.NewInvalidValue(&errors.Bubble{
			Where: "NewCoins",
			What:  "Coins list must be greater than zero, unique and alphanumeric only",
			Why: errors.Meta{
				"Coins": value,
			},
		})
	}

	return valueObj, nil
}
