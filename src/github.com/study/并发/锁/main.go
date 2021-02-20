package main

import (
	"fmt"
	"sync"
)

/**
锁
类似java中得writeLock
sync.Mutex声明锁 然后对公共访问得代码前Lock 用完锁得地方UnLock释放
*/
var x = 0

var wg sync.WaitGroup

var lock sync.Mutex

func add() {
	for i := 0; i < 500000; i++ {
		lock.Lock()
		x = x + 1
		lock.Unlock()
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
