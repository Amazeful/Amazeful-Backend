package util

import "go.uber.org/zap"

var logger *zap.Logger

//InitLogger initializes the logger
func InitLogger() error {
	lg, err := zap.NewProduction()
	if err != nil {
		return err
	}

	logger = lg
	return nil
}

//GetLogger returns global logger
func GetLogger() *zap.Logger {
	return logger
}
