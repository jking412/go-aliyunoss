package viper

import (
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestViper(t *testing.T) {
	a := assert.New(t)

	InitViper("../../")
	port := viper.GetString("web.port")

	t.Log(port)
	a.NotEmpty(port)
}
