package repository

import (
	"context"
	"database/sql"

	"github.com/go-kit/kit/log"
	entities "github.com/timoteoBone/final-project-microservice/grpc-service/entities"
)

type sqlRepo struct {
	DB     *sql.DB
	Logger log.Logger
}

func (sr *sqlRepo) CreateUser(ctx context.Context, user entities.User) error {

	return nil
}

func (sr *sqlRepo) GetUser(ctx context.Context, userId string) (entities.User, error) {

	return entities.User{}, nil
}
