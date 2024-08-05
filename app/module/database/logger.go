package database

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"gofiber/app/config"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"time"
)

type GormZapLogger struct {
	logger.Config
}

func (_GormZapLogger *GormZapLogger) LogMode(LogLevel logger.LogLevel) logger.Interface {
	_GormZapLogger.LogLevel = LogLevel
	if !config.Conf.Debug {
		_GormZapLogger.LogLevel = logger.Error
	}
	return _GormZapLogger
}

func (_GormZapLogger *GormZapLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	zap.L().Info("gorm", zap.String("msg", msg), zap.Any("data", data))
}

func (_GormZapLogger *GormZapLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	zap.L().Warn("gorm", zap.String("msg", msg), zap.Any("data", data))
}

func (_GormZapLogger *GormZapLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	zap.L().Error("gorm", zap.String("msg", msg), zap.Any("data", data))
}

func (_GormZapLogger *GormZapLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	msg := utils.FileWithLineNum() + "\t" + "Gorm"
	sql, rows := fc()

	slowLog := fmt.Sprintf("SLOW SQL >= %v", _GormZapLogger.SlowThreshold)
	if err != nil {
		zap.L().WithOptions(zap.WithCaller(false)).Error(msg, zap.Int64("rows", rows), zap.String("slow", slowLog), zap.Duration("time", time.Since(begin)), zap.String("sql", sql), zap.Error(err))
	} else {
		zap.L().WithOptions(zap.WithCaller(false)).Info(msg, zap.Int64("rows", rows), zap.String("slow", slowLog), zap.Duration("time", time.Since(begin)), zap.String("sql", sql))
	}
}
