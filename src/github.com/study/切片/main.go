package main

import "fmt"

func main() {

	//切片是一个拥有相同类型元素的可变长度的序列。它是基于数组类型做的一层封装。它非常灵活，支持自动扩容
	//切片是一个引用类型
	//切片的定义 之定义了存储的类型是什么没有定义长度
	var s1 []int
	var s2 []string
	fmt.Println(s1, s2)
	//切片的初始化
	s1 = []int{1, 2, 3}
	s2 = []string{"沙河", "看丹桥", "庞各庄"}
	fmt.Println(s1, s2)
	//切边只能跟nil比较
	var ss []string
	fmt.Println(ss == nil)
	ss = []string{"aaa"}
	fmt.Println(ss == nil)

	//切片的长度和容量
	fmt.Printf("len:%d cap:%d\n", len(s1), cap(s1))
	fmt.Printf("len:%d cap:%d\n", len(s2), cap(s2))

	//由数组得到切片
	a1 := []int{1, 3, 5, 7, 9, 11, 13}
	a2 := a1[0:4] //基于一个数组的切割，左包含右不包含,左闭右开
	fmt.Println(a2)
	a3 := a1[:4]
	a4 := a1[3:]
	a5 := a1[:]
	fmt.Println(a3, a4, a5)

	//长度和容量
	//切片的容量是指底层数组的容量 len:4 cap:7
	fmt.Printf("len:%d cap:%d\n", len(a3), cap(a3))
	//上边的总结对数组中间到最后切就不严谨 所以是底层数组从第一个元素到最后的元素数量 len:4 cap:4
	fmt.Printf("len:%d cap:%d\n", len(a4), cap(a4))
	//len:7 cap:7
	fmt.Printf("len:%d cap:%d\n", len(a5), cap(a5))

	//切片再切片
	a6 := a4[3:]
	//len:1 cap:1
	fmt.Printf("len:%d cap:%d\n", len(a6), cap(a6))

	//底层数据修改切片是否变更
	a1[6] = 1300
	//结论就是底层数组变更切片就变更 因为切片是引用传递不同于数组是值传递 因为切片底层是数组没有自己的值
	fmt.Println("a4:", a4)
	fmt.Println("a6:", a6)

	//var s1 []int
	//s1=[]int{1,2,3}
	//s3 := s1
	//s3[0]=3
	//fmt.Println(s1) //3,2,3
	//fmt.Println(s3) //3,2,3
	////s1和s3一样是因为切片不存值
	///**
	//	首先判断，如果新申请的容量大于2倍的旧容量，最终容量就是新申请的容量
	//	否则判断，如果切片长度小于1024，则最终容量就是就容量的2倍
	//	否则判断，如果旧切片长度大于等于1024 则最终容量从旧容量开始循环增加原来的1/4 直到最终容量大于等于新申请的容量
	//	如果最终容量计算溢出，则最终容量就是新申请的容量
	// */

}
