package main

import (
	"fmt"
	"unicode"
)

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func main() {
	s := "乔鹏aaa"

	for _,rr := range s {
		if unicode.Is(unicode.Han,rr){
			fmt.Printf("%c\n",rr)
		}

	}


	//
	//left := dispatchCoin(users, distribution)
	//fmt.Printf("还剩下:%v\n", left)
	//for k,v := range distribution{
	//	fmt.Printf("%s:%d\n",k,v)
	//}
}

func dispatchCoin(users []string, distribution map[string]int) int {
	for _, v := range users {
		sum := 0
		for _, i := range v {
			switch i {
			case 'e', 'E':
				coins -= 1
				sum += 1
			case 'i', 'I':
				coins -= 2
				sum += 2
			case 'o', 'O':
				coins -= 3
				sum += 3
			case 'u', 'U':
				coins -= 4
				sum += 4
			}
		}
		distribution[v] = sum
	}
	return coins
}
