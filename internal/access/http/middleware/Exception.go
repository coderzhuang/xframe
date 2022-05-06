package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"runtime/debug"
	"xframe/pkg"
)

func Exception(c *gin.Context) {
	defer func() {
		err := recover()
		if err == nil {
			return
		}
		log.Printf("%+v\n", err)
		log.Println(string(debug.Stack()))
		msg := "Internal Server Error"
		if v, ok := err.(error); ok {
			msg = v.Error()
		}
		c.JSON(200, pkg.Response{
			Code: 500,
			Msg:  msg,
			Data: nil,
		})
		c.Abort()
	}()
	c.Next()
}
