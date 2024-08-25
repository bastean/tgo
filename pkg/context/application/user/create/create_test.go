package create_test

import (
	"testing"

	"github.com/bastean/tgo/pkg/context/application/user/create"
	"github.com/bastean/tgo/pkg/context/domain/aggregate/user"
	"github.com/bastean/tgo/pkg/context/domain/usecase"
	"github.com/bastean/tgo/pkg/context/infrastructure/persistence"
	"github.com/stretchr/testify/suite"
)

type CreateTestSuite struct {
	suite.Suite
	sut        usecase.UserCreate
	repository *persistence.UserMock
}

func (suite *CreateTestSuite) SetupTest() {
	suite.repository = new(persistence.UserMock)
	suite.sut = create.New(suite.repository)
}

func (suite *CreateTestSuite) TestCreate() {
	random := user.Random()

	primitive := random.ToPrimitive()

	suite.repository.On("Create", random)

	suite.NoError(suite.sut.Run(primitive))

	suite.repository.AssertExpectations(suite.T())
}

func TestUnitCreateSuite(t *testing.T) {
	suite.Run(t, new(CreateTestSuite))
}
