package main

import (
	"github.com/coderzhuang/core"
	"xframe/config"
	"xframe/internal"
	"xframe/services"
)

func main() {
	services.Init()
	internal.Init()
	core.Run(config.Cfg)
}
