package common

import (
	"github.com/gin-gonic/gin"
	"xframe/config"
	"xframe/pkg"
)

func Version(c *gin.Context) {
	pkg.ResponseSuc(c, map[string]string{
		"BuildVersion": config.BuildVersion,
		"BuildAt":      config.BuildAt,
	})
}
