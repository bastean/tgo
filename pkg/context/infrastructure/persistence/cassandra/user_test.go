package cassandra_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/bastean/tgo/pkg/context/domain/aggregate/portfolio"
	"github.com/bastean/tgo/pkg/context/domain/aggregate/user"
	"github.com/bastean/tgo/pkg/context/domain/errors"
	"github.com/bastean/tgo/pkg/context/domain/repository"
	"github.com/bastean/tgo/pkg/context/infrastructure/persistence/cassandra"
	"github.com/stretchr/testify/suite"
)

type UserTestSuite struct {
	suite.Suite
	sut repository.User
}

func (suite *UserTestSuite) SetupTest() {
	auth := &cassandra.Auth{
		Hostname: os.Getenv("DATABASE_CASSANDRA_HOSTNAME"),
		Port:     os.Getenv("DATABASE_CASSANDRA_PORT"),
		Username: os.Getenv("DATABASE_CASSANDRA_USER"),
		Password: os.Getenv("DATABASE_CASSANDRA_PASSWORD"),
	}

	config := &cassandra.Config{
		Keyspace: os.Getenv("DATABASE_CASSANDRA_KEYSPACE"),
		Strategy: "SimpleStrategy",
		Factor:   1,
	}

	session, err := cassandra.Open(auth, config)

	if err != nil {
		panic(err)
	}

	table := "users_test"

	suite.sut, err = cassandra.OpenUser(session, table)

	if err != nil {
		panic(err)
	}
}

func (suite *UserTestSuite) TestSave() {
	expected := user.Random()

	suite.NoError(suite.sut.Save(expected))

	criteria := &repository.UserSearchCriteria{
		Username: expected.Username,
	}

	actual, err := suite.sut.Search(criteria)

	suite.NoError(err)

	suite.Equal(expected, actual)
}

func (suite *UserTestSuite) TestSaveErrDuplicateKey() {
	random := user.Random()

	suite.NoError(suite.sut.Save(random))

	err := suite.sut.Save(random)

	var actual *errors.ErrAlreadyExist

	suite.ErrorAs(err, &actual)

	expected := &errors.ErrAlreadyExist{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "Save",
		What:  fmt.Sprintf("%s already registered", random.Username.Value),
		Why: errors.Meta{
			"Username": random.Username.Value,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func (suite *UserTestSuite) TestUpdate() {
	expected := user.Random()

	suite.NoError(suite.sut.Save(expected))

	expected.Portfolio = portfolio.Random()

	suite.NoError(suite.sut.Update(expected))

	criteria := &repository.UserSearchCriteria{
		Username: expected.Username,
	}

	actual, err := suite.sut.Search(criteria)

	suite.NoError(err)

	suite.Equal(expected, actual)
}

func (suite *UserTestSuite) TestDelete() {
	random := user.Random()

	suite.NoError(suite.sut.Save(random))

	suite.NoError(suite.sut.Delete(random.Username))

	criteria := &repository.UserSearchCriteria{
		Username: random.Username,
	}

	_, err := suite.sut.Search(criteria)

	suite.Error(err)
}

func (suite *UserTestSuite) TestSearch() {
	expected := user.Random()

	suite.NoError(suite.sut.Save(expected))

	criteria := &repository.UserSearchCriteria{
		Username: expected.Username,
	}

	actual, err := suite.sut.Search(criteria)

	suite.NoError(err)

	suite.Equal(expected, actual)
}

func (suite *UserTestSuite) TestSearchErrNotFound() {
	random := user.Random()

	criteria := &repository.UserSearchCriteria{
		Username: random.Username,
	}

	_, err := suite.sut.Search(criteria)

	var actual *errors.ErrNotExist

	suite.ErrorAs(err, &actual)

	expected := &errors.ErrNotExist{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "Search",
		What:  fmt.Sprintf("%s not found", random.Username.Value),
		Why: errors.Meta{
			"Username": random.Username.Value,
		},
		Who: actual.Who,
	}}

	suite.EqualError(expected, actual.Error())
}

func TestIntegrationUserSuite(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}
