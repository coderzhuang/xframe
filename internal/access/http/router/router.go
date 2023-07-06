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
			type Response struct {
				BuildVersion string `json:"build_version"` //
				BuildAt      string `json:"build_at"`      //
			}
			core.ResponseSuc(c, Response{
				BuildVersion: "v1.0.0",
				BuildAt:      "2023-01-01",
			})
		})
		goodsGroup := e.Group("/goods")
		{
			goodsGroup.POST("/", HandlerGoods.Add)
			goodsGroup.GET("/", HandlerGoods.Info)
		}
	}
}
