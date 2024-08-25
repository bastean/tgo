package persistence

import (
	"github.com/bastean/tgo/pkg/context/domain/aggregate/user"
	"github.com/bastean/tgo/pkg/context/domain/repository"
	"github.com/stretchr/testify/mock"
)

type UserMock struct {
	mock.Mock
}

func (repository *UserMock) Create(user *user.User) error {
	repository.Called(user)
	return nil
}

func (repository *UserMock) Update(user *user.User) error {
	repository.Called(user)
	return nil
}

func (repository *UserMock) Delete(username *user.Username) error {
	repository.Called(username)
	return nil
}

func (repository *UserMock) Search(criteria *repository.UserSearchCriteria) (*user.User, error) {
	args := repository.Called(criteria)
	return args.Get(0).(*user.User), nil
}
