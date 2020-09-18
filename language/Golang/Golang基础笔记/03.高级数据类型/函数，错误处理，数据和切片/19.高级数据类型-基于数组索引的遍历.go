package main

import (
	"fmt"
)

func main() {

	var arr [5]int = [5]int{100, 80, 300, 500, 1024}

	/*
		基于数组的索引下标进行遍历
	*/
	for index := 0; index < len(arr); index++ {
		fmt.Printf("arr数组索引下标[%d]保存的值是: [%d]\n", index, arr[index])
	}

}
