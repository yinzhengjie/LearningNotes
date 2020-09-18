package main

import (
	"fmt"
)

/*
	函数的返回值是通过函数中的return语句获得的，return后面的值也可以是一个表达式,只要返回值类型和定义的返回值列表所匹配即可。
	Go语言支持多个返回值。
*/
func test() (x int, y float64, z string) {

	return 18, 3.14, "尹正杰"
}

func main() {

	/*
		如果函数定义了多个返回值，就需要使用多个变量来接收这些返回值
		可以使用匿名变量("_")来接收不使用的变量的值,因此我们无法将匿名变量的值取出来
	*/
	a, _, c := test()
	fmt.Println(a)
	fmt.Println(c)

}
