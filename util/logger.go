package util

import (
	"sync"

	"go.uber.org/zap"
)

var loggerOnce sync.Once
var logger *zap.Logger

//InitLogger initializes the logger
func InitLogger() error {
	var err error
	logger, err = zap.NewProduction()
	return err
}

//GetLogger returns global logger
func Logger() *zap.Logger {
	return logger
}
