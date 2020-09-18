package main

import (
	"fmt"
)

func main() {
	a := 28
	b := 18

	fmt.Printf("判断a和b是否相等，即a == b ,结果为:[%t]\n", a == b)

	fmt.Printf("判断a和b是否不相等，即a != b ,结果为:[%t]\n", a != b)

	fmt.Printf("判断a是否大于b，即a > b ,结果为:[%t]\n", a > b)

	fmt.Printf("判断a是否小于b，即a < b ,结果为:[%t]\n", a < b)

	fmt.Printf("判断a是否大于等于b，即a >= b ,结果为:[%t]\n", a >= b)

	fmt.Printf("判断a是否小于等于b，即a <= b ,结果为:[%t]\n", a <= b)

}
