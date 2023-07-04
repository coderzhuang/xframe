package router

import (
	"github.com/coderzhuang/core"
	"github.com/gin-gonic/gin"
	handlerGoods "xframe/internal/access/http/handler/goods"
)

func InitRoute(
	HandlerGoods *handlerGoods.HandlerGoods,
) core.Middle {
	return func(e *gin.Engine) {
		e.GET("/version", func(c *gin.Context) {
			c.JSON(200, map[string]string{
				"BuildVersion": "1111",
				"BuildAt":      "",
			})
		})
		goodsGroup := e.Group("/goods")
		{
			goodsGroup.POST("/", HandlerGoods.Add)
			goodsGroup.GET("/", HandlerGoods.Info)
		}
	}
}
