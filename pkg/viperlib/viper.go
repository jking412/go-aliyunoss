package viperlib

import (
	"aliyunoss/pkg/logger"
	"github.com/spf13/viper"
)

func InitViper(Config ...string) {
	viper.SetConfigType("yml")
	viper.AddConfigPath(Config[0])
	viper.SetConfigName(Config[1])
	err := viper.ReadInConfig()
	if err != nil {
		logger.Error(Config[0])
		logger.Error(Config[1])
		logger.Error(err)
	}
	viper.WatchConfig()
	viper.AutomaticEnv()
}
