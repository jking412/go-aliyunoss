package main

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

func init() {
	initViper()
}

func main() {
	client, err := oss.New(viper.GetString("endpoint"), viper.GetString("accessKeyID"), viper.GetString("accessKeySecret"))
	if err != nil {
		handleError(err)
	}
	logrus.Info("oss client init success")
	bucket, err := client.Bucket(viper.GetString("bucketName"))
	if err != nil {
		handleError(err)
	}
	logrus.Info("bucket:", bucket.BucketName)
}

func downloadFile(bucket *oss.Bucket, objectName string, fileName string) {
	err := bucket.GetObjectToFile(objectName, fileName)
	if err != nil {
		handleError(err)
	}
}

func uploadFile(bucket *oss.Bucket, objectName string, fileName string) {
	err := bucket.PutObjectFromFile(objectName, fileName)
	if err != nil {
		handleError(err)
	}
}

func listBuckets(bucket *oss.Bucket) {

	marker := ""
	for {
		lor, err := bucket.ListObjects(oss.Marker(marker))
		if err != nil {
			handleError(err)
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

func deleteFile(bucket *oss.Bucket, objectName string) {
	err := bucket.DeleteObject(objectName)
	if err != nil {
		handleError(err)
	}
}

func initViper() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	viper.WatchConfig()
}

func handleError(err error) {
	logrus.Error(err)
	os.Exit(-1)
}
