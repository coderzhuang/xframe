package router

import (
	"fmt"
	"github.com/coderzhuang/core"
	"github.com/gin-gonic/gin"
	"xframe/config"
	handlerGoods "xframe/internal/access/http/handler/goods"
)

func InitRoute(
	HandlerGoods *handlerGoods.HandlerGoods,
) core.Middle {
	return func(e *gin.Engine) {
		e.GET("/version", func(c *gin.Context) {
			fmt.Println(config.Cfg.Text)
			fmt.Println(config.Cfg.DB)
			fmt.Println(config.Cfg.Redis)
			c.JSON(200, map[string]string{
				"BuildVersion": "1111",
				"BuildAt":      config.Cfg.Text,
			})
		})
		goodsGroup := e.Group("/goods")
		{
			goodsGroup.POST("/", HandlerGoods.Add)
			goodsGroup.GET("/", HandlerGoods.Info)
		}
	}
}
