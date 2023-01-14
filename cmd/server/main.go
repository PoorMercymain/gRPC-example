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

	for i := 1; i < int(request.MessageLinesAmount)+1; i++ {
		resultMessage += fmt.Sprintf("Строка номер %d\n", i)
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

	l, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
