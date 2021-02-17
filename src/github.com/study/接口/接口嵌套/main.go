package main

import "fmt"

type mover interface {
	move()
}

type eater interface {
	eat(string)
}

type animal interface {
	mover
	eater
}

type cat struct {
	name string
	feet uint8
}

func (c *cat)eat(food string){
	fmt.Printf("猫吃%s",food)
}

func (c *cat)move(){
	fmt.Printf("猫用%d腿动",c.feet)
}