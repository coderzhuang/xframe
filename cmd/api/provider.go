package api

import (
	"go.uber.org/dig"
	"xframe/internal/access/grpc/server"
	handlerGoods "xframe/internal/access/http/handler/goods"
	repoGoods "xframe/internal/infrastructure/repository/goods"
	serviceGoods "xframe/internal/service/goods"
	"xframe/pkg"
)

var container *dig.Container

func GetContainer() *dig.Container {
	if container != nil {
		return container
	}
	container = dig.New()
	_ = container.Provide(pkg.NewMysql)
	_ = container.Provide(pkg.NewRedis)
	_ = container.Provide(repoGoods.New)
	_ = container.Provide(serviceGoods.New)
	_ = container.Provide(handlerGoods.New)
	_ = container.Provide(server.New)
	return container
}
