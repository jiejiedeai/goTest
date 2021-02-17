package main

import (
	"fmt"
	"time"
)

/**
时间
*/
func main() {
	now := time.Now()
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	//时间戳
	timestamp1 := now.Unix()
	timestamp2 := now.UnixNano()
	//秒
	fmt.Println(timestamp1)
	//纳秒
	fmt.Println(timestamp2)
	//时间戳转换成时间
	ret1 := time.Unix(timestamp1, 0)
	fmt.Println(ret1)
	fmt.Println(ret1.Date())
	//时间间隔
	second := time.Second //  1秒
	fmt.Println(second)
	//当前时间 1小时后
	fmt.Println(now.Add(time.Hour))
	//sub 求时间差值
	fmt.Println(now.Add(time.Hour).Sub(now))
	//获取1天前的时间
	fmt.Println(now.AddDate(0, 0, -1))
	//获取10分钟前的时间
	fmt.Println(now.Add(-time.Minute * 10))
	//比较时间用Equal After判断该改时间之后 Before判断在该时间之前
	//定时器 1秒中执行1次
	timer := time.Tick(time.Second)
	for t := range timer {
		if t.Second() < 10 {
			fmt.Println(t)
		} else {
			break
		}

	}
	//时间格式化 但用的不是常见的模板y-m-d H:m:s形式
	//使用的是2006-01-02 15:04:05.000 是Go语言诞生的时间
	fmt.Println(now.Format("20060102"))
	//字符转转换成时间
	timeParse, err := time.Parse("2006-01-02", "2021-02-11")
	if err != nil {
		fmt.Println("时间格式转换失败")
	}
	fmt.Println(timeParse)

}
