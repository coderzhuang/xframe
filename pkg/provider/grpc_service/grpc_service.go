package grpc_service

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"xframe/config"
	"xframe/internal/access/grpc/mall_server"
	grpcMall "xframe/internal/access/grpc/proto/mall"
	"xframe/pkg/application"
)

type GrpcService struct {
	e          *grpc.Server
	ServerMall *mall_server.Mall
}

func New(ServerMall *mall_server.Mall) application.Service {
	return &GrpcService{
		e:          grpc.NewServer(),
		ServerMall: ServerMall,
	}
}

func (s *GrpcService) Run() {
	lis, err := net.Listen("tcp", config.Conf.GrpcServer.Addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return
	}
	grpcMall.RegisterMallServer(s.e, s.ServerMall)
	log.Printf("server listening at %v", lis.Addr())
	go func() {
		if err := s.e.Serve(lis); err != nil {
			log.Fatalf("GrpcService Start failed. %+v", err)
			return
		}
	}()
}

func (s *GrpcService) Shutdown() {
	s.e.GracefulStop()
}
