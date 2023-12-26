package queue

import (
	"github.com/ThreeDotsLabs/watermill"
	"go.uber.org/zap"
)

func newWatermillZap(in digIn) *watermillZap {
	if in.Logger != nil && in.Logger.GetLogger() != nil {
		if z, ok := in.Logger.GetLogger().(*zap.Logger); ok {
			return &watermillZap{
				logger: z,
			}
		}
	}
	zapLog, _ := zap.NewProduction()
	return &watermillZap{
		logger: zapLog,
	}
}

type watermillZap struct {
	logger *zap.Logger
}

func (w watermillZap) Error(msg string, err error, fields watermill.LogFields) {
	w.logger.Error(msg, zap.Error(err), zap.Any("fields", fields))
}

func (w watermillZap) Info(msg string, fields watermill.LogFields) {
	w.logger.Info(msg, zap.Any("fields", fields))
}

func (w watermillZap) Debug(msg string, fields watermill.LogFields) {
	w.logger.Debug(msg, zap.Any("fields", fields))
}

func (w watermillZap) Trace(msg string, fields watermill.LogFields) {
	w.logger.Debug(msg, zap.Any("fields", fields)) // zap doesn't have a Trace level, so we use Debug instead
}

func (w watermillZap) With(fields watermill.LogFields) watermill.LoggerAdapter {
	newLogger := w.logger.With(zap.Any("fields", fields))
	return watermillZap{logger: newLogger}
}
