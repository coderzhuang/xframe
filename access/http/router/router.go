package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"xframe/access/http/handler/common"
	"xframe/access/http/handler/goods"
)

func InitRout(s *gin.Engine, c *dig.Container) {
	err := c.Invoke(func(HandlerGoods *goods.HandlerGoods) {
		commonGroup := s.Group("/")
		{
			commonGroup.GET("/version", common.Version)
		}
		goodsGroup := s.Group("/goods")
		{
			goodsGroup.POST("/", HandlerGoods.Add)
			goodsGroup.GET("/", HandlerGoods.Info)
		}
	})
	if err != nil {
		fmt.Println(err.Error())
	}
}
