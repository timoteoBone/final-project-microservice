package service_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/go-kit/kit/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/timoteoBone/final-project-microservice/grpc-service/entities"
	"github.com/timoteoBone/final-project-microservice/grpc-service/mocks"
	"github.com/timoteoBone/final-project-microservice/grpc-service/service"
)

func TestNewService(t *testing.T) {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "grpcUserService",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	srvc := service.NewService(logger, &mocks.RepoSitoryMock{mock.Mock{}})

	assert.False(t, srvc == nil)
}

func TestCreateUser(t *testing.T) {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "grpcUserService",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	user := entities.User{
		Name: "Timo",
		Id:   "Shimo",
		Pass: "123",
		Age:  19,
	}

	correctCreateUserRequest := entities.CreateUserRequest{
		Name: user.Name,
		Id:   user.Id,
		Pass: user.Pass,
		Age:  user.Age,
	}

	Succesfullresponse := entities.CreateUserResponse{
		Status: entities.Status{
			Message: " created successfully",
		}, UserId: user.Id,
	}

	testsCases := []struct {
		Name            string
		Identifier      string
		User            entities.User
		Request         entities.CreateUserRequest
		RepoReturn      error
		ServiceResponse entities.CreateUserResponse
		Error           error
	}{
		{
			Name:            "Create User Valid Case",
			Identifier:      "CreateUser",
			User:            user,
			Request:         correctCreateUserRequest,
			RepoReturn:      nil,
			ServiceResponse: Succesfullresponse,
		},

		{
			Name:       "Create User Empty Fields",
			Identifier: "CreateUser",
			User:       entities.User{Pass: "sdsd"},
			Request:    entities.CreateUserRequest{Pass: "sdsd"},
			RepoReturn: errors.New("change to errors implementation"),
			ServiceResponse: entities.CreateUserResponse{
				Status: entities.Status{
					Message: "Unable to create user",
				},
			},
		},
	}

	repo := new(mocks.RepoSitoryMock)
	srvc := service.NewService(logger, repo)

	for _, tc := range testsCases {
		t.Run(tc.Name, func(t *testing.T) {
			ctx := context.Background()
			repo.On(tc.Identifier, ctx, tc.User).Return(tc.RepoReturn)

			res, err := srvc.CreateUser(ctx, tc.Request)
			assert.ErrorIs(t, err, tc.RepoReturn)
			assert.Equal(t, tc.ServiceResponse, res)

		})

	}

}

func TestGetUser(t *testing.T) {

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "grpcUserService",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	user := entities.User{
		Name: "Timo",
		Id:   "Shimo",
		Pass: "123",
		Age:  19,
	}

	correctGetUserRequest := entities.GetUserRequest{
		UserID: user.Id,
	}

	correctGetUserResponse := entities.GetUserResponse{
		Name: user.Name,
		Id:   user.Id,
		Age:  user.Age,
	}

	repo := new(mocks.RepoSitoryMock)
	srvc := service.NewService(logger, repo)

	ctx := context.Background()

	testCases := []struct {
		Name                    string
		Identifier              string
		UserId                  string
		Request                 entities.GetUserRequest
		RepoReturn              entities.User
		RepoError               error
		ServiceExpectedResponse entities.GetUserResponse
		Error                   error
	}{
		{
			Name:                    "Get User Valid Request",
			Identifier:              "GetUser",
			UserId:                  user.Id,
			Request:                 correctGetUserRequest,
			RepoReturn:              user,
			RepoError:               nil,
			ServiceExpectedResponse: correctGetUserResponse,
			Error:                   nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			repo.Mock.On(tc.Identifier, ctx, tc.UserId).Return(tc.RepoReturn, tc.RepoError)

			res, err := srvc.GetUser(ctx, tc.Request)
			assert.Equal(t, tc.ServiceExpectedResponse, res)
			assert.ErrorIs(t, err, tc.Error)

		})
	}

}
