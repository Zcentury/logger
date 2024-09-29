package logger

import (
	"github.com/logrusorgru/aurora/v4"
	"github.com/sirupsen/logrus"
)

var (
	labels = map[logrus.Level]string{
		logrus.PanicLevel: aurora.Bold(aurora.BgRed(" PNC ")).String(),
		logrus.FatalLevel: aurora.Bold(aurora.BgRed(" FAT ")).String(),
		logrus.ErrorLevel: aurora.Bold(aurora.BgRed(" ERR ")).String(),
		logrus.WarnLevel:  aurora.Bold(aurora.BgYellow(" WRN ")).String(),
		logrus.InfoLevel:  aurora.Bold(aurora.BgGreen(" INF ")).String(),
		logrus.DebugLevel: aurora.Bold(aurora.BgBlue(" DBG ")).String(),
		logrus.TraceLevel: aurora.Bold(aurora.BgMagenta(" TRC ")).String(),
	}
)
