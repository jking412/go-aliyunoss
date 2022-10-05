package viper

import (
	"aliyunoss/pkg/logger"
	"github.com/spf13/viper"
)

func InitViper(configPath string) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configPath)
	err := viper.ReadInConfig()
	if err != nil {
		logger.Error(configPath, err)
	}
	viper.WatchConfig()
	viper.AutomaticEnv()
}
