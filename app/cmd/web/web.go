package web

import (
	"aliyunoss/pkg/logger"
	"aliyunoss/pkg/route"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var WebCmd = &cobra.Command{
	Use:   "serve",
	Short: "Use serve to start the web server",
	Run:   serve,
}

func serve(cmd *cobra.Command, args []string) {
	router := gin.Default()

	route.InitRouter(router)

	err := router.Run(":" + viper.GetString("web.port"))

	if err != nil {
		logger.Error(err)
	}
}
