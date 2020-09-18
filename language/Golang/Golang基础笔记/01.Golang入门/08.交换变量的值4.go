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
			使用多重赋值交换两个变量的值
		优点:
			没有引入任何新的变量和运算符就实现了变量值的互换,效率非常高，生产环境中推荐这种写法。
		缺点:
			暂无
	*/
	x, y = y, x

	fmt.Printf("变量发生交换之后,x的值为[%d],y的值为[%d]\n", x, y)
}
