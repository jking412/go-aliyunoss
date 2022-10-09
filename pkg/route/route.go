package route

import (
	"aliyunoss/app/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {
	router.Use(middleware.Cors(),
		middleware.Logger(),
		middleware.Recovery())

	registerRouter(router)
}
