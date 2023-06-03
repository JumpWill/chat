package v2

import (
	utils "chat/utils"

	"github.com/gin-gonic/gin"
)

func deploymentGet(c *gin.Context) {
	utils.Response(c, 200, "操作成功!", "")
}

func deploymentPost(c *gin.Context) {
	utils.Response(c, 200, "操作成功!", "")
}

func deploymentPut(c *gin.Context) {
	utils.Response(c, 200, "操作成功!", "")

}

func deploymentDelete(c *gin.Context) {
	utils.Response(c, 200, "操作成功!", "")

}

func DeploymentRouter(api *gin.RouterGroup) {

	v1 := api.Group("/v2/deploy")
	{
		v1.GET("", deploymentGet)
		v1.POST("", deploymentPost)
		v1.PUT("", deploymentPut)
		v1.DELETE("", deploymentDelete)
	}

}
