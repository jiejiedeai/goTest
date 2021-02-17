package main

import "fmt"

/**
go语言中没有异常(try catch)进行处理 但是使用panic和recover模式来处理错误
panic可以再任何地方引发，但是recover只有再defer函数中有效
panic和recover成对出现 这样防止程序崩溃不会退出程序
该用panic用panic recover能少用就少用 比如连接数据库失败了 那就不能recover panic就是让他失败
recover 只是恢复之前panic收集的资源
并且defer recover 必须在panic语句之前定义 因为一旦发生了panic以后如果defer在下边那么就永远执行不到
*/
func main() {
	funA()
	funB()
	funC()

}

func funA() {
	fmt.Println("func A")
}

func funB() {
	defer func() {
		err := recover()
		fmt.Println(err)
		fmt.Println("错误处理完毕")
	}()
	panic("panic in B")
	fmt.Println("b")
}

func funC() {
	fmt.Println("func C")
}
