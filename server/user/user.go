package user

import (
	"context"
	"github.com/StasJDM/proto-example/pkg/xmp/common"
	"github.com/StasJDM/proto-example/pkg/xmp/server/user"
)

type UserServer struct {
	user.UnimplementedUserServer
}

func (s *UserServer) Register(context.Context, *user.RegisterUserRequest) (*common.IdMessage, error) {
	return &common.IdMessage{
		Id: "test-user-id-1",
	}, nil
}
