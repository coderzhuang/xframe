package services

import (
	"github.com/coderzhuang/core"
	"go.uber.org/dig"
	"xframe/services/mysql"
	"xframe/services/redis"
)

func Init() {
	core.RegistryService(
		func(container *dig.Container) {
			//_ = container.Provide(telemetry.Init, dig.Group("middle"))
			//_ = container.Provide(swag.Init, dig.Group("middle"))
			_ = container.Provide(redis.New)
			_ = container.Provide(mysql.New)
		})
}
