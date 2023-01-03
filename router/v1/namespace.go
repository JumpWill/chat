package v1

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func namespaceGet(c *gin.Context) {
	fmt.Println("xxx")
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "get",
	})
}

func namespacePost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "post",
	})
}

func namespacePut(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "put",
	})
}

func namespaceDelete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "delete",
	})
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
