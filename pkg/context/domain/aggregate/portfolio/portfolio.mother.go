package portfolio

import (
	"github.com/bastean/tgo/pkg/context/domain/errors"
)

func Random() *Portfolio {
	currency := CurrencyWithValidValue()
	coins := CoinsWithValidValue()

	portfolio, err := New(&Primitive{
		Currency: currency.Value,
		Coins:    coins.Value,
	})

	if err != nil {
		errors.Panic(err.Error(), "Random")
	}

	return portfolio
}

func RandomPrimitive() *Primitive {
	return Random().ToPrimitive()
}
