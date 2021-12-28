package service

import (
	"context"

	"github.com/go-kit/kit/log"

	entities "github.com/timoteoBone/final-project-microservice/grpc-service/entities"
)

type Repository interface {
	GetUser(ctx context.Context, userId string)
	CreateUser(ctx context.Context, user entities.User)
}

type service struct {
	Repo   Repository
	Logger log.Logger
}

func (s *service) CreateUser(context.Context, entities.User) error {
	return nil
}

func (s *service) GetUser(ctx context.Context, userId string) (entities.User, error) {
	return entities.User{}, nil
}
