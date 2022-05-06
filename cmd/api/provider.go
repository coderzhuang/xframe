package api

import (
	"go.uber.org/dig"
	"xframe/internal/access/grpc/server"
	handlerGoods "xframe/internal/access/http/handler/goods"
	"xframe/internal/core"
	repoGoods "xframe/internal/infrastructure/repository/goods"
	serviceGoods "xframe/internal/service/goods"
	"xframe/pkg"
)

func GetContainer() *dig.Container {
	container := dig.New()
	_ = container.Provide(pkg.NewMysql)
	_ = container.Provide(pkg.NewRedis)
	_ = container.Provide(repoGoods.New)
	_ = container.Provide(serviceGoods.New)
	_ = container.Provide(handlerGoods.New)
	_ = container.Provide(server.New)
	_ = container.Provide(core.NewHttpServer)
	_ = container.Provide(core.NewGrpcServer)
	return container
}
