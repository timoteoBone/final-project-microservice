package endpoints

import (
	"context"

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

func MakeEndpoints(s Service) *Endpoints {

	return &Endpoints{
		createUs: MakeCreateUserEndpoint(s),
		getUs:    MakeGetUserEndpoint(s),
	}
}

func MakeCreateUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

	}
}

func MakeGetUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

	}
}
