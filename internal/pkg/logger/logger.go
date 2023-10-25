package logger

import (
	"time"

	"im/internal/pkg/config"
	"im/internal/pkg/consts"

	"go.uber.org/dig"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type digIn struct {
	dig.In

	Config *config.Config
}

func NewLogger(in digIn) *Logger {
	level := GetZapLevel(in.Config.LogConfig.Level)
	serverName := in.Config.LogConfig.Name
	env := in.Config.LogConfig.Env
	return newLogger(level, serverName, env)
}

// TODO
func newLogger(level zapcore.Level, serviceName string, env string) *Logger {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     timeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	encoding := "json"
	if env == "local" {
		encoding = "console"
	}
	outputPaths := []string{"stdout"}
	errorOutputPaths := []string{"stderr"}
	if env != "local" {
		outputPaths = append(outputPaths, "./im_info.log")
		errorOutputPaths = append(errorOutputPaths, "./im_error.log")
	}

	config := zap.Config{
		Level:             zap.NewAtomicLevelAt(level), // 日志级别
		DisableStacktrace: true,
		Development:       false,         // 开发模式，堆栈跟踪
		Encoding:          encoding,      // 输出格式 console 或 json
		EncoderConfig:     encoderConfig, // 编码器配置
		InitialFields: map[string]interface{}{
			"service": serviceName,
		}, // 初始化字段，如：添加一个服务器名称
		OutputPaths:      outputPaths, // 输出到指定文件 stdout（标准输出，正常颜色） stderr（错误输出，红色）
		ErrorOutputPaths: errorOutputPaths,
	}
	options := make([]zap.Option, 2)
	options[0] = zap.AddCallerSkip(1)
	options[1] = zap.AddStacktrace(zapcore.ErrorLevel)

	logger, err := config.Build(options...)
	if err != nil {
		panic(err)
	}

	return &Logger{
		logger: logger,
		level:  level,
	}
}

// 自定义时间编码器
func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(consts.TIME_FORMAT))
}
