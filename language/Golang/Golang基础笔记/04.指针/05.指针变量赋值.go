package main

import (
	"fmt"
)

func main() {
	/**
	为指针赋值有两种方法:
		1>.把同类型的变量地址赋值给指针;
		2>.使用new函数，就是申请一片内存空间，返回地址;
	*/
	var (
		p1 *int
		p2 *int
	)
	fmt.Println("p1的默认值是:", p1)
	fmt.Println("p2的默认值是:", p2)

	/**
	方案一:
		定义一个同类型的变量,然后取地址赋值给指针变量。
	*/
	age := 18
	p1 = &age
	fmt.Printf("p1的内存地址是:%p,p1的值是:%d\n", p1, *p1)

	/**
	方案二:
		使用new函数,比如"new(int)"表示创建一个int大小的内存空间，返回为*int
	*/
	p2 = new(int)
	fmt.Printf("p2的内存地址是:%p,p2的值是:%d\n", p2, *p2)
	*p2 = 2020 //为指针变量赋值
	fmt.Printf("p2的内存地址是:%p,p2的值是:%d\n", p2, *p2)
}
