package swag

import (
	"github.com/coderzhuang/core"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"xframe/docs"
)

func Init() core.Middle {
	return func(e *gin.Engine) {
		docs.SwaggerInfo.BasePath = "/"
		e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}
