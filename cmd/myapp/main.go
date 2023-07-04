package main

import (
	"github.com/coderzhuang/core"
	"xframe/internal"
	"xframe/services"
)

func main() {
	services.Init()
	internal.Init()
	core.Run()
}
