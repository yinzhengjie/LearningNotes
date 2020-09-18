package main

import (
	"fmt"
)

type Operation struct {
	Name string
	Age  uint8
}

func main() {
	//使用自动推导类型直接进行初始化赋值
	op := Operation{"尹正杰", 18}
	fmt.Println("op的数据为:", op)

	//定义一个结构体切片
	var s1 []Operation

	//对结构体切片进行初始化
	s1 = make([]Operation, 3)

	fmt.Print("s1结构体初始化值为:")
	fmt.Println(s1)
	//对结构体切片第二个元素进行赋值
	s1[1] = op
	//我们可以对结构体切片的元素对应的属性进行修改操作
	s1[1].Name = "Jason Yin"
	fmt.Print("s1结构体被修改后的值为:")
	fmt.Println(s1)
}
