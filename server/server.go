package server

import (
	"log"
	"net"

	postPkg "github.com/StasJDM/proto-example/pkg/xmp/server/post"
	projectPkg "github.com/StasJDM/proto-example/pkg/xmp/server/project"
	userPkg "github.com/StasJDM/proto-example/pkg/xmp/server/user"

	"github.com/StasJDM/go-grpc-example/server/post"
	"github.com/StasJDM/go-grpc-example/server/project"
	"github.com/StasJDM/go-grpc-example/server/user"
	"google.golang.org/grpc"
)

type AppServiceSet struct {
	UserService    *user.UserServer
	PostService    *post.PostServer
	ProjectService *project.ProjectServer
}

func Run() {
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Listener creation error: %v", err)
	}

	serviceSet := newAppServiceSet()
	grpcServer := newGrpcServer(serviceSet)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC: %v", err)
	}
}

func newAppServiceSet() *AppServiceSet {
	userService := &user.UserServer{}
	postService := &post.PostServer{}
	projectService := &project.ProjectServer{}

	return &AppServiceSet{
		UserService:    userService,
		PostService:    postService,
		ProjectService: projectService,
	}
}

func newGrpcServer(serviceSet *AppServiceSet) *grpc.Server {
	grpcServer := grpc.NewServer()

	userPkg.RegisterUserServer(grpcServer, serviceSet.UserService)
	postPkg.RegisterPostServer(grpcServer, serviceSet.PostService)
	projectPkg.RegisterProjectServer(grpcServer, serviceSet.ProjectService)

	return grpcServer
}
