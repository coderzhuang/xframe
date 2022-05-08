package provider

import (
	"go.uber.org/dig"
	"xframe/config"
	"xframe/internal/access/grpc/mall_server"
	handlerGoods "xframe/internal/access/http/handler/goods"
	"xframe/internal/access/http/router"
	"xframe/internal/cron"
	repoGoods "xframe/internal/repository/goods"
	serviceGoods "xframe/internal/service/goods"
	"xframe/pkg/application"
	"xframe/pkg/mysql"
	"xframe/pkg/provider/cron_service"
	"xframe/pkg/provider/grpc_service"
	"xframe/pkg/provider/http_service"
	"xframe/pkg/redis"
)

func GetContainer() *dig.Container {
	container := dig.New()
	// 以下为系统级别服务
	if config.Conf.HttpServer.Switch {
		_ = container.Provide(http_service.New, dig.Group("server"))
		_ = container.Provide(http_service.NewRouter)
		_ = container.Provide(router.InitRoute)
	}
	if config.Conf.GrpcServer.Switch {
		_ = container.Provide(grpc_service.New, dig.Group("server"))
	}
	if config.Conf.CronServer.Switch {
		_ = container.Provide(cron_service.New, dig.Group("server"))
		_ = container.Provide(cron_service.NewRouter)
		_ = container.Provide(cron.InitCron)
	}
	_ = container.Provide(application.New)
	_ = container.Provide(redis.New)
	_ = container.Provide(mysql.New)

	//  以下为业务-接入层
	_ = container.Provide(mall_server.New)
	_ = container.Provide(handlerGoods.New)
	//  以下为业务-服务层
	_ = container.Provide(serviceGoods.New)
	//  以下为业务-数据层
	_ = container.Provide(repoGoods.New)
	return container
}
