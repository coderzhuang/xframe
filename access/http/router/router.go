package router

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/dig"
	"xframe/access/http/handler/common"
	"xframe/access/http/handler/goods"
)

func InitRout(s *gin.Engine, c *dig.Container) {
	s.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	commonGroup := s.Group("/")
	{
		commonGroup.GET("/version", common.Version)
	}

	goodsGroup := s.Group("/goods")
	_ = c.Invoke(func(HandlerGoods *goods.HandlerGoods) {
		goodsGroup.POST("/", HandlerGoods.Add)
		goodsGroup.GET("/", HandlerGoods.Info)
	})
}
