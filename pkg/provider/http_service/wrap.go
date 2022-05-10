package http_service

import (
	"github.com/chenjiandongx/ginprom"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	b3prop "go.opentelemetry.io/contrib/propagators/b3"
	"go.opentelemetry.io/otel"
	"xframe/config"
	"xframe/docs"
	"xframe/pkg/common"
)

type ControllerClosure func(r *gin.Engine)

func NewRouter(fn ControllerClosure) *gin.Engine {
	gin.SetMode(config.Conf.HttpServer.Mode)
	e := gin.New()
	_ = e.SetTrustedProxies(config.Conf.HttpServer.TrustedProxies)

	// trace
	b3 := b3prop.New()
	otel.SetTextMapPropagator(b3)
	e.Use(otelgin.Middleware(config.Conf.HttpServer.Name, otelgin.WithPropagators(b3)))

	// 添加prometheus 监控
	e.Use(ginprom.PromMiddleware(&ginprom.PromOpts{ExcludeRegexEndpoint: "^/(swagger|metrics)"}))
	e.GET("/metrics", ginprom.PromHandler(promhttp.Handler()))

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
