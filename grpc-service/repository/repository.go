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

func NewSQL(db *sql.DB, log log.Logger) *sqlRepo {
	return &sqlRepo{db, log}
}

func (repo *sqlRepo) CreateUser(ctx context.Context, user entities.User) error {

	stmt, err := repo.DB.Prepare("INSERT INTO USER VALUES(?,?,?)")
	if err != nil {
		return err
	}

	res, err := stmt.Exec(user.Id, user.Name, user.Age, user.Pass)
	if err != nil {
		return err
	}

	repo.Logger.Log(res, "rows affected")
	return nil
}

func (sr *sqlRepo) GetUser(ctx context.Context, userId string) (entities.User, error) {

	return entities.User{}, nil
}
