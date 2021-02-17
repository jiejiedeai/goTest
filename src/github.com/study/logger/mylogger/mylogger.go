package mylogger

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
)

//往终端写日志内容
type LogLevel uint16

//日志级别
const (
	UNKNOWN LogLevel = iota
	DEBUG
	INFO
	WARNING
	ERROR
	FATAL
)

//日志结构体
type ConsoleLogger struct {
	Level LogLevel
}

func parseLogLevel(s string) (LogLevel, error) {
	s = strings.ToUpper(s)
	switch s {
	case "DEBUG":
		return DEBUG, nil
	case "INFO":
		return INFO, nil
	case "WARNING":
		return WARNING, nil
	case "ERROR":
		return ERROR, nil
	case "FATAL":
		return FATAL, nil
	default:
		err := errors.New("无效的日志级别")
		return UNKNOWN, err

	}
}

func getLogString(lv LogLevel) string {
	switch lv {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}

func getInfo(skip int) (funcName, fileName string, lineNo int) {
	//runtime 运行时自带的gc 原生内置的程序在进行垃圾回收，函数调用啊等相关操作
	//Caller可以拿到当前程序隔了多少层
	/**
	  ok 如果可以取到 ok就是true
	  pc 调用函数信息
	  file 当前执行的文件名
	  line 行数
	*/
	pc, file, lineNo, ok := runtime.Caller(skip)
	if !ok {
		fmt.Printf("runtime.Caller() failed\n")
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	funcName = strings.Split(funcName, ".")[1]
	fileName = path.Base(file)
	return
}
