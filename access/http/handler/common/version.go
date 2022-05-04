package common

import (
	"github.com/gin-gonic/gin"
	"test/config"
	"test/pkg"
)

func Version(c *gin.Context) {
	pkg.ResponseSuc(c, map[string]string{
		"BuildVersion": config.BuildVersion,
		"BuildAt":      config.BuildAt,
	})
}
