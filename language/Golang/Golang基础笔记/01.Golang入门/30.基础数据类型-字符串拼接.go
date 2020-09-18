package main

import (
	"fmt"
)

func main() {

	var Name string = "尹正杰"

	blog := "博客地址:https://www.cnblogs.com/yinzhengjie"

	fmt.Printf("Name的类型为:[%T],Name的值为:[%s]\n", Name, Name)
	fmt.Printf("blog的类型为:[%T],blog的值为:[%s]\n", blog, blog)

	/*
		+:
			可以将两个字符串的值进行拼接
		==:
			判断两个字符串内容是否相同
	*/
	res := Name + blog
	fmt.Printf("res的类型为:[%T],res的值为:[%s]\n", res, res)

	fmt.Printf("Name和blog的值是否相等:[%t]\n", Name == blog)
}
