package main

import (
	"fmt"
	"net"

	pb "github.com/nirajnagrale/grpc-go/proto"
	"google.golang.org/grpc"
)

type Server struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", ":8083")
	if err != nil {
		fmt.Println("Failed to listen:", err)
		return
	}

	grpcServer := grpc.NewServer()
	err = grpcServer.Serve(lis)
	if err != nil {
		fmt.Println("Failed to start gRPC server:", err)
		return
	}
	pb.RegisterGreetServiceServer(grpcServer, &Server{})
	fmt.Println("gRPC server is running on port 50051")

}
