package bootstrap

import (
	"github.com/1255177148/golangTask4/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

var Logger *zap.Logger

func InitLogger() {
	level := zapcore.InfoLevel
	switch config.Cfg.Log.Level {
	case "debug":
		level = zapcore.DebugLevel
	case "error":
		level = zapcore.ErrorLevel
	}
	env := os.Getenv("APP_ENV") // 从环境变量读取运行环境
	if env == "" {
		env = "dev" // 默认开发环境
	}

	writer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   config.Cfg.Log.Filename,
		MaxSize:    config.Cfg.Log.MaxSize,
		MaxBackups: config.Cfg.Log.MaxBackups,
		MaxAge:     config.Cfg.Log.MaxAge,
		Compress:   config.Cfg.Log.Compress,
	})
	var core zapcore.Core
	encoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	if env == "dev" {
		// 开发环境：控制台和文件同时输出
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, writer, zapcore.DebugLevel),
			zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel),
		)
	} else {
		core = zapcore.NewCore(encoder, writer, level)
	}
	Logger = zap.New(core, zap.AddCaller())
}
