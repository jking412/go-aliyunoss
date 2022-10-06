package route

import (
	"aliyunoss/app/controller"
	"github.com/gin-gonic/gin"
)

func registerRouter(router *gin.Engine) {
	router.GET("/ping", ping)

	ossGroup := router.Group("/oss")
	{
		ossGroup.GET("/show", controller.Show)
		ossGroup.POST("/upload", controller.Upload)
		ossGroup.POST("/download", controller.Download)
		ossGroup.POST("/delete", controller.Delete)
	}
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
