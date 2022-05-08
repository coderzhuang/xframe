package router

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"xframe/internal/access/http/handler/common"
	handlerGoods "xframe/internal/access/http/handler/goods"
	"xframe/pkg/provider/http_service"
)

func InitRoute(
	HandlerGoods *handlerGoods.HandlerGoods,
) http_service.ControllerClosure {
	return func(e *gin.Engine) {
		e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

		commonGroup := e.Group("/")
		{
			commonGroup.GET("/version", common.Version)
		}

		goodsGroup := e.Group("/goods")
		{
			goodsGroup.POST("/", HandlerGoods.Add)
			goodsGroup.GET("/", HandlerGoods.Info)
		}
	}
}
