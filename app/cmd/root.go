package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "oss",
	Short: "oss is a command line tool for oss",
	Run: func(cmd *cobra.Command, args []string) {
		initViper(ConfigPath)
	},
}

var ConfigPath string

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.Flags().StringVarP(&ConfigPath, "config", "c", "../", "config file path")
}

func initViper(configPath string) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configPath)
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Info(configPath)
		panic(err)
	}
	viper.WatchConfig()
}
