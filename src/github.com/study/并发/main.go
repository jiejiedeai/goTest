package main

import (
	"fmt"
	"time"
)

/**
go语言并发通过goroutine实现，goroutine类似线程，属于用户态线程
我们可以根据需要创建成千上万个goroutine并发工作
goroutine是由Go语言的运行时调度完成，而线程是由操作系统调度完成
go程序h会智能地将goroutine中的任务合理地分配给每个cpu
在go语言中我们不需要自己进程，线程，协程，你的技能包只有一个技能goroutine
当你需要为某个任务并发执行的时候，你只需要把这个任务包装成一个函数
开启一个goroutine去执行这个函数就可以了
go语言还提供channel在多个goroutine间进行通信
*/

func hello(i int) {
	fmt.Println("hello", i)
}

//程序启动后其实会创建一个main goroutine去执行main函数
func main() {
	for i := 0; i < 100; i++ {
		/**
		开启goroutine跑这个hello函数用 go关键字去执行函数
		for循环开启一堆线程执行 但是顺序是随机的 谁先跑起来不一定
		*/
		go hello(i)
		//匿名函数
		//go func(i int) {
		//	/**
		//	匿名函数加上参数就能避免for循环在执行时候把当前的i给到这个函数
		//	那么匿名函数内部执行的时候就可以用自己的i来操作
		//	如果不传进来参数直接使用for循环的可能会导致这个goroutine启动执行时候外边的for循环跑的够快已经又执行了几条命令 导致i可能变换了
		//	 */
		//	fmt.Println("匿名函数", i)
		//}(i)
	}
	fmt.Println("main")
	//main函数结束由main函数启动的goroutine也就都结束了
	time.Sleep(time.Second * 10)
}
