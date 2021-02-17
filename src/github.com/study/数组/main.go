package main

import "fmt"

//存放元素的容器必须指定存放元素的类型和容量（长度）
func main() {

	var a1 [3]bool
	var a2 [4]bool
	fmt.Printf("a1:%T a2:%T\n", a1, a2)
	//数组初始化 如果不初始化默认都是零值（布尔值：false 整形和浮点型都是0 字符串"")
	fmt.Println(a1, a2)
	//初始化方式1

	a1 = [3]bool{true, true, true}
	fmt.Println(a1)
	//初始化方式2 根据初始值自动推断数组长度
	a3 := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println(a3)
	//初始化方式3 根据索引初始化
	a4 := [...]int{0: 1, 14: 2}
	fmt.Println(a4)
	//数组遍历 根据索引方式便利
	citys := [...]string{"北京", "上海", "深圳"}
	for i := 0; i < len(citys); i++ {
		fmt.Println(citys[i])
	}
	//数组遍历 for range
	for i, v := range citys {
		fmt.Println(i, v)
	}
	//多维数组
	var a5 [3][2]int
	a5 = [3][2]int{
		{1, 2},
		{11, 12},
		{31, 32},
	}
	//多维数组遍历
	for _, v1 := range a5 {
		fmt.Printf("第一层:%d\n", v1)
		for _, v2 := range v1 {
			fmt.Printf("第二层:%d\n", v2)
		}
	}
	//数组是值类型 值传递
	b1 := [3]int{1, 2, 3}
	b2 := b1
	b1[0] = 100
	fmt.Println(b1, b2)
}
