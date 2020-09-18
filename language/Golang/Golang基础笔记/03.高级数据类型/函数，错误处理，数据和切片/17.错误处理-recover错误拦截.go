package main

import "fmt"

func test3(index int) {
	/*
		错误拦截要在产生错误前设置，因此建议大家把错误拦截的函数放在函数内部的首行定义。
	*/
	defer func() {
		/*
			运行时panic异常一旦被引发就会导致程序崩溃
			Go语言提供了专用于"拦截"运行时panic的内建函数"recover"。
			它可以使当前的程序从运行时panic的状态中恢复并重新获得流程控制权，欢聚话说，通过recover进行不可控的错误拦截，重新获取程序的控制权
		*/
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	/*
		定义容量为10的数组
	*/
	var arr [10]int

	if index >= 10 {
		panic("请注意，index > 10,出现了索引越界异常...(index的取值范围0~9)")
	}

	arr[index] = 123

	fmt.Println(arr)
}

func main() {
	test3(5)
	test3(12)
	fmt.Println("代码执行完毕")
}
