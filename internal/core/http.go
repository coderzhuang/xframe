package core

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
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
	return &HttpServer{
		Engine:       gin.New(),
		HandlerGoods: p.HandlerGoods,
	}
}
