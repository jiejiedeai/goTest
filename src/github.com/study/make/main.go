package main

import "fmt"

/**
make函数创造切片
*/
func main() {

	//第一个参数写什么类型的切片 第二个参数写长度 第三个参数写容量 如果不写容量默认就是和长度一致
	s1 := make([]int, 5)
	s2 := make([]int, 5, 10)
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2=%v len(s2)=%d cap(s2)=%d\n", s2, len(s2), cap(s2))
}
