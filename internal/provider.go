package internal

import (
	"github.com/coderzhuang/core"
	"go.uber.org/dig"
	"xframe/internal/access/grpc/mall_server"
	handlerGoods "xframe/internal/access/http/handler/goods"
	"xframe/internal/access/http/router"
	repoGoods "xframe/internal/repository/goods"
	serviceGoods "xframe/internal/service/goods"
)

func Init() {
	core.RegistryService(
		func(container *dig.Container) {
			_ = container.Provide(router.InitRoute, dig.Group("middle"))

			//  以下为业务-接入层
			_ = container.Provide(mall_server.RegisterServer, dig.Group("grpc_server"))
			_ = container.Provide(handlerGoods.New)
			//  以下为业务-服务层
			_ = container.Provide(serviceGoods.New)
			//  以下为业务-数据层
			_ = container.Provide(repoGoods.New)
		})
}
