package post

import (
	"context"
	"github.com/StasJDM/proto-example/pkg/xmp/common"
	"github.com/StasJDM/proto-example/pkg/xmp/server/post"
)

type PostServer struct {
	post.UnimplementedPostServer
}

func (s *PostServer) Create(context.Context, *post.CreatePostRequest) (*common.IdMessage, error) {
	return &common.IdMessage{
		Id: "test-post-id-1",
	}, nil
}
