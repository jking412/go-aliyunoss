package cmd

import (
	"aliyunoss/pkg/logger"
	"github.com/spf13/cobra"
)

var downloadCmd = &cobra.Command{
	Use:     "download",
	Aliases: []string{"get"},
	Short:   "download file from oss",
	Long:    "接受两个参数，第一个参数为OSS上的文件名，第二个参数为本地文件名",
	Args:    cobra.ExactArgs(2),
	Run:     downloadFile,
}

func init() {
	rootCmd.AddCommand(downloadCmd)
}

func downloadFile(command *cobra.Command, arg []string) {
	objectName := arg[0]
	fileName := arg[1]
	err := bucket.GetObjectToFile(objectName, fileName)
	if err != nil {
		logger.HandleError(err)
	}
}
