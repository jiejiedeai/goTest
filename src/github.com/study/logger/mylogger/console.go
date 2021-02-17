package mylogger

import (
	"fmt"
	"time"
)

//构造函数
func NewLog(levelStr string) ConsoleLogger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return ConsoleLogger{
		Level: level,
	}
}

/**
打印输入日志内容
*/
func (c ConsoleLogger) log(lv LogLevel, msg string) {
	if c.enable(lv) {
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3)
		format := now.Format("2006-01-02 15:04:05")
		fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n", format, getLogString(lv), funcName, fileName, lineNo, msg)
	}
}

/**
获取日志等级
*/
func (c ConsoleLogger) enable(level LogLevel) bool {
	return level >= c.Level
}

func (c ConsoleLogger) Debug(msg string) {
	c.log(DEBUG, msg)
}

func (c ConsoleLogger) Info(msg string) {
	c.log(INFO, msg)
}

func (c ConsoleLogger) Warning(msg string) {
	c.log(WARNING, msg)
}

func (c ConsoleLogger) Error(msg string) {
	c.log(ERROR, msg)
}

func (c ConsoleLogger) Fatal(msg string) {
	c.log(FATAL, msg)
}
