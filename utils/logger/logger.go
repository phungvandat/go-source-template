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
