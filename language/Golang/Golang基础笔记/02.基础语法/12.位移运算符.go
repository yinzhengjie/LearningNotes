package main

import (
	"fmt"
)

func main() {

	a := 5 //对应二进制为:	0000 0101

	fmt.Printf("变量a对应的二进制值是:[%b]\n", a)

	/*
		<<:
			左移，表示将对应的二进制数字向左移动相应的位数，比如 5 << 3,结果为40。

		>>:
			右移，表示将对应的二进制数字向右移动相应的位数，比如 5 >> 3,结果为1。
	*/

	fmt.Printf("5 << 3 = [%d]\n", 5<<3)
	fmt.Printf("5 >> 3 = [%d]\n", 5>>3)

}
