package main

import (
	"fmt"
)

/*
	形参和实参操作的数组不是同一个地址中的数据，所以修改被调函数中的元素的值并不会影响主函数中实参数组元素的值。
*/
func modify(arr [5]int) {

	arr[2] = 100 //修改数组元素的值

	fmt.Println("in modify: ", arr) //被掉函数中，数组元素被修改

	fmt.Printf("被调函数中形参数组的地址:%p\n", &arr) //打印被调函数的的参数地址:
}

func main() {
	arr := [5]int{9, 5, 2, 7, 1}

	modify(arr) //将数组作为函数参数，其实是值传递，形参不能改变实参的值

	fmt.Println("主函数: ", arr) //主函数中数组元素并未被修改

	fmt.Printf("主调函数中形参数组的地址:%p\n", &arr) //打印主函数的的参数地址:
}
