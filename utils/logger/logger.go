package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func init() {
	log.Out = os.Stdout
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
}

// Info logs a message at level Info on the standard logger.
func Info(format string, args ...interface{}) {
	log.Infof(format, args...)
}

// Error logs a message at level Error on the standard logger.
func Error(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

// Warning logs a message at level Warning on the standard logger.
func Warning(format string, args ...interface{}) {
	log.Warnf(format, args...)
}

// Panic logs a message at level Panic on the standard logger.
func Panic(format string, args ...interface{}) {
	log.Panicf(format, args...)
}

// Trace logs a message at level Trace on the standard logger.
func Trace(format string, args ...interface{}) {
	log.Warnf(format, args...)
}
