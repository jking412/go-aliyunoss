package main

import (
	"aliyunoss/app/model"
	"aliyunoss/app/model/utils"
	"aliyunoss/pkg/logger"
	"aliyunoss/pkg/request"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
)

type BaseRep struct {
	Message string `json:"message"`
}

type PingRep struct {
	BaseRep
}

type RegisterRep struct {
	BaseRep
	Token string `json:"token"`
}

type LoginRep struct {
	BaseRep
	Token string `json:"token"`
}

type ossRep struct {
	BaseRep
	Files []string `json:"files"`
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

func TestJWT(t *testing.T) {
	a := assert.New(t)

	registerRep, err := register()

	a.Nil(err)

	a.Equal("注册成功", registerRep.Message)

	a.NotEmpty(registerRep.Token)

	t.Log(registerRep.Token)

	loginRep, err := login()

	a.NotEmpty(loginRep.Token)

	t.Log(loginRep.Token)

	ossRep, err := showBucketList(loginRep.Token)

	t.Log(ossRep.Files)

	utils.DeleteUser(&model.User{Username: "test"})
}

func TestLogger(t *testing.T) {
	logger.Error("test", zap.String("test", "test"))
	logger.Info("test", zap.String("test", "test"))
}

func register() (*RegisterRep, error) {
	registerRep := &RegisterRep{}

	body := `{"username":"test","password":"123456"}`

	err := request.PostRequest("http://localhost:"+viper.GetString("web.port")+"/user/register", body, registerRep)

	return registerRep, err
}

func login() (*LoginRep, error) {
	loginRep := &LoginRep{}

	body := `{"username":"test","password":"123456"}`

	err := request.PostRequest("http://localhost:"+viper.GetString("web.port")+"/user/login", body, loginRep)

	return loginRep, err
}

func showBucketList(token string) (*ossRep, error) {
	ossRep := &ossRep{}

	err := request.GetRequestWithHeader("http://localhost:"+viper.GetString("web.port")+"/oss/show", ossRep, map[string]string{
		"Authorization": "Bearer " + token,
	})

	return ossRep, err
}
