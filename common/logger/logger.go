package logger

import (
	"go.uber.org/zap"
	"sync"
)

var once sync.Once

// Init initializes the logger once and returns it
func Init(isProduction bool) (*zap.Logger, *zap.SugaredLogger, error) {
	var l *zap.Logger
	var s *zap.SugaredLogger
	var err error

	once.Do(func() {
		if isProduction {
			l, err = zap.NewProduction()
		} else {
			l, err = zap.NewDevelopment()
		}

		if err == nil {
			// Sugared logger with caller skip for convenience
			s = l.Sugar().WithOptions(zap.AddCallerSkip(0))
			zap.ReplaceGlobals(l)
		}
	})

	return l, s, err
}

// Convenience functions using global loggers (optional)
var Logger *zap.Logger
var SugaredLogger *zap.SugaredLogger

func Info(msg string, fields ...zap.Field) {
	if Logger != nil {
		Logger.Info(msg, fields...)
	}
}

func Error(msg string, fields ...zap.Field) {
	if Logger != nil {
		Logger.Error(msg, fields...)
	}
}

func Warn(msg string, fields ...zap.Field) {
	if Logger != nil {
		Logger.Warn(msg, fields...)
	}
}

func Debug(msg string, fields ...zap.Field) {
	if Logger != nil {
		Logger.Debug(msg, fields...)
	}
}

func Fatal(msg string, fields ...zap.Field) {
	if Logger != nil {
		Logger.Fatal(msg, fields...)
	}
}

func Infof(template string, args ...interface{}) {
	if SugaredLogger != nil {
		SugaredLogger.Infof(template, args...)
	}
}

func Errorln(args ...interface{}) {
	if SugaredLogger != nil {
		SugaredLogger.Errorln(args...)
	}
}
