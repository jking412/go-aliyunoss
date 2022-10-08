package main

import (
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestMain(m *testing.M) {
	go main()
	m.Run()
}

func TestWeb(t *testing.T) {
	a := assert.New(t)

	resp, _ := http.Get("http://localhost:" + viper.GetString("web.port") + "/ping")

	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)

	a.Equal("{\"message\":\"pong\"}", string(data))
}

func TestRegister(t *testing.T){
	
}