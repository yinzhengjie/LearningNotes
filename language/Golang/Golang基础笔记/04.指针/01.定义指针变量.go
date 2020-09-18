package main

import (
	"fmt"
)

func main() {
	/**
	在Go语言中,一旦定义了一个变量就会在内存中开辟空间。因此,在Go语言中为了避免空间的浪费,定义变量若未使用就会编译报错。
	*/
	var age uint8 = 18
	/**
	我们可以使用取地址运算符("&")来获取数据的内存地址。
	%p:
		是一个占位符,表示输出一个十六进制地址格式。
	\n:
		表示换行符。
	*/
	fmt.Printf("age的值是:[%d],age的内存地址是:[%p]\n", age, &age)

	/*
		接下来我们定义一个uint8类型的指针变量p1，其实指针变量也是变量，只不过指针变量指向了变量的内存地址。
	*/
	var p1 *uint8

	/*
		对age变量取地址并将结果赋值给指针变量。
	*/
	p1 = &age
	fmt.Println(p1)
}
