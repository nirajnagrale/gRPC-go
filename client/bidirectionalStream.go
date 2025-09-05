package main

import (
	"context"
	"fmt"

	pb "github.com/nirajnagrale/grpc-go/proto"
)

func callBidirectionalStreaming(client pb.GreetServiceClient) {
	ctx := context.Background()
	stream, err := client.SayHelloBidirectionalStreaming(ctx)
	if err != nil {
		fmt.Println("Error while calling SayHelloBidirectionalStreaming:", err)
		return
	}
	names := []string{"Alice", "Bob", "Charlie", "David", "Eve", "Frank"}
	for _, name := range names {
		req := &pb.HelloRequest{
			HelloRequest: name,
		}
		if err := stream.Send(req); err != nil {
			fmt.Println("Error while sending request to stream:", err)
			return
		}
		msg, err := stream.Recv()
		if err != nil {
			fmt.Println("Error while receiving message from stream:", err)
			return
		}
		fmt.Println("Response from server:", msg.GetHelloResponse())
	}
	stream.CloseSend()
}
