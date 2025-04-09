package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.SugaredLogger

func InitLogger(isProduction bool) {
	var cfg zap.Config

	if isProduction {
		cfg = zap.NewProductionConfig()
	} else {
		cfg = zap.NewDevelopmentConfig()
		cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	logger, err := cfg.Build()
	if err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}

	log = logger.Sugar()
}

func GetLogger() *zap.SugaredLogger {
	if log == nil {
		InitLogger(false) // Default to development mode
	}
	return log
}

// Helper functions
func Info(args ...interface{}) {
	GetLogger().Info(args...)
}

func Error(args ...interface{}) {
	GetLogger().Error(args...)
}

func Debug(args ...interface{}) {
	GetLogger().Debug(args...)
}

func Warn(args ...interface{}) {
	GetLogger().Warn(args...)
}
