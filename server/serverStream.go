package main

import (
	pb "github.com/nirajnagrale/grpc-go/proto"
)

func (*Server) SayHelloServerStreaming(req *pb.NamesList, stream pb.GreetService_SayHelloServerStreamingServer) error {
	for _, name := range req.GetNames() {
		result := "Hello " + name
		stream.Send(&pb.HelloResponse{
			HelloResponse: result,
		})
	}
	return nil
}
