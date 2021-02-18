package main

import (
	"fmt"
	"math/rand"
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

func main() {
	for i := 0; i < 10; i++ {
		wc.Add(1)
		go f1(i)
	}
	//main函数结束由main函数启动的goroutine也就都结束了
	fmt.Println("main")
	wc.Wait()
}
