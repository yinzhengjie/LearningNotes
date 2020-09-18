package main

import "fmt"

func main() {
	var year int
	fmt.Printf("year的类型为:[%T],year默认初始化值为[%d]\n", year, year)

	year = 2020 //赋值运算符
	fmt.Printf("赋值运算符,year的值为[%d]\n", year)

	year += 10 //加法赋值运算符，等效于year = year + 10
	fmt.Printf("加法赋值运算符,year的值为[%d]\n", year)

	year -= 10 //减法赋值运算符，等效于year = year - 10
	fmt.Printf("减法赋值运算符,year的值为[%d]\n", year)

	year *= 10 //乘法赋值运算符，等效于year = year * 10
	fmt.Printf("乘法赋值运算符,year的值为[%d]\n", year)

	year /= 10 //除法赋值运算符，等效于year = year / 10
	fmt.Printf("除法赋值运算符,year的值为[%d]\n", year)

	year %= 10 //取余赋值运算符，等效于year = year % 10
	fmt.Printf("取余赋值运算符,year的值为[%d]\n", year)
}
