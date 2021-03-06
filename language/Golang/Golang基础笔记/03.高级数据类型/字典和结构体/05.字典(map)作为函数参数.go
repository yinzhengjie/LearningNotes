package main

import (
	"fmt"
)

func Rename(m map[string]string) {
	//为传递进来的map增加一个key.
	m["name"] = "Jason Yin"
	fmt.Printf("Rename函数中的m地址为:%p\n", m)
}

func main() {
	/*
		在Go语言中，数组作为参数进行传递时值传递，而切片作为参数进行传递时引用传递。
			值传递:
				方法调用时，实参数把他的值传递给对应的形式参数，方法执行中形式参数值的改变不会影响实际参数的值。
			引用传递(也称为传地址):
				函数调用时，实际参数的引用(地址，而不是参数的值)被传递给函数中相对应的形式参数(实参与形参指向了同一块存储区域);
				在函数执行时,对形式参数的操作实际上就是对实际参数的操作，方法执行中形式参数值的改变会影响实际参数的值。

		map作为函数参数传递实际上和切片传递一样,传递的是地址,也就是我们常说的引用传递.

		温馨提示:
			(1)说白了,只要先使用make进行初始化操作再使用的类型,在函数传递时基本上都是引用传递哟~
			(2)在我们日常开发中,常见引用传递的高级数据类型有切片,字典(map)和管道(channel).
	*/
	m1 := make(map[string]string)

	fmt.Println("调用前的m1数据为:", m1)
	fmt.Printf("main函数中的m1地址为:%p\n", m1)

	Rename(m1)

	fmt.Println("调用后的m1数据为:", m1)
}
