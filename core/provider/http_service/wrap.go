package http_service

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"xframe/pkg/config"
)

type ControllerClosure func(r *gin.Engine)

type Option func(e *gin.Engine)

type OptionGroup struct {
	dig.In

	Opts []Option `group:"middle"`
}

func NewRouter(opts OptionGroup) *gin.Engine {
	gin.SetMode(config.Conf.HttpServer.Mode)
	e := gin.New()
	_ = e.SetTrustedProxies(config.Conf.HttpServer.TrustedProxies)

	for _, opt := range opts.Opts {
		opt(e)
	}
	e.Use(gin.Recovery())
	return e
}
