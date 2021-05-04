package logs

import (
	"os"

	log "github.com/sirupsen/logrus"
)

// Logger enforces specific log message formats
type Logger struct {
	*log.Logger
}

// Log initializes the logger
func Log() *Logger {
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Error(err)
	}

	baseLogger := log.New()
	standardLogger := &Logger{baseLogger}
	standardLogger.SetFormatter(&log.TextFormatter{})
	standardLogger.SetOutput(file)

	return standardLogger
}
