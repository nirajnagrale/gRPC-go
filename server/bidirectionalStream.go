package main

import (
	"fmt"
	"io"

	pb "github.com/nirajnagrale/grpc-go/proto"
)

func (*Server) SayHelloBidirectionalStreaming(stream pb.GreetService_SayHelloBidirectionalStreamingServer) error {

	for {
		req := pb.HelloRequest{}
		err := stream.RecvMsg(&req)
		if err == io.EOF{
			fmt.Println("End of stream")
			break
		}
		if err != nil {
			fmt.Println("Error while receiving request from stream:", err)
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
