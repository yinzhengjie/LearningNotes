package main

import (
	"fmt"
)

/*
	不定参函数定义:
		计算N个整形数据的和
*/
func sum(arr ...int) int {
	value := 0
	/*
		使用数组下标进行遍历
	*/
	//for index := 0; index < len(arr); index++ {
	//	value += arr[index]
	//}

	/*
		使用range关键字进行范围遍历,range会从集合中返回两个数:
			第一个是对应的坐标,赋值给了匿名变量"_"
			第二个对应的是值，赋值给了变量"data"
	*/
	for _, data := range arr {
		value += data
	}
	return value
}

func main() {
	/*
		我们在调用函数时可以指定函数参数的个数不尽相同。
	*/
	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}
