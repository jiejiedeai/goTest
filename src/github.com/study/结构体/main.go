package main

import "fmt"

//结构体可以使用new来进行实例话开辟内存
//new开辟以后得到得都是内存地址 不同于make make只能开辟特定类型得 切片 map channle等 返回得就是对应类型

//结构体 一个结构体是一块连续的内存 每个字段得内存地址也都是连在一起的
type Person struct {
	age          int
	name, gender string
}

//因为go语言中传过来得永远都是拷贝 所以不会印象外边得p 虽然带入到了方法f中但是外边得源文件不会改变
func f1(x Person) {
	x.gender = "女"
}

//如果想修改里边得值 需要传入指针 这样如的就是内存地址
func f2(x *Person) {
	//根据内存地址找到变量 修改得就是原来得变量
	(*x).gender = "nv"
	//上面这个写法还可以用go得语法糖直接x.gender go会自己判断如果是指针类型变量就找指针对应变量
	//因为go是不能修改指针得也不能对指针进行操作 可以直接写 x.gender="nv"
}

func f3() {
	v := new(Person)
	v.name = "乔鹏"
	v.gender = "男"
	//返回得是*main.Person 指针
	fmt.Printf("%T\n", v)
	//打印得值也就是内存地址
	fmt.Printf("%p\n", v)
	//打印纸
	fmt.Println(*v)
	fmt.Printf("p2=%#v\n", v)
}

func main() {
	var p Person
	p.age = 10
	p.name = "乔鹏"
	p.gender = "男"
	f1(p)
	fmt.Println(p)
	//使用取址符号
	f2(&p)
	fmt.Println(p)
	f3()
	//结构体指针 生命变量类型同时初始化
	var p3 = Person{
		name:   "哈哈哈",
		gender: "男",
	}
	fmt.Printf("%#v\n", p3)
	//结构体直接值初始化 值得顺寻要跟生命得结构体顺寻一样
	p4 := Person{
		10,
		"嘿嘿额",
		"女",
	}
	fmt.Printf("%#v", p4)

	//构造函数
	p1 := newPerson(18, "乔鹏", "男")
	p2 := newPerson(20, "周杰", "女")
	//当结构体比较大的时候尽量使用结构体指针 减少结构内存开销
	// func newPerson() *person 这样person就不会来回拷贝 拷贝的都是那一个内存地址开销比较小
	fmt.Println(p1, p2)
	//嵌套结构体
	peo1 := people{
		name: "乔鹏",
		age:  10,
		addr: address{
			province: "北京",
			city:     "北京",
		},
		noNameAddress: noNameAddress{
			province: "北京",
			city:     "北京",
		},
	}
	fmt.Println(peo1, peo1.name, peo1.addr.city)
	//go 中语法糖 因为noNameAddress是person中嵌套结构体所以可以直接访问 但是嵌套结构体必须是匿名嵌套
	fmt.Println(peo1.city)
}

//约定成俗 构造函数用new开头
func newPerson(age int, name string, gender string) Person {
	return Person{
		age:    age,
		name:   name,
		gender: gender,
	}
}

//用于匿名嵌套person用
type noNameAddress struct {
	province string
	city     string
}

//嵌套结构体
type address struct {
	province string
	city     string
}

type people struct {
	name string
	age  int
	addr address
	noNameAddress
}

type company struct {
	name string
	addr address
}
