package main

import "fmt"

func main() {
	//此时map的内存地址是nil 没有初始化不能使用 用下边的make初始化
	var m1 map[string]int
	//用make初始化 要预估好map容量,避免在程序中动态扩容 超过范围后会自动扩容 最好提前算出来
	m1 = make(map[string]int, 1)
	m1["理想"] = 10
	m1["测试"] = 20
	fmt.Println(m1)
	//获取值
	fmt.Println(m1["理想"])
	//获取一个不存在的变量 go语言中一般约定成俗用ok来接收 用别的也可以
	value, ok := m1["xxx"]
	if !ok {
		fmt.Println("无此key")
	} else {
		fmt.Println(value)
	}
	//map的遍历
	for k, v := range m1 {
		fmt.Println(k, v)
	}
	//map的遍历只获取值
	for _, v := range m1 {
		fmt.Println(v)
	}

	//删除key
	delete(m1, "测试")
	fmt.Println(m1)

	//元素为map的切片
	var s1 = make([]map[int]string, 1)
	//map 初始化
	s1[0] = make(map[int]string, 1)
	s1[0][10] = "aaa"
	s1[0][11] = "bbb"
	fmt.Println(s1)

	//值为切片类型的map
	var s2 = make(map[string][]int, 10)
	s2["北京"] = []int{10, 20, 30}
	fmt.Println(s2)


}
