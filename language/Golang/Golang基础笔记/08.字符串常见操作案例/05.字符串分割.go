package main

import (
	"fmt"
	"strings"
)

func main() {

	s1 := "YinZhengjie           尹正杰           2020"
	fmt.Println(s1)
	fmt.Println()

	/**
	Split函数的签名如下所示:
		func Split(s, sep string) []string

	以下时对函数签名的相关参数说明:
		s:
			表示待操作的字符串
		sep:
			表示以那个字符串为切割
		[]string:
			返回值是一个字符串切片数组。

	温馨提示:
		Split的作用和join相反,是把字符串按照指定的子串切割成字符串切片
	*/
	result1 := strings.Split(s1, " ") //此处我们按照空格来进行切分
	fmt.Println(result1)

	/**
	如下所示，通过查看Fields函数源码,默认是按照指定的特定的分隔符进行切割的.
		"[256]uint8{'\t': 1, '\n': 1, '\v': 1, '\f': 1, '\r': 1, ' ': 1}"进行切割
	*/
	result2 := strings.Fields(s1)
	fmt.Println(result2)

}
