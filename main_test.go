package main

import (
	"aliyunoss/app/model"
	"aliyunoss/app/model/utils"
	"aliyunoss/pkg/oss"
	"aliyunoss/pkg/request"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"testing"
)

type PingRep struct {
	Message string `json:"message"`
}

type RegisterRep struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

func TestMain(m *testing.M) {
	go main()
	m.Run()
}

func TestWeb(t *testing.T) {
	a := assert.New(t)

	pingRep := &PingRep{}

	err := request.GetRequest("http://localhost:"+viper.GetString("web.port")+"/ping", pingRep)

	a.Nil(err)

	a.Equal("pong", pingRep.Message)
}

func TestRegister(t *testing.T) {
	a := assert.New(t)

	registerRep := &RegisterRep{}

	body := `{"username":"test","password":"123456"}`

	err := request.PostRequest("http://localhost:"+viper.GetString("web.port")+"/user/register", body, registerRep)

	a.Nil(err)

	a.Equal("注册成功", registerRep.Message)

	utils.DeleteUser(&model.User{Username: "test"})
}

func TestOss(t *testing.T) {
	a := assert.New(t)

	file, err := oss.ListFile()

	a.Nil(err)

	a.NotEqual(0, len(file))
	t.Log(file)
}



