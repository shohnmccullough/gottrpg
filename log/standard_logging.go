package log

import (
	"os"

	"github.com/sirupsen/logrus"
)

type Log struct {
	Id      LogLevel
	Message string
}

type LogLevel int

const (
	Info    = 1
	Warning = 2
	Error   = 3
)

var (
	invalidArgumentMessage = Log{Id: Error, Message: "Invalid arg %s"}
)

type StdLgr interface{}

type StandardLogger struct {
	*logrus.Logger
}

/*
Build a new logging instance
  - if environment variable "LOG_OUT" is set, attempts to write STDOUT to file
*/
func NewLogger() *StandardLogger {
	baseLogger := logrus.New()
	stdLogger := &StandardLogger{baseLogger}
	stdLogger.SetFormatter(&logrus.JSONFormatter{})

	if val, flag := os.LookupEnv("LOGGING_PATH"); flag {
		f, err := os.OpenFile(val, os.O_CREATE|os.O_RDWR, 0777)
		if err != nil {
			stdLogger.Errorf("failed to open file %s - logs will sent to STDOUT", val)
			return stdLogger
		}
		stdLogger.SetOutput(f)
	}

	return stdLogger
}

/*
Generic logging function
  - Log.id indicates the log level
  - Log.message indicates the pattern for the log message
  - args are any sequence of values required to fulfill the directives in Log.message
*/
func (l *StandardLogger) Log(log Log, args ...interface{}) {
	switch log.Id {
	case 1:
		l.Infof(log.Message, args...)
	case 2:
		l.Warnf(log.Message, args...)
	case 3:
		l.Errorf(log.Message, args...)
	}
}

func (l *StandardLogger) InvalidArgument(argumentName string) {
	l.Errorf(invalidArgumentMessage.Message, argumentName)
}

func (l *StandardLogger) ValidArgument(argumentName string) {
	l.Infof(argumentName)
}
