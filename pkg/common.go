package pkg

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"` //
	Msg  string      `json:"msg"`  //
	Data interface{} `json:"data"` //
}

func ResponseSuc(c *gin.Context, data interface{}) {
	c.JSON(200, Response{
		Code: 0,
		Msg:  "",
		Data: data,
	})
}
func ResponseErr(c *gin.Context, code int, msg string) {
	c.JSON(200, Response{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}
