package main

import (
	"log"
	"net"

	"github.com/StasJDM/go-grpc-example/pkg/server/post"
	"github.com/StasJDM/go-grpc-example/pkg/server/user"
	postServer "github.com/StasJDM/go-grpc-example/server/post"
	userServer "github.com/StasJDM/go-grpc-example/server/user"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Listener creation error: %v", err)
	}

	grpcServer := grpc.NewServer()
	userService := &userServer.UserServer{}
	postService := &postServer.PostServer{}

	user.RegisterUserServer(grpcServer, userService)
	post.RegisterPostServer(grpcServer, postService)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC: %v", err)
	}
}
