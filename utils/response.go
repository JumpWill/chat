package utils

import (
	"github.com/gin-gonic/gin"
)

func Response(c *gin.Context, code int, msg string, data interface{}) {

	c.JSON(code, map[string]interface{}{
		"code": code,
		"msg":  msg,
		"data": data,
	})

}
