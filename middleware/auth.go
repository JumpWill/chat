package middleware

import (
	"chat/global"
	"fmt"
	"github.com/gin-gonic/gin"
)

func AuthMid(c *gin.Context) {
	token_in_header := c.Request.Header.Get(global.Settings.Jwt.Header)
	token_in_query := c.Query(global.Settings.Jwt.Query)
	token_in_cookie, _ := c.Cookie(global.Settings.Jwt.Cookie)
	var token string
	switch {
	case token_in_header != "":
		token = token_in_header
	case token_in_query != "":
		token = token_in_query
	case token_in_cookie != "":
		token = token_in_cookie
	default:
		{
			// TODO 定义全局返回
			// c.Abort()
		}
	}
	// TODO 集成casbin,完成权限鉴定
	fmt.Println("token is", token)
	c.Next()

}
