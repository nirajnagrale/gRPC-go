package main

import (
	"fmt"

	pb "github.com/nirajnagrale/grpc-go/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8083"
)

func main() {
	conn, err := grpc.NewClient("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Error while connecting to server:", err)
		return
	}
	defer conn.Close()
	client := pb.NewGreetServiceClient(conn)
	callSayHello(client)
	//callSayHelloServerStreaming(client)

}
