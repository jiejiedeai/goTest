package main

import (
	"fmt"
	"sync"
)

/**
channel
单纯地将函数b并发执行时没有意义的。函数与函数之间需要交换数据才能体现并发执行函数的意义
虽然可以使用共享内存进行数据交换 但是共享内存在不同的goroutine中容易发生静态问题 为了保证数据交换的正确性
必须使用互斥量对内存进行加锁 这种做法势必会造成性能问题
go语言的并发模型是CSP提倡通过通信共享内存而不是通过共享内存实现通信
如果说goroutine是go程序的执行体那么channel就是他们之间的连接
channel是可以让一个goroutine发送特定值到另一个goroutine的通信机制
go语言中的通道是一种特殊的类型，就想一个通道或者队列总是先进先出规则保证发送数据的顺序
每一个通道都是一个具体类型的导管也就是声明channel的时候需要为其指定y元素类型
声明channel格式如下：
var 变量 chan 元素类型
channel需要初始化分配内存地址才可以使用
make(chan 元素类型,缓冲区大小)
*/

/**
通道操作
channel发送值都是用 <-
发送 ch1 <- 1
接收 x := <- ch1
关闭 用close函数来关闭通道
*/

var b chan int //需要指定通道中元素类型 并且初始化才可以使用 用make函数初始化
var a chan *string
var wg sync.WaitGroup

/**
带缓冲区的通道 传入的是内存地址
*/
func bufChannel() {
	a = make(chan *string, 16) //带缓冲区的初始化 这样初始化就有16个如果超过就需要等
	wg.Add(1)
	go func() {
		defer wg.Done()
		s := <-a
		fmt.Println("从通道A中获取数据", *s)
	}()
	value := "dddd"
	a <- &value
	fmt.Println(a)
	wg.Wait()
	close(a)
}

/**
不带缓冲区的通道
*/
func noBufChannel() {
	//通道初始化
	b = make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		x := <-b
		fmt.Println("从后台goroutine中去到了值:", x)
	}()
	//向通道b中发送数据
	b <- 10
	wg.Wait()
	fmt.Println(b)
	close(b)
}

/**
单向通道
func f1(ch1 chan <- int){}
确保暴露出去的通道只能放数据不能取
*/
func singleChannel(ch1 chan<- int) {

}

func main() {
	bufChannel()
}
