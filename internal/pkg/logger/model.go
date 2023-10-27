package logger

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"im/internal/pkg/consts"
)

type Logger interface {
	Level() string
	Debug(ctx context.Context, message string)
	Info(ctx context.Context, message string)
	Warn(ctx context.Context, message string)
	Error(ctx context.Context, err error)
	Panic(ctx context.Context, err error)
	GetLogger() any
}

type ZapLogger struct {
	logger *zap.Logger
	level  zapcore.Level
}

func (lg *ZapLogger) Level() string {
	return lg.level.String()
}

func (lg *ZapLogger) Debug(ctx context.Context, message string) {
	lg.logger.Debug(message, requestID(ctx))
}

func (lg *ZapLogger) Info(ctx context.Context, message string) {
	lg.logger.Info(message, requestID(ctx))
}

func (lg *ZapLogger) Warn(ctx context.Context, message string) {
	lg.logger.Warn(message, requestID(ctx))
}

func (lg *ZapLogger) Error(ctx context.Context, err error) {
	lg.logger.Error(err.Error(), requestID(ctx))
}

func (lg *ZapLogger) Panic(ctx context.Context, err error) {
	lg.logger.Panic(err.Error(), requestID(ctx))
}

func (lg *ZapLogger) GetLogger() any {
	return lg.logger
}

func requestID(ctx context.Context) zap.Field {
	requestID := ""
	if s := ctx.Value(consts.REQUEST_ID); s != nil {
		requestID, _ = s.(string)
	}

	return zap.String(consts.REQUEST_ID, requestID)
}

func GetZapLevel(l string) zapcore.Level {
	level, err := zapcore.ParseLevel(l)
	if err != nil { // default
		return zapcore.DebugLevel
	}
	return level
}
