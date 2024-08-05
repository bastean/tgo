package coins

import (
	"github.com/bastean/tgo/pkg/context/domain/errors"
)

type Coins struct {
	*Currency
	*List
}

type Primitive struct {
	Currency string
	List     []string
}

func create(primitive *Primitive) (*Coins, error) {
	currency, errCurrency := NewCurrency(primitive.Currency)
	list, errList := NewList(primitive.List)

	if err := errors.Join(errCurrency, errList); err != nil {
		return nil, errors.BubbleUp(err, "create")
	}

	return &Coins{
		Currency: currency,
		List:     list,
	}, nil
}

func (coins *Coins) ToPrimitive() *Primitive {
	return &Primitive{
		Currency: coins.Currency.Value,
		List:     coins.List.Value,
	}
}

func FromPrimitive(primitive *Primitive) (*Coins, error) {
	coins, err := create(primitive)

	if err != nil {
		return nil, errors.BubbleUp(err, "FromPrimitive")
	}

	return coins, nil
}

func New(primitive *Primitive) (*Coins, error) {
	coins, err := create(primitive)

	if err != nil {
		return nil, errors.BubbleUp(err, "New")
	}

	return coins, nil
}
