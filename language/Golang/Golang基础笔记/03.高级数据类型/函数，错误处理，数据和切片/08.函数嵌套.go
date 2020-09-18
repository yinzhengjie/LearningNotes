package main

import (
	"fmt"
)

func add1(x int, y int) int {
	fmt.Println("in add1...")
	return x + y
}

/*
	什么是嵌套函数:
		其实就是在一个函数中调用另外的函数
*/
func add2(x int, y int) int {
	fmt.Println("in add2...")
	return add1(x, y)
}

func main() {

	res := add2(100, 20)

	fmt.Println(res)
}
