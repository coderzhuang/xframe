package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
	"runtime"
	"xframe/config"
	"xframe/docs"
	"xframe/pkg/application"
	"xframe/pkg/provider"
)

var App = &cli.App{
	Version: fmt.Sprintf("%s|%s|%s|%s",
		runtime.GOOS, runtime.GOARCH, config.BuildVersion, config.BuildAt),
	Action: func(c *cli.Context) error {
		docs.SwaggerInfo.BasePath = "/"

		container := provider.GetContainer()
		return container.Invoke(func(app *application.Application) {
			app.Start()
		})
	},
}

func main() {
	// 启动服务
	if err := App.Run(os.Args); err != nil {
		panic(err)
	}
}
