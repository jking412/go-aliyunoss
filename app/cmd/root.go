package cmd

import (
	"aliyunoss/app/cmd/web"
	"aliyunoss/pkg/oss"
	"aliyunoss/pkg/viper"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "oss",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		viper.InitViper(ConfigPath)
		oss.InitOss()
	},
}

var ConfigPath string

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(web.WebCmd)
	rootCmd.Flags().StringVarP(&ConfigPath, "config", "c", ".", "config file path")
}
