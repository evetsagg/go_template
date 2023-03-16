package logger

import (
	"log"

	"go.uber.org/zap"
)

type Logger struct {
	log *zap.Logger
}

func NewLogger() *Logger {
	cfg := zap.NewDevelopmentConfig()
	cfg.OutputPaths = []string{
		"logs/system.log",
	}
	logger, err := cfg.Build((zap.AddCallerSkip(1)))

	if err != nil {
		log.Fatal(err)
	}
	defer logger.Sync()

	return &Logger{log: logger}
}

func (l *Logger) Error(e error) {
	l.log.Error(e.Error())
}
func (l *Logger) Info(msg string) {
	l.log.Info(msg)
}
func (l *Logger) Debug(msg string) {
	l.log.Debug(msg)
}
func (l *Logger) Fatal(e error) {
	l.log.Fatal(e.Error())
}
