package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

func Error(arg ...interface{}) {
	logrus.Error(arg)
	os.Exit(1)
}
