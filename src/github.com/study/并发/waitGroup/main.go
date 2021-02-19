package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

/**
替换time.Sleep等待 使用sync.WaitGroup方式等待goroutine执行完成之后关闭main函数或者继续往下执行
类似java中的CountDownLatch
在执行函数中使用defer wc.Done
外部函数中使用wc.Wait()
那么goroutine什么时候结束呢 当goroutine执行的函数结束时候 goroutine就随之结束
*/

/**
os线程（操作系统线程）一般都有固定的栈内存（通常为2MB）
一个goroutine在生命周期开始的时候很小只有2KB
goroutine的栈是不固定的它可以按需增大或减小
goroutine的栈大小限制可以达到1GB
在go语言中一次性创建10W个goroutine也是可以的不会占用太多的内存
*/

/**
goroutine的调度GMP的调度
G：里面除了存放goroutine的信息外还有与所在的P绑定的关系
M：是Go运行时对操作系统内核的线程虚拟，M与内核线程一般都是一一映射的关系，一个goroutine最终是要放在M上执行的
P：管理这一组goroutine队列，P里面会存储goroutine运行的上下文环境（函数指针，堆栈地址b边界）
P会对自己管理的goroutine队列做一些调度（比如把CPU时间长的goroutine暂停，运行后续的goroutine等）
当自己的队列消费完了就去全局队列里边取
如果全局队列里边页消费完了就去其他P的队列里抢任务
*/

/**
GOMAXPROCS
go运行时的调度器使用GOMAXPROCS参数来确定需要使用多少个OS线程来同时执行go代码
默认值是机器上的CPU核心数。例如在一个8核心的机器上调度器会把go代码同时调度到8个os线程上(GOMAXPROCS是m:n技术中的n）
go语言中可以通过runtime。GOMAXPROCS()函数设置当前程序并发时占用的CPU逻辑核心数 默认使用全部的CPU逻辑核心数
*/
var wc sync.WaitGroup

func f() {
	//保证每次随机数都不一样加一个种子
	rand.Seed(time.Now().UnixNano())
	//Int左包含右不包含 0<=0 x <10
	for i := 0; i < 5; i++ {
		r1 := rand.Int()
		r2 := rand.Intn(10)
		fmt.Println(r1, r2)
	}
}

func f1(i int) {
	defer wc.Done()
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
	fmt.Println(i)
}

func waitGroup() {
	for i := 0; i < 10; i++ {
		wc.Add(1)
		go f1(i)
	}
	//main函数结束由main函数启动的goroutine也就都结束了
	fmt.Println("main")
	wc.Wait()
}

func a() {
	defer wc.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("A:%d\n", i)
	}
}

func b() {
	defer wc.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("B:%d\n", i)
	}
}

func maxProcs1() {
	//在windows中无法演示出该效果
	runtime.GOMAXPROCS(1)
	wc.Add(2)
	go a()
	go b()
	wc.Wait()
}

func main() {
	//waitGroup()
	maxProcs1()
}
