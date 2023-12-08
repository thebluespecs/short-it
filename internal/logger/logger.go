package logger

import (
	"sync"

	"go.uber.org/zap"
)

type logger struct {
	log *zap.Logger
}

var (
	instance *logger
	one      sync.Once
)

func getInstance() *logger {
	one.Do(func() {
		zapLogger, err := zap.NewProduction()
		if err != nil {
			panic("Failed to init logger: " + err.Error())
		}
		instance = &logger{log: zapLogger}
	})
	return instance
}

func Info(msg string) {
	getInstance().log.Info(msg)
}

func Error(msg string) {
	getInstance().log.Error(msg)
}

func Debug(msg string) {
	getInstance().log.Debug(msg)
}
