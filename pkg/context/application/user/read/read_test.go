package read_test

import (
	"testing"

	"github.com/bastean/tgo/pkg/context/application/user/read"
	"github.com/bastean/tgo/pkg/context/domain/aggregate/user"
	"github.com/bastean/tgo/pkg/context/domain/repository"
	"github.com/bastean/tgo/pkg/context/domain/usecase"
	"github.com/bastean/tgo/pkg/context/infrastructure/persistence"
	"github.com/stretchr/testify/suite"
)

type ReadTestSuite struct {
	suite.Suite
	sut        usecase.UserRead
	repository *persistence.UserMock
}

func (suite *ReadTestSuite) SetupTest() {
	suite.repository = new(persistence.UserMock)
	suite.sut = read.New(suite.repository)
}

func (suite *ReadTestSuite) TestRead() {
	random := user.Random()

	criteria := &repository.UserSearchCriteria{
		Username: random.Username,
	}

	suite.repository.On("Search", criteria).Return(random)

	expected := random.ToPrimitive()

	actual, err := suite.sut.Run(random.Username.Value)

	suite.NoError(err)

	suite.repository.AssertExpectations(suite.T())

	suite.EqualValues(expected, actual)
}

func TestUnitReadSuite(t *testing.T) {
	suite.Run(t, new(ReadTestSuite))
}
