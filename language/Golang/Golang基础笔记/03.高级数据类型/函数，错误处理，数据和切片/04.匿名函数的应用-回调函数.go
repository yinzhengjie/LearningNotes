package main

import (
	"fmt"
)

/*
	函数回调:
		简称回调,英文名为:"Callback",即"call then back",被主函数调用运算后会返回主函数。
		是指通过函数参数传递到其它代码的，某一块可执行代码的引用

	匿名函数作为回调函数的设计在Go语言的系统包中是很常见的,比如strings包中又有着中实现，代码如下所示:
		func TrimFunc(s string, f func(rune) bool) string{
			return TrimRightFunc(TrimLeftFunc(s,f),f)
		}
*/
func callback(f func(int, int) int) int {
	return f(10, 20)
}

func add(x int, y int) int {
	return x + y
}

func main() {
	/*
		匿名函数(函数名本身是代码区的一个地址)的用途非常广泛,匿名函数本身是一种值,可以方便的保存在各种容器中实现回调函数和操作封装
	*/
	fmt.Println(add)

	/*
		函数回调操作
	*/
	fmt.Println(callback(add))
}
