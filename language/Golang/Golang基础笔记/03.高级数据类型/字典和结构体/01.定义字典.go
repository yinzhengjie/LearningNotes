package main

import (
	"fmt"
)

func main() {
	/*

		声明字典结构语法如下:
			var 字典 map[键类型]值类型

		定义字典结构使用map关键字,"[]"中指定的是键(key)的类型，后面紧跟着的是值(value)的类型。

		温馨提示:
			map中的key值除了切片，函数，复数(complex)以及包含切片的结构体都可以,换句话说,使用这些类型会造成编译错误。
			map在使用前也需要使用make函数进行初始化。
			map没有容量属性，map只有长度属性，长度表示的是map中key和value有多少对。
			map满足集合的特性,即key是不能重复的
	*/

	//声明一个字典类型
	var m1 map[string]string
	//map在使用前必须初始化空间，和切片类似的是map自身也没有空间哟~
	m1 = make(map[string]string)
	//注意,key和value都是字符串类型
	m1["Name"] = "Jason Yin"
	//注意,上一行已经定义"Name"这个key名称了,再次使用同名key会将上一个key对应的value进行覆盖
	m1["Name"] = "尹正杰"
	fmt.Printf("m1的数据类型是:%T,对应的长度是:%d\n", m1, len(m1))
	fmt.Println("m1的数据是:", m1)

	//使用自动推到类型并初始化空间
	m2 := make(map[string]int)
	//注意key是字符串类型,而vlaue是int类型
	m2["Age"] = 18
	fmt.Printf("m2的数据类型是:%T,对应的长度是:%d\n", m2, len(m2))
	fmt.Println("m2的数据是:", m2)

	//直接初始化空间并赋初始值
	m3 := map[string]rune{"first": '尹', "second": '正', "third": '杰'}
	fmt.Printf("m3的数据类型是:%T,对应的长度是:%d\n", m3, len(m3))
	fmt.Println("m3的数据是:", m3)
}
