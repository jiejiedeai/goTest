package main

import "fmt"

type student struct {
	id   int64
	name string
}

var (
	allStudent map[int64]*student //声明

)

func newStudent(id int64, name string) *student {
	return &student{
		id:   id,
		name: name,
	}
}

func addStudent() {
	//创建学生添加到allStudent中
	var (
		id   int64
		name string
	)
	fmt.Println("请输入学生的学号:")
	fmt.Scanln(&id)
	flag := studentIsExists(id)
	if !flag{
		fmt.Println("学生已存在")
	}
	fmt.Println("请输入学生的姓名:")
	fmt.Scanln(&name)
	s := newStudent(id, name)
	allStudent[id] = s

}

func delStudent() {
	var id int64
	fmt.Println("请输入要删除的学生的学号:")
	fmt.Scanln(&id)
	flag := studentIsExists(id)
	if !flag{
		fmt.Println("学生已存在")
	}
	delete(allStudent, id)
}

func getAllStudent() {
	for _, v := range allStudent {
		//此处v是指针 因为go的语法糖此处可以不写*student
		fmt.Printf("学号:%d 姓名:%s\n", v.id, v.name)
	}
}

func studentIsExists(id int64) bool{
	s := allStudent[id]
	if nil !=s {
		return false
	}
	return true
}

func main() {
	//初始化开辟内存空间 map自动扩容
	allStudent = make(map[int64]*student, 50)

	fmt.Println("欢迎光临学生管理系统")
	fmt.Println(`
		1.查看所有学生
		2.新增学生
		3.删除学生`)
	for {
		fmt.Println("请输入你要干啥:")
		var choice int
		fmt.Scanln(&choice)
		fmt.Printf("你选择了%d这个选项\n", choice)
		//执行对应函数
		switch choice {
		case 1:
			getAllStudent()
		case 2:
			addStudent()
		case 3:
			delStudent()
		default:
			fmt.Println("滚！！！")
		}
	}
}
