package router

import (
	"github.com/gin-gonic/gin"
	handlerGoods "xframe/internal/access/http/handler/goods"
	"xframe/pkg/provider/http_service"
)

func InitRoute(
	HandlerGoods *handlerGoods.HandlerGoods,
) http_service.ControllerClosure {
	return func(e *gin.Engine) {
		goodsGroup := e.Group("/goods")
		{
			goodsGroup.POST("/", HandlerGoods.Add)
			goodsGroup.GET("/", HandlerGoods.Info)
		}
	}
}
