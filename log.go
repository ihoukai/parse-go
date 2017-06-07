package parse

import (
	"fmt"
)

// LogLevel 日志级别
type LogLevel int

const (
	LogDebug   LogLevel = 0
	LogInfo    LogLevel = 1
	LogWarning LogLevel = 2
	LogError   LogLevel = 3
)

var (
	log       ILog
	logEnable bool
	logLevel  LogLevel
)

// ILog 日志模块
type ILog interface {
	Debug(format string, a ...interface{})
	Info(format string, a ...interface{})
	Error(format string, a ...interface{})
	Warning(format string, a ...interface{})
}

func init() {
	SetLog(&defaultLog{})
}

// SetLog 设置日志模块
func SetLog(ilog ILog) {
	log = ilog
	logEnable = true
	logLevel = LogDebug
}

// SetLogEnable 是否开启日志
func SetLogEnable(enable bool) {
	logEnable = enable
}

// SetLogLevel 设置日志级别
func SetLogLevel(level LogLevel) {
	logLevel = level
}

// Debug log
func Debug(format string, a ...interface{}) {
	if log != nil &&
		logEnable &&
		LogDebug >= logLevel {
		log.Debug(format, a...)
	}
}

// Info log
func Info(format string, a ...interface{}) {
	if log != nil &&
		logEnable &&
		LogInfo >= logLevel {
		log.Info(format, a...)
	}
}

// Error log
func Error(format string, a ...interface{}) {
	if log != nil &&
		logEnable &&
		LogError >= logLevel {
		log.Error(format, a...)
	}
}

// Warning log
func Warning(format string, a ...interface{}) {
	if log != nil &&
		logEnable &&
		LogWarning >= logLevel {
		log.Warning(format, a...)
	}
}

type defaultLog struct {
}

// Debug log
func (d *defaultLog) Debug(format string, a ...interface{}) {
	fmt.Println(fmt.Sprintf(format, a...))
}

// Info log
func (d *defaultLog) Info(format string, a ...interface{}) {
	fmt.Println(fmt.Sprintf(format, a...))
}

// Error log
func (d *defaultLog) Error(format string, a ...interface{}) {
	fmt.Println(fmt.Sprintf(format, a...))
}

// Warning log
func (d *defaultLog) Warning(format string, a ...interface{}) {
	fmt.Println(fmt.Sprintf(format, a...))
}
