package middleware

import (
	utils "chat/utils"

	"chat/global"

	"github.com/gin-gonic/gin"
)

func AuthMid(c *gin.Context) {
	// TODO 这样写不太好
	path := c.Request.URL.Path
	if path == "/api/v1/auth" {

	} else {
		token_in_header := c.Request.Header.Get(global.Settings.Jwt.AccessToken)
		token_in_cookie, _ := c.Cookie(global.Settings.Jwt.AccessToken)
		token := ""
		if token_in_header != "" {
			token = token_in_header
		} else if token_in_cookie != "" {
			token = token_in_cookie
		}
		if token == "" {
			utils.Response(c, 403, "请先完成认证!", nil)
			c.Abort()
		}

	}
	c.Next()
}
