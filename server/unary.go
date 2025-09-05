package main

import (
	"context"

	pb "github.com/nirajnagrale/grpc-go/proto"
)

func (*Server) SayHello(ctx context.Context, req *pb.NoParam) (res *pb.HelloResponse, err error) {
	return &pb.HelloResponse{
		HelloResponse: "Hello from gRPC server",
	}, nil
}
