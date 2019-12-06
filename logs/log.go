package logs

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

var levelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

func init()  {
	logger := InitLog()
	zap.ReplaceGlobals(logger) // 将配置应用于全局的zap
}

func getLoggerLevel(lvl string) zapcore.Level {
	if level, ok := levelMap[lvl]; ok {
		return level
	}
	return zapcore.InfoLevel
}

func InitLog() *zap.Logger {
	consoleLevel := getLoggerLevel("debug")
	fileLevel := getLoggerLevel("error")

	filePriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= fileLevel
	})
	consolePriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= consoleLevel
	})

	var allCore []zapcore.Core

	fileWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "logs/app.log",
		MaxSize:    128, // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 3,   // 日志文件最多保存多少个备份
		LocalTime:  true,
		// MaxAge:     7, // 文件最多保存多少天
		Compress: true, // 是否压缩
	})
	fileEncoder := zap.NewProductionEncoderConfig()
	fileEncoder.EncodeTime = zapcore.ISO8601TimeEncoder

	consoleWriter := zapcore.Lock(os.Stdout)
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())

	allCore = append(allCore, zapcore.NewCore(consoleEncoder, consoleWriter, consolePriority))
	allCore = append(allCore, zapcore.NewCore(zapcore.NewJSONEncoder(fileEncoder), fileWriter, filePriority))

	core := zapcore.NewTee(allCore...)

	return zap.New(core).WithOptions(zap.AddCaller())
}

