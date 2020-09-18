package main

import (
	"fmt"
)

func main() {
	//定义一个字符串数组
	var nameList [56]string
	fmt.Printf("nameLists数组的内存地址为%p\n", &nameList)
	fmt.Printf("nameList数组中第一个元素的地址为%p\n", &nameList[0])
	fmt.Printf("nameList数组中第二个元素的地址为%p\n", &nameList[1])
	fmt.Printf("nameList数组中第三个元素的地址为%p\n\n", &nameList[2])

	//定义一个整型数组
	var age_list [56]int
	fmt.Printf("age_list数组的内存地址为%p\n", &age_list)
	fmt.Printf("age_list数组中第一个元素的地址为%p\n", &age_list[0])
	fmt.Printf("age_list数组中第二个元素的地址为%p\n", &age_list[1])
	fmt.Printf("age_list数组中第三个元素的地址为%p\n\n", &age_list[2])

	//定义一个整型切片
	age_slice := make([]int, 56)
	/**
	age_slice切片变量保存的是存储切片容器的地址。他并不像数组那样直接指向数组的首元素地址,它类似于一个二级指针。
	当我们使用取地址运算符对age_slice变量取地址时并不是取的切片真正存储切片容器的地址,而是存储切片容器地址变量本身的地址哟~
	*/
	fmt.Printf("age_slice切片变量本身的内存地址为%p\n", &age_slice)

	/**
	通过下标访问第一个元素其实访问的是age_slice变量中保存的内存地址对应的下标。
	*/
	fmt.Printf("age_slice切片中第一个元素的地址为%p\n", &age_slice[0])
	fmt.Printf("age_slice切片中第二个元素的地址为%p\n", &age_slice[1])
	fmt.Printf("age_slice切片中第三个元素的地址为%p\n\n", &age_slice[2])
}
