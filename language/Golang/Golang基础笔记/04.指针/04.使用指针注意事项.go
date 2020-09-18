package main

import (
	"fmt"
)

func main() {

	/**
	温馨提示:
		在使用指针变量时,要注意定义指针默认值为nil,如果直接擦欧总指向nil的内存地址会报错;
		在其它语言中,比如c++，指针是允许运算的,但是在Go语言中，指针只有取地址运算符和取值运算符,不允许指针参与运算哟~
	*/
	var p *int

	/**
	错误代码:
		由于p指针变量默认值是nil,因此会报错"invalid memory address or nil pointer dereference"
	*/
	//*p = 123

	age := 18
	p = &age //所以,在使用指针变量时,一定要让指针变量有正确的指向。
	fmt.Printf("age的内存地址是:%p,age的值是:%d\n", p, *p)

	/**
	错误代码:
		在Go语言中，指针只有取地址运算符和取值运算符。
	*/
	//p2 := p + 100
}
