package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/yk2220s/go-grpc-sample/blog"
	"grpc.go4.org"
)

const (
	port = 50001
)

type blogServer struct {
	pb.UnimplementedBlogServer
}

func main() {
	fmt.Println("building server..")
	address := fmt.Sprintf("localhost:%d", port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	fmt.Println("server running on " + address)
	s.Serve(lis)
}
