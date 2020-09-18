package main

import (
	"fmt"
)

func main() {

	/*
		什么是切片(slice):
			切片是Go中一种比较特殊的数据结构，这种数据结构更便于使用和管理数据集合。
			切片是围绕动态数组的概念构建的，可以按需自动增长。

		切片的定义格式:
			语法一(定义空(nil)切片):
				var 切片名称 []数组类型

			语法二(通过make创建切片):
				var 切片名称 []数组类型 = make([]数据类型，长度，容量)

			语法三(自动推导类型初始化):
				切片名称 := []数组类型{元素1,元素2,...,元素N}
				切片名称 := make([]数据类型，长度，容量)

		切片的长度(len)与容量(cap):
			len(切片名称):
				计算切片有效数据个数,长度是已经初始化的空间,切片初始空间若为被赋值则使用对应数据类型的默认值。
			cap(切片名称):
				计算切片容量，容量是一家开辟的空间，包括一家初始化的空间和空闲的空间。

		常见数组类型的默认值:
			整型默认初始化值为0;
			浮点型默认初始化值为0;
			字符串类型默认初始化值为空串("");
			布尔类型默认初始化值为false;
	*/

	//定义一个空(nil)切片,空切片不能添加数据
	var slice1 []int
	fmt.Printf("slice1的数据为:%d,长度为:%d,容量为:%d\n", slice1, len(slice1), cap(slice1))

	//通过make创建切片
	var slice2 []int = make([]int, 3, 5)
	fmt.Printf("slice2的数据为:%d,长度为:%d,容量为:%d\n", slice2, len(slice2), cap(slice2))

	//基于切片索引进行赋值时，其索引大小不能等于或超过切片的长度,否则就会报错"index out of range"
	slice2[1] = 666
	fmt.Printf("slice2的数据为:%d,长度为:%d,容量为:%d\n", slice2, len(slice2), cap(slice2))

	//自动推导类型创建切片
	slice3 := []int{1, 2, 3, 4, 5}
	slice4 := make([]int, 2, 5) //make函数可以认为是给切片申请空间
	fmt.Printf("slice3的数据为:%d,长度为:%d,容量为:%d\n", slice3, len(slice3), cap(slice3))
	fmt.Printf("slice4的数据为:%d,长度为:%d,容量为:%d\n", slice4, len(slice4), cap(slice4))

}
