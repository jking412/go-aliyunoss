package oss

import (
	"aliyunoss/pkg/logger"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/spf13/viper"
)

var Client *oss.Client
var Bucket *oss.Bucket

func InitOss() {
	var err error
	Client, err = oss.New(viper.GetString("oss.endpoint"), viper.GetString("oss.accessKeyID"), viper.GetString("oss.accessKeySecret"))
	if err != nil {
		logger.Error(err)
	}
	Bucket, err = Client.Bucket(viper.GetString("oss.BucketName"))
	if err != nil {
		logger.Error(err)
	}
}

func UploadFile(objectName, filePath string) {
	err := Bucket.PutObjectFromFile(objectName, filePath)
	if err != nil {
		logger.Error(err)
	}
}

func DownloadFile(objectName, filePath string) {
	err := Bucket.GetObjectToFile(objectName, filePath)
	if err != nil {
		logger.Error(err)
	}
}

func ListFile() {
	marker := ""
	for {
		lor, err := Bucket.ListObjects(oss.Marker(marker))
		if err != nil {
			logger.Error(err)
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
func DeleteFile(objectName string) {
	err := Bucket.DeleteObject(objectName)
	if err != nil {
		logger.Error(err)
	}
}
