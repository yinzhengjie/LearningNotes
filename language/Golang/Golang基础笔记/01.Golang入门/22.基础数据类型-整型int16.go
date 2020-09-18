package main

import "fmt"

func main() {

	/*
			int16:
		　　　　占用两个字节，取值范围为-32768~32767，初始值为0。
			uint16:
		　　　　占用两个字节,取值范围为0~65535，初始化值为0。
	*/
	var a int16 = -32768
	var b uint16 = 65535

	fmt.Printf("a的类型为:[%T],a的值为:[%d]\n", a, a)
	fmt.Printf("b的类型为:[%T],b的值为:[%d]\n", b, b)
}
