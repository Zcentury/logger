package formatter

import (
	"fmt"
	"github.com/Zcentury/logger/levels"
	"github.com/logrusorgru/aurora/v4"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

// CLI 在控制台中输出日志
type CLI struct {
	TimeFormat string // 自定义日期格式化
	UseColor   bool   // 是否使用颜色
}

func (f *CLI) Format(entry *logrus.Entry) ([]byte, error) {
	level := levels.Level(entry.Level)
	var levelData string
	var layout string

	if f.UseColor {
		levelData = level.Color()
	} else {
		levelData = level.Acronym()
	}

	if f.TimeFormat != "" {
		layout = f.TimeFormat
	} else {
		layout = "2006-01-02 15:04:05"
	}

	tm := entry.Time.Format(layout)

	if f.UseColor {
		tm = aurora.Bold(aurora.Cyan(tm)).String()
	}

	var logMsg string

	if entry.HasCaller() {
		entry.Caller = getCaller()
		var ph string

		cwd, err := os.Getwd()
		if err != nil {
			ph = entry.Caller.Function
		} else {
			// 计算相对路径
			if relativePath, err := filepath.Rel(cwd, entry.Caller.File); err != nil {
				ph = entry.Caller.Function
			} else {
				ph = relativePath
			}
		}
		triggerPosition := fmt.Sprintf("%s:%d", ph, entry.Caller.Line)

		logMsg = fmt.Sprintf("|%s| %s | %s | %s\n", levelData, tm, aurora.Bold(aurora.Underline(aurora.Blue(triggerPosition))).String(), entry.Message)
	} else {
		logMsg = fmt.Sprintf("|%s| %s | %s\n", levelData, tm, entry.Message)
	}

	return []byte(logMsg), nil
}
