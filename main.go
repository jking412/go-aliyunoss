package main

import (
	"aliyunoss/boot"
	"aliyunoss/pkg/logger"
	"aliyunoss/pkg/route"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func init() {
	boot.Boot()
}

func main() {

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	route.InitRouter(router)

	err := router.Run(":" + viper.GetString("web.port"))

	if err != nil {
		logger.Error("web", zap.String("run", err.Error()))
	}
}
