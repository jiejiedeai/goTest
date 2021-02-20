package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/**
原子操作类似java中得Atomic 但是go语言中只有整数原子操作
不用再自己声明加锁和释放锁 Atomic也是无锁操作CAS
int类型用 atomic.AddInt64 AddInt32 AddUint64 AddUint32
此外还有Store方法设置值 Load方法读取值 Add方法修改值 Swap交换之 CompareAndSwapInt 比较并交换操作
*/
var x int64
var f float64
var wg sync.WaitGroup

func add() {
	defer wg.Done()
	atomic.AddInt64(&x, 1)
}

func main() {
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go add()
	}
	wg.Wait()
	fmt.Println(x)
}
