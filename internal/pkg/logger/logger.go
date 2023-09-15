package logger

import (
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	levelInfo    = "info"
	levelWarning = "warning"
	levelError   = "error"
)

func NewLogger(level string) (l *zap.SugaredLogger, err error) {
	cfg := zap.NewProductionConfig()
	cfg.DisableCaller = true

	switch strings.ToLower(level) {
	case levelInfo:
		cfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	case levelWarning:
		cfg.Level = zap.NewAtomicLevelAt(zap.WarnLevel)
	case levelError:
		cfg.Level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	default:
		cfg = zap.NewDevelopmentConfig()
		cfg.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	}

	basicLogger, err := cfg.Build()
	l = basicLogger.Sugar()

	return
}
