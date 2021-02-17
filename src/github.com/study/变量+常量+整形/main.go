package main

//导入语句
import "fmt"

//函数外只能放（变量\常量\helloworld\类型） 没法放逻辑语句

//变量声明
var name string
var isOk bool
var age int

//批量声明变量
//声明完变量后变量是空值
var (
	a string //""
	b bool   //false
	c int    //0
)

//常量使用 const 定义了常量是不可以修改的 在运行期间不可以改变
const f1 = 10
const f2 = "aaa"

//批量设置常量 如果值设置一个值后边的值都会跟第一个值一样
const (
	n1 = 100
	n2
	n3
)

//iota 常量计数器，只能在常量表达式中使用，再const关键字出现的地方将被设置为0，const中每新增一行常量声明将使iota计数一次
//iota 可以理解为const语句块中的索引，使用iota能简化定义，再枚举中经常会使用

//程序的入口函数
func main() {
	a = "理想"
	c = 16
	//非全局变量声明了必须要使用 不适用不能编译过
	fmt.Printf("name:%s", a)
	fmt.Print(c)
	//类型推导
	var s1 = "qp"
	fmt.Println(s1)
	//简短变量声明 相当于生变简写 并且只能再函数中用 不能再全局中用
	s2 := "zj"
	fmt.Println(s2)
	//匿名变量用_ 匿名变量不占用内存空间，不会分配内存，所以匿名变量之间不存在重复声明 可以忽略不想要的值
	_, y := foo()
	x, _ := foo()
	fmt.Println(x)
	fmt.Println(y)

	fmt.Println(f1)
	fmt.Println(f2)

	//10进制度
	var i1 = 101
	//10进制用%d
	fmt.Printf("%d\n", i1)
	//10进制转换成2进制 %b
	fmt.Printf("%b\n", i1)
	//10进制转换成8进制 %o
	fmt.Printf("%o\n", i1)
	//10进制转换成16进制 %x
	fmt.Printf("%x\n", i1)
	//8进制
	i2 := 077
	fmt.Printf("%d\n", i2)
	//16进制
	var i3 = 0xff
	fmt.Printf("%d\n", i3)
	//查看变量类型T%
	fmt.Printf("T%\n", i3)
}

func foo() (int, string) {
	return 10, "hhh"
}
