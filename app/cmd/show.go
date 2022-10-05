package cmd

import (
	"aliyunoss/pkg/logger"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "show file from oss",
	Long:  "接受一个参数，为OSS上的文件名",
	Args:  cobra.NoArgs,
	Run:   listBuckets,
}

func init() {
	rootCmd.AddCommand(showCmd)
}

func listBuckets(command *cobra.Command, args []string) {
	marker := ""
	for {
		lor, err := bucket.ListObjects(oss.Marker(marker))
		if err != nil {
			logger.HandleError(err)
		}
		for _, object := range lor.Objects {
			fmt.Println("Object:", object.Key)
		}
		if lor.IsTruncated {
			marker = lor.NextMarker
		} else {
			break
		}
	}
}
