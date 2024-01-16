package logger

import (
	"context"
	"log/slog"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var levelMap = map[slog.Level]zapcore.Level{
	slog.LevelDebug: zap.DebugLevel,
	slog.LevelInfo:  zap.InfoLevel,
	slog.LevelWarn:  zap.WarnLevel,
	slog.LevelError: zap.ErrorLevel,
}

var levelReverseMap = map[zapcore.Level]slog.Level{
	zap.DebugLevel:  slog.LevelDebug,
	zap.InfoLevel:   slog.LevelInfo,
	zap.WarnLevel:   slog.LevelWarn,
	zap.ErrorLevel:  slog.LevelError,
	zap.DPanicLevel: slog.LevelError,
	zap.PanicLevel:  slog.LevelError,
	zap.FatalLevel:  slog.LevelError,
}

func NewZapHandler(zapLogger *zap.Logger) *ZapHandler {
	level := levelReverseMap[zapLogger.Level()]
	return &ZapHandler{
		attrs:  make([]slog.Attr, 0),
		groups: make([]string, 0),
		level:  level,
		log:    zapLogger,
	}
}

type ZapHandler struct {
	attrs  []slog.Attr
	groups []string
	level  slog.Level
	log    *zap.Logger
}

func (h *ZapHandler) Enabled(_ context.Context, level slog.Level) bool {
	return level >= h.level
}

func (h *ZapHandler) Handle(ctx context.Context, r slog.Record) error {
	var fields []zap.Field
	r.Attrs(func(attr slog.Attr) bool {
		key := attr.Key
		if key == "!BADKEY" {
			if _, ok := attr.Value.Any().(error); ok {
				key = "error"
			}
		}
		fields = append(fields, zap.Any(key, attr.Value))
		return true
	})
	fields = append(fields, requestID(ctx))

	level := levelMap[r.Level]
	h.log.Log(level, r.Message, fields...)

	return nil
}

func (h *ZapHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	newHandler := *h
	for _, attr := range attrs {
		newHandler.attrs = append(newHandler.attrs, attr)
	}
	return &newHandler
}

func (h *ZapHandler) WithGroup(name string) slog.Handler {
	newHandler := *h
	newHandler.groups = append(newHandler.groups, name)
	return &newHandler
}
