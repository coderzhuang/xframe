package internal

import (
	"github.com/coderzhuang/core"
	"go.uber.org/dig"
	"xframe/internal/access/grpc/mall_server"
	handlerGoods "xframe/internal/access/http/handler/goods"
	"xframe/internal/access/http/router"
	repoGoods "xframe/internal/repository/goods"
	serviceGoods "xframe/internal/service/goods"
	"xframe/services/mysql"
	"xframe/services/redis"
)

func Init() {
	core.RegistryService(
		func(container *dig.Container) {
			_ = container.Provide(func() *core.Option {
				return &core.Option{
					Mode:           core.Conf.HttpServer.Mode,
					TrustedProxies: core.Conf.HttpServer.TrustedProxies,
					Addr:           core.Conf.HttpServer.Addr,
				}
			})
			_ = container.Provide(func() *core.OptionRpc {
				return &core.OptionRpc{
					Addr: core.Conf.GrpcServer.Addr,
				}
			})
			_ = container.Provide(router.InitRoute, dig.Group("middle"))
			//_ = container.Provide(telemetry.Init, dig.Group("middle"))
			//_ = container.Provide(swag.Init, dig.Group("middle"))
			_ = container.Provide(redis.New)
			_ = container.Provide(mysql.New)
			//  以下为业务-接入层
			_ = container.Provide(mall_server.RegisterServer, dig.Group("grpc_server"))
			_ = container.Provide(handlerGoods.New)
			//  以下为业务-服务层
			_ = container.Provide(serviceGoods.New)
			//  以下为业务-数据层
			_ = container.Provide(repoGoods.New)
		})
}
