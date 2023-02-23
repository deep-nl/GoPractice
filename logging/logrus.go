package logging

import (
	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func init() {
	Logger = logrus.New()
	Logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05Z07:00",
	})
	Logger.Info("info message")
	Logger.Warn("warn message")
	Logger.Error("error message")
}
