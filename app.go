package main

import (
	"chat/initialize"
	"chat/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	// init config for app
	initialize.InitConfig()
	// init log
	initialize.InitLog()
	// init db
	initialize.InitDB()
	// init redis
	initialize.InitRedis()

	r := gin.New()
	// add middleware
	r.Use(middleware.TimerMid, middleware.AuthMid, gin.Recovery())
	// init router
	initialize.InitRouter(r)

	r.Run()
}
