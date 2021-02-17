package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

type FileLogger struct {
	Level    LogLevel
	filePath string //日志保存路径
	fileName string //日志文件名
	maxSize  int64
	fileObj  *os.File
	errObj   *os.File
}

//构造函数
func NewFieLogger(levelStr, filePath, fileName string, maxSize int64) *FileLogger {
	logLevel, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	file := &FileLogger{
		Level:    logLevel,
		filePath: filePath,
		fileName: fileName,
		maxSize:  maxSize,
	}
	err = file.initFile() //按照文件路径和文件名将文件打开
	if err != nil {
		panic(err)
	}
	return file
}

//按照文件路径和文件名将文件打开
func (f *FileLogger) initFile() error {
	fileObj, err := os.OpenFile(path.Join(f.filePath, f.fileName), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("打开文件出错 %s\n", err)
		return err
	}
	errFileObj, err := os.OpenFile(path.Join(f.filePath, f.fileName), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("打开错误文件出错 %s\n", err)
		return err
	}
	//日志文件都已经打开了
	f.fileObj = fileObj
	f.errObj = errFileObj
	return nil
}

//关闭文件
func (f *FileLogger) Close() {
	f.fileObj.Close()
	f.errObj.Close()
}

/**
是否需要切割
*/
func (f *FileLogger) checkSize(file *os.File) bool {
	//获取文件大小
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("获取问价大小失败%s\n", err)
		return false
	}
	size := fileInfo.Size()
	return fileInfo.Size() >= size
}

/**
打印输入日志内容
*/
func (f *FileLogger) log(lv LogLevel, msg string) {
	if f.enable(lv) {
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3)
		format := now.Format("2006-01-02 15:04:05")
		if f.checkSize(f.fileObj) {
			/**
			切割日志文件
			1.关闭当前日志文件
			2.备份一下并且重命名
			3.打开一个新文件 将打开的新文件对象复制给f.fileObj
			*/
			f.Close()
			timeStamp := time.Now().Format("20060102150405")
			logName := path.Join(f.filePath, f.fileName)
			newLogName := fmt.Sprintf("%s.bak%s",logName,timeStamp)
			os.Rename(logName,newLogName)
			fileObj, err := os.OpenFile(newLogName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Printf("打开文件失败%s\n",err)
			}
			f.fileObj = fileObj
			fmt.Fprintf(f.fileObj, "[%s] [%s] [%s:%s:%d] %s\n", format, getLogString(lv), funcName, fileName, lineNo, msg)
		}else {
			fmt.Fprintf(f.fileObj, "[%s] [%s] [%s:%s:%d] %s\n", format, getLogString(lv), funcName, fileName, lineNo, msg)
		}
		if lv >= ERROR {
			if f.checkSize(f.errObj){
				/**
				切割日志文件
				1.关闭当前日志文件
				2.备份一下并且重命名
				3.打开一个新文件 将打开的新文件对象复制给f.fileObj
				*/
				f.Close()
				timeStamp := time.Now().Format("20060102150405")
				logName := path.Join(f.filePath, f.fileName)
				newLogName := fmt.Sprintf("%s.bak%s",logName,timeStamp)
				os.Rename(logName,newLogName)
				fileObj, err := os.OpenFile(newLogName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
				if err != nil {
					fmt.Printf("打开文件失败%s\n",err)
				}
				f.fileObj = fileObj
				//如果要记录的日志文件大于等于ERROR级别,我还要在err中记录一遍
				fmt.Fprintf(f.errObj, "[%s] [%s] [%s:%s:%d] %s\n", format, getLogString(lv), funcName, fileName, lineNo, msg)
			}
		}
	}
}

/**
获取日志等级
判断是否需要记录该日志
*/
func (f *FileLogger) enable(level LogLevel) bool {
	return level >= f.Level
}

func (f *FileLogger) Debug(msg string) {
	f.log(DEBUG, msg)
}

func (f *FileLogger) Info(msg string) {
	f.log(INFO, msg)
}

func (f *FileLogger) Warning(msg string) {
	f.log(WARNING, msg)
}

func (f *FileLogger) Error(msg string) {
	f.log(ERROR, msg)
}

func (f *FileLogger) Fatal(msg string) {
	f.log(FATAL, msg)
}
