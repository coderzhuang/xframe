package internal

import (
	"github.com/coderzhuang/core"
	"go.uber.org/dig"
	"xframe/internal/access/cron"
	"xframe/internal/access/http/handler"
	"xframe/internal/access/http/router"

	handlerGoods "xframe/internal/access/http/handler/goods"
	repoGoods "xframe/internal/repository/goods"
	serviceGoods "xframe/internal/service/goods"

	repoUser "xframe/internal/repository/user"
	serviceUser "xframe/internal/service/user"
)

func Init() {
	core.RegistryService(
		func(container *dig.Container) {
			//_ = container.Provide(mall_server.RegisterServer, dig.Group("grpc_server"))
			_ = container.Provide(router.InitRoute, dig.Group("middle"))
			_ = container.Provide(cron.InitCron)
			//
			_ = container.Provide(handlerGoods.New)
			_ = container.Provide(serviceGoods.New)
			_ = container.Provide(repoGoods.New)

			_ = container.Provide(handler.New)
			_ = container.Provide(serviceUser.New)
			_ = container.Provide(repoUser.New)
		})
}
