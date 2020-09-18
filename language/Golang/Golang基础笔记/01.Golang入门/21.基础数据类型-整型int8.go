package main

import (
	"fmt"
)

func main() {

	/*
			int8:
		　　　　占用一个字节大小，取值范围为-128~127,初始值为0。
		　　uint8:
		　　　　占用一个字节大小，取值范围为0~255,初始值为0。
	*/
	var a int8 = -128
	var b uint8 = 255

	fmt.Printf("a的类型为:[%T],a的值为:[%d]\n", a, a)
	fmt.Printf("b的类型为:[%T],b的值为:[%d]\n", b, b)
}
