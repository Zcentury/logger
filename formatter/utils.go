package formatter

import (
	"runtime"
	"strings"
	"sync"
)

var (
	minimumCallerDepth int
	callerInitOnce     sync.Once
)

const (
	loggerPackage          = "github.com/Zcentury/logger"
	logrusPackage          = "github.com/sirupsen/logrus"
	maximumCallerDepth int = 30
	knownLogrusFrames  int = 9
)

// getCaller 检索第一个非 logrus 调用函数的名称
func getCaller() *runtime.Frame {
	// 缓存该包的完全限定名称（函数名）
	callerInitOnce.Do(func() {
		// 创建一个存储调用栈信息的数组，最大深度为 maximumCallerDepth
		pcs := make([]uintptr, maximumCallerDepth)
		// 获取当前的调用栈信息
		_ = runtime.Callers(0, pcs)

		// 设置最小调用深度（跳过 knownLogrusFrames 层级）
		minimumCallerDepth = knownLogrusFrames
	})

	// 限制堆栈帧的回溯深度，以避免无止境的查找
	pcs := make([]uintptr, maximumCallerDepth)
	// 获取调用栈的深度，跳过前面的最低深度
	depth := runtime.Callers(minimumCallerDepth, pcs)
	// 获取堆栈帧的迭代器
	frames := runtime.CallersFrames(pcs[:depth])

	// 遍历堆栈帧
	for f, again := frames.Next(); again; f, again = frames.Next() {
		// 获取当前堆栈帧的包名
		pkg := getPackageName(f.Function)

		// 如果当前调用不在 logrus 包中，返回当前堆栈帧（即当前调用函数）
		if pkg != loggerPackage && pkg != logrusPackage {
			return &f //nolint:scopelint
		}
	}

	// 如果遍历完所有堆栈帧都没有找到非 logrus 包的调用，返回 nil
	return nil
}

func getPackageName(f string) string {
	for {
		lastPeriod := strings.LastIndex(f, ".")
		lastSlash := strings.LastIndex(f, "/")
		if lastPeriod > lastSlash {
			f = f[:lastPeriod]
		} else {
			break
		}
	}

	return f
}
