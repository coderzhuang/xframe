package router

import (
	"github.com/coderzhuang/core"
	"github.com/gin-gonic/gin"
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
				type Response struct {
					BuildVersion string `json:"build_version"` //
					BuildAt      string `json:"build_at"`      //
				}
				core.ResponseSuc(c, Response{
					BuildVersion: "v1.0.0",
					BuildAt:      "2023-01-01",
				})
			})
			base.POST("/register", HandlerUser.Register)     // 注册
			base.POST("/login/password", HandlerUser.Login)  // 密码登录
			base.GET("/get-captcha", HandlerUser.GetCaptcha) // 获取图形验证码
		}

		goodsGroup := e.Group("/goods")
		{
			goodsGroup.POST("/", HandlerGoods.Add)
			goodsGroup.GET("/", HandlerGoods.Info)
		}
	}
}
