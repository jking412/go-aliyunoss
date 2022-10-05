package oss

import (
	"aliyunoss/pkg/viper"
	"testing"
)

func TestOss(t *testing.T) {
	viper.InitViper("../../")
	InitOss()
	ListFile()
}
