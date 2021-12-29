package server

import (
	"context"

	gr "github.com/go-kit/kit/transport/grpc"

	"github.com/timoteoBone/final-project-microservice/grpc-service/endpoints"
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
			encodeResponse,
		),

		getUs: gr.NewServer(
			end.GetUser,
			decodeCreateUserRequest,
			encodeResponse,
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

}

func decodeCreateUserRequest(ctx context.Context, request interface{}) (interface{}, error) {

}

func encodeResponse(ctx context.Context, response interface{}) (interface{}, error) {

}
