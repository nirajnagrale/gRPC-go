package main

import (
	"fmt"
	"net"

	pb "github.com/nirajnagrale/grpc-go/proto"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedGreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", ":8084")
	if err != nil {
		fmt.Println("Failed to listen:", err)
		return
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &Server{})
	err = grpcServer.Serve(lis)
	if err != nil {
		fmt.Println("Failed to start gRPC server:", err)
		return
	}

	fmt.Println("gRPC server is running on port 50051")

}
