package middleware

import (
	utils "chat/utils"

	global "chat/global"

	"github.com/gin-gonic/gin"
)

func AuthMid(c *gin.Context) {
	// TODO 这样写不太好

	path := c.Request.URL.Path
	token := ""
	if path == "/api/v1/auth" {
		c.Next()
		return
	} else {
		token_in_header := c.Request.Header.Get(global.Settings.Jwt.AccessToken)
		token_in_cookie, _ := c.Cookie(global.Settings.Jwt.AccessToken)
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
	user_info, err := utils.JwtDecoder(token)
	if err != nil {
		utils.ResponseWithError(c, 403, err)
		return
	}

	c.Set("username", user_info.Payload.Name)
	c.Set("user_id", user_info.Payload.UserID)
	c.Next()
}
