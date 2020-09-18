package main

import (
	"fmt"
)

func main() {
	//定义除数
	a := 20

	//定义被除数
	b := 0

	//除数不能为0，编译时并不会报错，但是在代码运行时会报错哟~
	c := a / b

	//由于上面的代码执行报错啦，该行代码不会被执行
	fmt.Printf(" %d ➗ %d = %d\n", a, b, c)
}
