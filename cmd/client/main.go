package main

import (
	"context"
	"fmt"
	"gRPC/pkg/api"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}

	c := api.NewMessageClient(conn)
	res, err := c.GetMessage(context.Background(), &api.MessageRequest{MessageLinesAmount: 5})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(fmt.Sprintf("Получено от сервера:\n%s", res.GeneratedMessage))
}
