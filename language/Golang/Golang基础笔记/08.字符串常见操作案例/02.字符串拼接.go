package main

import (
	"fmt"
	"strings"
)

func main() {

	//定义一个字符串切片
	s1 := []string{
		"YinZhengjie",
		"尹正杰",
		"2020",
	}
	fmt.Println(s1)

	/**
	Join的函数签名如下所示:
		Join(a []string, sep string) string
	对函数签名个参数说明如下:
		a:
			表示待拼接的字符串切片.
		sep:
			表示指定拼接的链接符号，如果不想使用任何拼接符号可以使用空串("")
		string:
			返回值就是将字符串切片(a)按照拼接符号(sep)进行拼接的一个最终结果
	*/
	result := strings.Join(s1, "-")
	fmt.Println(result)
}
