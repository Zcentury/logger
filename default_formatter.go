package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

type DefaultFormatter struct {
}

func (f *DefaultFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	tm := entry.Time.Format("2006-01-02 15:04:05")
	logMsg := fmt.Sprintf("|%s| %s | %s\n", labels[entry.Level], tm, entry.Message)
	return []byte(logMsg), nil
}
