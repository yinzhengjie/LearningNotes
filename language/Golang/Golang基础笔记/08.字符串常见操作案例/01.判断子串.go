package main

import (
	"fmt"
	"strings"
)

func main() {

	s1 := "尹正杰2020"

	s2 := "2020"

	/**
	Contains函数的签名如下所示:
		func Contains(s, substr string) bool

	对函数签名个参数说明如下:
		s:
			表示一个字符串
		substr:
			表示需要和s进行对比的子串
		bool:
			表示返回值类型，如果s包含substr则返回true，如果不包含则返回false.
	*/
	result := strings.Contains(s1, s2)

	fmt.Printf("result = %t\n", result)

	if result {
		fmt.Println("s2是s1的子串")
	} else {
		fmt.Println("s2不是s1的子串")
	}
}
