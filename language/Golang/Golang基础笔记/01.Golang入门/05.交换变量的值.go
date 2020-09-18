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
			引入第三方变量temp来交换变量的值。
		优点:
			通俗易懂，生产环境中建议使用这种方式。
		缺点:
			多引入了一个变量，会占用内存。
	*/
	temp := x
	x = y
	y = temp

	fmt.Printf("变量发生交换之后,x的值为[%d],y的值为[%d]\n", x, y)
}
