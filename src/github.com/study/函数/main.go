package main

import "fmt"

//函数的定义有参数有返回值
func f0(x int, y int) (ret int) {
	return x + y
}

//没参数没返回值
func f2() {
	fmt.Println("f2")
}

//有参数没有返回值
func f1(x int) {
	fmt.Println(x)
}

//没有参数有返回值
func f3() int {
	ret := 3
	return ret
}

//返回值可以命名也可以不命名
//命名的返回值就相当于在函数中生命一个变量
//如果使用的是命名的返回值 那么return后边就可以不写
func f4(x int, y int) (ret int) {
	ret = x + y
	return
}

//多个返回值
func f5() (int, int) {
	return 1, 2
}

//参数类型简写
//当参数中连续多个参数的类型一致时，我们可以将非最后一个参数的类型省略
func f6(x, y int) int {
	return x + y
}

//可变参数
func f7(x string, y ...int) {
	fmt.Println(x)
	//多个参数时一个slice 是一个切片 []int 具体是什么类型的切片看...后边的类型
	//而且可变参数必须放在参数的最后
	fmt.Printf("y:%T", y)
}

func main() {
	ret := f0(1, 2)
	fmt.Println(ret)
	f1(2)
	f2()
	fmt.Println(f3())
	ret4 := f4(1, 2)
	fmt.Println(ret4)
	_, i2 := f5()
	fmt.Println(i2)
	ret7 := f6(1, 2)
	fmt.Println(ret7)
	f7("哈哈哈", 1, 2, 3)
}
