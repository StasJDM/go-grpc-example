package project

import (
	"context"

	"github.com/StasJDM/go-grpc-example/pkg/server/project"
)

type ProjectServer struct {
	project.UnimplementedProjectServer
}

func (s *ProjectServer) Create(context.Context, *project.CreateProjectRequest) (*project.CreateProjectResponse, error) {
	return &project.CreateProjectResponse{
		Id: "test-project-id-1",
	}, nil
}
