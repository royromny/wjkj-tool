package tool

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// logPath 日志文件路径
// logLevel 日志级别
func InitLogger(logPath string, logLevel string) *zap.SugaredLogger {
	hook := lumberjack.Logger{
		Filename:   logPath, // 日志文件路径
		MaxSize:    128,     // megabytes
		MaxBackups: 30,      // 最多保留30个备份
		MaxAge:     7,       // days
		Compress:   true,    // 是否压缩 disabled by default
	}

	w := zapcore.AddSync(&hook)

	// 设置日志级别,debug可以打印出info,debug,warn；info级别可以打印warn，info；warn只能打印warn
	// debug->info->warn->error
	var level zapcore.Level
	switch logLevel {
	case "debug":
		level = zap.DebugLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}
	encoderConfig := zap.NewProductionEncoderConfig()
	// 时间格式
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		w,
		level,
	)

	logger := zap.New(core)
	logger.Sugar().Infof("Wjkj Logger init success, in %s", logPath)

	return logger.Sugar()
}
