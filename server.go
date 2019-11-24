package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/yk2220s/go-grpc-sample/blog"
	"google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	status "google.golang.org/grpc/status"
)

const (
	port = 50001
)

type blogServer struct {
	pb.UnimplementedBlogServer
}

func (s *blogServer) GetPost(ctx context.Context, req *pb.GetRequest) (*pb.Post, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPost not implemented")
}

func (s *blogServer) SavePost(ctx context.Context, req *pb.SaveRequest) (*pb.Post, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SavePost not implemented")
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
	reflection.Register(s)

	fmt.Println("server running on " + address)
	s.Serve(lis)
}
