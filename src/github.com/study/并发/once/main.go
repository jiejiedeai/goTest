package main

import (
	"fmt"
	"sync"
)

/**
sync.Once
在编程很多场景下我们需要确保某些操作高并发得场景下只执行一次
例如只加载一次配置文件，只关闭一次通道等 使用sync.Once  该结构体只有一个方法 Do()
但是Do方法只接收一个没有参数没有返回值得方法 如果初始化得方法不匹配需要写闭包
*/

var loadOnce sync.Once

var onceMap map[string]string

//无参数无返回值函数
func initMap() {
	onceMap["哈哈哈"] = "第一条"
	onceMap["嘿嘿嘿"] = "第二条"
	onceMap["呵呵呵"] = "第三条"
}

//带参数带返回值得函数
func initMapByParam(param string) string {
	if onceMap == nil {
		onceMap = make(map[string]string, 16)
		onceMap["哈哈哈"] = "第一条"
		onceMap["嘿嘿嘿"] = "第二条"
		onceMap["呵呵呵"] = "第三条"
	}
	ret, ok := onceMap[param]
	if !ok {
		panic("参数不存在")
	}
	fmt.Println(ret)
	return ret
}

//闭包
func closePackage(param string) func() {

	tmp := func() {
		initMapByParam(param)
	}
	return tmp
}

func main() {
	//没有参数没有返回值得函数
	//loadOnce.Do(initMap)
	param := "哈哈哈"
	f := closePackage(param)
	loadOnce.Do(f)

}
