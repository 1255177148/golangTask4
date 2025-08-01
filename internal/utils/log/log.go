// 封装为全局包函数

package log

import (
	"go.uber.org/zap"
)

var Logger *zap.Logger
var Sugar *zap.SugaredLogger

func Init(l *zap.Logger) {
	Logger = l
	Sugar = l.Sugar()
}

func Debug(msg string, fields ...zap.Field) { Logger.Debug(msg, fields...) }
func Info(msg string, fields ...zap.Field)  { Logger.Info(msg, fields...) }
func Warn(msg string, fields ...zap.Field)  { Logger.Warn(msg, fields...) }
func Error(msg string, fields ...zap.Field) { Logger.Error(msg, fields...) }
func Fatal(msg string, fields ...zap.Field) {
	Logger.Fatal(msg, fields...) // Fatal 会调用 os.Exit(1)
}
func InfoF(template string, args ...interface{}) {
	Sugar.Infof(template, args...)
}

func WarnF(template string, args ...interface{}) {
	Sugar.Warnf(template, args...)
}

func ErrorF(template string, args ...interface{}) {
	Sugar.Errorf(template, args...)
}

func DebugF(template string, args ...interface{}) {
	Sugar.Debugf(template, args...)
}
