package log

import (
	"os"

	"github.com/sirupsen/logrus"
)

type Logger interface {
	Debug(args ...interface{})
	Debugf(format string, a ...interface{})
	Info(args ...interface{})
	Infof(format string, a ...interface{})
	Warn(args ...interface{})
	Warnf(format string, a ...interface{})
	Error(args ...interface{})
	Errorf(format string, a ...interface{})
	WithFields(fields logrus.Fields) *logrus.Entry
}

func NewLogger() Logger {
	return newLogger()
}

func newLogger() *logrus.Logger {
	l := logrus.New()
	l.SetLevel(getLevel(os.Getenv("LOG_LEVEL")))

	return l
}
