package main

import (
	"fmt"
	"github.com/study/logger/mylogger"
	"log"
	"os"
	"time"
)

func testLog() {
	for {
		log.Println("这是一条普通的日志")
		time.Sleep(time.Second * 3)
	}
}

func writeLogFile() {
	file, err := os.OpenFile("D:\\GoWorkspace\\log\\test.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open file faield ,%v", err)
		return
	}
	//打印在终端
	//log.SetOutput(os.Stdout)
	//往文件里写入内容
	log.SetOutput(file)
	testLog()
}

func main() {
	//writeLogFile()
	//log := mylogger.NewLog("debug")
	log := mylogger.NewFieLogger("INFO", "./", "D:\\GoWorkspace\\log\\test.log", 10*1024*1024)
	for {
		log.Debug("这是一条Debug日志")
		log.Info("这是一条Info日志")
		log.Warning("这是一条Warning日志")
		time.Sleep(time.Second * 5)
	}

}
