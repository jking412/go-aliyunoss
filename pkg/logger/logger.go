package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

func HandleError(arg ...interface{}) {
	logrus.Error(arg)
	os.Exit(1)
}
