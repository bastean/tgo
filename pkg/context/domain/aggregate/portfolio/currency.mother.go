package portfolio

import (
	"github.com/bastean/tgo/pkg/context/domain/errors"
	"github.com/bastean/tgo/pkg/context/domain/service"
)

func CurrencyWithValidValue() *Currency {
	value, err := NewCurrency(service.Create.RandomString([]string{"USD", "EUR", "RUB", "CNY", "JPY", "BRL"}))

	if err != nil {
		errors.Panic(err.Error(), "CurrencyWithValidValue")
	}

	return value
}

func CurrencyWithInvalidValue() (string, error) {
	value := "x"

	_, err := NewCurrency(value)

	return value, err
}
