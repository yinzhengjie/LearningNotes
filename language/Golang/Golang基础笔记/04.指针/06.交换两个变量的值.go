package main

import "fmt"

/*
	定义一个函数交换2个变量的值，接收的是2个int变量
*/
func swap(x int, y int) {
	x, y = y, x
}

/*
	定义一个函数交换2个变量的值，接收的是2个int指针变量
*/
func swapPointer(x *int, y *int) {
	*x, *y = *y, *x //需要改变的是地址上存储的内容
}

func main() {
	/*
		定义2个变量，一会用于在函数中交换两个变量的值
	*/
	var (
		a = 100
		b = 200
	)

	/*
		值传递，不能通过函数内部修改影响外部变量。
	*/
	swap(a, b)
	fmt.Printf("a = [%d], b = [%d]\n", a, b)

	/*
		指针变量作为函数参数传递的时候是引用传递
	*/
	swapPointer(&a, &b)
	fmt.Printf("a = [%d], b = [%d]\n", a, b)
}
