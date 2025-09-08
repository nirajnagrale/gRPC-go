package main

import (
	"fmt"
	"io"

	pb "github.com/nirajnagrale/grpc-go/proto"
)

func (*Server) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error {
	helloResponce := make([]string, 0)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("End of stream")
			break
		}
		if err != nil {
			fmt.Println("Error while receiving request from stream:", err)
			break
		}
		helloResponce = append(helloResponce, "Hello "+req.GetHelloRequest())
		fmt.Println("Request received: ", req.GetHelloRequest())
	}
	return stream.SendAndClose(&pb.MessageList{
		Messages: helloResponce,
	})

}
