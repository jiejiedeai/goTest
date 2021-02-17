package main

import (
	"fmt"
	"sort"
)

/**
append元素扩容策略
首先判断，如果新申请的容量大于2倍的旧容量，最终容量就是新申请的容量
否则判断，如果切片长度小于1024，则最终容量就是就容量的2倍
否则判断，如果旧切片长度大于等于1024 则最终容量从旧容量开始循环增加原来的1/4 直到最终容量大于等于新申请的容量
如果最终容量计算溢出，则最终容量就是新申请的容量
*/
func main() {
	s1 := []string{"北京", "上海", "深圳"}

	s1 = append(s1, "广州")
	fmt.Printf("s1=%v len:%d cap:%d\n", s1, len(s1), cap(s1))
	s1 = append(s1, "杭州", "成都")
	fmt.Printf("s1=%v len:%d cap:%d\n", s1, len(s1), cap(s1))

	ss := []string{"武汉", "西安", "苏州"}
	//...表示拆开
	s1 = append(s1, ss...)
	fmt.Printf("s1=%v len:%d cap:%d\n", s1, len(s1), cap(s1))

	//copy 是值传递
	a1 := []int{1, 3, 5, 7}
	a2 := a1
	a3 := make([]int, 3, 3)
	copy(a3, a1)
	fmt.Println(a1, a2, a3)
	a1[0] = 100
	fmt.Println(a1, a2, a3)

	//切片删除元素 go中没有直接从切片中删除元素得方法 将索引1得元素3删除掉
	a1 = append(a1[:1], a1[2:]...)
	fmt.Println(a1)

	//切片只能是连续内存所以切片修改得其实是底层数组得值
	x1 := [...]int{1, 3, 5, 7}
	xx1 := x1[:]
	fmt.Println(xx1, len(xx1), cap(xx1))
	fmt.Printf("%p\n", &xx1[0])
	xx1 = append(xx1[:], xx1[2:]...)
	fmt.Printf("%p\n", &xx1[0])
	fmt.Println(x1)

	var a = make([]int, 5, 10)
	for i := 0; i < 10; i++ {
		a = append(a, i)
	}
	fmt.Println(a)

	//排序功能
	var a11 = [...]int{1, 4, 7, 3, 8, 2, 9}
	sort.Ints(a11[:])
	fmt.Println(a11)

}
