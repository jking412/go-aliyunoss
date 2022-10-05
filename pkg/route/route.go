package route

import "github.com/gin-gonic/gin"

func InitRouter(router *gin.Engine) {
	registerRouter(router)
}
