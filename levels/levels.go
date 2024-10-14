package levels

import "github.com/logrusorgru/aurora/v4"

type Level uint32

const (
	PanicLevel Level = iota
	FatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
	TraceLevel
)

// String 全称，适用于json输出
func (l Level) String() string {
	return [...]string{"Panic", "Fatal", "Error", "Warn", "Info", "Debug", "Trace"}[l]
}

// Acronym 缩写，适用于控制台、文件等输出
func (l Level) Acronym() string {
	return [...]string{" PNC ", " FAT ", " ERR ", " WRN ", " INF ", " DBG ", " TRC "}[l]
}

// Color 控制台颜色输出
func (l Level) Color() string {
	switch l {
	case PanicLevel:
		return aurora.Bold(aurora.BgRed(l.Acronym())).String()
	case FatalLevel:
		return aurora.Bold(aurora.BgRed(l.Acronym())).String()
	case ErrorLevel:
		return aurora.Bold(aurora.BgRed(l.Acronym())).String()
	case WarnLevel:
		return aurora.Bold(aurora.BgYellow(l.Acronym())).String()
	case InfoLevel:
		return aurora.Bold(aurora.BgGreen(l.Acronym())).String()
	case DebugLevel:
		return aurora.Bold(aurora.BgBlue(l.Acronym())).String()
	case TraceLevel:
		return aurora.Bold(aurora.BgMagenta(l.Acronym())).String()
	default:
		return aurora.Bold(aurora.BgRed(" ??? ")).String()
	}
}
