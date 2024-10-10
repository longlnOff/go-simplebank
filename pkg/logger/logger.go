package logger

import (
	"os"

	"github.com/longln/go-simplebank/pkg/setting"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerZap struct {
	*zap.Logger
}

func NewLogger(config setting.LogConfig) *LoggerZap {
	loglevel := config.Level
	var level zapcore.Level
	switch loglevel {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	default:
		level = zapcore.InfoLevel
	}

	encoder := getEncoderLog()
	hook := lumberjack.Logger{
		Filename:   config.FileLogName,
		MaxSize:    config.MaxSize,
		MaxBackups: config.MaxBackups,
		MaxAge:     config.MaxAge,
		Compress:   config.Compress,
	}

	core := zapcore.NewCore(encoder,
	zapcore.NewMultiWriteSyncer(zapcore.AddSync(&hook),zapcore.AddSync(os.Stdout)),
	level)

	return &LoggerZap{
		Logger: zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel)),
	}

}


func getEncoderLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encodeConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encodeConfig.EncodeCaller = zapcore.FullCallerEncoder
	encodeConfig.TimeKey = "time"
	return zapcore.NewConsoleEncoder(encodeConfig)
}