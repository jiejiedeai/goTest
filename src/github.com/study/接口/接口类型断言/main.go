package main

import "fmt"

/**
类型断言
*/

func assign(a interface{}) {
	fmt.Printf("%T\n", a)
	switch t := a.(type) {
	case string:
		fmt.Printf("是一个字符串:%v", t)
	case int:
		fmt.Printf("是一个数字:%v", t)
	case bool:
		fmt.Printf("是一个布尔类型%v", t)
	}
}

func main() {
	//类型断言
	var a interface{}
	a = 100
	v, ok := a.(int8)
	if !ok {
		fmt.Println("猜对了", v)
	}
	//使用switch 做类型判断
	assign("asd")
}
