package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func deploymentGet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "get",
	})
}

func deploymentPost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "post",
	})
}

func deploymentPut(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "put",
	})
}

func deploymentDelete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "delete",
	})
}

func DeploymentRouter(api *gin.RouterGroup) {

	v1 := api.Group("/v1/deploy")
	{
		v1.GET("", deploymentGet)
		v1.POST("", deploymentPost)
		v1.PUT("", deploymentPut)
		v1.DELETE("", deploymentDelete)
	}

}
