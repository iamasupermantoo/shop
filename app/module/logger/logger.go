package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gofiber/app/config"
	"os"
)

// init 初始化日志
func init() {
	var logger *zap.Logger

	if config.Conf.Debug {
		// 开发环境配置
		cmf := zap.NewDevelopmentEncoderConfig()

		// 输出控制台添加颜色
		cmf.EncodeLevel = zapcore.CapitalColorLevelEncoder

		// 时间格式化
		cmf.EncodeTime = zapcore.ISO8601TimeEncoder

		// 显示全路径
		cmf.EncodeCaller = zapcore.FullCallerEncoder

		// 输出控制台格式配置
		enc := zapcore.NewConsoleEncoder(cmf)

		// 创建日志核心
		logger = zap.New(
			zapcore.NewCore(enc, zapcore.AddSync(os.Stdout), zapcore.InfoLevel),
			zap.AddCaller(),
		)
	} else {
		// 生产环境配置
		cmf := zap.NewProductionEncoderConfig()

		// 时间格式化
		cmf.EncodeTime = zapcore.ISO8601TimeEncoder

		// 显示全路径
		cmf.EncodeCaller = zapcore.FullCallerEncoder

		// 输出JSON格式配置
		enc := zapcore.NewJSONEncoder(cmf)

		// 写入文件
		appFile, _ := os.OpenFile("./logs/app.json", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 06666)

		// 创建日志核心
		logger = zap.New(
			zapcore.NewCore(enc, zapcore.AddSync(appFile), zapcore.WarnLevel),
			zap.AddCaller(),
			zap.AddStacktrace(zapcore.ErrorLevel),
		)
	}

	// 创建 Logger 对象
	defer logger.Sync()

	// 设置全局可用
	zap.ReplaceGlobals(logger)
}
