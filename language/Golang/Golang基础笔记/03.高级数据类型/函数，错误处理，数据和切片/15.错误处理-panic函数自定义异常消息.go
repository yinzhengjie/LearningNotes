package main

import (
	"fmt"
)

func myArr2(index int) {

	//定义数组
	var arr [10]int

	if index >= 10 {
		/*
			panic函数与errors接口不同:
				(1)使用panic时无需导入包;
				(2)若出现异常，使用panic封装异常消息时，会直接导致程序运行结束，换句话说,error是一般性错误，而panic函数返回的是让程序崩溃的错误

			温馨提示:
				一般而言，当panic异常发生时，程序会中断运行。随后，程序崩溃并输出日志消息。日志信息包括panic value和函数调用的堆栈跟踪信息
				当然，如果直接调用内置的panic函数也会引发panic异常，panic函数接收任何值作为参数。
		*/
		panic("请注意，index > 10,出现了索引越界异常...(index的取值范围0~9)")
	}

	//上面使用了panic接口封装了异常处理错误，因此如果代码到这一行说明没有索引越界异常
	arr[index] = 123

	fmt.Println(arr)
}

func main() {

	myArr2(12)

	fmt.Println("代码执行完毕")
}
