package router

import (
	"github.com/coderzhuang/core"
	"github.com/gin-gonic/gin"
	"xframe/config"
	"xframe/internal/access/http/dto"
	"xframe/internal/access/http/handler"
	handlerGoods "xframe/internal/access/http/handler/goods"
)

func InitRoute(
	HandlerGoods *handlerGoods.HandlerGoods,
	HandlerUser *handler.HandlerUser,
) core.Middle {
	return func(e *gin.Engine) {
		base := e.Group("/")
		{
			base.GET("/version", func(c *gin.Context) {
				core.ResponseSuc(c, dto.VersionResp{
					BuildVersion: config.BuildVersion,
					BuildAt:      config.BuildAt,
				})
			})
			base.POST("/register", HandlerUser.Register)     // 注册
			base.POST("/login/password", HandlerUser.Login)  // 密码登录
			base.GET("/get-captcha", HandlerUser.GetCaptcha) // 获取图形验证码
		}

		userGroup := e.Group("/user")
		{
			userGroup.GET("", HandlerUser.GetUser) // 获取用户信息
		}

		goodsGroup := e.Group("/goods")
		{
			goodsGroup.POST("/", HandlerGoods.Add)
			goodsGroup.GET("/", HandlerGoods.Info)
		}
	}
}
