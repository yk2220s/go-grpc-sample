package main

import (
	"log"

	"github.com/yk2220s/go-grpc-sample/api/application/di"
)

func main() {
	server, cleanup, err := di.InitializeGRPCServer()

	if err != nil {
		log.Fatalf("failed to build server: %v", err)
	}

	defer cleanup()

	server.Serve()
}
