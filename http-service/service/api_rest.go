package service

import (
	"github.com/go-kit/kit/log"
)

type Repository interface {
	CreateUser()
	GetUser()
}

type service struct {
	Repo   Repository
	Logger log.Logger
}

func (s *service) CreateUser() {

}

func (s *service) GetUser() {

}
