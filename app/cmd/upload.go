package cmd

import (
	"aliyunoss/pkg/logger"
	"github.com/spf13/cobra"
)

var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "upload file to oss",
	Long:  "接受两个参数，第一个参数为OSS上的文件名，第二个参数为本地文件名",
	Args:  cobra.ExactArgs(2),
	Run:   upload,
}

func init() {
	rootCmd.AddCommand(uploadCmd)
}

func upload(command *cobra.Command, arg []string) {
	objectName := arg[0]
	filePath := arg[1]
	err := bucket.PutObjectFromFile(objectName, filePath)
	if err != nil {
		logger.HandleError(err)
	}
}
