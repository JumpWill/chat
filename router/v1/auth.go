package v1

import (
	schema "chat/schema"
	utils "chat/utils"

	"github.com/gin-gonic/gin"
)

func authGet(c *gin.Context) {

	payload := schema.Payload{
		Name:   "hjy",
		UserID: "111",
		Role:   "die",
	}

	access_token, _ := utils.JwtEncoder(payload)
	token := schema.Token{
		RefreshToken: access_token,
		AccessToken:  access_token,
	}
	utils.Response(c, 200, "操作成功!", token)

}

func authPost(c *gin.Context) {

	// access_token, _ := utils.JwtEncoder("xxxx", "paa")
	// token := schema.Token{
	// 	RefreshToken: access_token,
	// 	AccessToken:  access_token,
	// }
	utils.Response(c, 200, "操作成功!", "token")

}

func authPut(c *gin.Context) {
	utils.Response(c, 200, "操作成功!", "hah")

}

func authDelete(c *gin.Context) {
	utils.Response(c, 200, "操作成功!", "hah")

}

func AuthRouter(api *gin.RouterGroup) {

	v1 := api.Group("/v1/auth")
	{
		v1.GET("", authGet)
		v1.POST("", authPost)
		// v1.PUT("", authPut)
		// v1.DELETE("", authDelete)
	}

}
