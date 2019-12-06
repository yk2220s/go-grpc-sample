// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"github.com/google/wire"
	"github.com/yk2220s/go-grpc-sample/api/application/server"
	"github.com/yk2220s/go-grpc-sample/api/usecase"
)

func InitializeGRPCServer() (*server.GRPCServer, func(), error) {
	wire.Build(
		server.NewGRPCServer,
		usecase.NewGetPost,
	)
	return nil, nil, nil
}
