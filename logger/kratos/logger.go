package kratos

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-volo/logger"
)

const defaultCallerSkipCount = 1

// defaultLogger is an implementation of log.Logger interface.
type defaultLogger struct {
	base logger.Logger
}

// New returns a new defaultLogger instance.
func New(base logger.Logger) log.Logger {
	l := &defaultLogger{
		base: base,
	}

	return l
}

func (l *defaultLogger) Log(level log.Level, v ...interface{}) error {
	base := l.base.WithCallDepth(defaultCallerSkipCount)
	switch level {
	case log.LevelDebug:
		base.Debug(v...)
	case log.LevelInfo:
		base.Info(v...)
	case log.LevelWarn:
		base.Warn(v...)
	case log.LevelError:
		base.Error(v...)
	case log.LevelFatal:
		base.Fatal(v...)
	default:
		base.Info(v...)
	}

	return nil
}
