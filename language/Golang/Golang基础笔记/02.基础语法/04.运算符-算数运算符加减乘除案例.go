package main

import (
	"fmt"
)

func main() {

	a := 23
	b := 8

	fmt.Printf("a的类型为:[%T],a的值为[%d]\n", a, a)
	fmt.Printf("b的类型为:[%T],b的值为[%d]\n", b, b)

	fmt.Println("a + b = ", a+b) //变量a和变量b进行加法运算

	fmt.Println("a - b = ", a-b) //变量a和变量b进行减法运算

	fmt.Println("a * b = ", a*b) //变量a和变量b进行乘法运算

	/*
		除法运算时注意事项:
			两个整数类型在进行触发计算时，结果为整数,规则是向下取整
			在除法中，除法不能为0，否则会抛除零异常(panic: runtime error: integer divide by zero)
	*/
	fmt.Println("a / b = ", a/b)

	/*
		取模(余)运算注意事项:
			只能在整型中使用
	*/
	fmt.Println("a % b = ", a%b)

}
