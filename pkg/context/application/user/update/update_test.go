package update_test

import (
	"testing"

	"github.com/bastean/tgo/pkg/context/application/user/update"
	"github.com/bastean/tgo/pkg/context/domain/aggregate/user"
	"github.com/bastean/tgo/pkg/context/domain/repository"
	"github.com/bastean/tgo/pkg/context/domain/usecase"
	"github.com/bastean/tgo/pkg/context/infrastructure/persistence"
	"github.com/stretchr/testify/suite"
)

type UpdateTestSuite struct {
	suite.Suite
	sut        usecase.UserUpdate
	repository *persistence.UserMock
}

func (suite *UpdateTestSuite) SetupTest() {
	suite.repository = new(persistence.UserMock)
	suite.sut = update.New(suite.repository)
}

func (suite *UpdateTestSuite) TestUpdate() {
	random := user.Random()

	primitive := random.ToPrimitive()

	criteria := &repository.UserSearchCriteria{
		Username: random.Username,
	}

	suite.repository.On("Search", criteria).Return(random)

	suite.repository.On("Update", random)

	suite.NoError(suite.sut.Run(primitive))

	suite.repository.AssertExpectations(suite.T())
}

func TestUnitUpdateSuite(t *testing.T) {
	suite.Run(t, new(UpdateTestSuite))
}
