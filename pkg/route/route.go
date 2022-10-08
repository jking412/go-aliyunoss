package route

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {
	registerRouter(router)

	router.Use(gin.Logger(),
		gin.Recovery())
}
