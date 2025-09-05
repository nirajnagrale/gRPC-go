package main

import (
	pb "github.com/nirajnagrale/grpc-go/proto"
)

func (*Server) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error {
	helloResponce := make([]string, 0)
	for {
		req, err := stream.Recv()
		if err != nil {
			break
		}
		helloResponce = append(helloResponce, "Hello "+req.GetHelloRequest())
	}
	return stream.SendAndClose(&pb.MessageList{
		Messages: helloResponce,
	})

}
