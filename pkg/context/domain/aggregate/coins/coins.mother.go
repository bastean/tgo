package coins

import (
	"github.com/bastean/tgo/pkg/context/domain/errors"
)

func Random() *Coins {
	currency := CurrencyWithValidValue()
	list := ListWithValidValue()

	coins, err := New(&Primitive{
		Currency: currency.Value,
		List:     list.Value,
	})

	if err != nil {
		errors.Panic(err.Error(), "Random")
	}

	return coins
}

func RandomPrimitive() *Primitive {
	return Random().ToPrimitive()
}
