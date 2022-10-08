package viperlib

import (
	"github.com/spf13/viper"
	"log"
)

func InitViper(Config ...string) {
	viper.SetConfigType("yml")
	viper.AddConfigPath(Config[0])
	viper.SetConfigName(Config[1])
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("viper.ReadInConfig() failed, err:%v", err)
		log.Printf("ConfigName: %s, ConfigPath: %s", Config[1], Config[0])
	}
	viper.WatchConfig()
	viper.AutomaticEnv()
}
