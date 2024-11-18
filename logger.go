package logger

import (
	"github.com/Zcentury/logger/formatter"
	"github.com/Zcentury/logger/levels"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

var logger = logrus.New()

func init() {
	// 设置日志级别为Info
	SetLevel(levels.InfoLevel)

	// 设置日志格式为JSON格式
	SetFormatter(&formatter.CLI{
		UseColor: true,
	})

	// 设置输出到标准错误输出
	SetOutput(os.Stdout)
}

func SetOutput(outputs ...io.Writer) {
	if len(outputs) == 0 {
		return
	}
	if len(outputs) > 1 {
		logger.SetOutput(io.MultiWriter(outputs...))
	} else {
		logger.SetOutput(outputs[0])
	}
}

// SetReportCaller 是否输出错误位置
func SetReportCaller(reportCaller bool) {
	logger.SetReportCaller(reportCaller)
}

func SetFormatter(formatter logrus.Formatter) {
	logger.SetFormatter(formatter)
}

func SetLevel(level levels.Level) {
	logger.SetLevel(logrus.Level(level))
}

// Panic 程序会 panic
func Panic(args ...interface{}) {
	logger.Panic(args...)
}

// Panicf 程序会 panic
func Panicf(format string, args ...interface{}) {
	logger.Panicf(format, args...)
}

// Fatal 程序会退出
func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

// Fatalf 程序会退出
func Fatalf(format string, args ...interface{}) {
	logger.Fatalf(format, args...)
}

func Error(args ...interface{}) {
	logger.Error(args...)
}
func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

func Warn(args ...interface{}) {
	logger.Warn(args...)
}
func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

func Info(args ...interface{}) {
	logger.Info(args...)
}
func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

func Debug(args ...interface{}) {
	logger.Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

func Trace(args ...interface{}) {
	logger.Trace(args...)
}
func Tracef(format string, args ...interface{}) {
	logger.Tracef(format, args...)
}
