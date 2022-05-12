package internal

import (
	"go.uber.org/dig"
	"xframe/internal/access/grpc/mall_server"
	handlerGoods "xframe/internal/access/http/handler/goods"
	repoGoods "xframe/internal/repository/goods"
	serviceGoods "xframe/internal/service/goods"
	"xframe/pkg/mysql"
	"xframe/pkg/redis"
)

func InitContainer(container *dig.Container) {
	//_ = container.Provide(telemetry.Init, dig.Group("middle"))
	//_ = container.Provide(swag.Init, dig.Group("middle"))
	_ = container.Provide(redis.New)
	_ = container.Provide(mysql.New)
	//  以下为业务-接入层
	_ = container.Provide(mall_server.New)
	_ = container.Provide(handlerGoods.New)
	//  以下为业务-服务层
	_ = container.Provide(serviceGoods.New)
	//  以下为业务-数据层
	_ = container.Provide(repoGoods.New)
}
