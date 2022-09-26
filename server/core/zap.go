package core

import (
	"GoLog/commen"
	"GoLog/utils"
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// 自定义Logger,需要 New(core zapcore.Core, options ...Options)
// zapcore.Core配置： Encoder, WriteSyncer, LogLevel

// 配置zapcore.Core 的Encoder
func getEncoder() zapcore.Encoder {
	if commen.GVA_CONFIG.Zap.Format == "json" {
		return zapcore.NewJSONEncoder(getEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(getEncoderConfig())
}

// 对Encoder中的config进一步细化
func getEncoderConfig() zapcore.EncoderConfig {
	config := zapcore.EncoderConfig{
		MessageKey:   "msg",
		LevelKey:     "level",
		TimeKey:      "time",
		NameKey:      "logger",
		LineEnding:   zapcore.DefaultLineEnding,
		EncodeLevel:  zapcore.LowercaseLevelEncoder,
		EncodeTime:   zapcore.ISO8601TimeEncoder,
		EncodeCaller: zapcore.FullCallerEncoder,
	}
	// 配置encoderLevel的形式
	switch {
	case commen.GVA_CONFIG.Zap.EncodeLevel == "LowercaseLevelEncoder": // 默认小写编码器
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	case commen.GVA_CONFIG.Zap.EncodeLevel == "LowercaseColorLevelEncoder":
		config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case commen.GVA_CONFIG.Zap.EncodeLevel == "CapitalLevelEncoder":
		config.EncodeLevel = zapcore.CapitalLevelEncoder
	case commen.GVA_CONFIG.Zap.EncodeLevel == "CapitalColorLevelEncoder":
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
	return config
}

// 配置zapcore.core 包括Encoder, WriteSyncer, LogLevel
func getEncoderCore(filename string, loglevel zapcore.LevelEnabler) zapcore.Core {
	writer := getWriterSyncer(filename)
	return zapcore.NewCore(getEncoder(), writer, loglevel)
}

// 使用lumberjack对日志进行切分，zap + lumberjack配合 可以参考简书文章 https://www.jianshu.com/p/98d911b29a38
// gopkg.in/natefinch/lumberjack.v2
func getWriterSyncer(file string) zapcore.WriteSyncer {
	hook := &lumberjack.Logger{
		Filename:   file, // 日志文件的位置
		MaxSize:    10,   // 在进行切割前日志的最大大小，（MB为单位）
		MaxBackups: 200,  // 保留旧文件的最大个数
		MaxAge:     30,   // 保留旧文件的最大天数
		Compress:   true, // 是否压缩/归档旧文件
	}
	if commen.GVA_CONFIG.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(hook))
	}
	return zapcore.AddSync(hook)
}

func Zap() (logger *zap.Logger) {
	// 首先，确认存在日志文件夹，不存在则创建
	if ok, _ := utils.Exists(commen.GVA_CONFIG.Zap.Directory); !ok {
		os.Mkdir(commen.GVA_CONFIG.Zap.Directory, os.ModePerm) // 第二个参数代表日志文件的权限
	}
	// 设置日志级别，根据日志级别输出多个级别的日志文件
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zap.InfoLevel
	})

	debugLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zap.DebugLevel
	})

	warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zap.WarnLevel
	})

	errorLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zap.ErrorLevel
	})
	cores := []zapcore.Core{
		getEncoderCore(fmt.Sprintf("./%s/info_log.log", commen.GVA_CONFIG.Zap.Directory), infoLevel),
		getEncoderCore(fmt.Sprintf("./%s/debug_log.log", commen.GVA_CONFIG.Zap.Directory), debugLevel),
		getEncoderCore(fmt.Sprintf("./%s/warn_log.log", commen.GVA_CONFIG.Zap.Directory), warnLevel),
		getEncoderCore(fmt.Sprintf("./%s/error_log.log", commen.GVA_CONFIG.Zap.Directory), errorLevel),
	}
	// cores, options，添加调用函数
	logger = zap.New(zapcore.NewTee(cores...), zap.AddCaller())
	return logger
}
