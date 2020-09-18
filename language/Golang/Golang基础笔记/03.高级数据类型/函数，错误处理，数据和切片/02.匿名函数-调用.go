package main

import (
	"fmt"
)

func main() {

	/*
		定义匿名函数，此时add是一个函数类型,只不过它是一个匿名函数。
	*/
	add := func(x int, y int) (z int) {
		z = x + y
		return z
	}
	fmt.Printf("add的类型为[%T]\n", add)

	/*
		我们可以通过函数类型add来多次调用匿名函数
	*/
	res1 := add(100, 200)
	res2 := add(300, 500)
	fmt.Printf("res1的类型为[%T],res1的值为[%d]\n", res1, res1)
	fmt.Printf("res2的类型为[%T],res2的值为[%d]\n", res2, res2)
}
