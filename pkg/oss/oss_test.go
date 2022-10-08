package oss

import (
	"aliyunoss/pkg/viperlib"
	"github.com/spf13/viper"
	"testing"
)

func TestOss(t *testing.T) {
	viperlib.InitViper("../../", "config.yml")
	InitOss(viper.GetString("oss.endpoint"),
		viper.GetString("oss.accessKeyID"),
		viper.GetString("oss.accessKeySecret"),
		viper.GetString("oss.BucketName"))
	file, err := ListFile()
	if err != nil {
		t.Error(err)
	}
	t.Log(file)
}
