package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/yk2220s/go-grpc-sample/blog"
	"google.golang.org/grpc"
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
	pb.RegisterBlogServer(s, &blogServer{})

	fmt.Println("server running on " + address)
	s.Serve(lis)
}
