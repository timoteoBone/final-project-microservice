package mocks

import (
	"context"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/mock"
	"github.com/timoteoBone/final-project-microservice/grpc-service/entities"
)

type RepoSitoryMock struct {
	mock.Mock
}

func (repo *RepoSitoryMock) CreateUser(ctx context.Context, user entities.User) error {
	args := repo.Called(ctx, user)

	return args.Error(0)
}

func (repo *RepoSitoryMock) GetUser(ctx context.Context, userId string) (entities.User, error) {
	args := repo.Called(ctx, userId)
	id := args[0]

	if id == nil {
		return entities.User{}, args.Error(1)
	}

	return id.(entities.User), nil

}
