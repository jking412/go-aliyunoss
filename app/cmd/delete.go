package cmd

import (
	"aliyunoss/pkg/logger"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete file from oss",
	Long:  "接受一个参数，为OSS上的文件名",
	Args:  cobra.ExactArgs(1),
	Run:   deleteFile,
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

func deleteFile(command *cobra.Command, arg []string) {
	objectName := arg[0]
	err := bucket.DeleteObject(objectName)
	if err != nil {
		logger.HandleError(err)
	}
}
