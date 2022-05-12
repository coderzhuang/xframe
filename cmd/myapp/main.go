package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"go.uber.org/dig"
	"os"
	"runtime"
	"xframe/core/application"
	"xframe/core/provider/cron_service"
	"xframe/core/provider/grpc_service"
	"xframe/core/provider/http_service"
	"xframe/internal"
	"xframe/internal/access/http/router"
	"xframe/internal/cron"
	"xframe/pkg/config"
)

func InitProvider() *dig.Container {
	container := dig.New()
	// 以下为系统级别服务
	_ = container.Provide(application.New)
	if config.Conf.HttpServer.Switch {
		_ = container.Provide(http_service.New, dig.Group("server"))
		_ = container.Provide(http_service.NewRouter)
		_ = container.Provide(router.InitRoute, dig.Group("middle"))
	}
	if config.Conf.GrpcServer.Switch {
		_ = container.Provide(grpc_service.New, dig.Group("server"))
	}
	if config.Conf.CronServer.Switch {
		_ = container.Provide(cron_service.New, dig.Group("server"))
		_ = container.Provide(cron_service.NewRouter)
		_ = container.Provide(cron.InitCron)
	}
	// 加载业务相关服务
	internal.InitContainer(container)
	return container
}

func main() {
	var App = &cli.App{
		Version: fmt.Sprintf("%s|%s|%s|%s",
			runtime.GOOS, runtime.GOARCH, config.BuildVersion, config.BuildAt),
		Action: func(c *cli.Context) error {
			container := InitProvider()
			return container.Invoke(func(app *application.Application) {
				app.Start()
			})
		},
	}
	// 启动服务
	if err := App.Run(os.Args); err != nil {
		panic(err)
	}
}
