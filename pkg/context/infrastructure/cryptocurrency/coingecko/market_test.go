package coingecko_test

import (
	"os"
	"testing"

	"github.com/bastean/tgo/pkg/context/domain/aggregate/user"
	"github.com/bastean/tgo/pkg/context/domain/errors"
	"github.com/bastean/tgo/pkg/context/domain/market"
	"github.com/bastean/tgo/pkg/context/infrastructure/cryptocurrency/coingecko"
	"github.com/stretchr/testify/suite"
)

type MarketTestSuite struct {
	suite.Suite
	sut market.Market
}

func (suite *MarketTestSuite) SetupTest() {
	session, err := coingecko.New(
		os.Getenv("TGO_API_COINGECKO_DEMO_KEY"),
		"https://api.coingecko.com/api/v3/",
	)

	if err != nil {
		errors.Panic(err.Error(), "SetupTest")
	}

	suite.sut = coingecko.NewMarket(session)
}

func (suite *MarketTestSuite) TestTracker() {
	random := user.Random()

	coins := random.Portfolio.Coins.Value

	actual, err := suite.sut.Tracker(random.Portfolio)

	suite.NoError(err)

	for _, expected := range coins {
		if _, ok := actual.Value[expected]; !ok {
			suite.Equal(expected, "")
		}
	}
}

func TestIntegrationMarketSuite(t *testing.T) {
	suite.Run(t, new(MarketTestSuite))
}
