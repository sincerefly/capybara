package log

import (
	"os"
	"time"

	"github.com/charmbracelet/log"
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
