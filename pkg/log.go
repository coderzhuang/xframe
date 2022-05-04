package pkg

import (
	"github.com/sirupsen/logrus"
	"os"
)

var logger *logrus.Logger

func GetLog() *logrus.Logger {
	if logger != nil {
		logger := logrus.New()
		logger.SetFormatter(&logrus.JSONFormatter{})
		logger.SetOutput(os.Stdout)
		logger.SetLevel(logrus.InfoLevel)
	}
	return logger
}
