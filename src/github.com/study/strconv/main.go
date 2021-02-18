package main

import (
	"fmt"
	"strconv"
)

/**
字符串转换
*/
func main() {
	s1 := "10000"
	parseInt, err := strconv.ParseInt(s1, 10, 64)
	if err != nil {
		fmt.Println("parseInt failed,err:", err)
		return
	}
	aToi, err := strconv.Atoi(s1)
	if err != nil {
		fmt.Println("aoti faield,err:", err)
		return
	}
	//一般用aToi将字符串转成int 上边这种parseInt是可以指定禁止和int 8 16 32 64类型等
	fmt.Println(aToi)
	fmt.Println(parseInt)
	i1 := 10000
	parseString := fmt.Sprintf("%#v", i1)
	iToa := strconv.Itoa(i1)
	fmt.Printf("%#v\n", parseString)
	fmt.Printf("%#v\n", iToa)

}
