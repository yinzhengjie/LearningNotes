package main

import (
	"fmt"
)

func main() {

	/*
		温馨提示:
			以0x开始表示该数字是十六进制
		%b:
			是一个占位符，表示一个二进制(Binary,缩写BIN)格式的数字
	*/
	year := 0x7e4
	fmt.Printf("十六进制0x7e4对应的二进制表示为:[%b]\n", year)

	/*
		%o:
			是一个占位符，表示一个八进制(Octal,缩写OCT)格式的数字
	*/
	fmt.Printf("十六进制0x7e4对应的八进制表示为:[%o]\n", year)

	/*
		%d:
			是一个占位符，表示一个十进制(Decimal，缩写DEC)格式的数字
	*/
	fmt.Printf("十六进制0x7e4对应的十进制表示为:[%d]\n", year)

}
