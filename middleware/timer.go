package middleware

import (
	"chat/constant"
	"chat/utils"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func TimerMid(c *gin.Context) {
	start := time.Now()
	c.Next()

	utils.Logger(constant.Info, map[string]string{
		"status": fmt.Sprint(c.Writer.Status()),
		"method": c.Request.Method,
		"path":   c.Request.URL.Path,
		"ip":     c.ClientIP(),
		"ua":     c.Request.UserAgent(),
		"time":   fmt.Sprint(time.Duration(time.Since(start).Microseconds())),
		"error":  c.Errors.ByType(gin.ErrorTypePrivate).String(),
	})

}
