package main

import (
	"fmt"
)

func myArr(index int) {

	//定义数组
	var arr [10]int

	//数组这里赋值时可能会存在一个风险，即索引越界异常
	arr[index] = 123

	fmt.Println(arr)
}

func main() {
	myArr(12)
}
