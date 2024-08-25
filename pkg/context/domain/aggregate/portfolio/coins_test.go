package portfolio_test

import (
	"testing"

	"github.com/bastean/tgo/pkg/context/domain/aggregate/portfolio"
	"github.com/bastean/tgo/pkg/context/domain/errors"
	"github.com/stretchr/testify/suite"
)

type CoinsValueObjectTestSuite struct {
	suite.Suite
}

func (suite *CoinsValueObjectTestSuite) SetupTest() {}

func (suite *CoinsValueObjectTestSuite) TestWithInvalidLength() {
	value, err := portfolio.CoinsWithInvalidLength()

	var actual *errors.ErrInvalidValue

	suite.ErrorAs(err, &actual)

	expected := &errors.ErrInvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewCoins",
		What:  "The amount of coins must be between 1 and 10, without repeating, and the names must be from 4 to 20 characters",
		Why: errors.Meta{
			"Coins": value,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func (suite *CoinsValueObjectTestSuite) TestWithInvalidRepeats() {
	value, err := portfolio.CoinsWithInvalidRepeats()

	var actual *errors.ErrInvalidValue

	suite.ErrorAs(err, &actual)

	expected := &errors.ErrInvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewCoins",
		What:  "The amount of coins must be between 1 and 10, without repeating, and the names must be from 4 to 20 characters",
		Why: errors.Meta{
			"Coins": value,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func (suite *CoinsValueObjectTestSuite) TestWithInvalidNamesLength() {
	value, err := portfolio.CoinsWithInvalidNamesLength()

	var actual *errors.ErrInvalidValue

	suite.ErrorAs(err, &actual)

	expected := &errors.ErrInvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewCoins",
		What:  "The amount of coins must be between 1 and 10, without repeating, and the names must be from 4 to 20 characters",
		Why: errors.Meta{
			"Coins": value,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func TestUnitCoinsValueObjectSuite(t *testing.T) {
	suite.Run(t, new(CoinsValueObjectTestSuite))
}
