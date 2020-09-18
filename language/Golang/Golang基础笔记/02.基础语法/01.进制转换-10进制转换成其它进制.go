package main

import (
	"fmt"
)

func main() {

	/*
		温馨提示:
			自动推导类型的变量默认是十进制(Decimal，缩写DEC)的数字,这可能和咱们生活中的习惯方式有关。
		%b:
			是一个占位符，表示一个二进制(Binary,缩写BIN)格式的数字
	*/
	year := 2020
	fmt.Printf("十进制2020对应的二进制表示为:[%b]\n", year)

	/*
		%o:
			是一个占位符，表示一个八进制(Octal,缩写OCT)格式的数字
	*/
	fmt.Printf("十进制2020对应的八进制表示为:[%o]\n", year)

	/*
		%x|%X:
			是一个占位符，表示一个十六进制(Hexadecimal,缩写HEX)格式的数字
	*/
	fmt.Printf("十进制2020对应的十六进制表示为:[%x]\n", year)
	fmt.Printf("十进制2020对应的十六进制表示为:[%X]\n", year)
}
