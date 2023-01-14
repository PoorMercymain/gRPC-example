package main

import (
	"context"
	"fmt"
	"gRPC/pkg/api"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}

	c := api.NewMessageClient(conn)
	res, err := c.GetMessage(context.Background(), &api.MessageRequest{MessageLinesAmount: 5})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res.GeneratedMessage)
}
