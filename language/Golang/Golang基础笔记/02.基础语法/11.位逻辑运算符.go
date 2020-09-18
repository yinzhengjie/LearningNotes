package main

import (
	"fmt"
)

func main() {

	a := 5 //对应二进制为:	0000 0101
	b := 7 //对应二进制为:	0000 0111

	fmt.Printf("变量a对应的二进制值是:[%b]\n", a)

	fmt.Printf("变量b对应的二进制值是:[%b]\n", b)

	/*
		&:
			位与运算符,表示AND(表示所有条件都得匹配),运算规则为相同位都是1时结果才为1，不同则为0。举个例子:如"5 & 7",结果为5。

		|:
			位或运算符,表示OR(表示有一个条件匹配即可),运算规则为相同位只要一个为1则为1。举个例子:如"5 | 7",结果为7。

		^:
			位异或运算符,表示XOR,运算规则为相同位不同则为1，相同则为0。举个例子:如"5 ^ 7",结果为2。

		&^:
			位清空运算符，表示AND NOT，运算规则为后数为0，则用前数对应位代替，后数为1则取0。举个例子:如"5 ^ 7",结果为0。
	*/
	fmt.Printf("a & b = [%d]\n", (a & b))
	fmt.Printf("a | b = [%d]\n", (a | b))
	fmt.Printf("a ^ b = [%d]\n", (a ^ b))
	fmt.Printf("a &^ b = [%d]\n", (a &^ b))

}
