package portfolio

import (
	"github.com/bastean/tgo/pkg/context/domain/errors"
)

func CoinsWithValidValue() *Coins {
	value, err := NewCoins([]string{"monero", "bitcoin", "ethereum"})

	if err != nil {
		errors.Panic(err.Error(), "CoinsWithValidValue")
	}

	return value
}

func CoinsWithInvalidLength() ([]string, error) {
	value := []string{}

	_, err := NewCoins(value)

	return value, err
}

func CoinsWithInvalidRepeats() ([]string, error) {
	value := []string{"monero", "monero", "bitcoin"}

	_, err := NewCoins(value)

	return value, err
}

func CoinsWithInvalidNamesLength() ([]string, error) {
	value := []string{"x"}

	_, err := NewCoins(value)

	return value, err
}
