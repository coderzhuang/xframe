package http_service

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"xframe/config"
	"xframe/docs"
	"xframe/pkg/common"
	"xframe/pkg/telemetry"
)

type ControllerClosure func(r *gin.Engine)

func NewRouter(fn ControllerClosure) *gin.Engine {
	gin.SetMode(config.Conf.HttpServer.Mode)
	e := gin.New()
	_ = e.SetTrustedProxies(config.Conf.HttpServer.TrustedProxies)

	// trace
	telemetry.InitTracer()
	e.Use(otelgin.Middleware(config.Conf.HttpServer.Name))

	// 添加prometheus 监控
	//e.Use(middleware.New(e).Middleware())
	//e.GET("/metrics", gin.WrapH(promhttp.Handler()))

	if config.Conf.Swagger.Switch {
		docs.SwaggerInfo.BasePath = "/"
		e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	e.GET("/version", func(c *gin.Context) {
		common.ResponseSuc(c, map[string]string{
			"BuildVersion": config.BuildVersion,
			"BuildAt":      config.BuildAt,
		})
	})

	e.Use(gin.Recovery())
	fn(e)
	return e
}
