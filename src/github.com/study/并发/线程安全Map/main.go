package main

import (
	"fmt"
	"strconv"
	"sync"
)

/**
线程安全Map sync.Map{} 类似Java中得currentHashMap
go内置得map也不是安全得 类似java中得HashMap
sync.Map也不需要初始化了 开箱即用
但是取值和获取值删除需要跟map有锁区分
syncMap.Store()设置值
syncMap.Load()加载值
syncMap.Range()遍历
syncMap.Delete()删除值
*/

var syncMap = sync.Map{}
var wg sync.WaitGroup

func main() {
	for i := 0; i < 21; i++ {
		go func(n int) {
			wg.Add(1)
			key := strconv.Itoa(n)
			syncMap.Store(key, n)
			value, ok := syncMap.Load(key)
			if !ok {
				panic("key不存在")
			}
			fmt.Printf("k=:%v,v:=%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
