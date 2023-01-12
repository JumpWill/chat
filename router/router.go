package router

import (
	"chat/router/v1"
	"github.com/gin-gonic/gin"
)

func Router(e *gin.Engine) {
	// add your router 
	routers := []func(*gin.RouterGroup){
		v1.DeploymentRouter,
		v1.NamespaceRouter,
	}

	api := e.Group("/api")
	{
		for _, add := range routers {
			add(api)
		}
	}

}
