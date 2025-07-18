package log

import "github.com/sirupsen/logrus"

type StandardLoggerEntry struct {
	*logrus.Entry
}

/*
Optional method to add default attributes with value(s) to the StandardLogger
*/
func (l *StandardLogger) NewLoggerEntry(fields logrus.Fields) *StandardLoggerEntry {
	return &StandardLoggerEntry{l.WithFields(fields)}
}

func (l *StandardLoggerEntry) Log(log Log, args ...interface{}) {
	switch log.Id {
	case 1:
		l.Infof(log.Message, args...)
	case 2:
		l.Warnf(log.Message, args...)
	case 3:
		l.Errorf(log.Message, args...)
	}
}

func (l *StandardLoggerEntry) InvalidArgument(argumentName string) {
	l.Errorf(invalidArgumentMessage.Message, argumentName)
}
