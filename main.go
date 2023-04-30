package main

import (
	"context"
	"log"
	"net"

	"github.com/StasJDM/go-grpc-example/user"
	"google.golang.org/grpc"
)

type UserService struct {
	user.UnimplementedUserServer
}

func (s UserService) Register(context.Context, *user.RegisterUserRequest) (*user.RegisterUserResponse, error) {
	return &user.RegisterUserResponse{
		Id: "test-id-1",
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Listener creation error: %v", err)
	}

	grpcServer := grpc.NewServer()
	userService := &UserService{}

	user.RegisterUserServer(grpcServer, userService)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC: %v", err)
	}
}
