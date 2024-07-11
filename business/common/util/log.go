package util

import (
	"dubbo.apache.org/dubbo-go/v3/logger/zap"
	"github.com/dubbogo/gost/log/logger"
)

type ZapLogger struct {
	logger *zap.Logger
}

func (l *ZapLogger) Debug(msg string, fields map[string]interface{}) {
	logger.Debug(msg, fields)
}

func (l *ZapLogger) Info(msg string, fields map[string]interface{}) {
	logger.Info(msg, fields)
}

func (l *ZapLogger) Warning(msg string, fields map[string]interface{}) {
	logger.Warn(msg, fields)
}

func (l *ZapLogger) Error(msg string, fields map[string]interface{}) {
	logger.Error(msg, fields)
}

func (l *ZapLogger) Fatal(msg string, fields map[string]interface{}) {
	logger.Fatal(msg, fields)
}

func (l *ZapLogger) Level(level string) {
	logger.SetLoggerLevel(level)
}

func (l *ZapLogger) OutputPath(path string) (err error) {
	logger.Info("output path is not supported")
	return nil
}
