package main

import (
	"fmt"
)

func main() {
	x := 10
	y := 20

	fmt.Printf("变量未发生交换之前,x的值为[%d],y的值为[%d]\n", x, y)

	/*
		思路:
			使用加法("+")和减法("-")运算符来交换变量的值。
		优点:
			没有引入任何新的变量就完成了变量值的互换。
		缺点:
			当x和y的值足够大时，容易超出x和y的数值范围哟~
	*/
	x = x + y
	y = x - y
	x = x - y

	fmt.Printf("变量发生交换之后,x的值为[%d],y的值为[%d]\n", x, y)
}
