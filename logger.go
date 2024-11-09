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

var (
	Panic  = logger.Panic  // 程序会 panic
	Panicf = logger.Panicf // 程序会 panic
	Fatal  = logger.Fatal  // 程序会退出
	Fatalf = logger.Fatalf // 程序会退出
	Error  = logger.Error
	Errorf = logger.Errorf
	Warn   = logger.Warn
	Warnf  = logger.Warnf
	Info   = logger.Info
	Infof  = logger.Infof
	Debug  = logger.Debug
	Debugf = logger.Debugf
	Trace  = logger.Trace
	Tracef = logger.Tracef
)
