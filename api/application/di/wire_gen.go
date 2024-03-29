// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package di

import (
	"github.com/yk2220s/go-grpc-sample/api/application/server"
	"github.com/yk2220s/go-grpc-sample/api/usecase"
)

// Injectors from wire.go:

func InitializeGRPCServer() (*server.GRPCServer, func(), error) {
	iGetPost := usecase.NewGetPost()
	grpcServer := server.NewGRPCServer(iGetPost)
	return grpcServer, func() {
	}, nil
}
