package main

import (
	"context"
	"fmt"

	pb "github.com/nirajnagrale/grpc-go/proto"
)

func callSayHelloClientStreaming(client pb.GreetServiceClient) {
	ctx := context.Background()
	stream, err := client.SayHelloClientStreaming(ctx)
	if err != nil {
		fmt.Println("Error while calling SayHelloClientStreaming:", err)
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
	}
	stream.CloseSend()
}
