package main

import (
	"fmt"
)

func main() {

	var Name string = "尹正杰"

	blog := "博客地址:https://www.cnblogs.com/yinzhengjie"

	/*
		len(t Type):
			用于计算数据类型的长度
	*/
	fmt.Printf("Name的长度为:[%d],Name的值为:[%s]\n", len(Name), Name)
	fmt.Printf("blog的长度为:[%d],blog的值为:[%s]\n", len(blog), blog)

}
