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
		What:  "Coins list must be greater than zero, unique and alphanumeric only",
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
		What:  "Coins list must be greater than zero, unique and alphanumeric only",
		Why: errors.Meta{
			"Coins": value,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func (suite *CoinsValueObjectTestSuite) TestWithInvalidAlphanumeric() {
	value, err := portfolio.CoinsWithInvalidAlphanumeric()

	var actual *errors.ErrInvalidValue

	suite.ErrorAs(err, &actual)

	expected := &errors.ErrInvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewCoins",
		What:  "Coins list must be greater than zero, unique and alphanumeric only",
		Why: errors.Meta{
			"Coins": value,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func TestUnitCoinsValueObjectSuite(t *testing.T) {
	suite.Run(t, new(CoinsValueObjectTestSuite))
}
