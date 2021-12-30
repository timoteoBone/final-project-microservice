package repository

import (
	"context"
	"database/sql"
	"fmt"

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

	stmt, err := repo.DB.Prepare("INSERT INTO USER VALUES(?,?,?,?)")
	if err != nil {
		return err
	}

	res, err := stmt.Exec(user.Id, user.Name, user.Age, user.Pass)
	if err != nil {
		return err
	}

	repo.Logger.Log(res, "rows affected")
	fmt.Println()
	return nil
}

func (repo *sqlRepo) GetUser(ctx context.Context, userId string) (entities.User, error) {

	stmt, err := repo.DB.Query("SELECT first_name, id, age, pass FROM USER WHERE ID = ?", userId)
	if err != nil {
		return entities.User{}, err
	}

	user := entities.User{}
	for stmt.Next() {
		err := stmt.Scan(*&user.Name, &user.Id, &user.Age)
		if err != nil {
			//TO DO -- LEARN HOW TO HANDLE THESE KIND OF ERRORS
		}
	}

	return user, nil
}
