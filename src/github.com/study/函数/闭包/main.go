package main

import "fmt"

func f1(f func()) {
	fmt.Println("this is f1")
	f()
}

func f2(x, y int) {
	fmt.Println("this is f2")
	fmt.Println(x + y)
}

func f3(f func(int,int),m,n int) func() {
	tmp := func(){
		f(m,n)
	}
	return tmp
}

func main() {
	//把f2当成参数传到f1中 现在是不可以因为f1得参数func 跟f2类型不一样 f1(f2)
	//此处就用闭包 将f2包装一下成f3 f3中参数有3个 一个f2函数当成参数 m和n是普通参数 再f3函数中执行f2函数又用到m和n
	ret := f3(f2,10,20)
	f1(ret)
}
