package log

import (
	"github.com/charmbracelet/log"
	"os"
	"time"
)

var logger *log.Logger

func InitLogger() {
	logger = log.NewWithOptions(os.Stderr, log.Options{
		ReportCaller:    false,
		ReportTimestamp: true,
		TimeFormat:      time.Kitchen,
		Prefix:          "Capybara 🍪 ",
	})
}

func SetLoggerDebugLevel() {
	logger.SetLevel(log.DebugLevel)
}

func GetDefaultLogger() *log.Logger {
	return logger
}
