package main

import "fmt"

func main(){
	//Go语言中不存在指针操作 只需要记住两个符号
	//&:取地址
	//*:根据地址取值
	//GO语言里边的指针只能只读不能修改 不能修改指针变量指向的地址
	var s1="aaa"
	fmt.Println(&s1)
	fmt.Println(*&s1)

	var s2=s1
	fmt.Println(&s2)
	fmt.Println(s2)
	s2="bbb"
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(&s2)

	//make和new都是用来申请内从得
	//new一班很少用，一般是用来给基本数据类型申请内存 string  int 还有结构体 返回的是对应类型的指针（*int *string）
	//make是用来给slice map chan申请内存的，make函数返回的是对应的三个类型本身
}