package route

import (
	"aliyunoss/app/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {
	registerRouter(router)

	router.Use(middleware.Logger(),
		middleware.Recovery())
}
