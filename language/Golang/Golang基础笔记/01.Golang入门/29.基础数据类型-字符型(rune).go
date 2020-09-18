package main

import "fmt"

func main() {

	/*
		rune代表一个UTF-8字符(支持中文编码)，rune类型是int32的别名，在builtin.go文件中可以看到如下的代码:
			type rune = int32

		温馨提示:
			自动推导类型创建字符默认均是rune;
			rune是可以向下兼容byte,毕竟int32的数值范围要比uint8大很多倍。
	*/

	//自动推导类型创建字符默认是rune类型;
	FirstName := '尹'

	fmt.Println(FirstName)

	fmt.Printf("FirstName的类型为:[%T],FirstName的值为:[%c]\n", FirstName, FirstName)
}
