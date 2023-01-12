package v1

import (
	. "chat/utils"
	"github.com/gin-gonic/gin"
)

func namespaceGet(c *gin.Context) {
	Response(c, 200, "操作成功!", nil)

}

func namespacePost(c *gin.Context) {
	Response(c, 200, "操作成功!", "hah")

}

func namespacePut(c *gin.Context) {
	Response(c, 200, "操作成功!", "hah")

}

func namespaceDelete(c *gin.Context) {
	Response(c, 200, "操作成功!", "hah")

}

func NamespaceRouter(api *gin.RouterGroup) {

	v1 := api.Group("/v1/ns")
	{
		v1.GET("", namespaceGet)
		v1.POST("", namespacePost)
		v1.PUT("", namespacePut)
		v1.DELETE("", namespaceDelete)
	}

}
