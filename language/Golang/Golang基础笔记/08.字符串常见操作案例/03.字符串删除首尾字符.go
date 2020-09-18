package main

import (
	"fmt"
	"strings"
)

func main() {

	s1 := "2020YinZhengjie"
	fmt.Println(s1)
	s2 := "尹正杰2020"
	fmt.Println(s2)
	s3 := "尹正杰2020到此一游"
	fmt.Println(s3)

	/**
	Trim函数的签名如下所示:
		func Trim(s string, cutset string) string

	以下是对函数签名相关参数的说明:
		s:
			表示待操作的字符串
		cutset:
			表示需要删除的首尾字符串
		string:
			返回值是将待操作的字符串(s)经过删除首尾字符串(cutset)后的字符串结果返回。

	温馨提示:
		通常使用该函数去除字符串中包含的多余空格。
	*/
	result1 := strings.Trim(s1, "2020")
	result2 := strings.Trim(s2, "2020")
	result3 := strings.Trim(s3, "2020")
	fmt.Println()

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)

}
