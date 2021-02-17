package main

import "fmt"

/**
%T 查看类型
%d 十进制数
%b 二进制数
%o 八进制数
%x 十六进制数小写 a-f
%X 十六进制数大写 A-F
%U 表示为Unicode格式
%c 字符
%s 字符串
%p 指针
%f 浮点数
%v 值的默认格式
%+v 类似%v 但输出结构体时会添加字段名
%#v 值的打印结果会比%v更详细(会带字符串 数据格式类型）
%% 格式化输出转义
%t 布尔值
%q 数字转成字符加一个单引号 'A' 字符串直接加双引号 "hahaha"
%f 默认宽度默认精度
%9f 宽度9 默认精度
%.2f 默认宽度 精度2
%9.2f 宽度9 精度2
%9.f 宽度9 精度0
 */
func main() {
	n := 12.34
	fmt.Printf("%f\n",n)
	fmt.Printf("%9f\n",n)
	fmt.Printf("%.2f\n",n)
	fmt.Printf("%9.2f\n",n)
	fmt.Printf("%9.f\n",n)
}
