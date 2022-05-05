package main

import (
	"os"
	"xframe/cmd/api"
	"xframe/config"
	"xframe/docs"
	"xframe/pkg"
)

func main() {
	docs.SwaggerInfo_swagger.BasePath = "/"
	// 加载配置
	config.Init()
	// 加载 telemetry
	shutdown := pkg.InitTracer()
	defer shutdown()
	// 启动服务
	if err := api.Stack().Run(os.Args); err != nil {
		panic(err)
	}
}
