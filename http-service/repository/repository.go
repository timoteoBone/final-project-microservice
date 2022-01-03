package repository

import (
	"github.com/go-kit/log"
	"google.golang.org/grpc"
)

type grpcClient struct {
	server *grpc.ClientConn
	logger log.Logger
}

func (repo *grpcClient) CreateUser() {

}

func (repo *grpcClient) GetUser() {

}
