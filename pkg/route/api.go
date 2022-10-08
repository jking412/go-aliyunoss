package route

import (
	"aliyunoss/app/controller/oss"
	"aliyunoss/app/controller/user"
	"github.com/gin-gonic/gin"
)

func registerRouter(router *gin.Engine) {
	router.GET("/ping", ping)

	ossGroup := router.Group("/oss")
	{
		ossGroup.GET("/show", oss.Show)
		ossGroup.POST("/upload", oss.Upload)
		ossGroup.POST("/download", oss.Download)
		ossGroup.POST("/delete", oss.Delete)
	}

	userGroup := router.Group("/user")
	{
		userGroup.POST("/login", user.Login)
		userGroup.POST("/register", user.Register)
	}
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
