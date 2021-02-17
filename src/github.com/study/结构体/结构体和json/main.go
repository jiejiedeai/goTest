package main

import (
	json2 "encoding/json"
	"fmt"
)

//结构体转变成json
type person struct {
	//属性小写只能在本包可见如果想在别的包可见需要大写
	//因为格式化功能是由json这个包的Marshal帮我们做的
	//小写的话再json包就拿不到 所以在结构体中加上反引号
	//`json:"name"` db:"name" ini:"name" 可以写多个 这样在这些包中就是小写
	Name string `json:"name"`
	Age  int	`json:"age"`
}

func newPerson(name string, age int) person {
	return person{
		Name: name,
		Age:  age,
	}
}

func main() {
	p := newPerson("乔鹏", 10)
	//导入go语言中js的包 然后将结构体序列化
	b, err := json2.Marshal(p)
	if err != nil {
		fmt.Printf("marshal failed,err:%v",err)
	}
	//可以直接强制转换成字符串 因为字符串本身就是由字节组成的
	personJson := string(b)
	fmt.Printf("%#v\n",personJson)
	//将json转换成结构体是反序列化 第一个参数是json字符串 第二个参数是拷贝到那里去
	var p2 person
	json2.Unmarshal([]byte(personJson),&p2)
	fmt.Printf("%#v\n",p2)// 由于go中都是值拷贝所以要传入指针
}
