package router

import (
	v1 "chat/router/v1"
	v2 "chat/router/v2"

	"github.com/gin-gonic/gin"
)

func Router(e *gin.Engine) {
	// add your router
	routers := []func(*gin.RouterGroup){
		v1.AuthRouter,
		v1.DeploymentRouter,
		v1.NamespaceRouter,

		v2.DeploymentRouter,
		v2.NamespaceRouter,
	}

	api := e.Group("/api")
	{
		for _, add := range routers {
			add(api)
		}
	}

}
