package logger

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	logger *zap.Logger
	level  zapcore.Level
}

func (lg *Logger) Level() string {
	return lg.level.String()
}

func (lg *Logger) Debug(ctx context.Context, message string) {
	lg.logger.Debug(message, requestID(ctx))
}

func (lg *Logger) Info(ctx context.Context, message string) {
	lg.logger.Info(message, requestID(ctx))
}

func (lg *Logger) Warn(ctx context.Context, message string) {
	lg.logger.Warn(message, requestID(ctx))
}

func (lg *Logger) Error(ctx context.Context, err error) {
	lg.logger.With(requestID(ctx)).Error(err.Error(), zap.Error(err))
}

func (lg *Logger) Panic(ctx context.Context, err error) {
	lg.logger.Panic(err.Error(), requestID(ctx))
}

func (lg *Logger) GetLogger() *zap.Logger {
	return lg.logger
}

func requestID(ctx context.Context) zap.Field {
	requestID := ""
	if s := ctx.Value("requestID"); s != nil {
		requestID, _ = s.(string)
	}

	return zap.String("requestID", requestID)
}

func GetZapLevel(l string) zapcore.Level {
	level, err := zapcore.ParseLevel(l)
	if err != nil { // default
		return zapcore.DebugLevel
	}
	return level
}
