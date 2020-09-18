package main

import (
	"fmt"
)

func main() {
	/*
		什么是变量:
			在程序运行过程中其值可以发生改变的量称为变量。

		在golang中有三种定义变量的方式，分别为"声明变量"，"变量赋值",“自动推导类型”。
			(1)声明变量语法格式:
				var 变量名称 数据类型
			(2)变量赋值语法格式:
				var 变量名称 数据类型 = 值
			(3)自动推导类型(根据值的类型确定变量名的类型)语法格式:
				变量名 := 值
	*/

	//1>.声明变量
	var name string
	name = "Jason Yin"
	fmt.Println(name)
	fmt.Printf("变量name的数据类型是:%T\n", name)

	//2>.变量赋值
	var age int = 18
	fmt.Println(age)
	fmt.Printf("变量age的数据类型是:%T\n", age)

	//3>.自动推导类型
	blog := "https://www.cnblogs.com/yinzhengjie/"
	fmt.Println(blog)
	fmt.Printf("变量blog的数据类型是:%T\n", blog)
}
