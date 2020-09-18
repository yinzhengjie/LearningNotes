package main

import (
	"fmt"
)

func main() {
	/*
		切片的截取:
			所谓截取就是从切片中获取指定的数据。

		切片常见的操作如下所示:
			"slice[n]":
				切片slice中索引位置的为n的选项。
			"slice[:]":
				从切片slice的索引位置0到len(slice)-1处获得的切片
			"slice[low:]":
				从切片slice的索引位置low到len(slice)-1处获得的切片
			"slice[:high]":
				从切片selice的索引位置0到high处所获得的切片，len=high-low
			"slice[low:high:max]":
				从切片slice的索引位置low到high处所获得的切片，len=high-low,cap=max-low
			"len(s)":
				表示切片s的长度，切片的长度总是小于等于(>=)切片容量的(cap(s))
			"cap(s)":
				表示切片s的容量，切片的容量总是大于等于(<=)切片的长度(len(s))
	*/

	slice1 := []rune{'尹', '正', '杰', '到', '此', '一', '游'}
	fmt.Printf("slice1的数据为:[%s],长度为:[%d],容量为:[%d]\n", string(slice1), len(slice1), cap(slice1))

	s1 := slice1[2] //获取切片索引为2的元素
	fmt.Printf("s1的数据为:[%s]\n", string(s1))

	s2 := slice1[:] //默认截取的长度和容量相等
	fmt.Printf("s2的数据为:[%s],长度为:[%d],容量为:[%d]\n", string(s2), len(s2), cap(s2))

	s3 := slice1[3:] //设置截取的起始索引，左闭右开
	fmt.Printf("s3的数据为:[%s],长度为:[%d],容量为:[%d]\n", string(s3), len(s3), cap(s3))

	s4 := slice1[:4] //设置截取的结束索引，左闭右开
	fmt.Printf("s4的数据为:[%s],长度为:[%d],容量为:[%d]\n", string(s4), len(s4), cap(s4))

	s5 := slice1[1:3] //设置截取的起始索引和结束索引,左闭右开
	fmt.Printf("s5的数据为:[%s],长度为:[%d],容量为:[%d]\n", string(s5), len(s5), cap(s5))

	s6 := slice1[1:3:6] //设置截取的起始索引和结束索引并指定容量
	fmt.Printf("s6的数据为:[%s],长度为:[%d],容量为:[%d]\n", string(s6), len(s6), cap(s6))
}
