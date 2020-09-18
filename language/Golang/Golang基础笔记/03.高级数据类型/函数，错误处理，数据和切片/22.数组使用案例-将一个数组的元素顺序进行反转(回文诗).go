package main

import (
	"fmt"
)

func main() {
	var arr []rune = []rune{'湾', '前', '过', '渡', '小', '舟', '虚'}

	fmt.Printf("为发生交换之前，arr的数据为:[%s]\n", string(arr)) //使用string进行强制类型转换，打印中文

	//最小值索引下标
	minIndex := 0

	//最大值下标
	maxIndex := len(arr) - 1

	for minIndex < maxIndex {

		//使用多重赋值，交换两个遍历的值
		arr[minIndex], arr[maxIndex] = arr[maxIndex], arr[minIndex]

		minIndex++ //将最小索引下标加上1

		maxIndex-- //将最大索引下标减去1
	}

	fmt.Printf("发生交换之后，arr的数据为:[%s]\n", string(arr))
}
