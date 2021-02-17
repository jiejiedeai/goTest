package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {

	/**
	判断字符串中汉字个数
	*/
	s1 := "富丰园小区hello"
	count1 := 0
	//依次拿到字符串中的字符
	for _, c := range s1 {
		//判断字符串中的字符是否是汉字
		if unicode.Is(unicode.Han, c) {
			count1++
		}
	}
	fmt.Println(count1)

	/**
	判断单次出现的次数
	*/
	s2 := "how du you du"
	s3 := strings.Split(s2, " ")
	m1 := make(map[string]int, len(s3))
	for _, c := range s3 {
		fmt.Println(c)
		if _, ok := m1[c]; !ok {
			m1[c] = 1
		} else {
			m1[c]++
		}
	}
	for key, value := range m1 {
		fmt.Println(key, value)
	}
}
