package main

import (
	"fmt"
)

func main() {

	/*
		定义匿名函数时直接调用
	*/
	res := func(x int, y int) (z int) {
		z = x + y
		return z
	}(100, 20)

	fmt.Printf("res的类型为[%T],res的值为[%d]\n", res, res)
}
