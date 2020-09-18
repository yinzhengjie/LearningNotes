package main

import (
	"fmt"
)

func main() {
	Name := "yinzhengjie"

	/*
		&:
			取地址运算符，表示取一个变量的内存地址，它同时也属于一元操作符(也叫单目运算符)。
	*/
	fmt.Println(&Name)
	fmt.Printf("point的类型为[%T],Name变量的内存地址为[%x]\n", Name, &Name)
	fmt.Printf("point的类型为[%T],Name变量的内存地址为[%X]\n", Name, &Name)

	point := &Name //注意，point此时是一个指针变量哟~

	/*
		*:
			取值运算符，表示取一个变量所指向内存地址中保存的数据。
	*/
	fmt.Printf("point的类型为[%T],point的值为[%s]\n", point, *point)
}
