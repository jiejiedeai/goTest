package main

import "fmt"

//方法其实就是特殊的函数 只不过在函数前添加接收者类型func (变量 类型)函数{}
//目的就是只能这个变量类型调用这个方法
// 结构体首字母大代表其他包可以用 小写的话就是不让其他包使用
type Dog struct {
	name string
}

type person struct {
	name,gender string
	age int
}

//构造函数
func newDog(name string) Dog {
	return Dog{
		name: name,
	}
}

//方法是做总用于特定类型的函数 其实就是限定那个类型可以调用
//接收者变量一般是用于类型的首字母小写 d -> Dog
func (d Dog)wang() {
	fmt.Printf("%s:汪汪汪~",d.name)
}

func newPerson(age int,name string) person{
	return person{
		name: name,
		age: age,
	}
}

func (p person) valueAddPerson(){
	p.age++
}

func (p *person) memoryAddPerson(){
	p.age++
}
func main() {
	d1 := newDog("高飞")
	d1.wang()
	//值作为方法的接收者和指针作为方法的接收者区别
	p1 := newPerson(18,"乔鹏")
	fmt.Println(p1)
	p1.valueAddPerson()
	fmt.Println(p1)
	//此处操作是把p1的内存地址传过去 然后使用指针接收者 不同于上边的value 传的是一个拷贝
	p1.memoryAddPerson()
	fmt.Println(p1)

}

