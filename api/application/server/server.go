package server

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/yk2220s/go-grpc-sample/blog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

const (
	port = 50001
)

// BlogServer implements grpc methods
type BlogServer struct {
	pb.UnimplementedBlogServer
}

// GetPost implements interface
func (s *BlogServer) GetPost(ctx context.Context, req *pb.GetRequest) (*pb.Post, error) {
	if req.GetId() == 1 {
		return &pb.Post{Id: 1, Title: "How to train the dragon", Text: "the amazing film"}, nil
	}

	return nil, status.Errorf(codes.NotFound, "Post is not found")
}

// SavePost implements interface
func (s *BlogServer) SavePost(ctx context.Context, req *pb.SaveRequest) (*pb.Post, error) {
	// write your logic
	return nil, status.Errorf(codes.Unimplemented, "method SavePost not implemented")
}

// GRPCServer serves a servixe
type GRPCServer struct {
	grpc *grpc.Server
}

// Serve start listen and serve
func (gs *GRPCServer) Serve() {
	address := fmt.Sprintf("localhost:%d", port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	pb.RegisterBlogServer(gs.grpc, &BlogServer{})
	reflection.Register(gs.grpc)

	fmt.Println("server running on " + address)
	gs.grpc.Serve(lis)
}

// NewGRPCServer is a provider
func NewGRPCServer() *GRPCServer {
	gs := grpc.NewServer()
	return &GRPCServer{grpc: gs}
}
