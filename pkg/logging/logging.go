package logging

import (
	log "github.com/sirupsen/logrus"
	"os"
)

var logger = log.Logger{}

func GetLogger() *log.Logger {
	return &logger
}

func init() {
	logger.SetOutput(os.Stdout)
	logger.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
		ForceColors:     true,
	})
	logLevel, err := log.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		logLevel = log.InfoLevel
	}

	logger.SetLevel(logLevel)
}
