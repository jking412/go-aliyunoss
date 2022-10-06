package controller

import (
	"aliyunoss/pkg/viperlib"
	"encoding/json"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

type OssRep struct {
	Message string   `json:"message"`
	Data    []string `json:"data"`
	Error   string   `json:"error"`
}

func TestOssShowController(t *testing.T) {
	viperlib.InitViper("../../", "config.yml")
	a := assert.New(t)

	ossRep := &OssRep{}
	rep, err := http.Get("http://localhost:" + viper.GetString("web.port") + "/oss/show")

	defer rep.Body.Close()

	a.Nil(err)
	a.Equal(200, rep.StatusCode)

	data, _ := ioutil.ReadAll(rep.Body)

	err = json.Unmarshal(data, ossRep)
	a.Nil(err)
	a.Equal("获取成功", ossRep.Message)
	t.Log(ossRep.Data)
	if ossRep.Error != "" {
		t.Log(ossRep.Error)
	}
}

func TestOssUploadController(t *testing.T) {
	viperlib.InitViper("../../", "config.yml")
	a := assert.New(t)

	ossRep := &OssRep{}
	ossReq := strings.NewReader(`{"oss_object_name":"test.txt","local_file":"oss_test.txt"}`)
	rep, err := http.Post("http://localhost:"+viper.GetString("web.port")+"/oss/upload", "application/json", ossReq)

	defer rep.Body.Close()

	a.Nil(err)
	a.Equal(200, rep.StatusCode)

	data, _ := ioutil.ReadAll(rep.Body)

	err = json.Unmarshal(data, ossRep)
	a.Nil(err)
	a.Equal("上传成功", ossRep.Message)
	if ossRep.Error != "" {
		t.Log(ossRep.Error)
	}
}
