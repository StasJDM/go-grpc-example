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

func (s *PostServer) FindOne(context.Context, *common.IdMessage) (*post.FindOneResponse, error) {
	postRes := &post.PostMessage{
		Id:      "test-post-id-1",
		Title:   "Test title 1",
		Content: "Test content",
	}

	return &post.FindOneResponse{Post: postRes}, nil
}

func (s *PostServer) FindMany(ctx context.Context, pagination *common.PaginationMessage) (*post.FindManyResponse, error) {
	return &post.FindManyResponse{
		Posts:      nil,
		Pagination: pagination,
	}, nil
}
