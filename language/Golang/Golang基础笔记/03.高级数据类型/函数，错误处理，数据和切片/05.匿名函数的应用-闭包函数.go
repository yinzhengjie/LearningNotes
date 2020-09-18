package main

import (
	"fmt"
)

/*
	什么是闭包函数:
		闭包：闭是封闭（函数内部函数），包是包含（该内部函数对外部作用域而非全局作用域的变量的引用）。
		闭包指的是：函数内部函数对外部作用域而非全局作用域的引用。

	Go语言支持匿名函数作为闭包。匿名函数是一个"内联"语句或表达式。

	下面实例中，创建了函数"getSequence()",返回另外一个匿名函数"func() int"。该函数的目的在闭包中递增number变量。
*/
func getSequence() func() int {
	number := 100
	return func() int {
		/*
			匿名函数的优越性在于可以直接使用函数内的变量，不必声明。
		*/
		number += 1
		return number
	}
}

func main() {
	/*
		f1为一个空参匿名函数类型，number变量的值依旧为100
	*/
	f1 := getSequence()

	/*
		调用f1函数，number变量自增1并返回
	*/
	fmt.Println(f1())
	fmt.Println(f1())
	fmt.Println(f1())

	fmt.Println("=====我是分割线=====")
	/*
		创建新的匿名函数 f2，并查看结果
	*/
	f2 := getSequence()
	fmt.Println(f2())
	fmt.Println(f2())
}
