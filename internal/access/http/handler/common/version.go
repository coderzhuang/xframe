package common

import (
	"github.com/gin-gonic/gin"
	"xframe/config"
	"xframe/pkg/common"
)

func Version(c *gin.Context) {
	common.ResponseSuc(c, map[string]string{
		"BuildVersion": config.BuildVersion,
		"BuildAt":      config.BuildAt,
	})
}
