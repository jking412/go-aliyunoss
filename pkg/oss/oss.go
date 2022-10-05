package oss

import (
	"aliyunoss/pkg/logger"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

var Client *oss.Client
var Bucket *oss.Bucket

func init() {
	var err error
	Client, err = oss.New(viper.GetString("endpoint"), viper.GetString("accessKeyID"), viper.GetString("accessKeySecret"))
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
	Bucket, err = Client.Bucket(viper.GetString("BucketName"))
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
}

func UploadFile(objectName, filePath string) {
	err := Bucket.PutObjectFromFile(objectName, filePath)
	if err != nil {
		logger.HandleError(err)
	}
}

func DownloadFile(objectName, filePath string) {
	err := Bucket.GetObjectToFile(objectName, filePath)
	if err != nil {
		logger.HandleError(err)
	}
}

func ListFile() {
	marker := ""
	for {
		lor, err := Bucket.ListObjects(oss.Marker(marker))
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
func DeleteFile(objectName string) {
	err := Bucket.DeleteObject(objectName)
	if err != nil {
		logger.HandleError(err)
	}
}
