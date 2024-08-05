package coins_test

import (
	"testing"

	"github.com/bastean/tgo/pkg/context/domain/aggregate/coins"
	"github.com/bastean/tgo/pkg/context/domain/errors"
	"github.com/stretchr/testify/suite"
)

type ListValueObjectTestSuite struct {
	suite.Suite
}

func (suite *ListValueObjectTestSuite) SetupTest() {}

func (suite *ListValueObjectTestSuite) TestWithInvalidLength() {
	value, err := coins.ListWithInvalidLength()

	var actual *errors.ErrInvalidValue

	suite.ErrorAs(err, &actual)

	expected := &errors.ErrInvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewList",
		What:  "List must be greater than zero, unique and alphanumeric only",
		Why: errors.Meta{
			"List": value,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func (suite *ListValueObjectTestSuite) TestWithInvalidRepeats() {
	value, err := coins.ListWithInvalidRepeats()

	var actual *errors.ErrInvalidValue

	suite.ErrorAs(err, &actual)

	expected := &errors.ErrInvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewList",
		What:  "List must be greater than zero, unique and alphanumeric only",
		Why: errors.Meta{
			"List": value,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func (suite *ListValueObjectTestSuite) TestWithInvalidAlphanumeric() {
	value, err := coins.ListWithInvalidAlphanumeric()

	var actual *errors.ErrInvalidValue

	suite.ErrorAs(err, &actual)

	expected := &errors.ErrInvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewList",
		What:  "List must be greater than zero, unique and alphanumeric only",
		Why: errors.Meta{
			"List": value,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func TestUnitListValueObjectSuite(t *testing.T) {
	suite.Run(t, new(ListValueObjectTestSuite))
}
