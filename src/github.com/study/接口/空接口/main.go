package main

import "fmt"

/**
空接口没有必要有名字 一般直接写成 interface{} 所有类型都实现了空接口 一般用于两种情况
1.方法的参数接收 我们不知道让人家传入什么类型时候可以将参数定义未空接口
2.类似java这中map类型 key是个字符串 值是一个空接口 可以接收任何类型的参数
*/
func show(a interface{}) {
	fmt.Printf("类型:%T value:%v\n", a, a)
}

func main() {
	var m1 map[string]interface{}
	m1 = make(map[string]interface{}, 16)
	m1["name"] = "qiao"
	m1["age"] = 18
	m1["gender"] = "男"
	m1["married"] = true
	m1["hobby"] = [...]string{"chouyan", "hejiu"}
	show(false)
	show(nil)
	show(m1)

}
