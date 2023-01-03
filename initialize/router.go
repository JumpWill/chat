package initialize

import (
	"chat/constant"
	"chat/router"
	"chat/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter(e *gin.Engine) {
	utils.Logger(constant.Info, map[string]string{"mes": "route init"})
	router.Router(e)
	utils.Logger(constant.Info, map[string]string{"mes": "route init ok"})

}
