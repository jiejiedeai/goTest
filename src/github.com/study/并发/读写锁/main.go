package main

import (
	"fmt"
	"sync"
	"time"
)

/**
读写互斥锁
类似java中得ReadWriteLock 读写互斥 写写互斥 读读不互斥
应用场景都是 读得场景都是大于写得场景 如果只有写场景用不上读写互斥锁
读写锁 sync.RWMutex  读锁还是Lock和UnLock 写锁方法是RLock和RUnLock
替换读写锁之前大概2-3秒
替换后1.8秒 总之就是读写锁效率更高时间更快
*/
var (
	x      = 0
	lock   sync.Mutex
	wg     sync.WaitGroup
	rwLock sync.RWMutex
)

func write() {
	defer wg.Done()
	rwLock.RLock()
	x = x + 1
	time.Sleep(time.Millisecond * 5)
	rwLock.RUnlock()
}

func read() {
	defer wg.Done()
	rwLock.Lock()
	fmt.Println(x)
	time.Sleep(time.Millisecond)
	rwLock.Unlock()
}

func main() {
	start := time.Now()
	for i := 0; i < 100; i++ {
		go write()
		wg.Add(1)
	}
	for i := 0; i < 1000; i++ {
		go read()
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start))
}
