package price_test

import (
	"testing"

	"github.com/bastean/tgo/pkg/context/application/portfolio/price"
	"github.com/bastean/tgo/pkg/context/domain/aggregate/portfolio"
	"github.com/bastean/tgo/pkg/context/domain/aggregate/user"
	"github.com/bastean/tgo/pkg/context/domain/repository"
	"github.com/bastean/tgo/pkg/context/domain/usecase"
	"github.com/bastean/tgo/pkg/context/infrastructure/cryptocurrency"
	"github.com/bastean/tgo/pkg/context/infrastructure/persistence"
	"github.com/stretchr/testify/suite"
)

type PriceTestSuite struct {
	suite.Suite
	sut        usecase.PortfolioPrice
	repository *persistence.UserMock
	market     *cryptocurrency.MarketMock
}

func (suite *PriceTestSuite) SetupTest() {
	suite.repository = new(persistence.UserMock)

	suite.market = new(cryptocurrency.MarketMock)

	suite.sut = price.New(suite.repository, suite.market)
}

func (suite *PriceTestSuite) TestCreate() {
	random := user.Random()

	criteria := &repository.UserSearchCriteria{
		Username: random.Username,
	}

	prices := portfolio.PricesWithRandomValue()

	suite.repository.On("Search", criteria).Return(random)

	suite.market.On("Tracker", random.Portfolio).Return(prices)

	expected := prices.Value

	actual, err := suite.sut.Run(random.Username.Value)

	suite.NoError(err)

	suite.repository.AssertExpectations(suite.T())

	suite.market.AssertExpectations(suite.T())

	suite.EqualValues(expected, actual)
}

func TestUnitPriceSuite(t *testing.T) {
	suite.Run(t, new(PriceTestSuite))
}
