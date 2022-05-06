package core

import (
	"go.uber.org/dig"
	"google.golang.org/grpc"
	"xframe/internal/access/grpc/server"
)

type GrpcServer struct {
	Engine     *grpc.Server
	ServerMall *server.Mall
}

type GrpcServerParam struct {
	dig.In

	ServerMall *server.Mall
}

func NewGrpcServer(p GrpcServerParam) *GrpcServer {
	return &GrpcServer{
		Engine:     grpc.NewServer(),
		ServerMall: p.ServerMall,
	}
}
