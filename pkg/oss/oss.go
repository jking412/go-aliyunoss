package oss

import (
	"aliyunoss/pkg/logger"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var Client *oss.Client
var Bucket *oss.Bucket

func InitOss(endpoint, accessKeyId, accessKeySecret, bucketName string) {
	var err error
	Client, err = oss.New(viper.GetString("oss.endpoint"), viper.GetString("oss.accessKeyID"), viper.GetString("oss.accessKeySecret"))
	if err != nil {
		logger.Error("oss", zap.String("initClient", err.Error()))
	}
	Bucket, err = Client.Bucket(viper.GetString("oss.BucketName"))
	if err != nil {
		logger.Error("oss", zap.String("initBucket", err.Error()))
	}
}

func UploadFile(objectName, filePath string) error {
	err := Bucket.PutObjectFromFile(objectName, filePath)
	return err
}

func DownloadFile(objectName, filePath string) error {
	err := Bucket.GetObjectToFile(objectName, filePath)
	return err
}

func ListFile() ([]string, error) {
	marker := ""
	lsRes := make([]string, 0)
	for {
		lor, err := Bucket.ListObjects(oss.Marker(marker))
		if err != nil {
			return nil, err
		}
		for _, object := range lor.Objects {
			lsRes = append(lsRes, object.Key)
		}
		if lor.IsTruncated {
			marker = lor.NextMarker
		} else {
			break
		}
	}
	return lsRes, nil
}
func DeleteFile(objectName string) error {
	err := Bucket.DeleteObject(objectName)
	return err
}
