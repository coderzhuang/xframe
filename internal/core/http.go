package core

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"xframe/config"
	handlerGoods "xframe/internal/access/http/handler/goods"
)

type HttpServer struct {
	Engine       *gin.Engine
	HandlerGoods *handlerGoods.HandlerGoods
}

type HttpServerParam struct {
	dig.In

	HandlerGoods *handlerGoods.HandlerGoods
}

func NewHttpServer(p HttpServerParam) *HttpServer {
	// 初始化服务，注册路由
	if config.Conf.Common.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	s := gin.New()
	_ = s.SetTrustedProxies(config.Conf.Server.TrustedProxies)
	return &HttpServer{
		Engine:       s,
		HandlerGoods: p.HandlerGoods,
	}
}
