package user

import (
	"context"

	"github.com/StasJDM/go-grpc-example/pkg/server/user"
)

type UserServer struct {
	user.UnimplementedUserServer
}

func (s *UserServer) Register(context.Context, *user.RegisterUserRequest) (*user.RegisterUserResponse, error) {
	return &user.RegisterUserResponse{
		Id: "test-user-id-1",
	}, nil
}
