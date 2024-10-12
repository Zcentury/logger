package logger

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

var logger = logrus.New()

func init() {
	// 设置日志级别为Info
	SetLevel(InfoLevel)

	// 设置日志格式为JSON格式
	SetFormatter(&DefaultFormatter{})

	// 设置输出到标准错误输出
	SetOutput(os.Stdout)
}

func SetOutput(output io.Writer) {
	logger.SetOutput(output)
}

func SetFormatter(formatter logrus.Formatter) {
	logger.SetFormatter(formatter)
}

func SetLevel(level int) {
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
