package main

import "fmt"

func main() {

	var a int = 10086
	var b int64 = 10010
	fmt.Printf("a的类型为%T,a的值为%d\n", a, a)
	fmt.Printf("b的类型为%T,b的值为%d\n", b, b)

	/*
		注意运算符的优先级，会先执行加法算数运算符，在执行逗号运算符

		int和int相关类型在进行计算式需要进行类型转换，如下所示。
	*/
	fmt.Printf("a + b = %d\n", int64(a)+b)
}
