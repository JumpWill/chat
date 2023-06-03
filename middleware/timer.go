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
		"Code":   fmt.Sprint(c.Writer.Status()),
		"Method": c.Request.Method,
		"Path":   c.Request.URL.Path,
		"IP":     c.ClientIP(),
		"UA":     c.Request.UserAgent(),
		"Time":   fmt.Sprint(time.Duration(time.Since(start).Microseconds())),
		"Error":  c.Errors.ByType(gin.ErrorTypePrivate).String(),
	})

}
