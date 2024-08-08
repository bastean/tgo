package portfolio_test

import (
	"testing"

	"github.com/bastean/tgo/pkg/context/domain/aggregate/portfolio"
	"github.com/bastean/tgo/pkg/context/domain/errors"
	"github.com/stretchr/testify/suite"
)

type CurrencyValueObjectTestSuite struct {
	suite.Suite
}

func (suite *CurrencyValueObjectTestSuite) SetupTest() {}

func (suite *CurrencyValueObjectTestSuite) TestWithInvalidValue() {
	value, err := portfolio.CurrencyWithInvalidValue()

	var actual *errors.ErrInvalidValue

	suite.ErrorAs(err, &actual)

	expected := &errors.ErrInvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewCurrency",
		What:  "Currency must be the ISO 4217 3-letter code only",
		Why: errors.Meta{
			"Currency": value,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func TestUnitCurrencyValueObjectSuite(t *testing.T) {
	suite.Run(t, new(CurrencyValueObjectTestSuite))
}
