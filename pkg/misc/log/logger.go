package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

// TODO 暂时直接使用原始日志

// Logger 日志
var Logger *logrus.Logger

func init() {
	Logger = logrus.New()
	Logger.SetFormatter(&logrus.JSONFormatter{})
	Logger.SetLevel(logrus.InfoLevel)
	Logger.SetOutput(os.Stderr)
	Logger.SetReportCaller(true)

}
