package main

import (
	"fmt"
)

func main() {
	/*
		温馨提示:
			以数字0开始表示该数字是八进制
		%b:
			是一个占位符，表示一个二进制(Binary,缩写BIN)格式的数字
	*/
	year := 03744
	fmt.Printf("八进制03744对应的二进制表示为:[%b]\n", year)

	/*
		%d:
			是一个占位符，表示一个十进制(Decimal，缩写DEC)格式的数字
	*/
	fmt.Printf("八进制03744对应的十进制表示为:[%d]\n", year)

	/*
		%x|%X:
			是一个占位符，表示一个十六进制(Hexadecimal,缩写HEX)格式的数字
	*/
	fmt.Printf("八进制03744对应的十六进制表示为:[%x]\n", year)
	fmt.Printf("八进制03744对应的十六进制表示为:[%X]\n", year)

}
