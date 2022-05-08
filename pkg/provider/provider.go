package provider

import (
	"go.uber.org/dig"
	"xframe/internal/access/grpc/server"
	handlerGoods "xframe/internal/access/http/handler/goods"
	"xframe/internal/access/http/router"
	"xframe/internal/cron"
	repoGoods "xframe/internal/repository/goods"
	serviceGoods "xframe/internal/service/goods"
	"xframe/pkg/application"
	"xframe/pkg/mysql"
	"xframe/pkg/provider/cron_service"
	"xframe/pkg/provider/grpc_service"
	"xframe/pkg/redis"

	"xframe/pkg/provider/http_service"
)

func GetContainer() *dig.Container {
	container := dig.New()
	// 以下为系统级别服务
	_ = container.Provide(http_service.New, dig.Group("server"))
	_ = container.Provide(grpc_service.New, dig.Group("server"))
	_ = container.Provide(cron_service.New, dig.Group("server"))
	_ = container.Provide(http_service.NewRouter)
	_ = container.Provide(cron_service.NewRouter)
	_ = container.Provide(router.InitRoute)
	_ = container.Provide(cron.InitCron)
	_ = container.Provide(application.New)
	_ = container.Provide(redis.New)
	_ = container.Provide(mysql.New)

	//  以下为业务级别服务
	_ = container.Provide(server.New)
	_ = container.Provide(repoGoods.New)
	_ = container.Provide(serviceGoods.New)
	_ = container.Provide(handlerGoods.New)
	return container
}
