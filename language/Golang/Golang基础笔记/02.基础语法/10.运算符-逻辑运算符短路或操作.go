package main

import (
	"fmt"
)

func main() {
	a := false
	b := false

	fmt.Printf("变量a的数据类型是:%T,其对应的值是:%t\n", a, a)
	fmt.Printf("变量b的数据类型是:%T,其对应的值是:%t\n", b, b)

	/*
		||:
			逻辑或，同假为假，其余为真
	*/
	fmt.Printf("a && b 的结果为[%t]\n", a && b)

	c := true
	d := false
	fmt.Printf("变量c的数据类型是:%T,其对应的值是:%t\n", c, c)
	fmt.Printf("变量d的数据类型是:%T,其对应的值是:%t\n", d, d)

	fmt.Printf("c && d 的结果为:[%t]\n", (c && d))
}
