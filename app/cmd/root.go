package cmd

import (
	"aliyunoss/app/cmd/web"
	"aliyunoss/pkg/oss"
	"aliyunoss/pkg/viperlib"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "oss",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		viperlib.InitViper(configPath, configName)
		oss.InitOss()
	},
}

var configPath string
var configName string

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(web.WebCmd)
	rootCmd.PersistentFlags().StringVarP(&configPath, "config", "c", ".", "config file path")
	rootCmd.PersistentFlags().StringVarP(&configName, "name", "n", "config", "config file name")
}
