package create_test

import (
	"testing"

	"github.com/bastean/tgo/pkg/context/application/create"
	"github.com/bastean/tgo/pkg/context/domain/aggregate/user"
	"github.com/bastean/tgo/pkg/context/domain/handlers"
	"github.com/bastean/tgo/pkg/context/domain/usecase"
	"github.com/bastean/tgo/pkg/context/infrastructure/communications"
	"github.com/bastean/tgo/pkg/context/infrastructure/persistence"
	"github.com/stretchr/testify/suite"
)

type CreateTestSuite struct {
	suite.Suite
	sut        handlers.Command[*create.Command]
	create     usecase.Create
	repository *persistence.UserMock
	broker     *communications.BrokerMock
}

func (suite *CreateTestSuite) SetupTest() {
	suite.broker = new(communications.BrokerMock)

	suite.repository = new(persistence.UserMock)

	suite.create = &create.Create{
		User: suite.repository,
	}

	suite.sut = &create.Handler{
		Create: suite.create,
		Broker: suite.broker,
	}
}

func (suite *CreateTestSuite) TestCreate() {
	command := create.RandomCommand()

	new, _ := user.New(&user.Primitive{
		Id:       command.Id,
		Email:    command.Email,
		Username: command.Username,
		Password: command.Password,
	})

	messages := new.Messages

	suite.repository.On("Save", new)

	suite.broker.On("PublishMessages", messages)

	suite.NoError(suite.sut.Handle(command))

	suite.repository.AssertExpectations(suite.T())

	suite.broker.AssertExpectations(suite.T())
}

func TestUnitCreateSuite(t *testing.T) {
	suite.Run(t, new(CreateTestSuite))
}
