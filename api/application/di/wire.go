// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"github.com/google/wire"
	"github.com/yk2220s/go-grpc-sample/api/application/server"
)

func InitializeGRPCServer() (*server.GRPCServer, func(), error) {
	wire.Build(
		server.NewGRPCServer,
	)
	return nil, nil, nil
}
