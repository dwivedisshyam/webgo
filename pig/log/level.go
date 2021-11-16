package log

import (
	"strings"

	"github.com/sirupsen/logrus"
)

func getLevel(level string) logrus.Level {
	switch strings.ToUpper(level) {
	case "INFO":
		return logrus.InfoLevel
	case "WARN":
		return logrus.WarnLevel
	case "FATAL":
		return logrus.FatalLevel
	case "DEBUG":
		return logrus.DebugLevel
	case "ERROR":
		return logrus.ErrorLevel
	default:
		return logrus.InfoLevel
	}
}
