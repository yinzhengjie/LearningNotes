package main

import (
	"fmt"
)

/*
	什么是递归函数:
		如果一个函数在内部不调用其它的函数，而是调用自己本身，这个函数就是递归函数。

	递归函数的应用场景:
		电商网站中的商品类别菜单的应用
		查找某个目录下的文件

	定义递归函数注意事项:
		(1)函数嵌套调用函数本身
		(2)使用return指定函数出口
*/
var total = 1

func factorial(num int) {
	/*
		递归函数需要定义递归函数的结束条件，否则会出现"死递归"的现象,出现"死递归"情况程序就自动会抛出"fatal error: stack overflow"异常。
	*/
	if num == 0 {
		return
	}
	total *= num

	/*
		函数内部自己调用自己,那么这个函数就是递归函数。
	*/
	factorial(num - 1)
}

func main() {

	factorial(5)

	fmt.Printf("5的阶乘是[%d]\n", total)
}
