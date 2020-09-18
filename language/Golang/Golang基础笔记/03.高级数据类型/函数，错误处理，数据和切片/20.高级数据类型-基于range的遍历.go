package main

import (
	"fmt"
)

func main() {

	var arr [5]int = [5]int{100, 80, 300, 500, 1024}

	/*
		通过range关键字遍历数组，可以同时得到数组的索引下标和改下表对应的数据，
		通过range实现范围遍历，这一点和Python中的enumerate方法有着异曲同工之妙。
	*/
	for index, value := range arr {
		fmt.Printf("arr数组索引下标[%d]保存的值是: [%d]\n", index, value)
	}
}
