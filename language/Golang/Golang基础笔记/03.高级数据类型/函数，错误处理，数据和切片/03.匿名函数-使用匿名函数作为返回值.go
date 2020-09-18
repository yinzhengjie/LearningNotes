package main

import (
	"fmt"
)

//使用type来定义一个匿名函数类型
type FUNCTYPE func(int, int) int

func demo() FUNCTYPE {
	/*
		demo的返回值为我们上面定义的匿名函数类型
	*/
	return func(x int, y int) int {
		res := x + y
		return res
	}
}

func main() {
	/*

		add的类型为(匿名)函数类型
	*/
	add := demo()
	fmt.Printf("add的类型为[%T],add匿名函数的内存地址是[%X]\n", add, add)

	/*
		我们可以通过函数类型add来多次调用匿名函数
	*/
	res1 := add(100, 200)
	res2 := add(300, 500)
	fmt.Printf("res1的类型为[%T],res1的值为[%d]\n", res1, res1)
	fmt.Printf("res2的类型为[%T],res2的值为[%d]\n", res2, res2)
}
