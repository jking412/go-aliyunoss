package oss

import (
	"aliyunoss/pkg/viperlib"
	"testing"
)

func TestOss(t *testing.T) {
	viperlib.InitViper("../../", "config.yml")
	InitOss()
	file, err := ListFile()
	if err != nil {
		t.Error(err)
	}
	t.Log(file)
}
