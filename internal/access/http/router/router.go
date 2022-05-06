package router

import (
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"xframe/internal/access/http/handler/common"
	"xframe/internal/core"
)

func InitRout(s *core.HttpServer) {
	s.Engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	commonGroup := s.Engine.Group("/")
	{
		commonGroup.GET("/version", common.Version)
	}

	goodsGroup := s.Engine.Group("/goods")
	{
		goodsGroup.POST("/", s.HandlerGoods.Add)
		goodsGroup.GET("/", s.HandlerGoods.Info)
	}
}
