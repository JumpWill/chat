package v1

import (
	. "chat/utils"
	"github.com/gin-gonic/gin"
)

func deploymentGet(c *gin.Context) {
	Response(c, 200, "操作成功!", "")
}

func deploymentPost(c *gin.Context) {
	Response(c, 200, "操作成功!", "")
}

func deploymentPut(c *gin.Context) {
	Response(c, 200, "操作成功!", "")

}

func deploymentDelete(c *gin.Context) {
	Response(c, 200, "操作成功!", "")

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
