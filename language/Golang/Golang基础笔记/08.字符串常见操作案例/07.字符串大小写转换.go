package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "YinZhengjie2020"
	fmt.Println(s1)

	//将字符串转换为大写
	upper := strings.ToUpper(s1)
	fmt.Println(upper)

	//将字符串转换成小写
	lower := strings.ToLower(s1)
	fmt.Println(lower)
}
