package project

import (
	"context"
	"github.com/StasJDM/proto-example/pkg/xmp/common"
	"github.com/StasJDM/proto-example/pkg/xmp/server/project"
)

type ProjectServer struct {
	project.UnimplementedProjectServer
}

func (s *ProjectServer) Create(context.Context, *project.CreateProjectRequest) (*common.IdMessage, error) {
	return &common.IdMessage{
		Id: "test-project-id-1",
	}, nil
}
