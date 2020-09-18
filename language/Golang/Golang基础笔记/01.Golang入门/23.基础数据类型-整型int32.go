package main

import "fmt"

func main() {

	/*
		int32:
		　　占用四个字节，取值范围为-2147483648 ~ 2147483647，初始值为0。
		uint32:
		　　占用四个字节，取值范围为0 ~ 4294967295(42亿),初始化为0。
	*/
	var a int32 = -2147483648
	var b uint32 = 4294967295

	fmt.Printf("a的类型为:[%T],a的值为:[%d]\n", a, a)
	fmt.Printf("b的类型为:[%T],b的值为:[%d]\n", b, b)
}
