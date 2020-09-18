package main

import (
	"fmt"
	"strings"
)

func main() {

	s1 := "YinZhengjie2020"
	fmt.Println(s1)
	fmt.Println()

	/**
	Replace函数的签名如下所示:
		func Replace(s, old, new string, n int) string

	以下是对函数签名相关参数说明:
		s:
			表示待操作的字符串
		old:
			表示在代操作的字符串中需要替换的子串
		new:
			表示将old子串替换成新的字符串(new)
		n:
			表示替换的个数，当n为正数时,仅替换相应个数的子串,当n为负数时,则不限制替换的个数,即全部匹配。
	*/
	result1 := strings.Replace(s1, "YinZhengjie", "尹正杰", 1)
	fmt.Println(result1)

	result2 := strings.Replace(s1, "e", "E", 2)
	fmt.Println(result2)

	result3 := strings.Replace(s1, "n", "N", -1)
	fmt.Println(result3)
}
