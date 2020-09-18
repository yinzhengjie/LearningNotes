package main

import "fmt"

func main() {
	a := 100
	b := 20
	fmt.Printf("变量a的数据类型是:%T,其对应的值是:%d\n", a, a)
	fmt.Printf("变量b的数据类型是:%T,其对应的值是:%d\n", b, b)

	/*
		！:
			逻辑非，非真为假，非假为真
	*/
	fmt.Printf("a > b 结果为:[%t]\n", (a > b))
	fmt.Printf("a < b 结果为:[%t]\n", !(a > b))

	c := true
	fmt.Printf("变量c的数据类型是:%T,其对应的值是:[%t]\n", c, c)
	fmt.Printf("变量c的数据类型是:%T,其对应的值是:[%t]\n", !c, !c)
}
