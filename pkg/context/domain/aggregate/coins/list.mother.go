package coins

import (
	"github.com/bastean/tgo/pkg/context/domain/errors"
)

func ListWithValidValue() *List {
	value, err := NewList([]string{"monero", "bitcoin", "ethereum"})

	if err != nil {
		errors.Panic(err.Error(), "ListWithValidValue")
	}

	return value
}

func ListWithInvalidLength() ([]string, error) {
	value := []string{}

	_, err := NewList(value)

	return value, err
}

func ListWithInvalidRepeats() ([]string, error) {
	value := []string{"monero", "monero", "bitcoin"}

	_, err := NewList(value)

	return value, err
}

func ListWithInvalidAlphanumeric() ([]string, error) {
	value := []string{"<></>"}

	_, err := NewList(value)

	return value, err
}
