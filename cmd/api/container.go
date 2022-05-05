package api

import (
	"go.uber.org/dig"
	"xframe/access/grpc/server"
	handlerGoods "xframe/access/http/handler/goods"
	repoGoods "xframe/infrastructure/repository/goods"
	"xframe/pkg"
	serviceGoods "xframe/service/goods"
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
