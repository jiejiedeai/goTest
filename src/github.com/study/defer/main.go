package main

import "fmt"

/**
Go语言中defer语句会将其后面紧跟随得语句延迟处理。在defer归属函数即将返回时，将延迟处理的语句按defer
定义的顺序逆序执行，也就是说先被defer的语句最后执行 最后被defer的语句先执行
比如关闭文件 关闭网络socket 关闭数据库链接避免忘记可以先用defer声明 在执行完成后将其关闭
如果defer 执行过程中函数中嵌套了函数 它会先把嵌套得函数执行完 然后再按照defer声明顺序逆序执行
*/
func main() {
	//fmt.Println("start")
	//defer fmt.Println("1")
	//defer fmt.Println("2")
	//defer fmt.Println("3")
	//defer fmt.Println("end")
	fmt.Println(f1()) //返回5 defer 在return之后执行 返回5是因为修改的是x不是返回值
	fmt.Println(f2()) //返回6 第一步返回值赋值 x指向了5 然后此时执行defer语句 x++ 变成了6
	fmt.Println(f3()) //返回5 x给了y 然后执行defer 此时修改的是x 对返回的y没有变只是x变更了
	fmt.Println(f4()) //返回值是5 函数传参走的是值传递改变的是函数副本
}

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f4() (x int){
	defer func(x int) {
		x++
	}(x)
	return 5
}
