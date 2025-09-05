package main

import (
	"fmt"

	pb "github.com/nirajnagrale/grpc-go/proto"
)

func (*Server) SayHelloBidirectionalStreaming(stream pb.GreetService_SayHelloBidirectionalStreamingServer) error {

	for {
		req := pb.HelloRequest{}
		err := stream.RecvMsg(&req)
		if err != nil {
			break
		}
		result := "Hello " + req.GetHelloRequest()
		fmt.Println("Request received: ", result)
		err = stream.Send(&pb.HelloResponse{
			HelloResponse: result,
		})
		if err != nil {
			break
		}
	}
	return nil
}
