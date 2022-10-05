package cmd

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var client *oss.Client
var bucket *oss.Bucket

var rootCmd = &cobra.Command{
	Use:   "oss",
	Short: "oss is a command line tool for oss",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	initViper()
	var err error
	client, err = oss.New(viper.GetString("endpoint"), viper.GetString("accessKeyID"), viper.GetString("accessKeySecret"))
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
	bucket, err = client.Bucket(viper.GetString("bucketName"))
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
}

func initViper() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	viper.WatchConfig()
}
