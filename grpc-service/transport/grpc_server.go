package server

import (
	"context"
	"errors"

	gr "github.com/go-kit/kit/transport/grpc"

	"github.com/timoteoBone/final-project-microservice/grpc-service/endpoints"
	"github.com/timoteoBone/final-project-microservice/grpc-service/entities"
	proto "github.com/timoteoBone/final-project-microservice/grpc-service/pb"
)

type gRPCSv struct {
	createUs gr.Handler
	getUs    gr.Handler
	proto.UnimplementedUserServiceServer
}

func NewGrpcServer(end endpoints.Endpoints) proto.UserServiceServer {

	return &gRPCSv{
		createUs: gr.NewServer(
			end.CreateUser,
			decodeCreateUserRequest,
			encodeCreateUserResponse,
		),

		getUs: gr.NewServer(
			end.GetUser,
			decodeGetUserRequest,
			encodeGetUserResponse,
		),
	}
}

func (g *gRPCSv) CreateUser(ctx context.Context, rq *proto.CreateUserRequest) (rs *proto.CreateUserResponse, err error) {
	_, resp, err := g.createUs.ServeGRPC(ctx, rq)

	if err != nil {
		return nil, err
	}

	return resp.(*proto.CreateUserResponse), nil

}

func (g *gRPCSv) GetUser(ctx context.Context, rq *proto.GetUserRequest) (rs *proto.GetUserResponse, err error) {
	_, resp, err := g.getUs.ServeGRPC(ctx, rq)

	if err != nil {
		return nil, err
	}

	return resp.(*proto.GetUserResponse), nil
}

func decodeCreateUserRequest(ctx context.Context, request interface{}) (interface{}, error) {
	res, err := request.(*proto.CreateUserRequest)

	if !err {
		return nil, errors.New("Unable to decode request")
	}

	return entities.CreateUserRequest{
		Name: res.Name,
		Age:  res.Age,
		Pass: res.Pass,
	}, nil

}

func encodeCreateUserResponse(ctx context.Context, response interface{}) (interface{}, error) {
	res := response.(entities.CreateUserResponse)
	status := &proto.Status{Message: res.Status.Message}
	protoResp := &proto.CreateUserResponse{User_Id: res.UserId, Status: status}

	return protoResp, nil
}

func decodeGetUserRequest(ctx context.Context, request interface{}) (interface{}, error) {
	res, err := request.(*proto.GetUserRequest)

	if !err {
		return nil, errors.New("Unable to decode request")
	}

	return entities.GetUserRequest{
		UserID: res.User_Id,
	}, nil

}

func encodeGetUserResponse(ctx context.Context, response interface{}) (interface{}, error) {
	res := response.(entities.GetUserResponse)
	protoResp := &proto.GetUserResponse{Id: res.Id, Name: res.Name, Age: res.Age}
	return protoResp, nil
}
