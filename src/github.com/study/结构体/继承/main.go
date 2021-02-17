package main

import "fmt"

//go中没有直接继承关系 可以用结构体嵌套模拟实现其他语言中的继承

type animal struct {
	name string
}

//匿名结构体嵌套
type dog struct {
	feet uint8
	animal
}

func (a animal) move() {
	fmt.Printf("%s会动!\n", a.name)
}

func (d dog) wang() {
	fmt.Printf("%s在叫:汪汪汪~\n", d.name)
}

func main() {
	d1 := dog{
		animal: animal{
			name: "尼克",
		},
		feet: 4,
	}
	fmt.Println(d1)
	d1.wang()
	//相当于继承 dog结构体完整的包裹了animal的结构体 所以animal能做的事儿 dog也可以做
	d1.move()
}
