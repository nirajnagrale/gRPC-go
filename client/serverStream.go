package main

import (
	"context"
	"fmt"
	"io"

	pb "github.com/nirajnagrale/grpc-go/proto"
)

func callSayHelloServerStreaming(client pb.GreetServiceClient) {
	ctx := context.Background()
	req := &pb.NamesList{
		Names: []string{"Alice", "Bob", "Charlie"},
	}
	stream, err := client.SayHelloServerStreaming(ctx, req)
	if err != nil {
		fmt.Println("Error while calling SayHelloServerStreaming:", err)
		return
	}
	for {
		msg, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error while receiving message from stream:", err)
			return
		}
		fmt.Println("Response from server:", msg.GetHelloResponse())
	}
}
