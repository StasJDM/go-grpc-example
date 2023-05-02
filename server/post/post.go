package post

import (
	"context"

	"github.com/StasJDM/go-grpc-example/pkg/server/post"
)

type PostServer struct {
	post.UnimplementedPostServer
}

func (s *PostServer) Create(context.Context, *post.CreatePostRequest) (*post.CreatePostResponse, error) {
	return &post.CreatePostResponse{
		Id: "test-post-id-1",
	}, nil
}
