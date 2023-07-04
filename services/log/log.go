package log

import (
	"github.com/coderzhuang/core"
	"github.com/sirupsen/logrus"
	"os"
)

var myLog *logrus.Logger

func GetLog() *logrus.Logger {
	if myLog != nil {
		return myLog
	}
	logLevel := logrus.ErrorLevel
	if core.Conf.Common.Debug {
		logLevel = logrus.DebugLevel
	}
	myLog = logrus.New()
	myLog.SetFormatter(&logrus.JSONFormatter{})
	myLog.SetOutput(os.Stdout)
	myLog.SetLevel(logLevel)
	return myLog
}
