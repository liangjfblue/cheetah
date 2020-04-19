/*
@Time : 2020/4/19 14:28
@Author : liangjiefan
*/
package logger

import (
	"github.com/liangjfblue/gglog"
)

type Logger struct {
}

var (
	_gglog gglog.GGLog
)

func Init(option ...gglog.Option) {
	_gglog = gglog.NewGGLog(option...)
	_gglog.Init()
}

func Debug(format string, args ...interface{}) {
	_gglog.Debug(format, args...)
}
func Info(format string, args ...interface{}) {
	_gglog.Info(format, args...)
}
func Warn(format string, args ...interface{}) {
	_gglog.Warn(format, args...)
}
func Error(format string, args ...interface{}) {
	_gglog.Error(format, args...)
}
func Access(format string, args ...interface{}) {
	_gglog.Access(format, args...)
}
func InterfaceAvgDuration(format string, args ...interface{}) {
	_gglog.InterfaceAvgDuration(format, args...)
}
