package utils

import (
	"chat/constant"
	"chat/global"
	"fmt"
	"go.uber.org/zap/zapcore"
	"runtime"
	"sort"
)

func getCaller(level int) string {
	pc, _, line, _ := runtime.Caller(level)
	name := runtime.FuncForPC(pc).Name()
	return fmt.Sprintf("%v:%v", name, line)
}

func Logger(level constant.LogLevel, maps map[string]string) {
	length := len(maps)
	fields := make([]zapcore.Field, length)
	// TODO maybe可以优化，
	keys := make([]string, length)
	step := 0
	for key := range maps {
		keys[step] = key
		step++
	}
	// 为了保证日志的一致，所以将map排了序，也可以使用有序字典 待测试
	sort.Strings(keys)

	for index, value := range keys {
		fields[index] = zapcore.Field{Key: value, Type: zapcore.StringType, String: maps[value]}
	}
	caller := getCaller(2)
	global.Log.Log(zapcore.Level(level), fmt.Sprintf("%-30v", caller), fields...)
}
