package main

import "fmt"

/**
接口值接收者和指针接收者
使用值接收者实现接口，结构体类型和结构体指针类型的变量都能存
使用指针接收者实现接口只能存储结构指针的变量
*/
type animal interface {
	move()
	eat(something string)
}

type cat struct {
	name string
	feet int
}

//值接收者
//func (c cat) move() {
//	fmt.Printf("猫用%d动\n", c.feet)
//}
//
//func (c cat) eat(food string) {
//	fmt.Printf("猫吃%s\n", food)
//}

//指针接收者
func (c *cat) move() {
	fmt.Printf("猫用%d动\n", c.feet)
}

func (c *cat) eat(food string) {
	fmt.Printf("猫吃%s\n", food)
}

func main() {
	var a1 animal
	//值类型接收者
	c1 := &cat{
		name: "蓝猫",
		feet: 4,
	}
	//指针类型接收者
	c2 := &cat{   //*cat cat类型的指针
		name: "红毛",
		feet: 5,
	}
	a1 = c1
	fmt.Println(a1)
	a1 = c2
	fmt.Println(a1)

}
