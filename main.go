package main

import (
	"os"
	"test/cmd/api"
	"test/config"
	"test/pkg"
)

func main() {
	// 加载配置
	config.Init()
	// 加载 telemetry
	shutdown := pkg.InitTracer()
	defer shutdown()
	// 启动服务
	if err := api.Stack().Run(os.Args); err != nil {
		panic(err)
	}
	select {}
}
