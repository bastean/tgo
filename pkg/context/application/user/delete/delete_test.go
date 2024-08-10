package delete_test

import (
	"testing"

	"github.com/bastean/tgo/pkg/context/application/user/delete"
	"github.com/bastean/tgo/pkg/context/domain/aggregate/user"
	"github.com/bastean/tgo/pkg/context/domain/repository"
	"github.com/bastean/tgo/pkg/context/domain/usecase"
	"github.com/bastean/tgo/pkg/context/infrastructure/persistence"
	"github.com/stretchr/testify/suite"
)

type DeleteTestSuite struct {
	suite.Suite
	sut        usecase.UserDelete
	repository *persistence.UserMock
}

func (suite *DeleteTestSuite) SetupTest() {
	suite.repository = new(persistence.UserMock)
	suite.sut = delete.New(suite.repository)
}

func (suite *DeleteTestSuite) TestDelete() {
	random := user.Random()

	criteria := &repository.UserSearchCriteria{
		Username: random.Username,
	}

	suite.repository.On("Search", criteria).Return(random)

	suite.repository.On("Delete", random.Username)

	suite.NoError(suite.sut.Run(random.Username.Value))

	suite.repository.AssertExpectations(suite.T())
}

func TestUnitDeleteSuite(t *testing.T) {
	suite.Run(t, new(DeleteTestSuite))
}
