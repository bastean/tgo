package portfolio

import (
	"github.com/bastean/tgo/pkg/context/domain/errors"
)

type Portfolio struct {
	*Currency
	*Coins
}

type Primitive struct {
	Currency string
	Coins    []string
}

func create(primitive *Primitive) (*Portfolio, error) {
	currency, errCurrency := NewCurrency(primitive.Currency)
	coins, errCoins := NewCoins(primitive.Coins)

	if err := errors.Join(errCurrency, errCoins); err != nil {
		return nil, errors.BubbleUp(err, "create")
	}

	return &Portfolio{
		Currency: currency,
		Coins:    coins,
	}, nil
}

func (portfolio *Portfolio) ToPrimitive() *Primitive {
	return &Primitive{
		Currency: portfolio.Currency.Value,
		Coins:    portfolio.Coins.Value,
	}
}

func FromPrimitive(primitive *Primitive) (*Portfolio, error) {
	portfolio, err := create(primitive)

	if err != nil {
		return nil, errors.BubbleUp(err, "FromPrimitive")
	}

	return portfolio, nil
}

func New(primitive *Primitive) (*Portfolio, error) {
	portfolio, err := create(primitive)

	if err != nil {
		return nil, errors.BubbleUp(err, "New")
	}

	return portfolio, nil
}
