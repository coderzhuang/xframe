package http_service

import (
	"github.com/gin-gonic/gin"
	"xframe/config"
	"xframe/internal/access/http/middleware"
)

type ControllerClosure func(r *gin.Engine)

func NewRouter(fn ControllerClosure) *gin.Engine {
	gin.SetMode(config.Conf.Server.Mode)
	e := gin.New()
	_ = e.SetTrustedProxies(config.Conf.Server.TrustedProxies)

	e.Use(gin.Recovery())
	e.Use(middleware.Exception)
	fn(e)
	return e
}
