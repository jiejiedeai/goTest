package main

import "fmt"

/**
select 多路复用
多个通道取值 同一时刻随机取值 那个通道有值就从那里边取
select{
	case <- ch1:
		xxxx
	case <- ch2:
		xxxxx
}
上述就可以随机取值不是按照顺序一个个判断
*/
func main() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:

		}
	}

}
