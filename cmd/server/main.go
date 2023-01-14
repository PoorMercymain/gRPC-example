package main

import (
	"context"
	"fmt"
	"gRPC/pkg/api"
	"google.golang.org/grpc"
	"log"
	"net"
)

type GRPCServer struct {
	api.MessageServer
}

func (G GRPCServer) GetMessage(ctx context.Context, request *api.MessageRequest) (*api.MessageResponse, error) {
	var resultMessage string
	var messageBuffer string

	for i := 1; i < int(request.MessageLinesAmount)+1; i++ {
		messageBuffer = fmt.Sprintf("Строка номер %d", i)
		resultMessage += messageBuffer + "\n"
		fmt.Println(messageBuffer, "добавлена для ответа клиенту")
	}

	return &api.MessageResponse{GeneratedMessage: resultMessage}, nil
}

func (G GRPCServer) mustEmbedUnimplementedMessageServer() {
	//это нам не нужно
	log.Fatal("Ты тего наделал(-а)...")
}

func main() {
	s := grpc.NewServer()
	srv := &GRPCServer{}
	api.RegisterMessageServer(s, srv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Сервер запущен")

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
