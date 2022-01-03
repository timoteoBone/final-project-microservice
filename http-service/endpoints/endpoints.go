package endpoints

import (
	"github.com/go-kit/kit/endpoint"
)

type Service interface {
	CreateUser()
	GetUser()
}

type Endpoints struct {
	createUs endpoint.Endpoint
	getUs    endpoint.Endpoint
}

func MakeEndpoints()
