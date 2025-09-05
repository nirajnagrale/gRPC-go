package main

import (
	"context"
	"fmt"

	pb "github.com/nirajnagrale/grpc-go/proto"
)

func callSayHello(client pb.GreetServiceClient) {
	ctx := context.Background()
	res, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil {
		fmt.Println("Error while calling SayHello:", err)
		return
	}
	fmt.Println("Response from SayHello:", res.GetHelloResponse())
}
