package service

import (
	"context"

	"github.com/go-kit/kit/log"

	entities "github.com/timoteoBone/final-project-microservice/grpc-service/entities"
	mapper "github.com/timoteoBone/final-project-microservice/grpc-service/mapper"
)

type Repository interface {
	GetUser(ctx context.Context, userId string)
	CreateUser(ctx context.Context, user entities.User) error
}

type service struct {
	Repo   Repository
	Logger log.Logger
}

func NewService(l log.Logger, r Repository) *service {
	return &service{r, l}
}

func (s *service) CreateUser(ctx context.Context, userReq entities.CreateUserRequest) (entities.CreateUserResponse, error) {
	s.Logger.Log(s.Logger, "request", "create user", "recevied")

	response := entities.CreateUserResponse{}

	user := mapper.CreateUserRequestToUser(userReq)
	err := s.Repo.CreateUser(ctx, user)

	if err != nil {
		response.Status.Message = "Unable to create user"
		return response, err
	}

	response.Status.Message = ("" + user.Name + " created successfully")

	return response, nil
}

func (s *service) GetUser(ctx context.Context, userId string) (entities.User, error) {
	return entities.User{}, nil
}
