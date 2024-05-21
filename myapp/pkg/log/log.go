package log

import (
	"myapp/internal/config"

	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

type ILogger interface {
	Error(message string)
	Info(message string)
	Warn(message string)
}

type Logger struct {
	*logrus.Logger
}

// NewLogger creates a new instance of Logger.
func NewLogger(cfg *config.Config) (ILogger, error) {
	logger := logrus.New()

	// Create a new LFS hook for log rotation
	lfsHook := lfshook.NewHook(
		lfshook.PathMap{
			logrus.InfoLevel:  cfg.LogFilePath + "-info.log",
			logrus.WarnLevel:  cfg.LogFilePath + "-warn.log",
			logrus.ErrorLevel: cfg.LogFilePath + "-error.log",
		},
		&logrus.JSONFormatter{},
	)

	// Set the log level to Debug for the logger
	logger.SetLevel(logrus.DebugLevel)

	// Add the LFS hook to the logger
	logger.AddHook(lfsHook)

	return &Logger{logger}, nil
}

// Error logs a message at error level.
func (l *Logger) Error(message string) {
	l.Logger.Error(message)
}

// Warn logs a message at warning level.
func (l *Logger) Warn(message string) {
	l.Logger.Warn(message)
}

// Info logs a message at info level.
func (l *Logger) Info(message string) {
	l.Logger.Info(message)
}
