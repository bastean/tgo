package portfolio

import (
	"fmt"

	"github.com/bastean/tgo/pkg/context/domain/errors"
	"github.com/bastean/tgo/pkg/context/domain/service"
)

const (
	CoinsMinAmount          = "1"
	CoinsMaxAmount          = "10"
	CoinMinCharactersLength = "4"
	CoinMaxCharactersLength = "20"
)

type Coins struct {
	Value []string `validate:"gte=1,lte=10,unique,dive,gte=4,lte=20"`
}

func NewCoins(value []string) (*Coins, error) {
	valueObj := &Coins{
		Value: value,
	}

	if service.IsValueObjectInvalid(valueObj) {
		return nil, errors.NewInvalidValue(&errors.Bubble{
			Where: "NewCoins",
			What:  fmt.Sprintf("The amount of coins must be between %s and %s, without repeating, and the names must be from %s to %s characters", CoinsMinAmount, CoinsMaxAmount, CoinMinCharactersLength, CoinMaxCharactersLength),
			Why: errors.Meta{
				"Coins": value,
			},
		})
	}

	return valueObj, nil
}
