package http_service

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"xframe/config"
	"xframe/internal/access/http/middleware"
	"xframe/pkg/common"
)

type ControllerClosure func(r *gin.Engine)

func NewRouter(fn ControllerClosure) *gin.Engine {
	gin.SetMode(config.Conf.Server.Mode)
	e := gin.New()
	_ = e.SetTrustedProxies(config.Conf.Server.TrustedProxies)

	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	e.GET("/version", func(c *gin.Context) {
		common.ResponseSuc(c, map[string]string{
			"BuildVersion": config.BuildVersion,
			"BuildAt":      config.BuildAt,
		})
	})

	e.Use(gin.Recovery())
	e.Use(middleware.Exception)
	fn(e)
	return e
}
