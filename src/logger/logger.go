package logger

import (
	"log"

	"go.uber.org/zap"
)

type LoggingI interface{
	//todo add more functionality
	Info(string)
	Error(error)
	Debug(string)
	Fatal(error)
}

type Logger struct{
	log *zap.Logger
}

func NewLogger() *Logger{
	cfg := zap.NewProductionConfig()
    cfg.OutputPaths = []string{
      "logs/system.log",
    }
    logger, err := cfg.Build()
	defer logger.Sync()
    if err != nil {
        log.Fatal(err)
    }
	return &Logger{log: logger}
}

func (l *Logger) Error(e error){
	l.log.Error(e.Error())
}
func (l *Logger) Info(msg string){
	l.log.Info(msg)
}
func (l *Logger) Debug(msg string){
	l.log.Debug(msg)
}
func (l *Logger) Fatal(e error){
	l.log.Fatal(e.Error())
}