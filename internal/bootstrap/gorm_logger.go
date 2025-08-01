package bootstrap

import (
	"context"
	"errors"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm/logger"
)

type ZapGormLogger struct {
	ZapLogger     *zap.Logger
	SlowThreshold time.Duration
	LogLevel      logger.LogLevel
}

func NewZapGormLogger(zapLogger *zap.Logger) logger.Interface {
	return &ZapGormLogger{
		ZapLogger:     zapLogger,
		SlowThreshold: 200 * time.Millisecond, // 200ms 以上算慢查询，可改配置
		LogLevel:      logger.Info,
	}
}

func (l *ZapGormLogger) LogMode(level logger.LogLevel) logger.Interface {
	newLogger := *l
	newLogger.LogLevel = level
	return &newLogger
}

func (l *ZapGormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Info {
		l.ZapLogger.Sugar().Infof(msg, data...)
	}
}

func (l *ZapGormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Warn {
		l.ZapLogger.Sugar().Warnf(msg, data...)
	}
}

func (l *ZapGormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Error {
		l.ZapLogger.Sugar().Errorf(msg, data...)
	}
}

func (l *ZapGormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel <= logger.Silent {
		return
	}

	elapsed := time.Since(begin)
	sql, rows := fc()

	switch {
	case err != nil && (!errors.Is(err, logger.ErrRecordNotFound)):
		l.ZapLogger.Error("SQL 执行错误",
			zap.Error(err),
			zap.Duration("耗时", elapsed),
			zap.Int64("行数", rows),
			zap.String("SQL", sql),
		)

	case elapsed > l.SlowThreshold:
		l.ZapLogger.Warn("慢查询",
			zap.Duration("耗时", elapsed),
			zap.Int64("行数", rows),
			zap.String("SQL", sql),
		)

	default:
		l.ZapLogger.Info("SQL 执行",
			zap.Duration("耗时", elapsed),
			zap.Int64("行数", rows),
			zap.String("SQL", sql),
		)
	}
}
