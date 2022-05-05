package main

import (
	"os"
	"xframe/cmd/api"
)

func main() {
	// 启动服务
	if err := api.App.Run(os.Args); err != nil {
		panic(err)
	}
}
