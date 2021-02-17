package main

import "fmt"

/**
一个接口变量可以接收任意类型
不要为了接口而定义接口 有两个以上结构体才需要定义接口
 */

type animal interface {
	move()
	eat(something string)
}

type cat struct {
	name string
	feet int
}

type chicken struct {
	name string
	feet int
}

func (c chicken) eat(food string) {
	fmt.Printf("鸡吃%s\n", food)
}

func (c chicken) move() {
	fmt.Printf("鸡用%d动\n", c.feet)
}

func (c cat) eat(food string) {
	fmt.Printf("猫吃%s\n", food)
}

func (c cat) move() {
	fmt.Printf("猫用%d动\n", c.feet)
}

func main() {

	var a1 animal

	bc := cat{
		name: "蓝猫",
		feet: 4,
	}
	a1 = bc
	a1.eat("小黄花")
	a1.move()
	fmt.Printf("%T\n", a1)
	kfc := chicken{
		name: "鸡",
		feet: 2,
	}
	a1 = kfc
	a1.eat("饲料")
	a1.move()
	fmt.Printf("%T\n", a1)
	/**
	上边先后把cat和chicken给了a1 但是a1还是animal
	可是通过上边fmt.Printf打印后 发现是main.chicken和main.cat类型 并不是animal类型
	因为a1的底层分成了两部分 一部分存了类型另一部分存值 但是初始的时候两个都是nil
	那么我们把cat和chicken结构体给到a1后其实是带了两部分一部分是结构体cat或者chicken 另一部分是结构体的值
	那么a1一部分存储的就是main.cat类型 另一部分存的是cat{"name:"蓝猫",feet:4,}
	*/
}
