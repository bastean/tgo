package portfolio

import (
	"strings"

	"github.com/bastean/tgo/pkg/context/domain/errors"
	"github.com/bastean/tgo/pkg/context/domain/service"
)

type Currency struct {
	Value string `validate:"iso4217"`
}

func NewCurrency(value string) (*Currency, error) {
	value = strings.TrimSpace(value)
	value = strings.ToUpper(value)

	valueObj := &Currency{
		Value: value,
	}

	if service.IsValueObjectInvalid(valueObj) {
		return nil, errors.NewInvalidValue(&errors.Bubble{
			Where: "NewCurrency",
			What:  "Currency must be the ISO 4217 3-letter code only",
			Why: errors.Meta{
				"Currency": value,
			},
		})
	}

	return valueObj, nil
}
