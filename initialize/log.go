package initialize

import (
	"chat/global"
	"os"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	logTmFmt = "2006-01-02 15:04:05"
)

func InitLog() {
	encoderConf := EncoderConfig()
	WriteSyncer := GetWriteSyncer()
	LevelEnabler := GetLevelEnabler()
	newCore := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConf), WriteSyncer, LevelEnabler),                   // 写入文件
		zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConf), zapcore.Lock(os.Stdout), zapcore.DebugLevel), // 写入控制台
	)
	log := zap.New(newCore, zap.AddCaller())

	global.Log = log
	zap.ReplaceGlobals(log)
}

// GetEncoder 自定义的Encoder
func EncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller_line",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		EncodeLevel:    cEncodeLevel,
		EncodeTime:     cEncodeTime,
		EncodeDuration: zapcore.SecondsDurationEncoder,
	}
}

// GetWriteSyncer 自定义的WriteSyncer
func GetWriteSyncer() zapcore.WriteSyncer {

	lumberJackLogger := &lumberjack.Logger{
		Filename:   global.Settings.Log,
		MaxSize:    200,
		MaxBackups: 10,
		MaxAge:     30,
	}
	return zapcore.AddSync(lumberJackLogger)
}

// GetLevelEnabler 自定义的LevelEnabler
func GetLevelEnabler() zapcore.Level {
	return zapcore.InfoLevel
}

// cEncodeLevel 自定义日志级别显示
func cEncodeLevel(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + level.CapitalString() + "]")
}

// cEncodeTime 自定义时间格式显示
func cEncodeTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + t.Format(logTmFmt) + "]")
}

