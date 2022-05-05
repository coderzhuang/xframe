package pkg

import (
	"github.com/sirupsen/logrus"
	"os"
	"xframe/config"
)

var myLog *logrus.Logger

func GetLog() *logrus.Logger {
	if myLog != nil {
		return myLog
	}
	logLevel := logrus.WarnLevel
	if config.Conf.Server.Debug {
		logLevel = logrus.DebugLevel
	}
	myLog = logrus.New()
	myLog.SetFormatter(&logrus.JSONFormatter{})
	myLog.SetOutput(os.Stdout)
	myLog.SetLevel(logLevel)
	return myLog
}
