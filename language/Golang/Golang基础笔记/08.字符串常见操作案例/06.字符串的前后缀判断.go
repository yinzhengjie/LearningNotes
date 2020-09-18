package main

import (
	"fmt"
	"strings"
)

func main() {

	s1 := "yinzhengjie2020"
	fmt.Println(s1)

	/**
	HasPrefix函数签名如下:
		func HasPrefix(s, prefix string) bool

	以下是对函数签名的相关参数说明:
		s:
			表示待匹配的字符串
		prefix:
			指定一个字符串，判断待匹配的字符串(s)是否以指定字符串(prefix)开头的
		bool:
			返回的结果是一个布尔值，如果是指定的字符串(prefix)开头，则返回true，否则返回false.
	*/
	result1 := strings.HasPrefix(s1, "yin")
	fmt.Printf("s1是以'yin'字符开头,结果为: %t\n", result1)

	/**
	HasSuffix函数签名如下:
		func HasSuffix(s, suffix string) bool

	以下是对函数签名的相关参数说明:
		s:
			表示待匹配的字符串
		suffix:
			指定一个字符串，判断待匹配的字符串(s)是否以指定字符串(suffix)结尾的
		bool:
			返回的结果是一个布尔值，如果是指定的字符串(suffix)结尾，则返回true，否则返回false.
	*/
	result2 := strings.HasSuffix(s1, "2020")
	fmt.Printf("s1是以'2020'字符结尾,结果为: %t\n", result2)

}
