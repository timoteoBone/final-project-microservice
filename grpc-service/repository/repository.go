package repository

import (
	"context"
	//proto "github.com/timoteoBone/final-project-microservice/grpc-service/pb"
	"database/sql"
	"log"

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
