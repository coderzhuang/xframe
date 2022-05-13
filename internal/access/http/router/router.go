package router

import (
	"github.com/coderzhuang/core/provider/http_service"
	"github.com/gin-gonic/gin"
	handlerGoods "xframe/internal/access/http/handler/goods"
	"xframe/pkg/common"
	"xframe/pkg/config"
)

func InitRoute(
	HandlerGoods *handlerGoods.HandlerGoods,
) http_service.Middle {
	return func(e *gin.Engine) {
		e.GET("/version", func(c *gin.Context) {
			common.ResponseSuc(c, map[string]string{
				"BuildVersion": config.BuildVersion,
				"BuildAt":      config.BuildAt,
			})
		})
		goodsGroup := e.Group("/goods")
		{
			goodsGroup.POST("/", HandlerGoods.Add)
			goodsGroup.GET("/", HandlerGoods.Info)
		}
	}
}
