package formatter

import (
	"fmt"
	"github.com/Zcentury/logger/levels"
	"github.com/logrusorgru/aurora/v4"
	"github.com/sirupsen/logrus"
)

// CLIAndFileFormatter 在控制台中输出好看的日志，在文件中输出没有颜色字符的日志
// 暂时不支持多输出自动判断，这也是下一步要解决的问题
type CLIAndFileFormatter struct {
	TimeFormat string // 自定义日期格式化
	UseColor   bool   // 是否使用颜色
}

func (f *CLIAndFileFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	isTerminal := checkTerminal(entry.Logger.Out)
	fmt.Println(isTerminal)

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

	logMsg := fmt.Sprintf("|%s| %s | %s\n", levelData, tm, entry.Message)
	return []byte(logMsg), nil
}
